package projetred

import (
	"fmt"
	"math/rand"
	"strings"
)

func CharacterCreation() Character {
	fmt.Println("\033[1;33m=== Création du personnage ===\033[0m")

	var name string
	for {
		fmt.Print("\033[36mQuel est ton nom, aventurier ? \033[0m")
		fmt.Scanln(&name)
		if name != "" {
			break
		}
		fmt.Println("\033[31mLe nom ne peut pas être vide.\033[0m")
	}
	name = strings.ToUpper(name[:1]) + strings.ToLower(name[1:])

	fmt.Println("\033[36mChoisis ta classe :\033[0m")
	fmt.Println("\033[36m1.\033[0m Humain")
	fmt.Println("\033[36m2.\033[0m Elfe")
	fmt.Println("\033[36m3.\033[0m Nain")
	fmt.Println("\033[36m4.\033[0m Orc")
	fmt.Println("\033[36m5.\033[0m Sorcier")
	fmt.Println("\033[36m6.\033[0m Barbare")
	fmt.Print("Ton choix : ")

	var choice int
	fmt.Scanln(&choice)

	var class string
	var maxHP int
	var initiative int

	switch choice {
	case 1:
		class = "Humain"
		maxHP = 100
		initiative = 5 + rand.Intn(5)
	case 2:
		class = "Elfe"
		maxHP = 80
		initiative = 8 + rand.Intn(5)
	case 3:
		class = "Nain"
		maxHP = 120
		initiative = 3 + rand.Intn(5)
	case 4:
		class = "Orc"
		maxHP = 130
		initiative = 4 + rand.Intn(5)
	case 5:
		class = "Sorcier"
		maxHP = 70
		initiative = 7 + rand.Intn(5)
	case 6:
		class = "Barbare"
		maxHP = 115
		initiative = 6 + rand.Intn(5)
	default:
		fmt.Println("\033[31mChoix invalide, Humain par défaut.\033[0m")
		class = "Humain"
		maxHP = 100
		initiative = 5 + rand.Intn(5)
	}

	c := InitCharacter(name, class, 1, maxHP, maxHP/2, map[string]int{}, 100)
	c.Initiative = initiative
	return c
}
