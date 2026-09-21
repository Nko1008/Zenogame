package main

import "fmt"

type Character struct {
	Name      string
	Class     string
	Level     int
	MaxHealth int
	Health    int
	Inventory []string
	Skills    []string 
}

func initCharacter() Character {
	player := Character{
		Name:      "Zeno",
		Class:     "Elfe",
		Level:     1,
		MaxHealth: 100,
		Health:    40,
		Inventory: []string{"Potion", "Potion", "Potion"},
		Skills:    []string{"Coup de poing"}, 
	}
	return player
}

func displayCharacter(player Character) {
	player.Name = "Zeno"
	player.Class = "Elfe"
	player.Level = 1
	player.MaxHealth = 100
	player.Health = 40
	player.Inventory = []string{"Potion", "Potion", "Potion"}
	fmt.Println("Nom :", player.Name)
	fmt.Println("Classe :", player.Class)
	fmt.Println("Niveau :", player.Level)
	fmt.Println("Santé maximale :", player.MaxHealth)
	fmt.Println("PV :", player.Health, "/", player.MaxHealth)
	fmt.Println("Inventaire :", player.Inventory)
}

func main() {
	player := initCharacter()

	displayCharacter(player)
}
