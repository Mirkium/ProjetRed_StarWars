package main

import (
	Combat "Game/Jeux/GamePlay/Fight"
	chevalierJedi "Game/Jeux/GamePlay/FirstMission/Jedi/ChevalierJedi"
	save "Game/Jeux/Sauvegarde"
	"fmt"

	//consulaireJedi "Game/Jeux/GamePlay/FirstMission/Jedi/ConsulaireJedi"
	assassinSith "Game/Jeux/GamePlay/FirstMission/Sith/AssassinSith"
	guerrierSith "Game/Jeux/GamePlay/FirstMission/Sith/GuerrierSith"
	forgeron "Game/Jeux/GamePlay/Forgeron"
	marchand "Game/Jeux/GamePlay/Marchand"
)

func main() {
	save.CreatePerso()
	switch save.Campagne.Name {
	case "Jedi Knight":
		chevalierJedi.JediKnight()
	case "Jedi Consular":
		//consulaireJedi.JediConsular()
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
