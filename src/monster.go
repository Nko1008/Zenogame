package main

import (
	"fmt"
	"math/rand"
)

type Monster struct {
	Name       string
	MaxHealth  int
	Health     int
	Attack     int
	Experience int
	Initiative int
	Gold       int
}

func initGoblin() Monster {
	return Monster{
		Name:       "Gobelin d'entrainement",
		MaxHealth:  40,
		Health:     40,
		Attack:     5,
		Experience: 20,
		Initiative: 5,
		Gold:       15,
	}
}

func goblinPattern(m *Monster, c *Character, turn int) {
	dmg := m.Attack
	if turn%3 == 0 {
		dmg *= 2
	}
	c.Health -= dmg
	if c.Health < 0 {
		c.Health = 0
	}
	fmt.Println(red(fmt.Sprintf("%s inflige à %s %d de dégâts", m.Name, c.Name, dmg)))
	fmt.Printf("%s PV : %s\n", c.Name, healthColor(c.Health, c.MaxHealth))
}

var goblinDrops = []string{"Fourrure de Loup", "Peau de Troll", "Cuir de Sanglier", "Plume de Corbeau"}

func dropMaterial(c *Character) {
	mat := goblinDrops[rand.Intn(len(goblinDrops))]
	if addInventory(c, mat) {
		fmt.Println(green(fmt.Sprintf("Le gobelin laisse tomber : %s", mat)))
	} else {
		fmt.Println(red("Le gobelin laisse tomber un objet, mais votre inventaire est plein !"))
	}
}

func initNoxar() Monster {
	return Monster{
		Name:       "Noxar",
		MaxHealth:  120,
		Health:     120,
		Attack:     12,
		Experience: 100,
		Initiative: 8,
		Gold:       150,
	}
}

func bossPattern(m *Monster, c *Character, turn int) {
	dmg := m.Attack
	attackName := "Attaque basique"
	if turn%4 == 0 {
		dmg *= 3
		attackName = "Attaque spéciale"
	}
	c.Health -= dmg
	if c.Health < 0 {
		c.Health = 0
	}
	fmt.Println(red(fmt.Sprintf("%s utilise %s et inflige à %s %d de dégâts", m.Name, attackName, c.Name, dmg)))
	fmt.Printf("%s PV : %s\n", c.Name, healthColor(c.Health, c.MaxHealth))
}

// healthColor colore l'affichage des PV : vert si en bonne santé,
// jaune si moyen, rouge si critique.
func healthColor(current, max int) string {
	text := fmt.Sprintf("%d / %d", current, max)
	if max <= 0 {
		return text
	}
	ratio := float64(current) / float64(max)
	switch {
	case ratio <= 0.25:
		return red(text)
	case ratio <= 0.5:
		return yellow(text)
	default:
		return green(text)
	}
}
