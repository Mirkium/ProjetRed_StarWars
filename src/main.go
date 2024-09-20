package main

import (
	"fmt"
	save "Game/Jeux/Sauvegarde"
	//game "Game/Jeux/GamePlay"
	Combat "Game/Jeux/GamePlay/Fight"
	chevalierJedi "Game/Jeux/GamePlay/FirstMission/Jedi/ChevalierJedi"
	consulaireJedi "Game/Jeux/GamePlay/FirstMission/Jedi/ConsulaireJedi"
	assassinSith "Game/Jeux/GamePlay/FirstMission/Sith/AssassinSith"
	guerrierSith "Game/Jeux/GamePlay/FirstMission/Sith/GuerrierSith"
	marchand "Game/Jeux/GamePlay/Marchand"
)

func main() {
	//objet1 := save.ItemPacket{save.Sabre_laser, 1}
	//objet2 := save.ItemPacket{save.Casque, 2}
	save.CreatePerso()
	/*save.Ajout_Item(objet1)
	save.Ajout_Item(objet1)
	save.Ajout_Item(objet2)
	save.RegardeStat()*/
	switch save.Campagne.Name {
	case "Jedi Knight":
		chevalierJedi.JediKnight()
	case "Jedi Consular":
		consulaireJedi.JediConsular()
	case "Sith Warrior":
		guerrierSith.SithWarrior()
		save.Personnage.CoteForce += guerrierSith.Arrive_2()
		guerrierSith.Arrive_3()
		WayClasseSithWarrior()
		choice_1 := guerrierSith.Menu()
		switch choice_1 {
		case "1":
			save.DisplayCharacter()
		case "2":
			save.DisplayInventaire()
		case "3":
			marchand.Merchantchoice()
		case "4" :

		case "5" :
			if Combat.Fight(&save.Personnage, &guerrierSith.LimaceKor_Rang1, true) {
				if Combat.Fight(&save.Personnage, &guerrierSith.LimaceKor_Rang2, true) {
					if Combat.Fight(&save.Personnage, &guerrierSith.LimaceKor_Rang3, false) {
						choice_2 := guerrierSith.Menu()
						switch choice_2 {
						case "1":
							save.DisplayCharacter()
						case "2":
							save.DisplayInventaire()
						case "3":
							marchand.Merchantchoice()
						case "4" :
						
						case "5" :

						}
					}
				}
			}
		}
	case "Sith Assassin":
		assassinSith.SithAssassin()

	}

	//Forgeron2 := forge.Forgeron{"vbfjimegobes", []save.Weapon{save.Weapon{"fbezhlif", 10, 10, "red", 1, 1, 1, "ceci est une decription"}}}
	//Forgeron2.DisplayProduct()
}

func WayClasseSithWarrior() {
	var choix string
	fmt.Println("Choose your Ability class :")
	fmt.Println("(1) Tank :",guerrierSith.ClassTank,"     (2) Burst : ",guerrierSith.ClassBurst.Name,"     (3) Dot : ",guerrierSith.ClassDot)
	fmt.Scanln(&choix) 
	switch choix {
	case "1" :
		save.Personnage.Classe = guerrierSith.ClassTank
	case "2" :
		save.Personnage.Classe = guerrierSith.ClassBurst
	case "3" :
		save.Personnage.Classe = guerrierSith.ClassDot
	default :
		save.ClearScreen()
		WayClasseSithWarrior()
	}
}

/*
			DA := save.Abilite{"attaqua basic", 0, 10, 0}
			DA2 := save.Abilite{"eclai de force", 0, 100, 0}
			MC := save.Perso{"Toto", 0, 100, 100, 1, 2, 100, 0, DA, save.Personnage.Classe}
			Mob := save.Mob{"Dark Vador", 50, 50, 0, []save.Abilite{DA2}, map[save.Item]int{}, 10}
			Fight(&MC, &Mob, false)
	=======
		}
		DA := save.Abilite{"attaqua basic", 0, 10, 0}
		DA2 := save.Abilite{"eclai de force", 0, 100, 0}
		MC := save.Perso{"Toto", 0, 100, 100, 1, 2, 100, 0, DA, save.Personnage.Classe}
		Mob := save.Mob{"Dark Vador", 50, 50, 0, []save.Abilite{DA2}, map[save.Item]int{}, 10}
		Fight(&MC, &Mob, false)
	>>>>>>> 483ff21fc78e6c7864dbb022c25543abae595241
*/
