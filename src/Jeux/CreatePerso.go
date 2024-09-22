package jeux

import (
	"fmt"
)


var Personnage Perso
var Arme Lightsaber
var ArmeSecondaire SecondaryLightsaber
var Armors Armor
var Campagne Classe
var Capacite Abilite

func CreatePerso() {
	var choix string
	fmt.Println(Yellow, "        ▄▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀█     █▀▀▀▀▀█        █▀▀▀▀▀▀▀▀▀▄         █▀▀▀▀█     █▀▀▀█     █▀▀▀▀█   █▀▀▀▀▀█        █▀▀▀▀▀▀▀▀▀▄      ▄▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀█")
	fmt.Println(Yellow, "       █                            █    █   ▄   █       █   █▀▀▄   █         █    █   █     █   █    █   █   ▄   █       █   █▀▀▄   █    █                █")
	fmt.Println(Yellow, "        ▀▄    ▀█▀▀▀▀▀▀▀▀▀█   █▀▀▀▀▀▀▀   █   █ █   █      █   ▀▀▀  ▄▀           █    █ █       █ █    █   █   █ █   █      █   ▀▀▀  ▄▀      ▀▄    ▀█▀▀▀▀▀▀▀▀▀")
	fmt.Println(Yellow, "          ▀▄    ▀▄       █   █         █   █▄▄▄█   █     █   ▄     ▀▄           █    ▀    ▄    ▀    █   █   █▄▄▄█   █     █   ▄     ▀▄       ▀▄    ▀▄")
	fmt.Println(Yellow, "      ▄▄▄▄▄▄█▄    ▀▄     █   █        █    ▄▄▄▄▄    █    █   █▀▄     ▀▄▄▄▄▄▄▄    █       █ █       █   █    ▄▄▄▄▄    █    █   █▀▄     ▀▄▄▄▄▄▄▄▄█▄    ▀▄")
	fmt.Println(Yellow, "      █             █    █   █       █    █     █    █   █   █  ▀▄          █     █     █   █     █   █    █     █    █   █   █  ▀▄                    █")
	fmt.Println(Yellow, "      █▄▄▄▄▄▄▄▄▄▄▄▄▀     █▄▄▄█      █▄▄▄▄█       █▄▄▄▄█  █▄▄▄█    ▀▄▄▄▄▄▄▄▄▄█      █▄▄▄█     █▄▄▄█   █▄▄▄▄█       █▄▄▄▄█  █▄▄▄█    ▀▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▀")
	fmt.Println("▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬")
	fmt.Println("                              ", gray, "▄", Reset, "                                      ", Yellow, "PROJET  RED", Reset, "                                      ", gray, "▄", Reset)
	fmt.Println("                              ", gray, "▓", Reset, "                                                                                           ", gray, "▓", Reset)
	fmt.Println("                              ", gray, "▓", Reset, "                                                                                           ", gray, "▓", Reset)
	fmt.Println("                              ", gray, "█", Reset, "                                ", Cyan, "1 -  ", Yellow, "NEW  GAME", Reset, "                                       ", gray, "█", Reset)
	fmt.Println("                              ", Red, "█", Reset, "                                                                                           ", Red, "█", Reset)
	fmt.Println("                              ", Red, "█", Reset, "                                                                                           ", Red, "█", Reset)
	fmt.Println("                              ", Red, "█", Reset, "                                ", Cyan, "2 -  ", Yellow, "SAVE GAME", Reset, "                                       ", Red, "█", Reset)
	fmt.Println("                              ", Red, "█", Reset, "                                                                                           ", Red, "█", Reset)
	fmt.Println("                              ", Red, "█", Reset, "                                                                                           ", Red, "█", Reset)
	fmt.Println("                              ", Red, "█", Reset, "                                ", Cyan, "3 -  ", Yellow, "  QUIT", Reset, "                                          ", Red, "█", Reset)
	fmt.Println("                              ", Red, "█                                                                                               █", Reset)
	fmt.Println("                              ", Red, "█                                                                                               █", Reset)
	fmt.Println("                              ", Red, "█                                                                                               █", Reset)
	fmt.Println("                              ", Red, "█                                                                                               █", Reset)
	fmt.Scanln(&choix)
	switch choix {
	case "1":
		ClearScreen()
		CaracteristiquePerso()
	case "2":
		ClearScreen()
		ChargeGame()
	case "3":
		var choixQuit string
		ClearScreen()
		fmt.Println("Do you want a quit game ?")
		fmt.Println("  1 - yes    2 - no")
		fmt.Scanln(&choixQuit)
		switch choixQuit {
		case "1":
			break
		case "2":
			ClearScreen()
			CreatePerso()
		}
	default:
		ClearScreen()
		CreatePerso()
	}
}

func CaracteristiquePerso() {
	var name string

	fmt.Print("Write your name : ")
	fmt.Scanln(&name)
	Personnage.Name = name
	ClearScreen()
	var choixWay string
	fmt.Println("    	       ", Yellow, "CHOOSE YOUR WAY !", Reset)
	fmt.Println(" (1) ", Cyan, "Republic", Reset, "                  (2) ", Red, "Empire", Reset)
	fmt.Scanln(&choixWay)
	switch choixWay {
	case "1":
		Republic()
	case "2":
		SithEmpire()
	default:
		ClearScreen()
		CaracteristiquePerso()
	}

}
