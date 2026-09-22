package main

import "fmt"

// Equipment représente les 3 emplacements équipables du personnage.
type Equipment struct {
	Head  string
	Torso string
	Feet  string
}

type Character struct {
	Name          string
	Class         string
	Level         int
	MaxHealth     int
	Health        int
	Inventory     []string
	Skills        []string
	Initiative    int
	Experience    int
	ExperienceMax int
	Mana          int
	MaxMana       int

	Money         int
	Equipment     Equipment
	InventoryMax  int
	UpgradeCount  int
	GiftReceived  bool
	QuestDone     bool
	NoxarDefeated bool
}

func initCharacter() Character {
	return Character{
		Name:          "Zeno",
		Class:         "Elfe",
		Level:         1,
		MaxHealth:     100,
		Health:        40,
		Inventory:     []string{"Potion de vie", "Potion de vie", "Potion de vie"},
		Skills:        []string{"Coup de poing"},
		Initiative:    10,
		Experience:    0,
		ExperienceMax: 50,
		Mana:          30,
		MaxMana:       30,
		Money:         100,
		InventoryMax:  10,
	}
}

func displayCharacter(player Character) {
	fmt.Println("Nom :", player.Name)
	fmt.Println("Classe :", player.Class)
	fmt.Println("Niveau :", player.Level)
	fmt.Println("PV :", player.Health, "/", player.MaxHealth)
	fmt.Println("Or :", player.Money)
	fmt.Println("Inventaire :", player.Inventory)
	fmt.Println("Sorts :", player.Skills)
	fmt.Printf("Expérience : %d / %d\n", player.Experience, player.ExperienceMax)
	fmt.Printf("Mana : %d / %d\n", player.Mana, player.MaxMana)
	fmt.Printf("Initiative : %d\n", player.Initiative)
	fmt.Printf("Equipement - Tête: %s | Torse: %s | Pieds: %s\n",
		orNone(player.Equipment.Head), orNone(player.Equipment.Torso), orNone(player.Equipment.Feet))
}

func orNone(s string) string {
	if s == "" {
		return "-"
	}
	return s
}
