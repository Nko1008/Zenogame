package main

import (
	"fmt"
	"time"
)

func takePot(c *Character) {
	if !removeInventory(c, "Potion de vie", 1) {
		fmt.Println("Vous n'avez pas de Potion de vie.")
		return
	}
	fmt.Println("Vous utilisez une Potion de vie.")
	c.Health += 50
	if c.Health > c.MaxHealth {
		c.Health = c.MaxHealth
	}
	fmt.Printf("PV : %d / %d\n", c.Health, c.MaxHealth)
}

func poisonPot(c *Character) {
	if !removeInventory(c, "Potion de poison", 1) {
		fmt.Println("Vous n'avez pas de Potion de poison.")
		return
	}
	fmt.Println("Vous buvez une Potion de poison...")
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		c.Health -= 10
		if c.Health < 0 {
			c.Health = 0
		}
		fmt.Printf("PV : %d / %d\n", c.Health, c.MaxHealth)
	}
	isDead(c)
}
