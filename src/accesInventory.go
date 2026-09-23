package projetred

import "fmt"

func (c *Character) AccessInventory() {
	if len(c.Inventory) == 0 {
		fmt.Println("\033[33mTon inventaire est vide\033[0m")
		return
	}

	fmt.Println("\033[1;33m=== INVENTAIRE ===\033[0m")

	i := 1
	var items []string
	for item, quantity := range c.Inventory {
		if quantity <= 0 {
			continue
		}
		fmt.Printf("\033[35m%d.\033[0m %s x%d\n", i, item, quantity)
		items = append(items, item)
		i++
	}
	fmt.Println("\033[35m0.\033[0m Fermer l'inventaire")
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
		fmt.Println("\033[31mCet objet ne peut pas être utilisé\033[0m")
	}
}
