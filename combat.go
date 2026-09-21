package main

import "fmt"

// Mission 1 : le personnage commence si son initiative est >= à celle du monstre.
func playerStarts(c *Character, m *Monster) bool {
	return c.Initiative >= m.Initiative
}

func characterTurn(c *Character, m *Monster) {
	for {
		fmt.Println("\n1. Attaquer")
		fmt.Println("2. Inventaire")
		switch readChoice("> ") {
		case 1:
			dmg := 5
			m.HP -= dmg
			if m.HP < 0 {
				m.HP = 0
			}
			fmt.Printf("%s utilise Attaque basique et inflige %d dégâts à %s\n", c.Name, dmg, m.Name)
			fmt.Printf("%s PV : %d / %d\n", m.Name, m.HP, m.MaxHP)
			return
		case 2:
			fmt.Println("0. Retour")
			for i, it := range c.Inventory {
				fmt.Printf("%d. %s\n", i+1, it)
			}
			choice := readChoice("> ")
			if choice < 1 || choice > len(c.Inventory) {
				continue // retour au menu de combat sans perdre son tour
			}
			item := c.Inventory[choice-1]
			fmt.Printf("Vous utilisez %s\n", item)
			useItem(c, item)
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func trainingFight(c *Character) {
	m := initGoblin()
	turn := 1
	for c.HP > 0 && m.HP > 0 {
		fmt.Printf("\n===== Tour %d =====\n", turn)
		if playerStarts(c, &m) {
			characterTurn(c, &m)
			if m.HP > 0 {
				goblinPattern(&m, c, turn)
			}
		} else {
			goblinPattern(&m, c, turn)
			if c.HP > 0 {
				characterTurn(c, &m)
			}
		}
		turn++
	}
	if c.HP <= 0 {
		fmt.Println("Vous avez été vaincu...")
	} else {
		fmt.Println("Victoire ! Le gobelin est vaincu.")
	}
}
