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
		Marchand_Arme := MarchantWeapon{"fernuiqgrh", []save.Weapon{}}
		MenuWeapon(&save.Personnage, &Marchand_Arme)
	case "2":
		fmt.Println("Vous avez choisi de rencontrer le marchand d'abilités.")
		Marchand_Capacite := MarchantAbiliti{"feznog", []save.Abilite{}}
		MenuAbilitie(&save.Personnage, &Marchand_Capacite)
	case "3":
		fmt.Println("Vous avez choisi de rencontrer le marchand d'armures.")
		Marchand_Armure := MarchantArmor{"geoiqhm", []save.Armor{}}
		MenuArmor(&save.Personnage, &Marchand_Armure)
	case "4":
		fmt.Println("Vous avez choisi de rencontrer le marchand d'Item.")
		Marchand_Item := MarchantItem{"fdp", []save.Item{}}
		MenuItem(&save.Personnage, &Marchand_Item)
	case "5":
		break
	default:
		fmt.Println("Choix invalide. Veuillez choisir un numéro entre 1 et 4.")
		save.ClearScreen()
	}
}
