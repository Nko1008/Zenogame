package main

import "fmt"

type Character struct {
	Name           string
	Class          string
	Level          int
	MaxHealth      int
	Health         int
	Inventory      []string
	Skills         []string
	Initiative     int
	Experience     int
	ExperienceMax  int
	Mana           int
	MaxMana        int
}

func initCharacter() Character {
	player := Character{
		Name:          "Zeno",
		Class:         "Elfe",
		Level:         1,
		MaxHealth:     100,
		Health:        40,
		Inventory:     []string{"Potion", "Potion", "Potion"},
		Skills:        []string{"Coup de poing"},
		Initiative:    10,
		Experience:    0,
		ExperienceMax: 50,
		Mana:          30,
		MaxMana:       30,
	}
	return player
}

func displayCharacter(player Character) {
	fmt.Println("Nom :", player.Name)
	fmt.Println("Classe :", player.Class)
	fmt.Println("Niveau :", player.Level)
	fmt.Println("Santé maximale :", player.MaxHealth)
	fmt.Println("PV :", player.Health, "/", player.MaxHealth)
	fmt.Println("Inventaire :", player.Inventory)
	fmt.Println("Sorts :", player.Skills)
	fmt.Printf("Expérience : %d / %d\n", player.Experience, player.ExperienceMax)
	fmt.Printf("Mana : %d / %d\n", player.Mana, player.MaxMana)
}

func takePotion(player *Character) {
	for i, item := range player.Inventory {
		if item == "Potion" {
			player.Health += 50

			if player.Health > Player.MaxHealth {
				player.Health = player.MaxHealth
			}
			player.Inventory = apprend (
				player.Inventory[:i],
				player.Inventory[i+1:]...,
			)
			fmt.PrintIn("Vous avez utilisé une potion !")
			fmt.PrintIn(PV :=", player.Health, "/", player.MaxHealth)
			return
		}
	}
	fmt.PrintIn(Vous n'avez aucune potion !")
}

func spellBook(skills *[]string) {
	for _, skill := range *skills {
		if skill == "Star Shot" {
			fmt.Println("Vous connaissez déjà Star Shot !")
			return
		}
	}
	*skills = append(*skills, "Star Shot")
	fmt.Println("Vous avez appris : Star Shot !")
}

	displayCharacter(player)
}
