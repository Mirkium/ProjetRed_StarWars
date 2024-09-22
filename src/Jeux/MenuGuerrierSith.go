package jeux

import (
	"fmt"
	"time"
)

func Menu() string {
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
		return "1"
	case "2":
		return "2"
	case "3":
		return "3"
	case "4":
		return "4"
	case "5":
		return "5"
	case "0":
		return "0"
	default:
		return "0"
	}
}
