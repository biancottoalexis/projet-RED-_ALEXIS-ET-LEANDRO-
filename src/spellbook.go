package projetred

import "fmt"

func (c *Character) SpellBook() {
	for _, s := range c.Skill {
		if s == "Boule de Feu" {
			fmt.Println("\033[33mTu connais déjà ce sort\033[0m")
			return
		}
	}

	c.Skill = append(c.Skill, "Boule de Feu")
	fmt.Println("\033[32mTu as appris : Boule de Feu\033[0m")
}
