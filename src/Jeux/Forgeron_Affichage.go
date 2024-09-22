package jeux

import (
	"fmt"
)

func MenuForgeron() {
	var choix string
	fmt.Println("/===================FORGE===================\\")
	fmt.Println("")
	fmt.Println("             1. Weapon Forge")
	fmt.Println("")
	fmt.Println("             2. Armor Forge")
	fmt.Println("")
	fmt.Println("             0. Exit")
	fmt.Println("")
	fmt.Println("\\===========================================/")
	fmt.Println("")
	fmt.Print("Your choice : ")
	fmt.Scanln(&choix)
	switch choix {
	case "1" :

	case "2" :

	case "0" :
		break
	default :
		clearScreen()
		MenuForgeron()
	}
}