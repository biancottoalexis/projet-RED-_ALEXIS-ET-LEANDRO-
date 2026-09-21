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
		fmt.Println("\033[31mTu as atteint le nombre maximum d'améliorations d'inventaire (3)\033[0m")
		return
	}

	c.InventoryLimit += 10
	c.InventoryUpgrades++
	fmt.Printf("\033[32mTon inventaire peut maintenant contenir %d objets\033[0m\n", c.InventoryLimit)
}
