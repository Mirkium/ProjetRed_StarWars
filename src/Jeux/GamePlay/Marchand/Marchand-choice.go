package marchand

import (
	save "Game/Jeux/Sauvegarde"
	"fmt"
)

type mar struct {
	name    string
	product []Weapon
}

type play struct {
	credit    int
	coteForce int
	bag       []Weapon
}

func Merchantchoice(u *save.Perso, m *mar) {
	fmt.Println("Quel marchand voulez-vous rencontrer ?")
	fmt.Println("1 - Marchand d'armes")
	fmt.Println("2 - Marchand d'abilités")
	fmt.Println("3 - Marchand d'armures")
	fmt.Println("4 - Quitter")
	var choix int

	fmt.Print("Entrez le numéro correspondant à votre choix : ")
	_, err := fmt.Scan(&choix)

	if err != nil {
		fmt.Println("Erreur lors de la saisie. Veuillez entrer un nombre.")
		return
	}

	switch choix {
	case 1:
		fmt.Println("Vous avez choisi de rencontrer le marchand d'armes.")
		m := marcht{"fernuiqgrh", []Weapon{}}
		p := player{u.Credit, u.CoteForce, u.Weapon}
		MenuWeapon(&p, &m)
	case 2:
		fmt.Println("Vous avez choisi de rencontrer le marchand d'abilités.")
		MenuAbiliti(u, m)
	case 3:
		fmt.Println("Vous avez choisi de rencontrer le marchand d'armures.")
		MenuArmor(u, m)
	case 4:
		fmt.Println("Au revoir !")
	default:
		fmt.Println("Choix invalide. Veuillez choisir un numéro entre 1 et 4.")
	}
}
