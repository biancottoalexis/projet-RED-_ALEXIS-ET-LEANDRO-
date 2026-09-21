package projetred

import "fmt"

func (c *Character) TakePot() {
	if c.Inventory["Potion de vie"] <= 0 {
		fmt.Println("\033[31mTu n'as pas de Potion de vie dans ton inventaire\033[0m")
		return
	}

	c.Inventory["Potion de vie"]--
	c.CurrentHP += 50
	if c.CurrentHP > c.MaxHP {
		c.CurrentHP = c.MaxHP
	}

	fmt.Printf("\033[32mTu utilises une Potion de vie. PV : %d/%d\033[0m\n", c.CurrentHP, c.MaxHP)
}
