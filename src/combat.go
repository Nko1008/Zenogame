package main

import "fmt"

func playerStarts(c *Character, m *Monster) bool {
	return c.Initiative >= m.Initiative
}

func characterTurn(c *Character, m *Monster) {
	fmt.Println("\n--- Votre tour ---")
	fmt.Println("1. Attaquer")
	fmt.Println("2. Inventaire")
	fmt.Println("3. Sorts")

	switch readChoice("> ") {
	case 1:
		dmg := 5
		m.Health -= dmg
		if m.Health < 0 {
			m.Health = 0
		}
		fmt.Println("Vous utilisez Attaque basique")
		fmt.Printf("%s inflige %d dégâts à %s\n", c.Name, dmg, m.Name)
		fmt.Printf("%s PV : %d / %d\n", m.Name, m.Health, m.MaxHealth)

	case 2:
		if len(c.Inventory) == 0 {
			fmt.Println("Votre inventaire est vide.")
			return
		}
		fmt.Println("Inventaire :")
		for i, it := range c.Inventory {
			fmt.Printf("%d. %s\n", i+1, it)
		}
		choice := readChoice("> ")
		if choice < 1 || choice > len(c.Inventory) {
			fmt.Println("Choix invalide.")
			return
		}
		item := c.Inventory[choice-1]
		useItem(c, item)

	case 3:
		if len(c.Skills) == 0 {
			fmt.Println("Vous ne connaissez aucun sort.")
			return
		}
		fmt.Println("Sorts :")
		for i, s := range c.Skills {
			fmt.Printf("%d. %s (%d mana)\n", i+1, s, spellCosts[s])
		}
		choice := readChoice("> ")
		if choice < 1 || choice > len(c.Skills) {
			fmt.Println("Choix invalide.")
			return
		}
		castSpell(c, m, c.Skills[choice-1])

	default:
		fmt.Println("Choix invalide.")
	}
}

func trainingFight(c *Character) {
	m := initGoblin()
	turn := 1

	fmt.Println("\n=== Un Gobelin d'entrainement apparaît ! ===")

	for m.Health > 0 {
		fmt.Printf("\n----- Tour %d -----\n", turn)

		if playerStarts(c, &m) {
			characterTurn(c, &m)
			if m.Health > 0 {
				goblinPattern(&m, c, turn)
				isDead(c)
			}
		} else {
			goblinPattern(&m, c, turn)
			isDead(c)
			characterTurn(c, &m)
		}
		turn++
	}

	fmt.Println("\nVictoire ! Le gobelin est vaincu.")
	gainExperience(c, m.Experience)
	fmt.Println("Retour au menu principal.")
}

func bossFight(c *Character) {
	boss := initNoxar()
	turn := 1

	fmt.Printf("\n=== %s apparaît ! ===\n", boss.Name)

	for boss.Health > 0 {
		fmt.Printf("\n----- Tour %d -----\n", turn)

		if playerStarts(c, &boss) {
			characterTurn(c, &boss)
			if boss.Health > 0 {
				bossPattern(&boss, c, turn)
				isDead(c)
			}
		} else {
			bossPattern(&boss, c, turn)
			isDead(c)
			characterTurn(c, &boss)
		}
		turn++
	}

	fmt.Printf("\n%s est vaincu ! Vous avez gagné le combat final !\n", boss.Name)
	gainExperience(c, boss.Experience)
	fmt.Println("Retour au menu principal.")
}
