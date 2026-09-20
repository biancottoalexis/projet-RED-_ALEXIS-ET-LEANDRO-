package main

import projetred "projetred/src"

func main() {
	c := projetred.InitCharacter("Alexis", "Elfe", 1, 100, 40, map[string]int{"Potion de vie": 3})
	c.DisplayInfo()
	c.AccessInventory()
}
