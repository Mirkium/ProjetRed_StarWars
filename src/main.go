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
)

func main() {
	objet1 := save.ItemPacket{save.Sabre_laser, 1}
	objet2 := save.ItemPacket{save.Casque, 2}
	save.CreatePerso()
	save.Ajout_Item(objet1)
	save.Ajout_Item(objet1)
	save.Ajout_Item(objet2)
	save.RegardeStat()
	switch save.Campagne.Name {
	case "Jedi Knight":
		chevalierJedi.JediKnight()
	case "Jedi Consular":
		consulaireJedi.JediConsular()
	case "Sith Warrior":
		guerrierSith.SithWarrior()
	case "Sith Assassin":
		assassinSith.SithAssassin()
	}
}
