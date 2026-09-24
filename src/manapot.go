package projetred

import "fmt"

func (c *Character) TakeManaPot() {
	if c.Inventory["Potion de mana"] <= 0 {
		fmt.Println("\033[31mTu n'as pas de Potion de mana.\033[0m")
		return
	}

	c.Inventory["Potion de mana"]--
	c.Mana += 30
	if c.Mana > c.ManaMax {
		c.Mana = c.ManaMax
	}

	fmt.Printf("\033[32mTu bois une Potion de mana ! Mana : %d/%d\033[0m\n", c.Mana, c.ManaMax)
}
