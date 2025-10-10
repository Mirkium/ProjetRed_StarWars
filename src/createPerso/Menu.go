package createperso

import (
	"fmt"
	models "game/models"
	utils "game/utils"
)

var (
	options = []string{"NEW  GAME", "SAVE GAME", "QUIT"}
	index   = 0
)

func Menu() {
	for {
		models.ClearScreen()
		printHeader() // ton ASCII inchangé

		// Affichage des 3 lignes interactives
		printOption(0, "NEW  GAME", models.Gray, models.Red)
		printOption(1, "SAVE GAME", models.Gray, models.Red)
		printOption(2, "QUIT", models.Gray, models.Red)

		fmt.Println("  ", models.Red, "█                                                                                                 █ ", models.Reset)
		fmt.Println("  ", models.Red, "█                                                                                                 █ ", models.Reset)
		fmt.Println("  ", models.Red, "█                                                                                                 █ ", models.Reset)
		key := utils.VerticalKey()
		switch key {
		case "up":
			if index > 0 {
				index--
			}
		case "down":
			if index < len(options)-1 {
				index++
			}
		case "enter":
			models.ClearScreen()
			switch index {
			case 0:
				CreatePerso()
				return
			case 1:
				ChargeGame()
				return
			case 2:
				fmt.Println(models.Cyan, "Vous quittez le jeu", models.Reset)
				return
			}
		}
	}
}

func printHeader() {
	fmt.Println(models.Yellow, "    _______.___________.    ___      .______         ____    __    ____  ___      .______          _______.")
	fmt.Println("    /       |           |   /   \\     |   _  \\        \\   \\  /  \\  /   / /   \\     |   _  \\        /       |")
	fmt.Println("   |   (----`---|  |----`  /  ^  \\    |  |_)  |        \\   \\/    \\/   / /  ^  \\    |  |_)  |      |   (----`")
	fmt.Println("    \\   \\       |  |      /  /_\\  \\   |      /          \\            / /  /_\\  \\   |      /        \\   \\  ")
	fmt.Println(".----)   |      |  |     /  _____  \\  |  |\\  \\----.      \\    /\\    / /  _____  \\  |  |\\  \\----.----)   |  ")
	fmt.Println("|_______/       |__|    /__/     \\__\\ | _| `._____|       \\__/  \\__/ /__/     \\__\\ | _| `._____|_______/  ")
	fmt.Println("")
	fmt.Println(models.Gray, "   █                                                                                                 █ ", models.Reset)
	fmt.Println(models.Gray, "   █                                                                                                 █ ", models.Reset)
	fmt.Println(models.Red, "   █                                                                                                 █ ", models.Reset)
}

func printOption(optIndex int, text string, color1, color2 string) {
	textColor := models.Cyan
	if optIndex == index {
		textColor = models.Yellow // texte sélectionné en jaune fluo
	}

	borderColor := models.Red // bordure rouge vif

	totalWidth := 98             // largeur totale de la ligne
	innerWidth := totalWidth - 2 // zone entre les bordures

	leftPadding := (innerWidth - len(text)) / 2
	rightPadding := innerWidth - len(text) - leftPadding

	fmt.Print("    ", borderColor, "█", models.Reset) // bordure gauche

	for i := 0; i < leftPadding; i++ {
		fmt.Print(" ")
	}

	fmt.Print(textColor, text, models.Reset)

	for i := 0; i < rightPadding; i++ {
		fmt.Print(" ")
	}

	fmt.Println(borderColor, "█", models.Reset) // bordure droite
}
