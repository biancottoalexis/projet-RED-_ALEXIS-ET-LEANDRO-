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

	c.Chapter4()
}

func (c *Character) Chapter4() {
	fmt.Println()
	fmt.Println("\033[1;35m=== CHAPITRE 4 : Le Pacte du Passeur ===\033[0m")
	fmt.Print("\033[3;37m")
	storyTypeWriter("Le Passeur réapparaît. Il t'avoue une vérité troublante :")
	storyTypeWriter("c'est lui qui a brisé ton destin en deux, \"pour te sauver\".")
	storyTypeWriter("Il te propose un pacte : une part de tes souvenirs, contre de la force.")
	fmt.Print("\033[0m")

	fmt.Println("\033[35m1.\033[0m Accepter le pacte")
	fmt.Println("\033[35m2.\033[0m Refuser")
	fmt.Print("Ton choix : ")

	var choice int
	fmt.Scanln(&choice)

	pactAccepted := false

	switch choice {
	case 1:
		pactAccepted = true
		c.MaxHP += 20
		c.CurrentHP += 20
		fmt.Println("\033[35mUn frisson glacé te parcourt. Tu sens une force nouvelle en toi. (+20 PV max)\033[0m")
	default:
		fmt.Println("\033[35mTu refuses. Le Passeur hoche la tête, presque déçu.\033[0m")
	}

	c.Chapter5(pactAccepted)
}

func (c *Character) Chapter5(pactAccepted bool) {
	fmt.Println()
	fmt.Println("\033[1;35m=== CHAPITRE 5 : Les Marais Maudits ===\033[0m")
	fmt.Print("\033[3;37m")
	storyTypeWriter("Un gardien de brume bloque l'unique passage vers le seuil final.")
	fmt.Print("\033[0m")

	fmt.Println("\033[35m1.\033[0m L'affronter")
	fmt.Println("\033[35m2.\033[0m Le contourner en silence")
	fmt.Print("Ton choix : ")

	var choice int
	fmt.Scanln(&choice)

	if choice == 1 {
		guardian := InitGuardian()
		turn := 1
		for {
			fmt.Println()
			fmt.Println("\033[1;35m=== Tour", turn, "===\033[0m")
			c.CharacterTurn(&guardian)
			if guardian.CurrentHP <= 0 {
				fmt.Println("\033[32mTu as vaincu le", guardian.Name, "!\033[0m")
				c.GainXP(guardian.XPReward)
				break
			}
			guardian.GoblinPattern(c, turn)
			if c.IsDead() {
				break
			}
			turn++
		}
	} else {
		fmt.Println("\033[35mTu te faufiles dans la brume, invisible.\033[0m")
	}

	c.Chapter6(pactAccepted)
}

func (c *Character) Chapter6(pactAccepted bool) {
	fmt.Println()
	fmt.Println("\033[1;35m=== CHAPITRE 6 : Le Seuil du Destin ===\033[0m")
	fmt.Print("\033[3;37m")
	storyTypeWriter("Face à toi, une silhouette qui a exactement ta voix.")
	storyTypeWriter("C'est elle. C'est toi.")
	fmt.Print("\033[0m")

	fmt.Println("\033[35m1.\033[0m Y aller avec confiance")
	fmt.Println("\033[35m2.\033[0m Y aller avec prudence")
	fmt.Print("Ton choix : ")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		c.Initiative += 3
		fmt.Println("\033[35mTa confiance t'aiguise les sens. (+3 Initiative)\033[0m")
	default:
		c.CurrentHP = c.MaxHP
		fmt.Println("\033[35mTu inspires profondément et soignes tes blessures. (PV restaurés)\033[0m")
	}

	c.ShadowFight(pactAccepted)
}

func (c *Character) ShadowFight(pactAccepted bool) {
	fmt.Println()
	fmt.Println("\033[1;35m=== COMBAT FINAL : Ton Ombre ===\033[0m")

	shadow := InitShadow()
	turn := 1

	if c.Initiative >= shadow.Initiative {
		fmt.Println("\033[35mTu frappes en premier.\033[0m")
	} else {
		fmt.Println("\033[35mTon Ombre est plus rapide que toi et attaque en premier !\033[0m")
	}

	for {
		fmt.Println()
		fmt.Println("\033[1;35m=== Tour", turn, "===\033[0m")

		if c.Initiative >= shadow.Initiative {
			c.CharacterTurn(&shadow)
			if shadow.CurrentHP <= 0 {
				break
			}
			shadow.GoblinPattern(c, turn)
			if c.IsDead() {
				fmt.Println("\033[31mTon Ombre t'a submergé... Le combat est terminé.\033[0m")
				return
			}
		} else {
			shadow.GoblinPattern(c, turn)
			if c.IsDead() {
				fmt.Println("\033[31mTon Ombre t'a submergé... Le combat est terminé.\033[0m")
				return
			}
			c.CharacterTurn(&shadow)
			if shadow.CurrentHP <= 0 {
				break
			}
		}

		turn++
	}

	c.GainXP(shadow.XPReward)
	c.printEnding(pactAccepted)
}

func (c *Character) printEnding(pactAccepted bool) {
	fmt.Println()
	if pactAccepted {
		fmt.Println("\033[1;35m=== FIN DE LA FUSION ===\033[0m")
		fmt.Print("\033[3;37m")
		storyTypeWriter("Tu ne détruis pas ton Ombre. Tu l'absorbes.")
		storyTypeWriter("Tu as gagné... mais un doute plane sur qui tu es vraiment, désormais.")
		fmt.Print("\033[0m")
	} else {
		fmt.Println("\033[1;35m=== FIN DE LA RÉCONCILIATION ===\033[0m")
		fmt.Print("\033[3;37m")
		storyTypeWriter("Tu comprends que ton Ombre n'était pas ton ennemie, juste ta peur.")
		storyTypeWriter("Le Bracelet du Destin se répare. Une lumière chaude t'enveloppe.")
		fmt.Print("\033[0m")
	}
	fmt.Println()
}
