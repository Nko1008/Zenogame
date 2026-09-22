package main

import "fmt"

func isDead(c *Character) bool {
	if c.Health <= 0 {
		fmt.Printf("%s est mort...\n", c.Name)
		c.Health = c.MaxHealth / 2
		fmt.Printf("%s ressuscite avec %d / %d PV !\n", c.Name, c.Health, c.MaxHealth)
		return true
	}
	return false
}
