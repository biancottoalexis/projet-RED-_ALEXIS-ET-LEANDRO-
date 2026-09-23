package projetred

import (
	"fmt"
	"time"
)

func storyTypeWriter(text string) {
	for _, ch := range text {
		fmt.Print(string(ch))
		time.Sleep(25 * time.Millisecond)
	}
	fmt.Println()
}

func (c *Character) StartStory() {
	fmt.Println()
	fmt.Println("\033[1;35m=== PROLOGUE ===\033[0m")
	fmt.Print("\033[3;37m")
	storyTypeWriter("Tu te réveilles dans les Terres Ombreuses, sans souvenir de comment tu es arrivé là.")
	storyTypeWriter("À ton poignet, un Bracelet du Destin fissuré murmure des bribes de futur.")
	storyTypeWriter("Un vieil homme borgne, le Passeur, t'explique : le jour où tu as failli mourir,")
	storyTypeWriter("ton destin s'est brisé en deux. Une partie de toi est restée ici, et marche.")
	storyTypeWriter("On l'appelle... l'Ombre.")
	fmt.Print("\033[0m")

	c.Chapter1()
}

func (c *Character) Chapter1() {
	fmt.Println()
	fmt.Println("\033[1;35m=== CHAPITRE 1 : Le Bracelet Brisé ===\033[0m")
	fmt.Print("\033[3;37m")
	storyTypeWriter("Le bracelet vibre violemment près d'une ruine effondrée.")
	storyTypeWriter("Il pourrait te montrer une vision de ce qui t'attend plus loin...")
	fmt.Print("\033[0m")

	fmt.Println("\033[35m1.\033[0m Écouter la vision")
	fmt.Println("\033[35m2.\033[0m Ignorer et avancer")
	fmt.Print("Ton choix : ")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		c.Initiative += 2
		fmt.Println("\033[35mUn murmure glacé reste avec toi... mais tu sens le danger venir. (+2 Initiative)\033[0m")
	default:
		fmt.Println("\033[35mTu avances, serein, à l'aveugle.\033[0m")
	}

	c.Chapter2()
}

func (c *Character) Chapter2() {
	fmt.Println()
	fmt.Println("\033[1;35m=== CHAPITRE 2 : La Forêt qui Doute ===\033[0m")
	fmt.Print("\033[3;37m")
	storyTypeWriter("Une forêt où les arbres se penchent selon ce que tu ressens.")
	storyTypeWriter("Tu trouves une créature blessée, à moitié translucide, couchée dans les feuilles.")
	fmt.Print("\033[0m")

	fmt.Println("\033[35m1.\033[0m La soigner")
	fmt.Println("\033[35m2.\033[0m T'en méfier et partir")
	fmt.Print("Ton choix : ")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		c.Gold += 10
		fmt.Println("\033[35mLa créature guérit et te mène à un petit trésor caché. (+10 Or)\033[0m")
	default:
		fmt.Println("\033[35mTu poursuis ta route, méfiant.\033[0m")
	}

	c.Chapter3()
}

func (c *Character) Chapter3() {
	fmt.Println()
	fmt.Println("\033[1;35m=== CHAPITRE 3 : Le Village des Reflets ===\033[0m")
	fmt.Print("\033[3;37m")
	storyTypeWriter("Chaque habitant de ce village a une version \"ombre\" de lui-même,")
	storyTypeWriter("visible seulement au crépuscule.")
	fmt.Print("\033[0m")

	fmt.Println("\033[35m1.\033[0m Aider un villageois")
	fmt.Println("\033[35m2.\033[0m Écouter les reflets")
	fmt.Print("Ton choix : ")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		c.Gold += 15
		fmt.Println("\033[35mLe villageois te remercie chaleureusement. (+15 Or)\033[0m")
	case 2:
		c.Initiative += 2
		fmt.Println("\033[35mLes reflets te révèlent une faiblesse de ton Ombre. (+2 Initiative)\033[0m")
	default:
		fmt.Println("\033[35mTu restes silencieux et poursuis ta route.\033[0m")
	}

	fmt.Println()
	fmt.Println("\033[1;35m(La suite de l'histoire arrive bientôt...)\033[0m")
}
