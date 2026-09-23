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
	Mana              int
	ManaMax           int
}

func InitCharacter(name, class string, level, maxHP, currentHP int, inventory map[string]int, gold int) Character {
	var c Character
	c.Name = name
	c.Class = class
	c.Level = level
	c.MaxHP = maxHP
	c.CurrentHP = currentHP
	c.Inventory = inventory
	c.Skill = []string{"Coup de poing"}
	c.Gold = gold
	c.Equipment = Equipment{}
	c.InventoryLimit = 10
	c.InventoryUpgrades = 0
	c.Initiative = 0
	c.XP = 0
	c.XPMax = 50
	c.Mana = 30
	c.ManaMax = 30
	return c
}

func (c *Character) DisplayInfo() {
	fmt.Println("\033[1;33m=== Fiche du personnage ===\033[0m")
	fmt.Printf("\t Nom      : %s\n", c.Name)
	fmt.Printf("\t Classe   : %s\n", c.Class)
	fmt.Printf("\t Niveau   : %d\n", c.Level)
	fmt.Printf("\t PV       : %d\n", c.CurrentHP)
	fmt.Printf("\t PV max   : %d\n", c.MaxHP)
	fmt.Printf("\033[34m\t Mana     : %d/%d\033[0m\n", c.Mana, c.ManaMax)
	fmt.Printf("\033[35m\t XP       : %d/%d\033[0m\n", c.XP, c.XPMax)
	fmt.Printf("\033[33m\t Or       : %d\033[0m\n", c.Gold)
}
