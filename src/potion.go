package main

import (
	"fmt"
	"time"
)

func takePot(c *Character) {
	if !removeInventory(c, "Potion de vie", 1) {
		fmt.Println(red("Vous n'avez pas de Potion de vie."))
		return
	}
	fmt.Println(green("Vous utilisez une Potion de vie."))
	c.Health += 50
	if c.Health > c.MaxHealth {
		c.Health = c.MaxHealth
	}
	fmt.Printf("PV : %s\n", healthColor(c.Health, c.MaxHealth))
}

func manaPot(c *Character) {
	if !removeInventory(c, "Potion de mana", 1) {
		fmt.Println(red("Vous n'avez pas de Potion de mana."))
		return
	}
	fmt.Println(blue("Vous utilisez une Potion de mana."))
	c.Mana += 30
	if c.Mana > c.MaxMana {
		c.Mana = c.MaxMana
	}
	fmt.Printf("Mana : %s\n", blue(fmt.Sprintf("%d / %d", c.Mana, c.MaxMana)))
}

func poisonPot(c *Character) {
	if !removeInventory(c, "Potion de poison", 1) {
		fmt.Println(red("Vous n'avez pas de Potion de poison."))
		return
	}
	fmt.Println(red("Vous buvez une Potion de poison..."))
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		c.Health -= 10
		if c.Health < 0 {
			c.Health = 0
		}
		fmt.Printf("PV : %s\n", healthColor(c.Health, c.MaxHealth))
	}
	isDead(c)
}
