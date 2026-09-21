package main

import (
	"bufio"
	"fmt"
)

func playerStarts(c *Character, m *Monster) bool {
	return c.Initiative >= m.Initiative
}

func characterTurn(c *Character, m *Monster, reader *bufio.Reader) {
	fmt.Println("\n--- Votre tour ---")
	fmt.Println("1. Attaquer")
	fmt.Println("2. Inventaire")
	fmt.Println("3. Sorts")
	fmt.Print("Choix : ")

	choice, _ := reader.ReadString('\n')
	choice = trimNewline(choice)

	switch choice {
	case "1":
		dmg := 5
		m.Health -= dmg
		if m.Health < 0 {
			m.Health = 0
		}
		fmt.Println("Vous utilisez Attaque basique")
		fmt.Printf("%s inflige %d dégâts à %s\n", c.Name, dmg, m.Name)
		fmt.Printf("%s PV : %d / %d\n", m.Name, m.Health, m.MaxHealth)

	case "3":
		if len(c.Skills) == 0 {
			fmt.Println("Vous ne connaissez aucun sort.")
			return
		}
		fmt.Println("Sorts :")
		for i, s := range c.Skills {
			fmt.Printf("%d. %s (%d mana)\n", i+1, s, spellCosts[s])
		}
		fmt.Print("Choisissez un sort : ")
		spellChoice, _ := reader.ReadString('\n')
		spellChoice = trimNewline(spellChoice)

		index := -1
		fmt.Sscanf(spellChoice, "%d", &index)
		if index < 1 || index > len(c.Skills) {
			fmt.Println("Choix invalide.")
			return
		}

		if !castSpell(c, m, c.Skills[index-1]) {
			return
		}

	case "2":
		if len(c.Inventory) == 0 {
			fmt.Println("Votre inventaire est vide.")
			return
		}
		fmt.Println("Inventaire :")
		for i, item := range c.Inventory {
			fmt.Printf("%d. %s\n", i+1, item)
		}
		fmt.Print("Choisissez un objet : ")
		itemChoice, _ := reader.ReadString('\n')
		itemChoice = trimNewline(itemChoice)

		index := -1
		fmt.Sscanf(itemChoice, "%d", &index)
		if index < 1 || index > len(c.Inventory) {
			fmt.Println("Choix invalide.")
			return
		}

		item := c.Inventory[index-1]
		useItem(c, item)
		c.Inventory = append(c.Inventory[:index-1], c.Inventory[index:]...)

	default:
		fmt.Println("Choix invalide.")
	}
}

func useItem(c *Character, item string) {
	switch item {
	case "Potion":
		fmt.Println("Vous utilisez Potion de vie")
		c.Health += 50
		if c.Health > c.MaxHealth {
			c.Health = c.MaxHealth
		}
		fmt.Printf("%s PV : %d / %d\n", c.Name, c.Health, c.MaxHealth)
	case "Livre de Sort : Star Shot":
		spellBook(&c.Skills)
	default:
		fmt.Printf("%s n'a aucun effet en combat.\n", item)
	}
}

func trimNewline(s string) string {
	for len(s) > 0 && (s[len(s)-1] == '\n' || s[len(s)-1] == '\r') {
		s = s[:len(s)-1]
	}
	return s
}

func trainingFight(c *Character, reader *bufio.Reader) {
	m := initGoblin()
	turn := 1

	fmt.Println("\n=== Un Gobelin d'entrainement apparaît ! ===")

	for c.Health > 0 && m.Health > 0 {
		fmt.Printf("\n----- Tour %d -----\n", turn)

		if playerStarts(c, &m) {
			characterTurn(c, &m, reader)
			if m.Health > 0 {
				goblinPattern(&m, c, turn)
			}
		} else {
			goblinPattern(&m, c, turn)
			if c.Health > 0 {
				characterTurn(c, &m, reader)
			}
		}
		turn++
	}

	if c.Health <= 0 {
		fmt.Println("\nVous avez été vaincu...")
	} else {
		fmt.Println("\nVictoire ! Le gobelin est vaincu.")
		gainExperience(c, m.Experience)
	}
	fmt.Println("Retour au menu principal.")
}

func bossFight(c *Character, reader *bufio.Reader) {
	boss := initNoxar()
	turn := 1

	fmt.Printf("\n=== %s apparaît ! ===\n", boss.Name)

	for c.Health > 0 && boss.Health > 0 {
		fmt.Printf("\n----- Tour %d -----\n", turn)

		if playerStarts(c, &boss) {
			characterTurn(c, &boss, reader)
			if boss.Health > 0 {
				bossPattern(&boss, c, turn)
			}
		} else {
			bossPattern(&boss, c, turn)
			if c.Health > 0 {
				characterTurn(c, &boss, reader)
			}
		}
		turn++
	}

	if c.Health <= 0 {
		fmt.Println("\nVous avez été vaincu...")
	} else {
		fmt.Printf("\n%s est vaincu ! Vous avez gagné le combat final !\n", boss.Name)
		gainExperience(c, boss.Experience)
	}
	fmt.Println("Retour au menu principal.")
}
