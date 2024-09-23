package main

import (
	game "Game/Jeux"
	"fmt"
	"time"
)

func main() {
	game.CreatePerso()
	switch game.Campagne.Name {
	case "Jedi Knight":
		game.CampaingJediKnight(&game.Personnage)
	case "Jedi Consular":
		JediConsular()
	case "Sith Warrior":
		game.SithWarrior()
		game.Arrive_1()
		game.Personnage.CoteForce += game.Arrive_2()
		game.Arrive_3()
		WayClasseSithWarrior()
		game.Menu()
		if game.Fight(&game.Personnage, &game.LimaceKor_Rang1, true) {
			game.Menu()
		}
	case "Sith Assassin":
		game.SithAssassin()
	}
}

func WayClasseSithWarrior() {
	var choix string
	fmt.Println("Choose your Ability class :")
	fmt.Println("(1) Tank :", game.ClassTank.Name, "     (2) Burst : ", game.ClassBurst.Name, "     (3) Dot : ", game.ClassDot.Name)
	fmt.Scanln(&choix)
	switch choix {
	case "1":
		game.Personnage.Classe = game.ClassTank
	case "2":
		game.Personnage.Classe = game.ClassBurst
	case "3":
		game.Personnage.Classe = game.ClassDot
	default:
		game.ClearScreen()
		WayClasseSithWarrior()
	}
}

// Jedi consulaire
func JediConsular() {
	var choix_Intro string
	fmt.Println("Do you want skip Intro ?")
	fmt.Println(" (1) yes     (2) no")
	fmt.Scanln(&choix_Intro)
	switch choix_Intro {
	case "1":
		ChoixClasse()
	case "2":
		game.IntroConsulaire()
		ChoixClasse()

	default:
		fmt.Println("Bad Input")
		time.Sleep(2 * time.Second)
		game.ClearScreen()
		JediConsular()
	}
}

func ChoixClasse() {
	game.ClearScreen()
	for i, element := range game.ClasseConsulaireList {
		fmt.Printf("%s (%d): \n\tEnergie : %d\n\tAbilite : \n", element.Name, i+1, element.Energie)
		for _, ele := range element.Abilite {
			fmt.Printf("\t\t - %s, Cost Energie : %d, Dammage : %d, Heal : %d, Dot duartion : %d, Dot dammage : %d.\n", ele.Name, ele.EnergieCost, ele.Dammage, ele.Heal, ele.DotCompteur, ele.DotDammage)
		}
	}
	time.Sleep(2 * time.Second)
	fmt.Printf("Which class do you want to chosse ?")
	var choix string
	Ok := false
	for !Ok {
		fmt.Scan(&choix)
		switch choix {
		case "1":
			game.Personnage.Classe = game.ClasseConsulaire1
			Ok = true
		case "2":
			game.Personnage.Classe = game.ClasseConsulaire2
			Ok = true
		case "3":
			game.Personnage.Classe = game.ClasseConsulaire3
			Ok = true
		default:
			Ok = false
		}
	}
	game.Menu()
}
