package main

import (
	"fmt"

	projetred "projetred/src"
)

func main() {
	c := projetred.InitCharacter("Alexis", "Elfe", 1, 100, 40, map[string]int{"Potion de vie": 3})

	for {
		fmt.Println()
		fmt.Println("=== MENU ===")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder à l'inventaire")
		fmt.Println("3. Marchand")
		fmt.Println("4. Quitter")
		fmt.Print("Ton choix : ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			c.DisplayInfo()
		case 2:
			c.AccessInventory()
		case 3:
			c.OpenMerchant()
		case 4:
			fmt.Println("À bientôt !")
			return
		default:
			fmt.Println("Choix invalide, réessaie.")
		}
	}
}
