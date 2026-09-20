package projetred

import "fmt"

func (c *Character) craftItem(name string, cost int, materials map[string]int) {
	if !c.CanAddItem() {
		fmt.Println("Inventaire plein ! (10 objets maximum)")
		return
	}
	if c.Gold < cost {
		fmt.Println("Pas assez d'or pour fabriquer", name)
		return
	}
	for material, qty := range materials {
		if c.Inventory[material] < qty {
			fmt.Println("Il te manque des matériaux pour fabriquer", name)
			return
		}
	}

	c.Gold -= cost
	for material, qty := range materials {
		c.Inventory[material] -= qty
	}
	c.Inventory[name]++
	fmt.Println("Tu as fabriqué :", name)
}

func (c *Character) OpenForgeron() {
	fmt.Println("=== FORGERON ===")
	fmt.Println("1. Chapeau de l'aventurier (5 PO + 1 Plume de Corbeau + 1 Cuir de Sanglier)")
	fmt.Println("2. Tunique de l'aventurier (5 PO + 2 Fourrures de Loup + 1 Peau de Troll)")
	fmt.Println("3. Bottes de l'aventurier (5 PO + 1 Fourrure de Loup + 1 Cuir de Sanglier)")
	fmt.Println("0. Quitter le forgeron")
	fmt.Print("Ton choix : ")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		c.craftItem("Chapeau de l'aventurier", 5, map[string]int{"Plume de Corbeau": 1, "Cuir de Sanglier": 1})
	case 2:
		c.craftItem("Tunique de l'aventurier", 5, map[string]int{"Fourrure de Loup": 2, "Peau de Troll": 1})
	case 3:
		c.craftItem("Bottes de l'aventurier", 5, map[string]int{"Fourrure de Loup": 1, "Cuir de Sanglier": 1})
	case 0:
		return
	default:
		fmt.Println("Choix invalide")
	}
}
