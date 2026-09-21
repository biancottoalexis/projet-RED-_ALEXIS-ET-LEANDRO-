package projetred

import (
	"fmt"
	"time"
)

func (c *Character) PoisonPot() {
	if c.Inventory["Potion de poison"] <= 0 {
		fmt.Println("\033[31mTu n'as pas de Potion de poison dans ton inventaire\033[0m")
		return
	}

	c.Inventory["Potion de poison"]--
	fmt.Println("\033[35mTu bois la Potion de poison...\033[0m")

	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		c.CurrentHP -= 10
		fmt.Printf("\033[31mPV : %d/%d\033[0m\n", c.CurrentHP, c.MaxHP)
	}

	c.IsDead()
}
