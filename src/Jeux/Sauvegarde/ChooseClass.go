package Sauvegarde

import (
	"encoding/json"
	"fmt"
	"os"
)

func Republic() {
	var choix string
	fmt.Println("                                                ", Blue, "GALACTIC REPUBLIC", Reset)
	fmt.Println("                                             ")
	fmt.Println("               (1) Jedi Knight                                                   (2) Jedi Consular")
	fmt.Println("                                             ")
	fmt.Println("       Follow the voice of the Jedi Knight                               The Voice of the Jedi Consular brings         ")
	fmt.Println("       to join the Jedi Council and bring                                you knowledge of the Jedi archives        ")
	fmt.Println("       peace to the galaxy.                                              and the mysteries of the galaxy.")
	fmt.Scanln(&choix)
	switch choix {
	case "1":
		Campagne.Name = "Jedi Knight"
		Campagne.Energie = 0
		Arme.Name = "electric saber"
		Arme.DamageBonus = 5
		Arme.PvBonus = 0
		Arme.Color = "Jaune"
		Armor.Name = "Tenue de Padawan"
		Armor.DamageBonus = 0
		Armor.PvBonus = 10
		Personnage.PV_max = 1000 + Arme.PvBonus + Armor.PvBonus
		Personnage.PV_actuelle = 1000 + Arme.PvBonus + Armor.PvBonus
		Personnage.Force = 10 + Arme.DamageBonus + Armor.DamageBonus
		Personnage.Level = 1
		Personnage.CoteForce = 0
		Personnage.Credit = 0
	case "2":
		Campagne.Name = "Jedi Consular"
		Campagne.Energie = 10
		Arme.Name = "electric stick"
		Arme.DamageBonus = 5
		Arme.PvBonus = 0
		Arme.Color = "Jaune"
		Armor.Name = "Tenue de Padawan"
		Armor.DamageBonus = 0
		Armor.PvBonus = 10
		Personnage.PV_max = 1000 + Arme.PvBonus + Armor.PvBonus
		Personnage.PV_actuelle = 1000 + Arme.PvBonus + Armor.PvBonus
		Personnage.Force = 10 + Arme.DamageBonus + Armor.DamageBonus
		Personnage.Level = 1
		Personnage.CoteForce = 0
		Personnage.Credit = 0
	}
	filePath := "Save.json"

	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
		return
	}

	err = json.Unmarshal(file, &Personnage)
	if err != nil {
		fmt.Println("Erreur lors du décodage JSON:", err)
		return
	}
}

func SithEmpire() {
	var choix string
	fmt.Println("                                             ", Red, "Sith Empire")
	fmt.Println(" ")
	fmt.Println("         (1) Sith Warrior                                                 (2) Sith Assassin")
	fmt.Println(" ")
	fmt.Println(" ", Yellow, "Become the Emperor's Hand as a ", Red, "Sith Warrior", Yellow, ",                    Explore the depths of the Dark Side with ")
	fmt.Println(" to strike down the Empire's enemies.                            the ", Red, "Sith Assassin", Yellow, "'s Path, and destroy ")
	fmt.Println("                                                                 those who stand in your way.", Reset)
	fmt.Scanln(&choix)
	switch choix {
	case "1":
		Campagne.Name = "Sith Warrior"
		Campagne.Energie = 0
		Arme.Name = "electric saber"
		Arme.DamageBonus = 5
		Arme.PvBonus = 0
		Arme.Color = "Jaune"
		Armor.Name = "Tenue d'Acolyte"
		Armor.DamageBonus = 0
		Armor.PvBonus = 10
		Personnage.PV_max = 1000 + Arme.PvBonus + Armor.PvBonus
		Personnage.PV_actuelle = 1000 + Arme.PvBonus + Armor.PvBonus
		Personnage.Force = 10 + Arme.DamageBonus + Armor.DamageBonus
		Personnage.Level = 1
		Personnage.CoteForce = 0
		Personnage.Credit = 0
	case "2":
		Campagne.Name = "Sith Assassin"
		Campagne.Energie = 10
		Arme.Name = "electric stick"
		Arme.DamageBonus = 5
		Arme.PvBonus = 0
		Arme.Color = "Jaune"
		Armor.Name = "Tenue d'Acolyte"
		Armor.DamageBonus = 0
		Armor.PvBonus = 10
		Personnage.PV_max = 1000 + Arme.PvBonus + Armor.PvBonus
		Personnage.PV_actuelle = 1000 + Arme.PvBonus + Armor.PvBonus
		Personnage.Force = 10 + Arme.DamageBonus + Armor.DamageBonus
		Personnage.Level = 1
		Personnage.CoteForce = 0
		Personnage.Credit = 0
	}
	filePath := "Save.json"

	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
		return
	}
	fmt.Println(Personnage)
	err = json.Unmarshal(file, &Personnage)
	if err != nil {
		fmt.Println("Erreur lors du décodage JSON:", err)
		return
	}
}
