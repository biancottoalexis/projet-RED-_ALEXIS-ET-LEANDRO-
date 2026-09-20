package projetred

import "fmt"

func (c *Character) OpenMerchant() {
	fmt.Println("=== MARCHAND ===")
	fmt.Println("1. Potion de vie (gratuit)")
	fmt.Println("2. Potion de poison (gratuit)")
	fmt.Println("3. Livre de Sort : Boule de Feu (gratuit)")
	fmt.Println("0. Quitter le marchand")
	fmt.Print("Ton choix : ")

	var choice int
	fmt.Scanln(&choice)

	if choice == 1 || choice == 2 {
		if !c.CanAddItem() {
			fmt.Println("Inventaire plein ! (10 objets maximum)")
			return
		}
	}

	switch choice {
	case 1:
		c.Inventory["Potion de vie"]++
		fmt.Println("Tu as acheté : Potion de vie")
	case 2:
		c.Inventory["Potion de poison"]++
		fmt.Println("Tu as acheté : Potion de poison")
	case 3:
		c.SpellBook()
	case 0:
		return
	default:
		fmt.Println("Choix invalide")
	}
}
