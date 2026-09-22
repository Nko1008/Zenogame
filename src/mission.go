package main

import "fmt"

func gainExperience(c *Character, amount int) {
	fmt.Printf("%s gagne %d points d'expérience\n", c.Name, amount)
	c.Experience += amount

	for c.Experience >= c.ExperienceMax {
		c.Experience -= c.ExperienceMax
		c.Level++
		c.MaxHealth += 10
		c.MaxMana += 5
		c.ExperienceMax = int(float64(c.ExperienceMax) * 1.5)
		fmt.Printf("%s passe niveau %d ! (+10 PV max, +5 Mana max)\n", c.Name, c.Level)
	}

	fmt.Printf("Expérience : %d / %d\n", c.Experience, c.ExperienceMax)
}

func gainGold(c *Character, amount int) {
	c.Money += amount
	fmt.Printf("%s gagne %d pièces d'or (Or : %d)\n", c.Name, amount, c.Money)
}

var spellCosts = map[string]int{
	"Coup de poing": 5,
	"Star Shot":     15,
}

var spellDamage = map[string]int{
	"Coup de poing": 8,
	"Star Shot":     18,
}

func knowsSpell(c *Character, spell string) bool {
	for _, s := range c.Skills {
		if s == spell {
			return true
		}
	}
	return false
}

func castSpell(c *Character, m *Monster, spell string) bool {
	if !knowsSpell(c, spell) {
		fmt.Println("Vous ne connaissez pas ce sort.")
		return false
	}

	cost, ok := spellCosts[spell]
	if !ok {
		fmt.Println("Sort inconnu.")
		return false
	}

	if c.Mana < cost {
		fmt.Println("Mana insuffisant pour lancer", spell)
		return false
	}

	c.Mana -= cost
	dmg := spellDamage[spell]
	m.Health -= dmg
	if m.Health < 0 {
		m.Health = 0
	}

	fmt.Printf("%s lance %s (-%d mana) et inflige %d dégâts à %s\n", c.Name, spell, cost, dmg, m.Name)
	fmt.Printf("%s PV : %d / %d\n", m.Name, m.Health, m.MaxHealth)
	fmt.Printf("%s Mana : %d / %d\n", c.Name, c.Mana, c.MaxMana)

	return true
}
