package projetred

import (
	"fmt"
	"math/rand"
)

type Monster struct {
	Name       string
	MaxHP      int
	CurrentHP  int
	Attack     int
	Initiative int
	XPReward   int
}

func InitGoblin() Monster {
	return Monster{
		Name:       "Gobelin",
		MaxHP:      40,
		CurrentHP:  40,
		Attack:     5,
		Initiative: 5 + rand.Intn(5),
		XPReward:   20,
	}
}

func InitGuardian() Monster {
	return Monster{
		Name:       "Gardien des Marais",
		MaxHP:      30,
		CurrentHP:  30,
		Attack:     6,
		Initiative: 4 + rand.Intn(5),
		XPReward:   15,
	}
}

func InitShadow() Monster {
	return Monster{
		Name:       "Ton Ombre",
		MaxHP:      60,
		CurrentHP:  60,
		Attack:     10,
		Initiative: 6 + rand.Intn(6),
		XPReward:   50,
	}
}

func (m *Monster) GoblinPattern(target *Character, turn int) {
	damage := m.Attack
	if turn%3 == 0 {
		damage *= 2
		fmt.Printf("\033[1;31m%s entre en rage et double son attaque !\033[0m\n", m.Name)
	}
	target.CurrentHP -= damage
	fmt.Printf("\033[31m%s attaque %s et inflige %d dégâts\033[0m\n", m.Name, target.Name, damage)
	fmt.Printf("PV de %s : %d/%d\n", target.Name, target.CurrentHP, target.MaxHP)
}
