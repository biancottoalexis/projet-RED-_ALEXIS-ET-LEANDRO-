package projetred

import "fmt"

type Character struct {
	Name              string
	Class             string
	Level             int
	MaxHP             int
	CurrentHP         int
	Inventory         map[string]int
	Skill             []string
	Gold              int
	Equipment         Equipment
	InventoryLimit    int
	InventoryUpgrades int
	Initiative        int
	XP                int
	XPMax             int
}

func InitCharacter(name, class string, level, maxHP, currentHP int, inventory map[string]int, gold int) Character {
	return Character{
		Name:              name,
		Class:             class,
		Level:             level,
		MaxHP:             maxHP,
		CurrentHP:         currentHP,
		Inventory:         inventory,
		Skill:             []string{"Coup de poing"},
		Gold:              gold,
		Equipment:         Equipment{},
		InventoryLimit:    10,
		InventoryUpgrades: 0,
		Initiative:        0,
		XP:                0,
		XPMax:             50,
	}
}

func (c *Character) DisplayInfo() {
	fmt.Println("\033[1;33m=== Fiche du personnage ===\033[0m")
	fmt.Printf("\t Nom      : %s\n", c.Name)
	fmt.Printf("\t Classe   : %s\n", c.Class)
	fmt.Printf("\t Niveau   : %d\n", c.Level)
	fmt.Printf("\t PV       : %d\n", c.CurrentHP)
	fmt.Printf("\t PV max   : %d\n", c.MaxHP)
	fmt.Printf("\033[35m\t XP : %d/%d\033[0m\n", c.XP, c.XPMax)
	fmt.Printf("\033[33m\t Or : %d\033[0m\n", c.Gold)
}
