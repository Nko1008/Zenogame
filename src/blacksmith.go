package main

import (
	"fmt"
	"sort"
	"strings"
)

const forgeCost = 5

type recipe struct {
	Result    string
	Materials map[string]int
}

var recipes = []recipe{

	{itemHat, map[string]int{"Plume de Corbeau": 1, "Cuir de Sanglier": 1}},
	{itemTunic, map[string]int{"Fourrure de Loup": 2, "Peau de Troll": 1}},
	{itemBoots, map[string]int{"Fourrure de Loup": 1, "Cuir de Sanglier": 1}},

	{itemCap, map[string]int{"Plume de Corbeau": 2, "Cuir de Sanglier": 1}},
	{itemJacket, map[string]int{"Cuir de Sanglier": 2, "Fourrure de Loup": 1}},
	{itemSneakers, map[string]int{"Fourrure de Loup": 1, "Plume de Corbeau": 1, "Cuir de Sanglier": 1}},

	{itemCrown, map[string]int{"Plume de Corbeau": 2, "Peau de Troll": 1}},
	{itemCloak, map[string]int{"Peau de Troll": 2, "Plume de Corbeau": 1}},
	{itemVoidBoots, map[string]int{"Peau de Troll": 2, "Cuir de Sanglier": 1}},
}

func effectLabel(item string) string {
	info, ok := equipData[item]
	if !ok {
		return ""
	}
	effects := []string{}
	if info.HP != 0 {
		effects = append(effects, fmt.Sprintf("+%d PV max", info.HP))
	}
	if info.Initiative != 0 {
		effects = append(effects, fmt.Sprintf("+%d initiative", info.Initiative))
	}
	if len(effects) == 0 {
		return ""
	}
	return cyan(" -> " + strings.Join(effects, ", "))
}

func blacksmith(c *Character) {
	for {
		fmt.Println(cyan(fmt.Sprintf("\n--- Forgeron --- (Or : %d, coût : %d)", c.Money, forgeCost)))
		for i, r := range recipes {
			mats := make([]string, 0, len(r.Materials))
			for mat := range r.Materials {
				mats = append(mats, mat)
			}
			sort.Strings(mats)
			parts := []string{}
			for _, mat := range mats {
				parts = append(parts, fmt.Sprintf("%d %s", r.Materials[mat], mat))
			}
			fmt.Printf("%d. %s (%s)%s\n", i+1, r.Result, strings.Join(parts, ", "), effectLabel(r.Result))
		}
		fmt.Println("0. Retour")
		choice := readChoice("> ")
		if choice == 0 {
			return
		}
		if choice < 1 || choice > len(recipes) {
			fmt.Println(red("Choix invalide."))
			continue
		}
		craft(c, recipes[choice-1])
	}
}

func craft(c *Character, r recipe) {
	if c.Money < forgeCost {
		fmt.Println(red("Pas assez d'or pour la fabrication."))
		return
	}
	total := 0
	for mat, qty := range r.Materials {
		if countItem(c, mat) < qty {
			fmt.Println(red(fmt.Sprintf("Ressource manquante : %s (x%d requis)", mat, qty)))
			return
		}
		total += qty
	}
	if len(c.Inventory)-total+1 > c.InventoryMax {
		fmt.Println(red("Pas assez de place dans l'inventaire."))
		return
	}
	for mat, qty := range r.Materials {
		removeInventory(c, mat, qty)
	}
	c.Money -= forgeCost
	fmt.Println(green(fmt.Sprintf("Vous avez fabriqué : %s", r.Result)))

	addInventory(c, r.Result)
	equip(c, r.Result)
}
