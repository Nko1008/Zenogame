package main

import "fmt"

const (
	sanctuaryCost = 10
	questReward   = 20
)

// Nekomata : le chat blanc qui guide le joueur.
func guide(c *Character) {
	for {
		fmt.Println("\n--- Nekomata, le chat blanc ---")
		fmt.Println("1. Un conseil")
		fmt.Println("2. Cadeau de bienvenue")
		fmt.Println("3. Conseils de combat")
		fmt.Println("4. Objectifs")
		fmt.Printf("5. Sanctuaire (soin complet : %d or)\n", sanctuaryCost)
		fmt.Println("6. Qui sont-ils ?")
		fmt.Println("0. Retour")

		switch readChoice("> ") {
		case 1:
			advice(c)
		case 2:
			gift(c)
		case 3:
			combatTips()
		case 4:
			quest(c)
		case 5:
			sanctuary(c)
		case 6:
			whoAreThey()
		case 0:
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

// owns : vrai si l'objet est dans l'inventaire ou déjà équipé.
func owns(c *Character, item string) bool {
	return countItem(c, item) > 0 ||
		c.Equipment.Head == item || c.Equipment.Torso == item || c.Equipment.Feet == item
}

func canCraft(c *Character, r recipe) bool {
	if c.Money < forgeCost {
		return false
	}
	for mat, qty := range r.Materials {
		if countItem(c, mat) < qty {
			return false
		}
	}
	return true
}

// Idée 1 : conseil selon la situation
func advice(c *Character) {
	if c.Health < c.MaxHealth/2 {
		fmt.Println("Nekomata : « Tes PV sont bas. Bois une potion ou passe au sanctuaire avant de te battre. »")
		return
	}
	if len(c.Inventory) >= c.InventoryMax {
		fmt.Println("Nekomata : « Ton inventaire est plein. Pense à l'Augmentation d'inventaire chez le marchand. »")
		return
	}
	for _, r := range recipes {
		if canCraft(c, r) {
			fmt.Printf("Nekomata : « Tu as de quoi fabriquer : %s. Va voir le forgeron. »\n", r.Result)
			return
		}
	}
	fmt.Println("Nekomata : « Tout va bien. Récolte des ressources et fais un tour chez le marchand. »")
}

// Idée 3 : cadeau unique
func gift(c *Character) {
	if c.GiftReceived {
		fmt.Println("Nekomata : « Je t'ai déjà fait mon cadeau. »")
		return
	}
	if !addInventory(c, "Potion de vie") {
		fmt.Println("Nekomata : « Fais de la place dans ton inventaire, puis reviens. »")
		return
	}
	c.GiftReceived = true
	fmt.Println("Nekomata te donne : Potion de vie")
}

// Idée 4 : conseils de combat
func combatTips() {
	fmt.Println("Nekomata : « Celui qui a le plus d'initiative joue en premier. »")
	fmt.Println("Nekomata : « Le gobelin frappe deux fois plus fort tous les 3 tours. Prépare-toi ! »")
}

// Idée 5 : objectif avec récompense
func quest(c *Character) {
	if c.QuestDone {
		fmt.Println("Nekomata : « Tu as déjà accompli mon objectif, bravo ! »")
		return
	}
	items := []string{itemHat, itemTunic, itemBoots}
	have := 0
	for _, it := range items {
		if owns(c, it) {
			have++
		}
	}
	fmt.Printf("Objectif : fabriquer les 3 équipements de l'aventurier (%d/3)\n", have)
	if have == 3 {
		c.Money += questReward
		c.QuestDone = true
		fmt.Printf("Objectif accompli ! Récompense : %d or (or total : %d)\n", questReward, c.Money)
	}
}

// Idée 6 : sanctuaire
func sanctuary(c *Character) {
	if c.Health >= c.MaxHealth {
		fmt.Println("Nekomata : « Tu es déjà en pleine forme. »")
		return
	}
	if c.Money < sanctuaryCost {
		fmt.Println("Nekomata : « Il te faut", sanctuaryCost, "or pour te soigner. »")
		return
	}
	c.Money -= sanctuaryCost
	c.Health = c.MaxHealth
	fmt.Printf("Nekomata te soigne. PV : %d / %d (or restant : %d)\n", c.Health, c.MaxHealth, c.Money)
}
