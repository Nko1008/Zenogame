package main

import "fmt"

func playerStarts(c *Character, m *Monster) bool {
	return c.Initiative >= m.Initiative
}

func characterTurn(c *Character, m *Monster) bool {
	fmt.Println(cyan("\n--- Votre tour ---"))
	fmt.Println("1. Attaquer")
	fmt.Println("2. Inventaire")
	fmt.Println("3. Sorts")
	fmt.Println("4. Fuir le combat")

	switch readChoice("> ") {
	case 1:
		dmg := 5
		m.Health -= dmg
		if m.Health < 0 {
			m.Health = 0
		}
		fmt.Println("Vous utilisez Attaque basique")
		fmt.Println(green(fmt.Sprintf("%s inflige %d dégâts à %s", c.Name, dmg, m.Name)))
		fmt.Printf("%s PV : %s\n", m.Name, healthColor(m.Health, m.MaxHealth))

	case 2:
		if len(c.Inventory) == 0 {
			fmt.Println(red("Votre inventaire est vide."))
			return false
		}
		fmt.Println("Inventaire :")
		for i, it := range c.Inventory {
			fmt.Printf("%d. %s\n", i+1, it)
		}
		choice := readChoice("> ")
		if choice < 1 || choice > len(c.Inventory) {
			fmt.Println(red("Choix invalide."))
			return false
		}
		item := c.Inventory[choice-1]
		useItem(c, item)

	case 3:
		if len(c.Skills) == 0 {
			fmt.Println(red("Vous ne connaissez aucun sort."))
			return false
		}
		fmt.Println("Sorts :")
		for i, s := range c.Skills {
			fmt.Printf("%d. %s (%d mana)\n", i+1, s, spellCosts[s])
		}
		choice := readChoice("> ")
		if choice < 1 || choice > len(c.Skills) {
			fmt.Println(red("Choix invalide."))
			return false
		}
		castSpell(c, m, c.Skills[choice-1])

	case 4:
		fmt.Println(yellow("Vous prenez la fuite..."))
		return true

	default:
		fmt.Println(red("Choix invalide."))
	}

	return false
}

func trainingFight(c *Character) {
	m := initGoblin()
	turn := 1

	fmt.Println(yellow("\n=== Un Gobelin d'entrainement apparaît ! ==="))
	displayClassASCII("Gobelin d'entrainement")

	for m.Health > 0 {
		fmt.Println(cyan(fmt.Sprintf("\n----- Tour %d -----", turn)))

		var fled bool
		if playerStarts(c, &m) {
			fled = characterTurn(c, &m)
			if !fled && m.Health > 0 {
				goblinPattern(&m, c, turn)
				isDead(c)
			}
		} else {
			goblinPattern(&m, c, turn)
			isDead(c)
			fled = characterTurn(c, &m)
		}
		if fled {
			fmt.Println(yellow("\nVous avez fui le combat."))
			fmt.Println("Retour au menu principal.")
			return
		}
		turn++
	}

	fmt.Println(green("\nVictoire ! Le gobelin est vaincu."))
	displayClassASCII("WINNER")
	gainExperience(c, m.Experience)
	gainGold(c, m.Gold)
	dropMaterial(c)
	fmt.Println("Retour au menu principal.")
}

func bossDeath(c *Character, boss *Monster) bool {
	if c.Health <= 0 {
		displayClassASCII("LOSER")
		fmt.Println(red(fmt.Sprintf("\n%s succombe face à %s...", c.Name, boss.Name)))
		fmt.Println("Retour au menu principal.")
		return true
	}
	return false
}

func bossFight(c *Character) {
	boss := initNoxar()
	turn := 1

	fmt.Println(red(fmt.Sprintf("\n=== %s apparaît ! ===", boss.Name)))
	displayClassASCII("Noxar apparaît !")

	for boss.Health > 0 {
		fmt.Println(cyan(fmt.Sprintf("\n----- Tour %d -----", turn)))

		var fled bool

		if playerStarts(c, &boss) {
			fled = characterTurn(c, &boss)

			if !fled && boss.Health > 0 {
				bossPattern(&boss, c, turn)

				if bossDeath(c, &boss) {
					return
				}
			}
		} else {
			bossPattern(&boss, c, turn)

			if bossDeath(c, &boss) {
				return
			}

			fled = characterTurn(c, &boss)
		}

		if fled {
			fmt.Println(yellow("\nVous avez fui le combat."))
			fmt.Println("Retour au menu principal.")
			return
		}

		turn++
	}

	fmt.Println(green(fmt.Sprintf("\n%s est vaincu ! Vous avez gagné le combat final !", boss.Name)))

	c.NoxarDefeated = true
	gainExperience(c, boss.Experience)
	gainGold(c, boss.Gold)

	fmt.Println("Retour au menu principal.")
}
