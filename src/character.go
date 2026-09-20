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

func InitCharacter(name, class string, level, maxHP, currentHP int, inventory map[string]int) Character {
	return Character{
		Name:      name,
		Class:     class,
		Level:     level,
		MaxHP:     maxHP,
		CurrentHP: currentHP,
		Inventory: inventory,
	}
}

func (c *Character) DisplayInfo() {
	fmt.Println("=== Fiche du personnage ===")
	fmt.Printf("\t Nom      : %s\n", c.Name)
	fmt.Printf("\t Classe   : %s\n", c.Class)
	fmt.Printf("\t PV       : %d\n", c.CurrentHP)
	fmt.Printf("\t PV max   : %d\n", c.MaxHP)
}
