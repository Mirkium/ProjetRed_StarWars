package jeux

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func SithWarrior() int {
	var choix_Intro string
	fmt.Println("Do you want skip Intro ?")
	fmt.Println(" (1) yes     (2) no")
	fmt.Scanln(&choix_Intro)
	switch choix_Intro {
	case "1":
		return 1
	case "2":
		return 2
	}
	return 2
}

func Arrive_1() {
	var choix string
	fmt.Println("Inspector Tremel: Welcome to Korriban Acolyte! I brought you here because")
	fmt.Println("                  you were our last hope of restoring purity to the Sith Order.")
	fmt.Println("")
	fmt.Println("         1. I'm descended from a great line of Sith.")
	fmt.Println("         2. I can't wait to start killing.")
	fmt.Println("         3. I'm honored.")
	fmt.Scanln(&choix)
	switch choix {
	case "1":
		clearScreen()
		fmt.Println("Inspector Tremel: I'm sure that with you, the purity of the Siths will be respected.")
		time.Sleep(time.Second * 2)
		clearScreen()
	case "2":
		clearScreen()
		fmt.Println("Inspector Tremel: Note that it's forbidden to kill other acolytes inside the academy.")
		time.Sleep(time.Second * 2)
		clearScreen()

	case "3":
		clearScreen()
	}
}

func Arrive_2() int {
	var Return int
	var choix string
	fmt.Println("Inspector Tremel: Join the group behind me, but watch out for Vemrin, he's your rival,")
	fmt.Println("                  and he'll try to kill you.")
	fmt.Println(" ")
	fmt.Println("                  1. All right, excellence.")
	fmt.Println("                  2. I thought acolytes weren't supposed to kill each other?")
	fmt.Println("                  3. Violence isn't always the answer. (Light point)")
	fmt.Scanln(&choix)
	switch choix {
	case "1":
		clearScreen()
		Return = 0
	case "2":
		clearScreen()
		fmt.Println("Inspector Tremel: In the academy, but in the tombs, it's different. So be careful.")
		time.Sleep(time.Second * 2)
		Return = 0
	case "3":
		clearScreen()
		fmt.Println("Inspecteur Tremel: This is not the vision of a Sith.")
		time.Sleep(time.Second * 2)
		Return = 50
		clearScreen()
	}
	return Return
}

func Arrive_3() {
	fmt.Println("Inspecteur Tremel: Dear Acolytes, you've all done a great job, ")
	fmt.Println("                   and now you deserve to go and get your Sith blade.")
	time.Sleep(time.Second * 2)
	clearScreen()
}

func Vemrin_1() {

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
