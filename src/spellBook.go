package main

import "fmt"

func spellBook(skills *[]string) {
	for _, skill := range *skills {
		if skill == "Star Shot" {
			fmt.Println(yellow("Vous connaissez déjà Star Shot !"))
			return
		}
	}

	*skills = append(*skills, "Star Shot")
	fmt.Println(green("Vous avez appris : Star Shot !"))
}
