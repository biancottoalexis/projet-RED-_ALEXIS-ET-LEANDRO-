package projetred

import "fmt"

func (c *Character) TakePot() {
	if c.Inventory["Potion de vie"] <= 0 {
		fmt.Println("Tu n'as pas de Potion de vie dans ton inventaire")
		return
	}

	c.Inventory["Potion de vie"]--
	c.CurrentHP += 50
	if c.CurrentHP > c.MaxHP {
		c.CurrentHP = c.MaxHP
	}

	fmt.Printf("Tu utilises une Potion de vie. PV : %d/%d\n", c.CurrentHP, c.MaxHP)
}
