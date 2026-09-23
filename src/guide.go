package main

import "fmt"

const (
	sanctuaryCost = 10
	questReward   = 20
)

func guide(c *Character) {
	displayClassASCII("Nekomata")

	for {
		fmt.Println(cyan("\n--- Nekomata, le chat blanc ---"))
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
			fmt.Println(red("Choix invalide."))
		}
	}
}

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

func nekomata(s string) string {
	return cyan("Nekomata : « " + s + " »")
}

func advice(c *Character) {
	if c.Health < c.MaxHealth/2 {
		fmt.Println(nekomata("Tes PV sont bas. Bois une potion ou passe au sanctuaire avant de te battre."))
		return
	}
	if len(c.Inventory) >= c.InventoryMax {
		fmt.Println(nekomata("Ton inventaire est plein. Pense à l'Augmentation d'inventaire chez le marchand."))
		return
	}
	for _, r := range recipes {
		if canCraft(c, r) {
			fmt.Println(cyan(fmt.Sprintf("Nekomata : « Tu as de quoi fabriquer : %s. Va voir le forgeron. »", r.Result)))
			return
		}
	}
	fmt.Println(nekomata("Tout va bien. Récolte des ressources et fais un tour chez le marchand."))
}

func gift(c *Character) {
	if c.GiftReceived {
		fmt.Println(nekomata("Je t'ai déjà fait mon cadeau."))
		return
	}
	if !addInventory(c, "Potion de vie") {
		fmt.Println(nekomata("Fais de la place dans ton inventaire, puis reviens."))
		return
	}
	c.GiftReceived = true
	fmt.Println(green("Nekomata te donne : Potion de vie"))
}

func combatTips() {
	fmt.Println(nekomata("Celui qui a le plus d'initiative joue en premier."))
	fmt.Println(nekomata("Le gobelin frappe deux fois plus fort tous les 3 tours. Prépare-toi !"))
}

func quest(c *Character) {
	if c.QuestDone {
		fmt.Println(nekomata("Tu as déjà accompli mon objectif, bravo !"))
		return
	}
	if !c.NoxarDefeated {
		fmt.Println(yellow("Objectif : vaincre Noxar (0/1)"))
		fmt.Println(nekomata("Prépare-toi bien avant d'affronter Noxar : équipement, potions et sorts !"))
		return
	}
	fmt.Println(green("Objectif : vaincre Noxar (1/1)"))
	c.Money += questReward
	c.QuestDone = true
	fmt.Println(yellow(fmt.Sprintf("Objectif accompli ! Récompense : %d or (or total : %d)", questReward, c.Money)))
}

func sanctuary(c *Character) {
	if c.Health >= c.MaxHealth {
		fmt.Println(nekomata("Tu es déjà en pleine forme."))
		return
	}
	if c.Money < sanctuaryCost {
		fmt.Println(nekomata(fmt.Sprintf("Il te faut %d or pour te soigner.", sanctuaryCost)))
		return
	}
	c.Money -= sanctuaryCost
	c.Health = c.MaxHealth
	fmt.Println(green(fmt.Sprintf("Nekomata te soigne. PV : %d / %d (or restant : %d)", c.Health, c.MaxHealth, c.Money)))
}
