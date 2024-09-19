package guerriersith

import (
	"fmt"
)

func Menu() string{
	var choix string
	fmt.Println("/===============================\\")
	fmt.Println("")
	fmt.Println("           1. character")
	fmt.Println("")
	fmt.Println("           2. Inventory")
	fmt.Println("")
	fmt.Println("           3. Marchand")
	fmt.Println("")
	fmt.Println("           4. Forgeron")
	fmt.Println("")
	fmt.Println("           5. continue")
	fmt.Println("")
	fmt.Println("           0. exit")
	fmt.Println("")
	fmt.Println("\\===============================/")
	fmt.Print("Your choise : ")
	fmt.Scanln(&choix)
	switch choix {
	case "1" :
		return "1"
	case "2" :
		return "2"
	case "3" :
		return "3"
	case "4" :
		return "4"
	case "5" :
		return "5"
	case "0" :
		return "0"
	default :
		return "0"
	}
}