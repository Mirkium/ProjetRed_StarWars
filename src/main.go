package main

import (
	game "Game/Jeux/GamePlay/FirstMission/Jedi/ChevalierJedi"
	save "Game/Jeux/Sauvegarde"
	//marchand "Game/Jeux/GamePlay/Marchand"
	//forgeron "Game/Jeux/GamePlay/Forgeron"
)

func main() {
	//objet1 := save.ItemPacket{save.Sabre_laser, 1}
	//objet2 := save.ItemPacket{save.Casque, 2}
	//save.CreatePerso()
	/*save.Ajout_Item(objet1)
	save.Ajout_Item(objet1)
	save.Ajout_Item(objet2)
	save.RegardeStat()*/
	//marchand.Marchand(save.Personnage.Credit, save.Personnage.CoteForce)
	//marchand.MarchandArmor(save.Personnage.Credit, save.Personnage.CoteForce)
	/*
		switch save.Campagne.Name {
		case "Jedi Knight" :
			chevalierJedi.JediKnight()
		case "Jedi Consular" :
			consulaireJedi.JediConsular()
		case "Sith Warrior" :
			guerrierSith.SithWarrior()
			save.Personnage.CoteForce += guerrierSith.Arrive_2()
			guerrierSith.Arrive_3()
			choice_1 := guerriersith.Menu()
			switch choice_1 {
			case 1 :
				save.DisplayCharacter()
			case 2 :
				save.DisplayInventaire()
			case 3 :

			}
		case "Sith Assassin" :
			assassinSith.SithAssassin()

		}
	*/
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
	DA := save.Abilite{"attaqua basic", 0, 10, 0, 1, 0, 2, ""}
	DA2 := save.Abilite{"eclai de force", 0, 100, 0, 1, 0, 2, ""}
	C := save.Classe{"Force", 100, []save.Abilite{DA2}}
	MC := save.Perso{"Toto", 0, 100, 100, 1, 2, 100, 0, DA, []save.Weapon{}, []save.Weapon{}, []save.Armor{}, C}
	game.Quete1(&MC)
}
