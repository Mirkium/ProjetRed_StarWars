package Sauvegarde

import (
	"fmt"
	"strings"
)

func DisplayCharacter() {
	var choix string
	fmt.Println("/======================================\\")
	fmt.Println("           ",Personnage.Name," Lv :", Personnage.Level)
	fmt.Println("")
	fmt.Println("  Class : ",Personnage.Classe)
	fmt.Println("")
	fmt.Println("        PV    : ",Personnage.PV_actuelle,"/",Personnage.PV_max)
	fmt.Println("        Force : ",Personnage.Force)
	fmt.Println("     _________________________________")
	fmt.Println("             (1) Abiliti")
	fmt.Println("")
	fmt.Println("             (2) Weapon")
	fmt.Println("")
	fmt.Println("             (3) Armor")
	fmt.Println("")
	fmt.Println("             (0) Exit")
	fmt.Println("\\======================================/")
	fmt.Println("")
	fmt.Print("Your choice : ")
	fmt.Scanln(&choix)
	switch choix {
	case "1" :
		MenuAbilitie()
	case "2" :
		MenuWeapon()
	case "3" :
		MenuArmor()
	default :
		break
	}
}

func MenuAbilitie() {
	var choix int
	fmt.Println("/======================================\\")
	fmt.Println("                Abiliti")
	for i,k := range Personnage.AbilitieDefault {
		fmt.Println("/= ",k.Name)
		fmt.Println(i,"  ", k.EnergieCost)
		fmt.Println("\\=")
	}
	fmt.Println("              (0) Exit")
	fmt.Println("\\======================================/")
	fmt.Println(" ")
	fmt.Print("Your choice : ")
	fmt.Scanln(&choix)
	for elem,L := range Personnage.AbilitieDefault {
		if choix == elem {
			ClearScreen()
			var Exit string
			fmt.Println("==========================Abiliti============================")
			fmt.Println("|--------------------------Name-----------------------------|")
			fmt.Println("|                                                           |")
			fmt.Println("|                    " + Formatage(L.Name, 16) + "                     |")
			fmt.Println("|                                                           |")
			fmt.Println("|-----------------------Description-------------------------|")
			fmt.Println("|                                                           |")
			fmt.Println("|   " + L.Description)
			fmt.Println("|                                                           |")
			fmt.Println("|------------------------Damage/Heal------------------------|")
			fmt.Println("|                                                           |")
			fmt.Println("|            Heal   : " + formatagePrix(L.Heal) + "                    |")
			fmt.Println("|            Damage : " + formatagePrix(L.Dammage) + "                    |")
			fmt.Println("|                                                           |")
			fmt.Println("|                        (0) Exit                           |")
			fmt.Println("=============================================================")
			fmt.Scanln(&Exit)
			switch Exit {
			default :
				ClearScreen()
				MenuAbilitie()
			}
		}
	}
}

func MenuWeapon() {
	var Exit string
	fmt.Println("/======================================\\")
	fmt.Println("                Weapon")
	fmt.Println("   Name         : ", Arme.Name)
	fmt.Println("   Pv Bonus     : ", Arme.PvBonus)
	fmt.Println("   Damage Bonus : ", Arme.DamageBonus)
	fmt.Println("")
	if strings.Contains(Arme.Name, "Double") || strings.Contains(Arme.Name, "Stick") {
		fmt.Println(Arme.Color,"<===========================================",Reset,"]#####[]#####[",Arme.Color,"===========================================>",Reset)
	} else if Personnage.IsSecondaryWeapon {
		fmt.Println("[]####[",Arme.Color,"===========================================>",Reset)
		fmt.Println("")
		fmt.Println("       Secondary Weapon")
		fmt.Println("   Name         : ", ArmeSecondaire.Name)
		fmt.Println("   Pv Bonus     : ", ArmeSecondaire.PvBonus)
		fmt.Println("   Damage Bonus : ", ArmeSecondaire.DamageBonus)
		fmt.Println("")
		fmt.Println(ArmeSecondaire.Color,"<===========================================",Reset,"]####[]")
	} else {
		fmt.Println("[]####[",Arme.Color,"===========================================>",Reset)
	}
	fmt.Println("")
	fmt.Println("               (0) Exit")
	fmt.Println("")
	fmt.Println("\\======================================/")
	fmt.Println("")
	fmt.Print("Your choice : ")
	fmt.Scanln(&Exit)
	ClearScreen()
	DisplayCharacter()
}

func MenuArmor() {

	fmt.Println("/======================================\\")
	fmt.Println("                 Armor")
	fmt.Println("   Name         : ", Armors.Name)
	fmt.Println("   Pv Bonus     : ", Armors.PvBonus)
	fmt.Println("   Damage Bonus : ", Armors.DamageBonus)
	fmt.Println("  -----------------------------------")
	fmt.Println("")
}
