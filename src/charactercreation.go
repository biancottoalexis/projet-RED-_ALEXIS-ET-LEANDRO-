package projetred

import (
	"fmt"
	"strings"
)

func CharacterCreation() Character {
	var name string

	for {
		fmt.Print("Entre ton nom (lettres uniquement) : ")
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
		fmt.Println("Nom invalide, réessaie.")
	}

	name = strings.ToUpper(name[:1]) + strings.ToLower(name[1:])

	fmt.Println("Choisis ta classe :")
	fmt.Println("1. Humain")
	fmt.Println("2. Elfe")
	fmt.Println("3. Nain")
	fmt.Println("4. Orc")
	fmt.Println("5. Sorcier")
	fmt.Println("6. Barbare")

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
