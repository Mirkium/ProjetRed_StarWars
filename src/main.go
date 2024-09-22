package main

import (
	Combat "Game/Jeux/GamePlay/Fight"
	chevalierJedi "Game/Jeux/GamePlay/FirstMission/Jedi/ChevalierJedi"
	save "Game/Jeux/Sauvegarde"
	"fmt"

	consulaireJedi "Game/Jeux/GamePlay/FirstMission/Jedi/ConsulaireJedi"
	assassinSith "Game/Jeux/GamePlay/FirstMission/Sith/AssassinSith"
	guerrierSith "Game/Jeux/GamePlay/FirstMission/Sith/GuerrierSith"
	forgeron "Game/Jeux/GamePlay/Forgeron"
	marchand "Game/Jeux/GamePlay/Marchand"
	"time"
)

func main() {
	save.CreatePerso()
	switch save.Campagne.Name {
	case "Jedi Knight":
		chevalierJedi.JediKnight()
	case "Jedi Consular":
		JediConsular()
	case "Sith Warrior":
		guerrierSith.SithWarrior()
		save.Personnage.CoteForce += guerrierSith.Arrive_2()
		guerrierSith.Arrive_3()
		WayClasseSithWarrior()
		MenuGuerrierSith1()
	case "Sith Assassin":
		assassinSith.SithAssassin()

	}
}

func MenuGuerrierSith1() {
	choice_1 := guerrierSith.Menu()
	switch choice_1 {
	case "1":
		save.DisplayCharacter()
		MenuGuerrierSith1()
	case "2":
		save.DisplayInventaire()
		MenuGuerrierSith1()
	case "3":
		marchand.Merchantchoice()
		MenuGuerrierSith1()
	case "4":
		forgeron.SmithchoiceAffichage()
		MenuGuerrierSith1()
	case "5":
		if Combat.Fight(&save.Personnage, &guerrierSith.LimaceKor_Rang1, true) {
			if Combat.Fight(&save.Personnage, &guerrierSith.LimaceKor_Rang2, true) {
				if Combat.Fight(&save.Personnage, &guerrierSith.LimaceKor_Rang3, false) {

				}
			}
		}
	}
}

func MenuGuerrierSith2() {
	choice_1 := guerrierSith.Menu()
	switch choice_1 {
	case "1":
		save.DisplayCharacter()
		MenuGuerrierSith1()
	case "2":
		save.DisplayInventaire()
		MenuGuerrierSith1()
	case "3":
		marchand.Merchantchoice()
		MenuGuerrierSith1()
	case "4":
		forgeron.SmithchoiceAffichage()
		MenuGuerrierSith1()
	case "5":

	}
}

func WayClasseSithWarrior() {
	var choix string
	fmt.Println("Choose your Ability class :")
	fmt.Println("(1) Tank :", guerrierSith.ClassTank.Name, "     (2) Burst : ", guerrierSith.ClassBurst.Name, "     (3) Dot : ", guerrierSith.ClassDot.Name)
	fmt.Scanln(&choix)
	switch choix {
	case "1":
		save.Personnage.Classe = guerrierSith.ClassTank
	case "2":
		save.Personnage.Classe = guerrierSith.ClassBurst
	case "3":
		save.Personnage.Classe = guerrierSith.ClassDot
	default:
		save.ClearScreen()
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
		consulaireJedi.IntroConsulaire()
		ChoixClasse()
	default:
		fmt.Println("Bad Input")
		time.Sleep(2 * time.Second)
		save.ClearScreen()
		JediConsular()
	}
}

func ChoixClasse() {
	save.ClearScreen()
	for i, element := range consulaireJedi.ClasseConsulaireList {
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
			save.Personnage.Classe = consulaireJedi.ClasseConsulaire1
			Ok = true
		case "2":
			save.Personnage.Classe = consulaireJedi.ClasseConsulaire2
			Ok = true
		case "3":
			save.Personnage.Classe = consulaireJedi.ClasseConsulaire3
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
		fmt.Print("Your choise : \n")
		fmt.Scanln(&choix)
		switch choix {
		case "1":
			Ok = true
			save.DisplayCharacter()
		case "2":
			save.DisplayInventaire()
		case "3":
			marchand.Merchantchoice()
		case "4":
			Ok = true
		case "5":
			if consulaireJedi.Quete == 0 {
				consulaireJedi.Quete1(&save.Personnage)
			} else if consulaireJedi.Quete == 1 {
				consulaireJedi.Quete2(&save.Personnage)
			} else {
				consulaireJedi.Final()
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
