package projetred

import "fmt"

func (c *Character) OpenMerchant() {
	fmt.Println("=== MARCHAND ===")
	fmt.Println("1. Potion de vie (gratuit)")
	fmt.Println("0. Quitter le marchand")
	fmt.Print("Ton choix : ")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		c.Inventory["Potion de vie"]++
		fmt.Println("Tu as acheté : Potion de vie")
	case 0:
		return
	default:
		fmt.Println("Choix invalide")
	}
}
