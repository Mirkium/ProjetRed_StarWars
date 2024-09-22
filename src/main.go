package main

import (
	game "Game/Jeux"
	"fmt"
	"time"
)

func main() {
	game.PlayMusic(7)
	game.CreatePerso()
	switch game.Campagne.Name {
	case "Jedi Knight":
		game.JediKnight()
	case "Jedi Consular":
		JediConsular()
	case "Sith Warrior":
		game.SithWarrior()
		game.Personnage.CoteForce += game.Arrive_2()
		game.Arrive_3()
		WayClasseSithWarrior()
		MenuGuerrierSith1()
	case "Sith Assassin":
		game.SithAssassin()

	}
}

func MenuGuerrierSith1() {
	choice_1 := game.Menu()
	switch choice_1 {
	case "1":
		game.ClearScreen()
		game.DisplayCharacter()
		MenuGuerrierSith1()
	case "2":
		game.ClearScreen()
		game.DisplayInventaire()
		MenuGuerrierSith1()
	case "3":
		game.ClearScreen()
		game.Marchantchoice()
		MenuGuerrierSith1()
	case "4":
		game.ClearScreen()

		MenuGuerrierSith1()
	case "5":
		game.ClearScreen()
		if game.Fight(&game.Personnage, &game.LimaceKor_Rang1, true) {
			fmt.Println("")
			MenuGuerrierSith2()
		}
	}
}

func MenuGuerrierSith2() {
	choice_1 := game.Menu()
	switch choice_1 {
	case "1":
		game.ClearScreen()
		game.DisplayCharacter()
		MenuGuerrierSith1()
	case "2":
		game.ClearScreen()
		game.DisplayInventaire()
		MenuGuerrierSith1()
	case "3":
		game.ClearScreen()
		MenuGuerrierSith1()
	case "4":
		game.ClearScreen()
		MenuGuerrierSith1()
	case "5":

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
	Menu()

}

func Menu() {
	var choix string
	Ok := false
	for !Ok {
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
		case "1":
			Ok = true
			game.DisplayCharacter()
		case "2":
			game.DisplayInventaire()
		case "3":
			game.Marchantchoice()
		case "4":
			Ok = true
		case "5":
			if game.Quete == 0 {
				game.Quete1(&game.Personnage)
				Menu()
			} else if game.Quete == 1 {
				game.Quete2(&game.Personnage)
				Menu()
			} else {
				game.Final()
				return
			}
		case "0":
			return
		default:
			Menu()
		}
	}
	Menu()
}
