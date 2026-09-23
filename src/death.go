package main

import "fmt"

func isDead(c *Character) bool {
	if c.Health <= 0 {
		fmt.Println(red(fmt.Sprintf("%s est mort...", c.Name)))
		c.Health = c.MaxHealth / 2
		fmt.Println(yellow(fmt.Sprintf("%s ressuscite avec %d / %d PV !", c.Name, c.Health, c.MaxHealth)))
		return true
	}
	return false
}
