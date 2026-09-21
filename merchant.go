package main

import "fmt"

type shopItem struct {
	Name  string
	Price int
}

var shopItems = []shopItem{
	{"Potion de vie", 3},
	{"Potion de poison", 6},
	{"Livre de Sort : Star Shot", 25},
	{"Fourrure de Loup", 4},
	{"Peau de Troll", 7},
	{"Cuir de Sanglier", 3},
	{"Plume de Corbeau", 1},
	{itemUpg, 30},
}

func merchant(c *Character) {
	for {
		fmt.Printf("\n--- Marchand --- (Or : %d)\n", c.Money)
		for i, it := range shopItems {
			fmt.Printf("%d. %s (%d or)\n", i+1, it.Name, it.Price)
		}
		fmt.Println("0. Retour")
		choice := readChoice("> ")
		if choice == 0 {
			return
		}
		if choice < 1 || choice > len(shopItems) {
			fmt.Println("Choix invalide.")
			continue
		}
		it := shopItems[choice-1]
		if c.Money < it.Price {
			fmt.Println("Pas assez d'or !")
			continue
		}
		if !addInventory(c, it.Name) {
			continue
		}
		c.Money -= it.Price
		fmt.Printf("Vous avez acheté : %s (or restant : %d)\n", it.Name, c.Money)
	}
}
