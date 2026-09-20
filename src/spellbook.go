package projetred

import "fmt"

func (c *Character) SpellBook() {
	for _, s := range c.Skill {
		if s == "Boule de Feu" {
			fmt.Println("Tu connais déjà ce sort")
			return
		}
	}

	c.Skill = append(c.Skill, "Boule de Feu")
	fmt.Println("Tu as appris : Boule de Feu")
}
