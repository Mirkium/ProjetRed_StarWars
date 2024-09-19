package marchand

import (
	save "Game/Jeux/Sauvegarde"
	"fmt"
	"time"
)

func MenuWeapon() {
	MarchandWeapon()
}

func MarchandWeapon() {
	var choix string
	item01 := save.Weapon{
		Name:        "LaserSaber Sith",
		PvBonus:     50,
		DamageBonus: 0,
		Color:       "Red",
		Quantity:    5,
		Price:       45,
		CoteForce:   -50,
		Description: "melee weapon, This weapon requires training and its use is greatly improved when its user uses the Force."}
	item02 := save.Weapon{
		Name:        "Light Saber Jedi",
		PvBonus:     50,
		DamageBonus: 10,
		Color:       "Blue",
		Quantity:    1,
		Price:       45,
		CoteForce:   50,
		Description: " melee weapon, this weapon requires training and its use is greatly improved when its user uses the Force."}
	item03 := save.Weapon{
		Name:        "E-11 Blaster Rifle",
		PvBonus:     30,
		DamageBonus: 10,
		Color:       "Black and White",
		Quantity:    1,
		Price:       30,
		CoteForce:   0,
		Description: "A powerful, light and compact weapon, the E-11 was used widely throughout the galaxy for nearly a century and a half."}
	item04 := save.Weapon{
		Name:        "Electro-Proton Bomb",
		PvBonus:     100,
		DamageBonus: 0,
		Color:       "Purle in explosion",
		Quantity:    3,
		Price:       20000,
		CoteForce:   0,
		Description: "Invented on behalf of the Grand Army of the Republic by Doctor Sionver Boll, the Electro-Proton Bomb released upon its explosion an electromagnetic pulse capable of destroying hundreds of Battle Droids in a few seconds.."}
	fmt.Println("Marchand")
	fmt.Println("")
	fmt.Println("1- ", item01.Name, " : ", item01.Description, item01.PvBonus, item01.DamageBonus, item01.Color)
	fmt.Println("2- ", item02.Name, " : ", item02.Description, item02.PvBonus, item02.DamageBonus, item02.Color)
	fmt.Println("3- ", item03.Name, " : ", item03.Description, item03.PvBonus, item03.DamageBonus, item03.Color)
	fmt.Println("4- ", item04.Name, " : ", item04.Description, item04.PvBonus, item04.DamageBonus, item04.Color)
	fmt.Println("")
	fmt.Println("credit : ", save.Personnage.Credit, " | way Force : ", save.Personnage.CoteForce)
	fmt.Println("")
	fmt.Scanln(&choix) //input qui va prendre en considération l'objet voulu
	if choix == "1" {
		if save.Personnage.CoteForce <= -50 { //vérifie si on a assez de point Obscur
			save.ClearScreen()
			AchatWeapon(item01) // achète l'item01
		} else {
			fmt.Println("Tu ne crois pas assez au Côté Obscur")
			time.Sleep(time.Second * 2)
			save.ClearScreen()
			MarchandWeapon()
		}
	} else if choix == "2" {
		if save.Personnage.CoteForce >= 50 { //vérifie si on a assez de point lumineux
			save.ClearScreen()
			AchatWeapon(item02)
		}
	} else if choix == "3" {
		save.ClearScreen()
		AchatWeapon(item03)
	} else if choix == "4" {
		save.ClearScreen()
		AchatWeapon(item04)
	} else if choix == "0" {
		fmt.Println("Tu sors")
	} else {
		save.ClearScreen()
		MarchandWeapon()
	}
	MenuWeapon()
}

func AchatWeapon(BuyItem save.Weapon) {
	var choix string
	fmt.Println("Are you sure to buy ", BuyItem.Name)
	fmt.Println(" (1) yes   (2) no")
	fmt.Scanln(&choix)
	switch choix {
	case "1":
		save.Personnage.Credit -= BuyItem.Price
		save.Ajout_Weapon(BuyItem, 1)
	case "2":
		save.ClearScreen()
		MarchandWeapon()
	default:
		save.ClearScreen()
		AchatWeapon(BuyItem)
	}
}
