package Sauvegarde

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

const (
	Reset   = "\033[0m"
	Red     = "\033[91m"
	Green   = "\033[32m"
	Yellow  = "\033[93m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[96m"
	White   = "\033[37m"
	gray    = "\033[90m"
)

var Personnage Perso

func CreatePerso() {
	fmt.Println(" ▓  █ ▄ ▀ ■  ║")
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
	fmt.Println("                              ", Red, "█", Reset, "                                ", Cyan, "2 -  ",  Yellow, "SAVE GAME", Reset, "                                       ", Red, "█", Reset)
	fmt.Println("                              ", Red, "█", Reset, "                                                                                           ", Red, "█", Reset)
	fmt.Println("                              ", Red, "█", Reset, "                                                                                           ", Red, "█", Reset)
	fmt.Println("                              ", Red, "█", Reset, "                                ", Cyan, "3 -  ",  Yellow, "  QUIT", Reset, "                                          ", Red, "█", Reset)
	fmt.Println("                              ", Red, "█                                                                                               █", Reset)
	fmt.Println("                              ", Red, "█                                                                                               █", Reset)
	fmt.Println("                              ", Red, "█                                                                                               █", Reset)
	fmt.Println("                              ", Red, "█                                                                                               █", Reset)
}

func clearScreen() {
	var cmd *exec.Cmd
	// Détecter le système d'exploitation
	switch runtime.GOOS {
	case "windows":
		// Commande pour Windows
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		// Commande pour les systèmes Unix-like
		cmd = exec.Command("clear")
	}
	// Définir la sortie de la commande sur Stdout
	cmd.Stdout = os.Stdout
	cmd.Run()
}
