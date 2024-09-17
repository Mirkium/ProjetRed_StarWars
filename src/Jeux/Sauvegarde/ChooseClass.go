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
		Arme.DamageBonus = 5
		Arme.PvBonus = 0
		Arme.Color = "Jaune"
		Armors.Name = "Tenue de Padawan"
		Armors.DamageBonus = 0
		Armors.PvBonus = 10
		Personnage.PV_max = 1000 + Arme.PvBonus + Armors.PvBonus
		Personnage.PV_actuelle = 1000 + Arme.PvBonus + Armors.PvBonus
		Personnage.Force = 10 + Arme.DamageBonus + Armors.DamageBonus
		Personnage.Level = 1
		Personnage.CoteForce = 0
		Personnage.Credit = 0
	case "2":
		Campagne.Name = "Jedi Consular"
		Campagne.Energie = 10
		Arme.DamageBonus = 5
		Arme.PvBonus = 0
		Arme.Color = "Jaune"
		Armors.Name = "Tenue de Padawan"
		Armors.DamageBonus = 0
		Armors.PvBonus = 10
		Personnage.PV_max = 1000 + Arme.PvBonus + Armors.PvBonus
		Personnage.PV_actuelle = 1000 + Arme.PvBonus + Armors.PvBonus
		Personnage.Force = 10 + Arme.DamageBonus + Armors.DamageBonus
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
		Armors.Name = "Tenue d'Acolyte"
		Armors.DamageBonus = 0
		Armors.PvBonus = 10

		Personnage.Level = 1
		Personnage.CoteForce = 0
		Personnage.Credit = 0
	case "2":
		Campagne.Name = "Sith Assassin"
		Campagne.Energie = 10
		Armors.Name = "Tenue d'Acolyte"
		Armors.DamageBonus = 0
		Armors.PvBonus = 10
		Personnage.Level = 1
		Personnage.CoteForce = 0
		Personnage.Credit = 0
	}
	ClearScreen()
	var choix_weapon string
	fmt.Println("Choose your class weapon :")
	fmt.Println("")
	fmt.Println("(1) One blade :")
	fmt.Println(" []####[===========================================>")
	fmt.Println("")
	fmt.Println("(2) Two blade : ")
	fmt.Println("[]####[===========================================>")
	fmt.Println("<===========================================]####[]")
	fmt.Println("")
	fmt.Println("(3) Double blade : ")
	fmt.Println("<===========================================]#####[]#####[===========================================>")
	fmt.Scanln(&choix_weapon)
	switch choix_weapon {
	case "1":
		Arme.Name = "Trainning blade"
		Arme.DamageBonus = 10
		Arme.PvBonus = 0
		Arme.Color = "jaune"
		Personnage.PV_max = 1000 + Arme.PvBonus + Armors.PvBonus
		Personnage.PV_actuelle = 1000 + Arme.PvBonus + Armors.PvBonus
		Personnage.Force = 10 + Arme.DamageBonus + Armors.DamageBonus
	case "2":
		Arme.Name = "Trainning blade"
		Arme.DamageBonus = 10
		Arme.PvBonus = 0
		Arme.Color = "jaune"
		ArmeSecondaire.Name = "Trainning blade"
		ArmeSecondaire.DamageBonus = 10
		ArmeSecondaire.PvBonus = 0
		ArmeSecondaire.Color = "jaune"
		Personnage.PV_max = 1000 + Arme.PvBonus + Armors.PvBonus + ArmeSecondaire.PvBonus
		Personnage.PV_actuelle = 1000 + Arme.PvBonus + Armors.PvBonus + ArmeSecondaire.PvBonus
		Personnage.Force = 10 + Arme.DamageBonus + Armors.DamageBonus + ArmeSecondaire.DamageBonus
	case "3":

		Arme.Name = "electric training stick"
		Arme.DamageBonus = 10
		Arme.PvBonus = 0
		Arme.Color = "jaune"
		Personnage.PV_max = 1000 + Arme.PvBonus + Armors.PvBonus
		Personnage.PV_actuelle = 1000 + Arme.PvBonus + Armors.PvBonus
		Personnage.Force = 10 + Arme.DamageBonus + Armors.DamageBonus
	}
}
