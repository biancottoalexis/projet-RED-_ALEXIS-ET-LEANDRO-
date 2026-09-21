package projetred

import "fmt"

func (c *Character) CharacterTurn(monster *Monster) {
	fmt.Println("=== TON TOUR ===")
	fmt.Println("1. Attaquer")
	fmt.Println("2. Inventaire")
	fmt.Print("Ton choix : ")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		damage := 5
		monster.CurrentHP -= damage
		fmt.Printf("%s utilise Attaque basique, inflige %d dégâts à %s\n", c.Name, damage, monster.Name)
		fmt.Printf("PV de %s : %d/%d\n", monster.Name, monster.CurrentHP, monster.MaxHP)
	case 2:
		c.AccessInventory()
	default:
		fmt.Println("Choix invalide")
	}
}

func (c *Character) TrainingFight() {
	monster := InitGoblin()
	turn := 1

	for {
		fmt.Println()
		fmt.Println("=== Tour", turn, "===")

		c.CharacterTurn(&monster)

		if monster.CurrentHP <= 0 {
			fmt.Println("Tu as vaincu", monster.Name, "!")
			break
		}

		monster.GoblinPattern(c, turn)

		if c.IsDead() {
			fmt.Println("Le combat est terminé.")
			break
		}

		turn++
	}
}
