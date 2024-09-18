package Sauvegarde

import (
	"fmt"
	"strings"
)

func DisplayCharacter() {
	var choix string
	fmt.Println("/======================================\\")
	fmt.Println("           ", Personnage.Name, " Lv :", Personnage.Level)
	fmt.Println("")
	fmt.Println("  Class : ", Personnage.Classe)
	fmt.Println("")
	fmt.Println("        PV    : ", Personnage.PV_actuelle, "/", Personnage.PV_max)
	fmt.Println("        Force : ", Personnage.Force)
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
	case "1":
		MenuAbilitie()
	case "2":
		MenuWeapon()
	case "3":
		MenuArmor()
	default:
		break
	}
}

func MenuAbilitie() {
	var choix int
	fmt.Println("/======================================\\")
	fmt.Println("                Abiliti")
	for i, k := range Personnage.AbilitieDefault {
		fmt.Println("/= ", k.Name)
		fmt.Println(i, "  ", k.EnergieCost)
		fmt.Println("\\=")
	}
	fmt.Println("              (0) Exit")
	fmt.Println("\\======================================/")
	fmt.Println(" ")
	fmt.Print("Your choice : ")
	fmt.Scanln(&choix)
	for elem, L := range Personnage.AbilitieDefault {
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
			default:
				ClearScreen()
				MenuAbilitie()
			}
		}
	}
}

//===================================================WEAPON===============================================================

func MenuWeapon() {
	var Exit string
	fmt.Println("/======================================\\")
	fmt.Println("                Weapon")
	fmt.Println("   Name         : ", Arme.Name)
	fmt.Println("   Pv Bonus     : ", Arme.PvBonus)
	fmt.Println("   Damage Bonus : ", Arme.DamageBonus)
	fmt.Println("")
	if strings.Contains(Arme.Name, "Double") || strings.Contains(Arme.Name, "Stick") {
		fmt.Println(Arme.Color, "<===========================================", Reset, "]#####[]#####[", Arme.Color, "===========================================>", Reset)
	} else if Personnage.IsSecondaryWeapon {
		fmt.Println("[]####[", Arme.Color, "===========================================>", Reset)
		fmt.Println("")
		fmt.Println("       Secondary Weapon")
		fmt.Println("   Name         : ", ArmeSecondaire.Name)
		fmt.Println("   Pv Bonus     : ", ArmeSecondaire.PvBonus)
		fmt.Println("   Damage Bonus : ", ArmeSecondaire.DamageBonus)
		fmt.Println("")
		fmt.Println(ArmeSecondaire.Color, "<===========================================", Reset, "]####[]")
	} else {
		fmt.Println("[]####[", Arme.Color, "===========================================>", Reset)
	}
	fmt.Println("")
	fmt.Println("               (1) Modified Armor")
	fmt.Println("")
	fmt.Println("               (0)     Exit")
	fmt.Println("")
	fmt.Println("\\======================================/")
	fmt.Println("")
	fmt.Print("Your choice : ")
	fmt.Scanln(&Exit)
	switch Exit {
	case "1":
		ClearScreen()
		ModdiefWeapon()
	case "0" :
		ClearScreen()
		DisplayCharacter()
	default :
		ClearScreen()
		MenuWeapon()
	}
}

func ModdiefWeapon() {
	var choix string
	fmt.Println("/======================================\\")
	fmt.Println("                Weapon")
	if strings.Contains(Arme.Name, "Double") || strings.Contains(Arme.Name, "Stick") {
		fmt.Println(Arme.Color, "<===========================================", Reset, "]#####[]#####[", Arme.Color, "===========================================>", Reset)
		fmt.Println("")
		fmt.Println("          1. Change Weapon")
		fmt.Println("")
		fmt.Println("          2. Change Cristal")
		fmt.Println("")
		fmt.Println("          0. Exit")
		fmt.Println("\\======================================/")
		fmt.Println("")
		fmt.Print("Your choice : ")
		fmt.Scanln(&choix)
		switch choix {
		case "1" :
			ClearScreen()
			ChangeWeapon()
		case "2" :
			ClearScreen()
			ChangeCristal()
		case "0" :
			ClearScreen()
			MenuWeapon()
		default :
			ClearScreen()
			ModdiefWeapon()
		}
	} else if Personnage.IsSecondaryWeapon {
		fmt.Println("Primary Weapon :")
		fmt.Println("[]####[", Arme.Color, "===========================================>", Reset)
		fmt.Println("Secondary Weapon :")
		fmt.Println(ArmeSecondaire.Color, "<===========================================", Reset, "]####[]")
		fmt.Println("")
		fmt.Println("          1. Change primary weapon")
		fmt.Println("")
		fmt.Println("          2. Change primary weapon cristal")
		fmt.Println("")
		fmt.Println("          3. Change secondary weapon ")
		fmt.Println("")
		fmt.Println("          4. Change secondary weapon cristal")
		fmt.Println("")
		fmt.Println("          0. Exit")
		fmt.Println("\\======================================/")
		fmt.Println("")
		fmt.Print("Your choice : ")
		fmt.Scanln(&choix)
		switch choix {
		case "1" :
			ClearScreen()
			ChangeWeapon()
		case "2" :
			ClearScreen()
			ChangeCristal()
		case "3" :
			ClearScreen()
			ChangeSecondaryWeapon()
		case "4" :
			ClearScreen()
			ChangeSecondaryWeaponCristal()
		case "0" :
			ClearScreen()
			MenuWeapon()
		default :
			ClearScreen()
			ModdiefWeapon()
		}
	} else {
		fmt.Println("[]####[", Arme.Color, "===========================================>", Reset)
		fmt.Println("")
		fmt.Println("          1. Change Weapon")
		fmt.Println("")
		fmt.Println("          2. Change Cristal")
		fmt.Println("")
		fmt.Println("          0. Exit")
		fmt.Println("\\======================================/")
		fmt.Println("")
		fmt.Print("Your choice : ")
		fmt.Scanln(&choix)
		switch choix {
		case "1" :
			ClearScreen()
			ChangeWeapon()
		case "2" :
			ClearScreen()
			ChangeCristal()
		case "0" :
			ClearScreen()
			MenuWeapon()
		default :
			ClearScreen()
			ModdiefWeapon()
		}
	}
}

func ChangeWeapon() {

}

func ChangeCristal() {

}

func ChangeSecondaryWeapon() {

}

func ChangeSecondaryWeaponCristal() {
	
}

//=======================================================================================================================


//===================================================ARMOR================================================================

func MenuArmor() {
	var Choix string
	fmt.Println("/======================================\\")
	fmt.Println("                Armor")
	fmt.Println("   Name         : ", Armors.Name)
	fmt.Println("   Pv Bonus     : ", Armors.PvBonus)
	fmt.Println("   Damage Bonus : ", Armors.DamageBonus)
	fmt.Println("   Armor        : ", Armors.StatArmor)
	fmt.Println("  -----------------------------------")
	fmt.Println("          (1) Change armor")
	fmt.Println("")
	fmt.Println("          (0)    Exit")
	fmt.Println("")
	fmt.Println("\\======================================/")
	fmt.Println("")
	fmt.Print("Your choice : ")
	fmt.Scanln(&Choix)
	switch Choix {
	case "1" :
		ClearScreen()
		ChangeArmor()
	default :
		ClearScreen()
		MenuArmor()
	}
}

func ChangeArmor() {

}
