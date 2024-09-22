package jeux

import "fmt"

func WeaponShop() {
	var choix string

	item01 := Lightsaber{
		Name:        "LaserSaber Sith",
		PvBonus:     50,
		DamageBonus: 0,
		Color:       Cristal{"Cristal rouge", "\033[91m"},
		Quantity:    5,
		Price:       45,
		CoteForce:   -50,
		Description: "melee weapon, This weapon requires training and its use is greatly improved when its user uses the Force."}
	item02 := Lightsaber{
		Name:        "Light Saber Jedi",
		PvBonus:     50,
		DamageBonus: 10,
		Color:       Cristal{"Cristal Bleue", "\033[34m"},
		Quantity:    1,
		Price:       45,
		CoteForce:   50,
		Description: " melee weapon, this weapon requires training and its use is greatly improved when its user uses the Force."}
	item03 := Weapon{
		Name:        "E-11 Blaster Rifle",
		PvBonus:     30,
		DamageBonus: 10,
		Quantity:    1,
		Price:       30,
		CoteForce:   0,
		Description: "A powerful, light and compact weapon, the E-11 was used widely throughout the galaxy for nearly a century and a half."}
	item04 := Weapon{
		Name:        "Electro-Proton Bomb",
		PvBonus:     100,
		DamageBonus: 0,
		Quantity:    3,
		Price:       20000,
		CoteForce:   0,
		Description: "Invented on behalf of the Grand Army of the Republic by Doctor Sionver Boll, the Electro-Proton Bomb released upon its explosion an electromagnetic pulse capable of destroying hundreds of Battle Droids in a few seconds.."}

	fmt.Println("/=================WEAPON SHOP=================\\")
	fmt.Println("")
	fmt.Println(" 1. ",item01.Name)
	fmt.Println("   Damage Bonus :", item01.DamageBonus," | Pv Bonus : ",item01.PvBonus)
	fmt.Println("   Price : ",item01.Price,"  |  Way Force : ", item01.CoteForce)
	fmt.Println(" ____________________________________________")
	fmt.Println("")
	fmt.Println(" 1. ",item02.Name)
	fmt.Println("   Damage Bonus :", item02.DamageBonus," | Pv Bonus : ",item02.PvBonus)
	fmt.Println("   Price : ",item02.Price,"  |  Way Force : ", item02.CoteForce)
	fmt.Println(" ____________________________________________")
	fmt.Println("")
	fmt.Println(" 1. ",item03.Name)
	fmt.Println("   Damage Bonus :", item03.DamageBonus," | Pv Bonus : ",item03.PvBonus)
	fmt.Println("   Price : ",item03.Price,"  |  Way Force : ", item03.CoteForce)
	fmt.Println(" ____________________________________________")
	fmt.Println("")
	fmt.Println(" 1. ",item04.Name)
	fmt.Println("   Damage Bonus :", item04.DamageBonus," | Pv Bonus : ",item04.PvBonus)
	fmt.Println("   Price : ",item04.Price,"  |  Way Force : ", item04.CoteForce)
	fmt.Println(" ____________________________________________")
	fmt.Println("")
	fmt.Println("                 (0) Exit")
	fmt.Println("")
	fmt.Println("\\=============================================/")
	fmt.Println("")
	fmt.Print("Your choice : ")
	fmt.Scanln(&choix)
	switch choix {
	case "1" :
		if Personnage.CoteForce <= item01.CoteForce {
			clearScreen()
			item01.BuyLightSaber()
		}
	case "2" :
		if Personnage.CoteForce <= item02.CoteForce {
			clearScreen()
			item02.BuyLightSaber()
		}

	case "3" :
		if Personnage.CoteForce <= item03.CoteForce {
			clearScreen()
			item03.BuyWeapon()
		}
	case "4" :
		if Personnage.CoteForce <= item04.CoteForce {
			clearScreen()
			item04.BuyWeapon()
		}
	case "0" :
		break
	default :
		clearScreen()
		WeaponShop()
	}
}

func (arme Weapon) BuyWeapon() {
	if Personnage.Credit >= arme.Price {
		Personnage.Credit -= arme.Price
		Ajout_Weapon(arme, 1)
	}
}

func (sabreLaser Lightsaber) BuyLightSaber() {
	if Personnage.Credit >= sabreLaser.Price {
		Personnage.Credit -= sabreLaser.Price
		Ajout_LightSaber(sabreLaser, 1)
	}
}
