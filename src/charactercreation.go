package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isLettersOnly(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}

	return true
}

func formatName(s string) string {
	s = strings.ToLower(s)

	if len(s) == 0 {
		return s
	}

	return strings.ToUpper(s[:1]) + s[1:]
}

func chooseName() string {
	for {
		raw := readLine("Choisissez le nom de votre personnage (lettres uniquement) : ")

		if !isLettersOnly(raw) {
			fmt.Println("Le nom ne doit contenir que des lettres, réessayez.")
			continue
		}

		return formatName(raw)
	}
}

func chooseClass() (string, int) {
	for {
		fmt.Println("\n===== CHOIX DE LA CLASSE =====")
		fmt.Println("1. Humain (100 PV max)")
		fmt.Println("2. Elfe (80 PV max)")
		fmt.Println("3. Nain (120 PV max)")

		switch readChoice("> ") {

		case 1:
			fmt.Println("\nVous avez choisi la classe Humain.")
			displayClassASCII("Humain")
			return "Humain", 100

		case 2:
			fmt.Println("\nVous avez choisi la classe Elfe.")
			displayClassASCII("Elfe")
			return "Elfe", 80

		case 3:
			fmt.Println("\nVous avez choisi la classe Nain.")
			displayClassASCII("Nain")
			return "Nain", 120

		default:
			fmt.Println("Choix invalide.")
		}
	}
}

// characterCreation permet au joueur de créer son personnage.
func characterCreation() Character {
	c := initCharacter()

	c.Name = chooseName()

	class, maxHP := chooseClass()

	c.Class = class
	c.Level = 1
	c.MaxHealth = maxHP
	c.Health = maxHP / 2
	c.Skills = []string{"Coup de poing"}

	fmt.Printf(
		"\nBienvenue, %s le %s ! (PV : %d / %d)\n",
		c.Name,
		c.Class,
		c.Health,
		c.MaxHealth,
	)

	return c
}
