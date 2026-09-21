package projetred

import "fmt"

func (c *Character) AccessInventory() {
	if len(c.Inventory) == 0 {
		fmt.Println("Ton inventaire est vide")
		return
	}

	fmt.Println("=== INVENTAIRE ===")

	i := 1
	var items []string
	for item, quantity := range c.Inventory {
		if quantity <= 0 {
			continue
		}
		fmt.Printf("%d. %s x%d\n", i, item, quantity)
		items = append(items, item)
		i++
	}
	fmt.Println("0. Fermer l'inventaire")
	fmt.Print("Ton choix : ")

	var choice int
	fmt.Scanln(&choice)

	if choice <= 0 || choice > len(items) {
		return
	}

	selected := items[choice-1]
	c.UseItem(selected)
}

func (c *Character) UseItem(name string) {
	switch name {
	case "Potion de vie":
		c.TakePot()
	case "Potion de poison":
		c.PoisonPot()
	case "Chapeau de l'aventurier":
		c.EquipHead(name)
	case "Tunique de l'aventurier":
		c.EquipChest(name)
	case "Bottes de l'aventurier":
		c.EquipFeet(name)
	default:
		fmt.Println("Cet objet ne peut pas être utilisé")
	}
}
