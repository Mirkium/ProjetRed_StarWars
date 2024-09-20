package marchand

import (
	save "Game/Jeux/Sauvegarde"
	"fmt"
)

func Merchantchoice() {
	fmt.Println("Which merchant do you want to meet?")
	fmt.Println("1 - WeaponMerchant")
	fmt.Println("2 - Ability Merchant")
	fmt.Println("3 - Armor Merchant")
	fmt.Println("4-  Item Merchant")
	fmt.Println("5 - Exit")
	var choix string

	fmt.Print("Enter the number corresponding to your choice: ")
	_, err := fmt.Scan(&choix)

	if err != nil {
		fmt.Println("Error during entry. Please enter a number.")
		return
	}

	switch choix {
	case "1":
		fmt.Println("You chose to meet the weapon merchant ")
		MenuWeapon()
	case "2":

		fmt.Println("You have chosen to meet the ability merchant ")
		MenuAbilitie()
	case "3":

		fmt.Println("You chose to meet the armor merchant ")
		MenuArmor()
	case "4":

		fmt.Println("You chose to meet the Item merchant ")
		MenuItem()
	case "5":
		break
	default:
		fmt.Println("Invalid choice. Please choose a number between 1 and 4.")
		save.ClearScreen()
	}
}
