package projetred

import (
	"fmt"
	"time"
)

func (c *Character) PoisonPot() {
	if c.Inventory["Potion de poison"] <= 0 {
		fmt.Println("Tu n'as pas de Potion de poison dans ton inventaire")
		return
	}

	c.Inventory["Potion de poison"]--
	fmt.Println("Tu bois la Potion de poison...")

	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		c.CurrentHP -= 10
		fmt.Printf("PV : %d/%d\n", c.CurrentHP, c.MaxHP)
	}

	c.IsDead()
}
