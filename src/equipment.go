package projetred

import "fmt"

type Equipment struct {
	Head  string
	Chest string
	Feet  string
}

func (c *Character) EquipHead(name string) {
	if c.Inventory[name] <= 0 {
		fmt.Println("Tu n'as pas cet objet dans ton inventaire")
		return
	}
	c.Inventory[name]--

	if c.Equipment.Head != "" {
		c.Inventory[c.Equipment.Head]++
		c.MaxHP -= 10
	}

	c.Equipment.Head = name
	c.MaxHP += 10
	fmt.Println("Tu équipes :", name)
}

func (c *Character) EquipChest(name string) {
	if c.Inventory[name] <= 0 {
		fmt.Println("Tu n'as pas cet objet dans ton inventaire")
		return
	}
	c.Inventory[name]--

	if c.Equipment.Chest != "" {
		c.Inventory[c.Equipment.Chest]++
		c.MaxHP -= 25
	}

	c.Equipment.Chest = name
	c.MaxHP += 25
	fmt.Println("Tu équipes :", name)
}

func (c *Character) EquipFeet(name string) {
	if c.Inventory[name] <= 0 {
		fmt.Println("Tu n'as pas cet objet dans ton inventaire")
		return
	}
	c.Inventory[name]--

	if c.Equipment.Feet != "" {
		c.Inventory[c.Equipment.Feet]++
		c.MaxHP -= 15
	}

	c.Equipment.Feet = name
	c.MaxHP += 15
	fmt.Println("Tu équipes :", name)
}
