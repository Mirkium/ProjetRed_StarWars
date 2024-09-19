package marchand

import (
	save "Game/Jeux/Sauvegarde"
	"fmt"
)

func Merchantchoice(u *save.Perso) {
	fmt.Println("Quel marchand voulez-vous rencontrer ?")
	fmt.Println("1 - Marchand d'armes")
	fmt.Println("2 - Marchand d'abilités")
	fmt.Println("3 - Marchand d'armures")
	fmt.Println("4-  Marchand d'Item")
	fmt.Println("5 - Quitter")
	var choix string

	fmt.Print("Entrez le numéro correspondant à votre choix : ")
	_, err := fmt.Scan(&choix)

	if err != nil {
		fmt.Println("Erreur lors de la saisie. Veuillez entrer un nombre.")
		return
	}

	switch choix {
	case "1":

		fmt.Println("Vous avez choisi de rencontrer le marchand d'armes.")
		MenuWeapon()
	case "2":
	
		fmt.Println("Vous avez choisi de rencontrer le marchand d'abilités.")
		MenuAbilitie()
	case "3":
		
		fmt.Println("Vous avez choisi de rencontrer le marchand d'armures.")
		MenuArmor()
	case "4":

		fmt.Println("Vous avez choisi de rencontrer le marchand d'Item ")
		MenuItem()
	case "5":
		break
	default:
		fmt.Println("Choix invalide. Veuillez choisir un numéro entre 1 et 4.")
		save.ClearScreen()
	}
}
