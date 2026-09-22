package projetred

import "fmt"

func (c *Character) GainXP(amount int) {
	fmt.Printf("\033[35mTu gagnes %d points d'expérience !\033[0m\n", amount)
	c.XP += amount

	for c.XP >= c.XPMax {
		c.XP -= c.XPMax
		c.Level++
		c.XPMax += 20
		c.MaxHP += 15
		c.CurrentHP += 15
		fmt.Printf("\033[1;35m✨ Niveau supérieur ! Tu es maintenant niveau %d (PV max : %d)\033[0m\n", c.Level, c.MaxHP)
	}
}
