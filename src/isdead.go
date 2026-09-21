package projetred

import "fmt"

func (c *Character) IsDead() bool {
	if c.CurrentHP <= 0 {
		fmt.Println("\033[31m" + c.Name + " est mort...\033[0m")
		c.CurrentHP = c.MaxHP / 2
		fmt.Printf("\033[32mIl ressuscite avec %d/%d PV\033[0m\n", c.CurrentHP, c.MaxHP)
		return true
	}
	return false
}
