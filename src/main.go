package main

import (
	save "Game/Jeux/Sauvegarde"
	//game "Game/Jeux/GamePlay"
	//marchand "Game/Jeux/GamePlay/Marchand"
	//forgeron "Game/Jeux/GamePlay/Forgeron"
	chevalierJedi "Game/Jeux/GamePlay/FirstMission/Jedi/ChevalierJedi"
	consulaireJedi "Game/Jeux/GamePlay/FirstMission/Jedi/ConsulaireJedi"
	assassinSith "Game/Jeux/GamePlay/FirstMission/Sith/AssassinSith"
	guerrierSith "Game/Jeux/GamePlay/FirstMission/Sith/GuerrierSith"
	guerriersith "Game/Jeux/GamePlay/FirstMission/Sith/GuerrierSith"
)

func main() {
	//objet1 := save.ItemPacket{save.Sabre_laser, 1}
	//objet2 := save.ItemPacket{save.Casque, 2}
	save.CreatePerso()
	/*save.Ajout_Item(objet1)
	save.Ajout_Item(objet1)
	save.Ajout_Item(objet2)
	save.RegardeStat()*/
	//marchand.Marchand(save.Personnage.Credit, save.Personnage.CoteForce)
	//marchand.MarchandArmor(save.Personnage.Credit, save.Personnage.CoteForce)
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
}
