package projetred

import "fmt"

type Character struct {
	Name      string
	Class     string
	Level     int
	MaxHP     int
	CurrentHP int
	Inventory map[string]int
}

func (c *Character) Displaylnfo() {
	fmt.Println("=== Fiche du personnage ===")
	fmt.Printf("\t Nom      : %s\n", c.Name)
	fmt.Printf("\t Classe   : %s\n", c.Class)
	fmt.Printf("\t PV       : %d\n", c.CurrentHP)
	fmt.Printf("\t PV max   : %d\n", c.MaxHP)
}
