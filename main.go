package main

import (
	"fmt"
	"time"

	projetred "projetred/src"
)

func typeWriter(text string) {
	for _, ch := range text {
		fmt.Print(string(ch))
		time.Sleep(30 * time.Millisecond)
	}
	fmt.Println()
}

func printBanner() {
	fmt.Println("\033[33m════════════════════════════════════════════\033[0m")
	fmt.Println()
	fmt.Print("\033[1;31m")
	fmt.Println(` ████ █   █  ███  ████   ███  █   █
█     █   █ █   █ █   █ █   █ █   █
 ███  █████ █████ █   █ █   █ █ █ █
    █ █   █ █   █ █   █ █   █ ██ ██
████  █   █ █   █ ████   ███  █   █`)
	fmt.Print("\033[0m")
	fmt.Println()
	fmt.Print("\033[1;33m")
	fmt.Println(` ███  █████   ████  █████  ████ █████ █████ █   █ █   █
█   █ █       █   █ █     █       █     █   ██  █  █ █ 
█   █ ███     █   █ ███    ███    █     █   █ █ █   █  
█   █ █       █   █ █         █   █     █   █  ██   █  
 ███  █       ████  █████ ████    █   █████ █   █   █  `)
	fmt.Print("\033[0m")
	fmt.Println()
	fmt.Print("\033[3;36m")
	typeWriter("        ~ Une quête au cœur des Terres Ombreuses ~")
	fmt.Print("\033[0m")
	fmt.Println()
	fmt.Println("\033[33m════════════════════════════════════════════\033[0m")
	fmt.Println()
	fmt.Print("\033[3;37m")
	typeWriter("Une flamme pour te guider, une ombre pour te perdre...")
	fmt.Print("\033[0m")
	fmt.Println()
}

func printCredits() {
	fmt.Println()
	fmt.Println("\033[1;33m=== CRÉDITS ===\033[0m")
	fmt.Println("Projet RED - Shadow of Destiny")
	fmt.Println("Développé par Alexis & Leandro")
	fmt.Println("Aix Ynov Campus - Ymmersion")
	fmt.Println()
}
func playGame() {
	c := projetred.CharacterCreation()

	for {
		fmt.Println()
		fmt.Println("\033[1;33m=== MENU ===\033[0m")
		fmt.Println("\033[36m1.\033[0m Afficher les informations du personnage")
		fmt.Println("\033[36m2.\033[0m Accéder à l'inventaire")
		fmt.Println("\033[36m3.\033[0m Marchand")
		fmt.Println("\033[36m4.\033[0m Forgeron")
		fmt.Println("\033[36m5.\033[0m Entrainement")
		fmt.Println("\033[36m6.\033[0m Quitter")
		fmt.Print("Ton choix : ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			c.DisplayInfo()
		case 2:
			c.AccessInventory()
		case 3:
			c.OpenMerchant()
		case 4:
			c.OpenForgeron()
		case 5:
			c.TrainingFight()
		case 6:
			fmt.Println("\033[33mÀ bientôt !\033[0m")
			return
		default:
			fmt.Println("\033[31mChoix invalide, réessaie.\033[0m")
		}
	}
}

func main() {
	printBanner()

	for {
		fmt.Println("\033[36m[1]\033[0m JOUER")
		fmt.Println("\033[36m[2]\033[0m CRÉDITS")
		fmt.Println("\033[36m[3]\033[0m QUITTER")
		fmt.Print("Entre ton choix : ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			playGame()
			return
		case 2:
			printCredits()
		case 3:
			fmt.Println("\033[33mÀ bientôt, aventurier...\033[0m")
			return
		default:
			fmt.Println("\033[31mChoix invalide, réessaie.\033[0m")
		}
	}
}
