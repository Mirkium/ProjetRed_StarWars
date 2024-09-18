package forgeron

import (
	"fmt"
	save "Game/Jeux/Sauvegarde"
)
func Merchantchoice(u *save.Perso, m *mar, w *save.Weapon) {
	fmt.Println("Quel forgeron voulez-vous rencontrer ?")
	fmt.Println("1 - Forgeron d'armes")
	fmt.Println("2 - Forgeron d'armure")
	fmt.Println("3 - Quitter")
	var choix int

	fmt.Print("Entrez le numéro correspondant à votre choix : ")
	_, err := fmt.Scan(&choix)

	if err != nil {
		fmt.Println("Erreur lors de la saisie. Veuillez entrer un nombre.")
		return
	}

	switch choix {
	case 1:
		fmt.Println("Vous avez choisi de rencontrer le forgeron d'armes.")
		m := Marcht{"fernuiqgrh", []save.Weapon{}}
		p := Player{u.Credit, u.CoteForce, u.Weapon}
		MenuWeapon(&p, &m)
	case 2:
		fmt.Println("Vous avez choisi de rencontrer le forgeron d'armures.")
		m3 := Merchant{"geoiqhm", []save.Armor{}}
		p3 := Plyer{u.Credit, u.CoteForce, u.Armure}
		MenuArmor(&p3, &m3)
	case 3:
		fmt.Println("Au revoir !")
	default:
		fmt.Println("Choix invalide. Veuillez choisir un numéro entre 1 et 4.")
	}
}
