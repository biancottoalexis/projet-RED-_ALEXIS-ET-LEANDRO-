package projetred

func (c *Character) CanAddItem() bool {
	total := 0
	for _, quantity := range c.Inventory {
		total += quantity
	}
	return total < 10
}
