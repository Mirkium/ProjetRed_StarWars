package marchand

import (
	save "Game/Jeux/Sauvegarde"
	"fmt"
	"time"
)

func MenuArmor() {
	MarchandArmor()
}

func MarchandArmor() {
	var choix string
	item01 := save.Armor{
		Name:        "Dark Vader Helmet",
		PvBonus:     10000,
		DamageBonus: 10000,
		Quantity:    1,
		Price:       100000,
		CoteForce:   -10000000,
		Description: "Having been badly burned on Moustafar, Darth Vader had no choice but to wear a breathing helmet to survive."}
	item02 := save.Armor{
		Name:        "Jedi Battle Armor",
		PvBonus:     20000,
		DamageBonus: 1000,
		Quantity:    1,
		Price:       1000000,
		CoteForce:   10000000,
		Description: " As surprising as it may seem, the Jedi had not always worn their symbolic togas and robes. There was a time when the dark times forced the Jedi to don more adequate protections such as true battle armor."}
	item03 := save.Armor{
		Name:        "Gungan Personal Energy Shield",
		PvBonus:     500,
		DamageBonus: 0,
		Quantity:    3,
		Price:       50000,
		CoteForce:   0,
		Description: "Produced during the Galactic Republic era by the Otoh Gunga Defense League on Naboo, the Gungan Personal Energy Shield was a means of protection used by the soldiers and scouts of General Tobler Ceel's Grand Army."}
	item04 := save.Armor{
		Name:        "Coruscant Guard Armor",
		PvBonus:     1000,
		DamageBonus: 0,
		Quantity:    2,
		Price:       20000,
		CoteForce:   0,
		Description: "The Coruscant Guard Armor equipped the elite Stormtrooper unit responsible for policing Coruscant and the Core Worlds."}
	fmt.Println("Merchant")
	fmt.Println("")
	fmt.Println("1- ", item01.Name, " : ", item01.Description, item01.PvBonus, item01.DamageBonus)
	fmt.Println("2- ", item02.Name, " : ", item02.Description, item02.PvBonus, item02.DamageBonus)
	fmt.Println("3- ", item03.Name, " : ", item03.Description, item03.PvBonus, item03.DamageBonus)
	fmt.Println("4- ", item04.Name, " : ", item04.Description, item04.PvBonus, item04.DamageBonus)
	fmt.Println("")
	fmt.Println("credit : ", save.Personnage.Credit, " | way Force : ", save.Personnage.CoteForce)
	fmt.Println("")
	fmt.Scanln(&choix) //input qui va prendre en considération l'objet voulu
	if choix == "1" {
		if save.Personnage.CoteForce <= -10000000 { //vérifie si on a assez de point Obscur
			save.ClearScreen()
			AchatArmor(item01) // achète l'item01
		} else {
			fmt.Println("You don't believe enough in the Dark Side")
			time.Sleep(time.Second * 2)
			save.ClearScreen()
			MarchandArmor()
		}
	} else if choix == "2" {
		if save.Personnage.CoteForce >= 10000000 { //vérifie si on a assez de point lumineux
			save.ClearScreen()
			AchatArmor(item02)
		}
	} else if choix == "3" {
		save.ClearScreen()
		AchatArmor(item03)
	} else if choix == "4" {
		save.ClearScreen()
		AchatArmor(item04)
	} else if choix == "0" {
		fmt.Println("You go out")
	} else {
		save.ClearScreen()
		MarchandArmor()
	}
	MenuArmor()
}

func AchatArmor(BuyItem save.Armor) {
	var choix string
	fmt.Println("Are you sure to buy ", BuyItem.Name)
	fmt.Println(" (1) yes   (2) no")
	fmt.Scanln(&choix)
	switch choix {
	case "1":
		save.Personnage.Credit -= BuyItem.Price
		save.Ajout_Armur(BuyItem, 1)
	case "2":
		save.ClearScreen()
		MarchandArmor()
	default:
		save.ClearScreen()
		AchatArmor(BuyItem)
	}
}
