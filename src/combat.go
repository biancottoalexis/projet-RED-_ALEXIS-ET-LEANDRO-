package projetred

import (
	"fmt"
	"math/rand"
)

func (c *Character) CharacterTurn(monster *Monster) {
	fmt.Println("\033[1;33m=== TON TOUR ===\033[0m")
	fmt.Println("\033[35m1.\033[0m Attaquer")
	fmt.Println("\033[35m2.\033[0m Inventaire")
	fmt.Print("Ton choix : ")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		c.ChooseSkill(monster)
	case 2:
		c.AccessInventory()
	default:
		fmt.Println("\033[31mChoix invalide\033[0m")
	}
}

func (c *Character) ChooseSkill(monster *Monster) {
	fmt.Println("\033[1;33m=== Choisis ta compétence ===\033[0m")
	fmt.Printf("\033[34mMana : %d/%d\033[0m\n", c.Mana, c.ManaMax)
	for i, skill := range c.Skill {
		fmt.Printf("\033[35m%d.\033[0m %s\n", i+1, skill)
	}
	fmt.Print("Ton choix : ")

	var choice int
	fmt.Scanln(&choice)

	if choice < 1 || choice > len(c.Skill) {
		fmt.Println("\033[31mChoix invalide\033[0m")
		return
	}

	skillName := c.Skill[choice-1]
	var damage, manaCost int

	switch skillName {
	case "Coup de poing":
		damage = 8
		manaCost = 0
	case "Boule de Feu":
		damage = 18
		manaCost = 10
	default:
		damage = 5
		manaCost = 0
	}

	if c.Mana < manaCost {
		fmt.Println("\033[31mPas assez de mana pour utiliser", skillName, "!\033[0m")
		return
	}
	c.Mana -= manaCost

	if rand.Intn(100) < 15 {
		damage *= 2
		fmt.Println("\033[1;31m✨ Le Destin s'abat sur", monster.Name, "! Coup critique !\033[0m")
	}

	monster.CurrentHP -= damage
	fmt.Printf("\033[32m%s utilise %s, inflige %d dégâts à %s\033[0m\n", c.Name, skillName, damage, monster.Name)
	fmt.Printf("PV de %s : %d/%d\n", monster.Name, monster.CurrentHP, monster.MaxHP)
}

func (c *Character) TrainingFight() {
	monster := InitGoblin()
	turn := 1

	if c.Initiative >= monster.Initiative {
		fmt.Println("\033[35mTon instinct te pousse à agir en premier !\033[0m")
	} else {
		fmt.Println("\033[31mLe", monster.Name, "est plus rapide que toi et attaque en premier !\033[0m")
	}

	for {
		fmt.Println()
		fmt.Println("\033[1;33m=== Tour", turn, "===\033[0m")

		if c.Initiative >= monster.Initiative {
			c.CharacterTurn(&monster)
			if monster.CurrentHP <= 0 {
				fmt.Println("\033[32mTu as vaincu", monster.Name, "!\033[0m")
				c.GainXP(monster.XPReward)
				break
			}
			monster.GoblinPattern(c, turn)
			if c.IsDead() {
				fmt.Println("\033[31mLe combat est terminé.\033[0m")
				break
			}
		} else {
			monster.GoblinPattern(c, turn)
			if c.IsDead() {
				fmt.Println("\033[31mLe combat est terminé.\033[0m")
				break
			}
			c.CharacterTurn(&monster)
			if monster.CurrentHP <= 0 {
				fmt.Println("\033[32mTu as vaincu", monster.Name, "!\033[0m")
				c.GainXP(monster.XPReward)
				break
			}
		}

		turn++
	}
}
