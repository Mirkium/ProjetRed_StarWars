package Sauvegarde

import (
	"fmt"
)

func DisplayCharacter() {
	var choix string
	fmt.Println("/======================================\\")
	fmt.Println(Cyan,"           ", Personnage.Name,Yellow, " Lv :",Cyan, Personnage.Level)
	fmt.Println("")
	fmt.Println(Yellow,"  Class : ",Cyan, Personnage.Classe.Name)
	fmt.Println("")
	fmt.Println(Yellow,"        PV        : ",Cyan, Personnage.PV_actuelle, "/", Personnage.PV_max)
	fmt.Println(Yellow,"        Force     : ",Cyan, Personnage.Force)
	fmt.Println(Yellow,"        Way Force : ",Cyan, Personnage.CoteForce)
	fmt.Println("     _________________________________")
	fmt.Println(Yellow,"             (1)",Cyan," Abiliti")
	fmt.Println("")
	fmt.Println(Yellow,"             (2)",Cyan," Weapon")
	fmt.Println("")
	fmt.Println(Yellow,"             (3)",Cyan," Armor")
	fmt.Println("")
	fmt.Println(Yellow,"             (0)",Cyan," Exit",Reset)
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
	case "0" :
		break
	default:
		ClearScreen()
		DisplayCharacter()
	}
}

func MenuAbilitie() {
	var choix int
	fmt.Println("/======================================\\")
	fmt.Println(Cyan,"                Abiliti")
	for i, k := range Personnage.AbilitieDefault {
		fmt.Println("/= ", k.Name)
		fmt.Println(i, "  ", k.EnergieCost)
		fmt.Println("\\=")
	}
	fmt.Println(Yellow,"              (0)",Cyan," Exit",Reset)
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
	fmt.Println(Cyan,"                Weapon")
	fmt.Println(Yellow,"   Name         : ",Cyan, Arme.Name)
	fmt.Println(Yellow,"   Pv Bonus     : ",Cyan, Arme.PvBonus)
	fmt.Println(Yellow,"   Damage Bonus : ",Cyan, Arme.DamageBonus)
	fmt.Println("")
	if Personnage.IsDoubleBlade {
		fmt.Println(Arme.Color, "<===========================================",gray,"]#####[]#####[",Arme.Color,"===========================================>", Reset)
	} else if Personnage.IsSecondaryWeapon {
		fmt.Println(gray,"[]####[",Arme.Color,"===========================================>", Reset)
		fmt.Println("")
		fmt.Println("       Secondary Weapon")
		fmt.Println(Yellow,"   Name         : ",Cyan, ArmeSecondaire.Name)
		fmt.Println(Yellow,"   Pv Bonus     : ",Cyan, ArmeSecondaire.PvBonus)
		fmt.Println(Yellow,"   Damage Bonus : ",Cyan, ArmeSecondaire.DamageBonus)
		fmt.Println("")
		fmt.Println(ArmeSecondaire.Color, "<===========================================",gray,"]####[]",Reset)
	} else {
		fmt.Println(gray,"[]####[",Arme.Color,"===========================================>", Reset)
	}
	fmt.Println("")
	fmt.Println(Yellow,"               (1)",Cyan," Modified Armor")
	fmt.Println("")
	fmt.Println(Yellow,"               (0)",Cyan,"     Exit",Reset)
	fmt.Println("")
	fmt.Println("\\======================================/")
	fmt.Println("")
	fmt.Print("Your choice : ")
	fmt.Scanln(&Exit)
	switch Exit {
	case "1":
		ClearScreen()
		ModdiefWeapon()
	case "0":
		ClearScreen()
		DisplayCharacter()
	default:
		ClearScreen()
		MenuWeapon()
	}
}

