package projetred

import (
	"fmt"
	"math/rand"
)

func (c *Character) CharacterTurn(monster *Monster) {
	fmt.Println("\033[1;33m=== TON TOUR ===\033[0m")
	fmt.Println("\033[36m1.\033[0m Attaquer")
	fmt.Println("\033[36m2.\033[0m Inventaire")
	fmt.Print("Ton choix : ")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		damage := 5
		if rand.Intn(100) < 15 {
			damage *= 2
			fmt.Println("\033[1;31m✨ Le Destin s'abat sur", monster.Name, "! Coup critique !\033[0m")
		}
		monster.CurrentHP -= damage
		fmt.Printf("\033[32m%s utilise Attaque basique, inflige %d dégâts à %s\033[0m\n", c.Name, damage, monster.Name)
		fmt.Printf("PV de %s : %d/%d\n", monster.Name, monster.CurrentHP, monster.MaxHP)
	case 2:
		c.AccessInventory()
	default:
		fmt.Println("\033[31mChoix invalide\033[0m")
	}
}

func (c *Character) TrainingFight() {
	monster := InitGoblin()
	turn := 1

	for {
		fmt.Println()
		fmt.Println("\033[1;33m=== Tour", turn, "===\033[0m")

		c.CharacterTurn(&monster)

		if monster.CurrentHP <= 0 {
			fmt.Println("\033[32mTu as vaincu", monster.Name, "!\033[0m")
			break
		}

		monster.GoblinPattern(c, turn)

		if c.IsDead() {
			fmt.Println("\033[31mLe combat est terminé.\033[0m")
			break
		}

		turn++
	}
}
