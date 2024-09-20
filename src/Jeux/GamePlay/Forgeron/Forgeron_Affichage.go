package forgeron

import (
	save "Game/Jeux/Sauvegarde"
	"fmt"
)

func SmithchoiceAffichage() {
	fmt.Println("Which blacksmith do you want to meet?")
	fmt.Println("1 - Weaponsmith")
	fmt.Println("2 - Armorsmith")
	fmt.Println("3 - Exit")
	var choix string

	fmt.Print("Enter the number corresponding to your choice: ")
	_, err := fmt.Scan(&choix)

	if err != nil {
		fmt.Println("Error during entry. Please enter a number.")
		return
	}

	switch choix {
	case "1":
		fmt.Println("You have chosen to meet the weaponsmith")
		MenuFWeapon()
	case "2":

		fmt.Println("You have chosen to meet the armorsmith")
		MenuFArmor()
	case "3":
		break
	default:
		fmt.Println("Invalid choice. Please choose a number between 1 and 3.")
		save.ClearScreen()
		SmithchoiceAffichage()
	}
}
