package projetred

import "fmt"

func (c *Character) buyItem(name string, price int) {
	if !c.CanAddItem() {
		fmt.Println("Inventaire plein ! (10 objets maximum)")
		return
	}
	if c.Gold < price {
		fmt.Println("Pas assez d'or pour acheter", name)
		return
	}

	c.Gold -= price
	c.Inventory[name]++
	fmt.Println("Tu as acheté :", name)
}

func (c *Character) OpenMerchant() {
	fmt.Println("=== MARCHAND ===")
	fmt.Println("1. Potion de vie - 3 PO")
	fmt.Println("2. Potion de poison - 6 PO")
	fmt.Println("3. Livre de Sort : Boule de Feu - 25 PO")
	fmt.Println("4. Fourrure de Loup - 4 PO")
	fmt.Println("5. Peau de Troll - 7 PO")
	fmt.Println("6. Cuir de Sanglier - 3 PO")
	fmt.Println("7. Plume de Corbeau - 1 PO")
	fmt.Println("0. Quitter le marchand")
	fmt.Print("Ton choix : ")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		c.buyItem("Potion de vie", 3)
	case 2:
		c.buyItem("Potion de poison", 6)
	case 3:
		if c.Gold < 25 {
			fmt.Println("Pas assez d'or pour acheter le Livre de Sort")
			return
		}
		c.Gold -= 25
		c.SpellBook()
	case 4:
		c.buyItem("Fourrure de Loup", 4)
	case 5:
		c.buyItem("Peau de Troll", 7)
	case 6:
		c.buyItem("Cuir de Sanglier", 3)
	case 7:
		c.buyItem("Plume de Corbeau", 1)
	case 0:
		return
	default:
		fmt.Println("Choix invalide")
	}
}
