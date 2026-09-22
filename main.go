package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/mp3"
	"github.com/gopxl/beep/v2/speaker"

	projetred "projetred/src"
)

const termWidth = 130

var musicCtrl *beep.Ctrl
var musicPlaying bool
var musicAvailable bool

func startMusic(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		return err
	}

	if err := speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10)); err != nil {
		return err
	}

	loop := beep.Loop(-1, streamer)
	musicCtrl = &beep.Ctrl{Streamer: loop, Paused: false}

	speaker.Play(musicCtrl)
	musicPlaying = true
	musicAvailable = true
	return nil
}

func toggleMusic() {
	if musicCtrl == nil || !musicAvailable {
		fmt.Println("\033[31mAucune musique chargée.\033[0m")
		return
	}
	speaker.Lock()
	musicCtrl.Paused = !musicCtrl.Paused
	musicPlaying = !musicCtrl.Paused
	speaker.Unlock()
}

func drawMusicStatus() {
	if !musicAvailable {
		return
	}

	label := " Musique : ON   (tape 0 pour couper)"
	color := "\033[1;35m"
	if !musicPlaying {
		label = " Musique : OFF  (tape 0 pour relancer)"
		color = "\033[1;30m"
	}

	fmt.Print(color)
	printCentered(label)
	fmt.Print("\033[0m")
	fmt.Println()
}

func typeWriter(text string) {
	for _, ch := range text {
		fmt.Print(string(ch))
		time.Sleep(30 * time.Millisecond)
	}
	fmt.Println()
}

func printCentered(text string) {
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		padding := (termWidth - len([]rune(line))) / 2
		if padding < 0 {
			padding = 0
		}
		fmt.Println(strings.Repeat(" ", padding) + line)
	}
}

func typeWriterCentered(text string) {
	padding := (termWidth - len([]rune(text))) / 2
	if padding < 0 {
		padding = 0
	}
	fmt.Print(strings.Repeat(" ", padding))
	typeWriter(text)
}

func printBanner() {
	fmt.Print("\033[33m")
	printCentered(`════════════════════════════════════════════`)
	fmt.Print("\033[0m")
	fmt.Println()

	fmt.Print("\033[1;31m")
	printCentered(` ████ █   █  ███  ████   ███  █   █
█     █   █ █   █ █   █ █   █ █   █
 ███  █████ █████ █   █ █   █ █ █ █
    █ █   █ █   █ █   █ █   █ ██ ██
████  █   █ █   █ ████   ███  █   █`)
	fmt.Print("\033[0m")
	fmt.Println()

	fmt.Print("\033[1;33m")
	printCentered(` ███  █████   ████  █████  ████ █████ █████ █   █ █   █
█   █ █       █   █ █     █       █     █   ██  █  █ █ 
█   █ ███     █   █ ███    ███    █     █   █ █ █   █  
█   █ █       █   █ █         █   █     █   █  ██   █  
 ███  █       ████  █████ ████    █   █████ █   █   █  `)
	fmt.Print("\033[0m")
	fmt.Println()

	drawMusicStatus()

	fmt.Print("\033[3;36m")
	typeWriterCentered("~ Une quête au cœur des Terres Ombreuses ~")
	fmt.Print("\033[0m")
	fmt.Println()

	fmt.Print("\033[33m")
	printCentered(`════════════════════════════════════════════`)
	fmt.Print("\033[0m")
	fmt.Println()

	fmt.Print("\033[3;37m")
	typeWriterCentered("Une flamme pour te guider, une ombre pour te perdre...")
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
		drawMusicStatus()
		fmt.Println("\033[36m1.\033[0m Afficher les informations du personnage")
		fmt.Println("\033[36m2.\033[0m Accéder à l'inventaire")
		fmt.Println("\033[36m3.\033[0m Marchand")
		fmt.Println("\033[36m4.\033[0m Forgeron")
		fmt.Println("\033[36m5.\033[0m Entrainement")
		fmt.Println("\033[36m6.\033[0m Quitter")
		fmt.Println("\033[36m0.\033[0m Musique on/off")
		fmt.Print("Ton choix : ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 0:
			toggleMusic()
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
	if err := startMusic("assets/theme.mp3"); err != nil {
		fmt.Println("\033[31m(Musique indisponible : " + err.Error() + ")\033[0m")
		musicAvailable = false
	}

	printBanner()

	for {
		fmt.Println("\033[36m[1]\033[0m JOUER")
		fmt.Println("\033[36m[2]\033[0m CRÉDITS")
		fmt.Println("\033[36m[3]\033[0m QUITTER")
		fmt.Println("\033[36m[0]\033[0m MUSIQUE ON/OFF")
		fmt.Print("Entre ton choix : ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 0:
			toggleMusic()
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
