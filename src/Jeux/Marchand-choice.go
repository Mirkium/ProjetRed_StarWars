package jeux

import (
	"fmt"
)

func Marchantchoice() {
	var choix string
	fmt.Println("/=================SHOP=================\\")
	fmt.Println("")
	fmt.Println(Yellow, "          1.", Cyan, " Weapon")
	fmt.Println("")
	fmt.Println(Yellow, "          2.", Cyan, " Ability")
	fmt.Println("")
	fmt.Println(Yellow, "          3.", Cyan, " Armor")
	fmt.Println("")
	fmt.Println(Yellow, "          4.", Cyan, " Item")
	fmt.Println("")
	fmt.Println(Yellow, "          0.", Cyan, " Exit", Reset)
	fmt.Println("")
	fmt.Println("\\======================================/")
	fmt.Println("")
	fmt.Print("Your choice : ")
	fmt.Scanln(&choix)
	switch choix {
	case "1":
		ClearScreen()
		WeaponShop()
		clearScreen()
		Marchantchoice()
	case "2":
		clearScreen()
		AbilityShop()
		clearScreen()
		Marchantchoice()
	case "3":
		clearScreen()
		ArmorShop()
		clearScreen()
		Marchantchoice()
	case "4":
		clearScreen()
		ItemShop()
		clearScreen()
		Marchantchoice()
	case "0":
		break
	default:
		clearScreen()
		Marchantchoice()
	}
}
