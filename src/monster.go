package projetred

import "fmt"

type Monster struct {
	Name      string
	MaxHP     int
	CurrentHP int
	Attack    int
}

func InitGoblin() Monster {
	return Monster{
		Name:      "Gobelin d'entrainement",
		MaxHP:     40,
		CurrentHP: 40,
		Attack:    5,
	}
}

func (m *Monster) GoblinPattern(target *Character, turn int) {
	damage := m.Attack
	if turn%3 == 0 {
		damage = m.Attack * 2
	}
	target.CurrentHP -= damage
	fmt.Printf("\033[31m%s inflige à %s %d de dégâts\033[0m\n", m.Name, target.Name, damage)
	fmt.Printf("PV : %d/%d\n", target.CurrentHP, target.MaxHP)
}
