package projetred

import "fmt"

func (c *Character) CanAddItem() bool {
	total := 0
	for _, quantity := range c.Inventory {
		total += quantity
	}
	return total < c.InventoryLimit
}

func (c *Character) UpgradeInventorySlot() {
	if c.InventoryUpgrades >= 3 {
		fmt.Println("Tu as atteint le nombre maximum d'améliorations d'inventaire (3)")
		return
	}

	c.InventoryLimit += 10
	c.InventoryUpgrades++
	fmt.Printf("Ton inventaire peut maintenant contenir %d objets\n", c.InventoryLimit)
}
