package main

import "fmt"

func whoAreThey() {
	fmt.Println(cyan("\nDeux artistes sont cachés dans les tâches du projet :"))
	fmt.Println(yellow("- Partie 2 (Économie/Fabrication) : ABBA"))
	fmt.Println(yellow("- Partie 3 (Combat) : Steven Spielberg"))
}

func main() {
	fmt.Println(bold(magenta("Bienvenue à Skylandia !")))

	player := characterCreation()

	for {
		fmt.Println(cyan("\n=== MENU ==="))
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder à l'inventaire")
		fmt.Println("3. Marchand")
		fmt.Println("4. Forgeron")
		fmt.Println("5. Nekomata")
		fmt.Println("6. Entrainement")
		fmt.Println("7. Boss final")
		fmt.Println("8. Qui sont-ils ?")
		fmt.Println("9. Quitter")

		switch readChoice("> ") {
		case 1:
			displayCharacter(player)
		case 2:
			accessInventory(&player)
		case 3:
			merchant(&player)
		case 4:
			blacksmith(&player)
		case 5:
			guide(&player)
		case 6:
			trainingFight(&player)
		case 7:
			bossFight(&player)
		case 8:
			whoAreThey()
		case 9:
			fmt.Println(magenta("À bientôt !"))
			return
		default:
			fmt.Println(red("Choix invalide."))
		}
	}
}
