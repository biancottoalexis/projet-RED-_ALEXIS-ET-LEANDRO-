package projetred

import (
	"fmt"
	"strings"
)

func CharacterCreation() Character {
	var name string

	for {
		fmt.Print("\033[36mEntre ton nom (lettres uniquement) : \033[0m")
		fmt.Scanln(&name)

		valid := true
		for _, r := range name {
			if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')) {
				valid = false
				break
			}
		}

		if valid && name != "" {
			break
		}
		fmt.Println("\033[31mNom invalide, réessaie.\033[0m")
	}

	name = strings.ToUpper(name[:1]) + strings.ToLower(name[1:])

	fmt.Println()
	fmt.Println("\033[1;33m=== Choisis ta classe ===\033[0m")
	fmt.Println("\033[36m1.\033[0m Humain")
	fmt.Println("\033[36m2.\033[0m Elfe")
	fmt.Println("\033[36m3.\033[0m Nain")
	fmt.Println("\033[36m4.\033[0m Orc")
	fmt.Println("\033[36m5.\033[0m Sorcier")
	fmt.Println("\033[36m6.\033[0m Barbare")
	fmt.Print("\033[36mTon choix : \033[0m")

	var classChoice int
	fmt.Scanln(&classChoice)

	var class string
	var maxHP int

	switch classChoice {
	case 1:
		class = "Humain"
		maxHP = 100
	case 2:
		class = "Elfe"
		maxHP = 80
	case 3:
		class = "Nain"
		maxHP = 120
	case 4:
		class = "Orc"
		maxHP = 130
	case 5:
		class = "Sorcier"
		maxHP = 70
	case 6:
		class = "Barbare"
		maxHP = 115
	default:
		fmt.Println("Choix invalide, Humain par défaut")
		class = "Humain"
		maxHP = 100
	}

	currentHP := maxHP / 2

	return InitCharacter(name, class, 1, maxHP, currentHP, map[string]int{}, 100)
}
