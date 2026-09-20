package projetred

import "fmt"

func (c *Character) IsDead() bool {
	if c.CurrentHP <= 0 {
		fmt.Println(c.Name, "est mort...")
		c.CurrentHP = c.MaxHP / 2
		fmt.Printf("Il ressuscite avec %d/%d PV\n", c.CurrentHP, c.MaxHP)
		return true
	}
	return false
}
