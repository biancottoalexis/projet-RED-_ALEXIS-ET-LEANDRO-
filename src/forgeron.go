package projetred

import "fmt"

func (c *Character) craftItem(name string, cost int, materials map[string]int) {
	if !c.CanAddItem() {
		fmt.Println("\033[31mInventaire plein ! (", c.InventoryLimit, "objets maximum)\033[0m")
		return
	}
	if c.Gold < cost {
		fmt.Println("\033[31mPas assez d'or pour fabriquer\033[0m", name)
		return
	}
	for material, qty := range materials {
		if c.Inventory[material] < qty {
			fmt.Println("\033[31mIl te manque des matériaux pour fabriquer\033[0m", name)
			return
		}
	}

	c.Gold -= cost
	for material, qty := range materials {
		c.Inventory[material] -= qty
	}
	c.Inventory[name]++
	fmt.Println("\033[32mTu as fabriqué :\033[0m", name)
}

func (c *Character) OpenForgeron() {
	fmt.Println("\033[1;33m=== FORGERON ===\033[0m")
	fmt.Println("\033[35m1.\033[0m Chapeau de l'aventurier (5 PO + 1 Plume de Corbeau + 1 Cuir de Sanglier)")
	fmt.Println("\033[35m2.\033[0m Tunique de l'aventurier (5 PO + 2 Fourrures de Loup + 1 Peau de Troll)")
	fmt.Println("\033[35m3.\033[0m Bottes de l'aventurier (5 PO + 1 Fourrure de Loup + 1 Cuir de Sanglier)")
	fmt.Println("\033[35m0.\033[0m Quitter le forgeron")
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
		fmt.Println("\033[31mChoix invalide\033[0m")
	}
}
