package jeux

import (
	
	"fmt"
)

func Marchantchoice() {
	var choix string
	fmt.Println("/=================SHOP=================\\")
	fmt.Println("")
	fmt.Println("              1. Weapon")
	fmt.Println("")
	fmt.Println("              2. Ability")
	fmt.Println("")
	fmt.Println("              3. Armor")
	fmt.Println("")
	fmt.Println("              4. Item")
	fmt.Println("")
	fmt.Println("              0. Exit")
	fmt.Println("")
	fmt.Println("\\======================================/")
	fmt.Println("")
	fmt.Print("Your choice : ")
	fmt.Scanln(&choix)
	switch choix {
	case "1" :
		ClearScreen()
		WeaponShop()
		clearScreen()
		Marchantchoice()
	case "2" :

	case "3" :

	case "4" :

	case "0" :
		break
	default :

	}
}
