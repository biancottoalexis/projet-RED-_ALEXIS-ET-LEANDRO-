package projetred

import "fmt"

func (c *Character) AccessInventory() {
	if len(c.Inventory) == 0 {
		fmt.Println("Ton inventaire est vide")
		return
	}

	fmt.Println("=== INVENTAIRE ===")

	i := 1
	for item, quantity := range c.Inventory {
		fmt.Printf("%d. %s x%d\n", i, item, quantity)
		i++
	}

	fmt.Println("0. Fermer l'inventaire")
}