func ModdiefWeapon() {
	var choix string
	fmt.Println("/======================================\\")
	fmt.Println("                Weapon")
	if Personnage.IsDoubleBlade {
		fmt.Println(Arme.Color, "<===========================================", Reset, "]#####[]#####[", Arme.Color, "===========================================>", Reset)
		fmt.Println("")
		fmt.Println(Yellow,"          1.",Cyan," Change Weapon")
		fmt.Println("")
		fmt.Println(Yellow,"          2.",Cyan," Change Cristal")
		fmt.Println("")
		fmt.Println(Yellow,"          0.",Cyan," Exit",Reset)
		fmt.Println("\\======================================/")
		fmt.Println("")
		fmt.Print("Your choice : ")
		fmt.Scanln(&choix)
		switch choix {
		case "1":
			ClearScreen()
			ChangeWeapon()
		case "2":
			ClearScreen()
			ChangeCristal()
		case "0":
			ClearScreen()
			MenuWeapon()
		default:
			ClearScreen()
			ModdiefWeapon()
		}
	} else if Personnage.IsSecondaryWeapon {
		fmt.Println("Primary Weapon :")
		fmt.Println("[]####[", Arme.Color, "===========================================>", Reset)
		fmt.Println("Secondary Weapon :")
		fmt.Println(ArmeSecondaire.Color, "<===========================================", Reset, "]####[]")
		fmt.Println("")
		fmt.Println(Yellow,"          1.",Cyan," Change primary weapon")
		fmt.Println("")
		fmt.Println(Yellow,"          2.",Cyan," Change primary weapon cristal")
		fmt.Println("")
		fmt.Println(Yellow,"          3.",Cyan," Change secondary weapon ")
		fmt.Println("")
		fmt.Println(Yellow,"          4.",Cyan," Change secondary weapon cristal")
		fmt.Println("")
		fmt.Println(Yellow,"          0.",Cyan," Exit",Reset)
		fmt.Println("\\======================================/")
		fmt.Println("")
		fmt.Print("Your choice : ")
		fmt.Scanln(&choix)
		switch choix {
		case "1":
			ClearScreen()
			ChangeWeapon()
		case "2":
			ClearScreen()
			ChangeCristal()
		case "3":
			ClearScreen()
			ChangeSecondaryWeapon()
		case "4":
			ClearScreen()
			ChangeSecondaryWeaponCristal()
		case "0":
			ClearScreen()
			MenuWeapon()
		default:
			ClearScreen()
			ModdiefWeapon()
		}
	} else {
		fmt.Println("[]####[", Arme.Color, "===========================================>", Reset)
		fmt.Println("")
		fmt.Println(Yellow,"          1.",Cyan," Change Weapon")
		fmt.Println("")
		fmt.Println(Yellow,"          2.",Cyan," Change Cristal")
		fmt.Println("")
		fmt.Println(Yellow,"          0.",Cyan," Exit",Reset)
		fmt.Println("\\======================================/")
		fmt.Println("")
		fmt.Print("Your choice : ")
		fmt.Scanln(&choix)
		switch choix {
		case "1":
			ClearScreen()
			ChangeWeapon()
		case "2":
			ClearScreen()
			ChangeCristal()
		case "0":
			ClearScreen()
			MenuWeapon()
		default:
			ClearScreen()
			ModdiefWeapon()
		}
	}
}

func ChangeWeapon() {
	fmt.Println("/======================================\\")
	fmt.Println("          Weapon in inventory")
	
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
	fmt.Println(Cyan,"                Armor")
	fmt.Println(Yellow,"   Name         : ",Cyan, Armors.Name)
	fmt.Println(Yellow,"   Pv Bonus     : ",Cyan, Armors.PvBonus)
	fmt.Println(Yellow,"   Damage Bonus : ",Cyan, Armors.DamageBonus)
	fmt.Println(Yellow,"   Armor        : ",Cyan, Armors.StatArmor)
	fmt.Println(Cyan,"  -----------------------------------")
	fmt.Println(Yellow,"          (1)",Cyan," Change armor")
	fmt.Println("")
	fmt.Println(Yellow,"          (0)",Cyan,"    Exit")
	fmt.Println("")
	fmt.Println("\\======================================/")
	fmt.Println("")
	fmt.Print("Your choice : ")
	fmt.Scanln(&Choix)
	switch Choix {
	case "1":
		ClearScreen()
		ChangeArmor()
	default:
		ClearScreen()
		MenuArmor()
	}
}

func ChangeArmor() {

}
