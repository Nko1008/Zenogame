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
	fmt.Printf("%s inflige à %s %d de dégâts\n", m.Name, c.Name, dmg)
	fmt.Printf("%s PV : %d / %d\n", c.Name, c.Health, c.MaxHealth)
}

var goblinDrops = []string{"Fourrure de Loup", "Peau de Troll", "Cuir de Sanglier", "Plume de Corbeau"}


func dropMaterial(c *Character) {
	mat := goblinDrops[rand.Intn(len(goblinDrops))]
	if addInventory(c, mat) {
		fmt.Printf("Le gobelin laisse tomber : %s\n", mat)
	} else {
		fmt.Println("Le gobelin laisse tomber un objet, mais votre inventaire est plein !")
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
	fmt.Printf("%s utilise %s et inflige à %s %d de dégâts\n", m.Name, attackName, c.Name, dmg)
	fmt.Printf("%s PV : %d / %d\n", c.Name, c.Health, c.MaxHealth)
}
