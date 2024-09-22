package jeux


import (
	"fmt"
	"time"
)

func Menu() {
	choix := " "
	fmt.Println("/===============================\\")
	fmt.Println("")
	fmt.Println(Yellow,"       1.",Cyan," character")
	fmt.Println("")
	fmt.Println(Yellow,"       2.",Cyan," Inventory")
	fmt.Println("")
	fmt.Println(Yellow,"       3.",Cyan," Marchand")
	fmt.Println("")
	fmt.Println(Yellow,"       4.",Cyan," Forgeron")
	fmt.Println("")
	fmt.Println(Yellow,"       5.",Cyan," continue")
	fmt.Println("")
	fmt.Println(Yellow,"       0.",Cyan," exit",Reset)
	fmt.Println("")
	fmt.Println("\\===============================/")
	time.Sleep(time.Second * 1)
	fmt.Print("Your choise : ")
	fmt.Scanln(&choix)
	switch choix {
	case "1":
		clearScreen()
		DisplayCharacter()
		clearScreen()
		Menu()
	case "2":
		clearScreen()
		DisplayInventaire()
		clearScreen()
		Menu()
	case "3":
		clearScreen()
		Marchantchoice()
		clearScreen()
		Menu()
	case "4":
		
	case "5":
		break
	case "0":
		clearScreen()
		CreatePerso()
	default:
		clearScreen()
		Menu()
	}
}
