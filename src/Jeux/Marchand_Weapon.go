package jeux

import "fmt"

func WeaponShop() {
	var choix string

	item01 := Lightsaber{
		Name:        "LightSaber Sith",
		PvBonus:     50,
		DamageBonus: 0,
		Type:        1,
		Color:       Cristal{"Cristal rouge", Red, 666},
		Quantity:    5,
		Price:       666,
		CoteForce:   -50,
		Description: "melee weapon, This weapon requires training and its use is greatly improved when its user uses the Force.",
	}
	item02 := Lightsaber{
		Name:        "LightSaber Jedi",
		PvBonus:     50,
		DamageBonus: 10,
		Type:        1,
		Color:       Cristal{"Cristal Bleue", Cyan, 66},
		Quantity:    1,
		Price:       66,
		CoteForce:   50,
		Description: " melee weapon, this weapon requires training and its use is greatly improved when its user uses the Force.",
	}
	item03 := Weapon{
		Name:        "E-11 Blaster Rifle",
		PvBonus:     30,
		DamageBonus: 10,
		Quantity:    1,
		Price:       30,
		CoteForce:   0,
		Description: "A powerful, light and compact weapon, the E-11 was used widely throughout the galaxy for nearly a century and a half.",
	}
	item04 := Weapon{
		Name:        "Electro-Proton Bomb",
		PvBonus:     100,
		DamageBonus: 0,
		Quantity:    3,
		Price:       20000,
		CoteForce:   0,
		Description: "Invented on behalf of the Grand Army of the Republic by Doctor Sionver Boll, the Electro-Proton Bomb released upon its explosion an electromagnetic pulse capable of destroying hundreds of Battle Droids in a few seconds..",
	}
	item05 := Cristal{
		Name:  "violet",
		Color: Violet,
		Price: 200,
	}

	fmt.Println("/=================WEAPON SHOP=================\\")
	fmt.Println("")
	fmt.Println(Yellow, " 1. ", Cyan, item01.Name)
	fmt.Println(gray, "   []####[", item01.Color.Color, "===========================================>", Reset)
	fmt.Println(Yellow, "   Damage Bonus :", item01.DamageBonus, " |", Yellow, " Pv Bonus : ", Cyan, item01.PvBonus)
	fmt.Println(Yellow, "   Price : ",Green, item01.Price, "  |", Yellow, "  Way Force : ", Cyan, item01.CoteForce)
	fmt.Println(" ____________________________________________")
	fmt.Println("")
	fmt.Println(Yellow, " 2. ", Cyan, item02.Name)
	fmt.Println(gray, "   []####[", item02.Color.Color, "===========================================>", Reset)
	fmt.Println(Yellow, "   Damage Bonus :", item02.DamageBonus, " |", Yellow, " Pv Bonus : ", Cyan, item02.PvBonus)
	fmt.Println(Yellow, "   Price : ",Green, item02.Price, "  |", Yellow, "  Way Force : ", Cyan, item02.CoteForce, Reset)
	fmt.Println(" ____________________________________________")
	fmt.Println("")
	fmt.Println(Yellow, " 3. ", Cyan, item03.Name)
	fmt.Println(Yellow, "   Damage Bonus :", Cyan, item03.DamageBonus, " |", Yellow, " Pv Bonus : ", Cyan, item03.PvBonus)
	fmt.Println(Yellow, "   Price : ", Green, item03.Price, "  |", Yellow, "  Way Force : ", Cyan, item03.CoteForce, Reset)
	fmt.Println(" ____________________________________________")
	fmt.Println("")
	fmt.Println(Yellow, " 4. ", item04.Name)
	fmt.Println(Yellow, "   Damage Bonus :", Cyan, item04.DamageBonus, " |", Yellow, " Pv Bonus : ", Cyan, item04.PvBonus)
	fmt.Println(Yellow, "   Price : ", Green, item04.Price, "  |", Yellow, "  Way Force : ", Cyan, item04.CoteForce, Reset)
	fmt.Println(" ____________________________________________")
	fmt.Println("")
	fmt.Println(Yellow, " 5. ", Cyan, item05.Name)
	fmt.Println(Yellow, "   Price : ", Green, item05.Price, Reset)
	fmt.Println(" ____________________________________________")
	fmt.Println("")
	fmt.Println(Yellow, "                 (0)", Cyan, " Exit",Reset)
	fmt.Println("")
	fmt.Println("\\=============================================/")
	fmt.Println("")
	fmt.Print("Your choice : ")
	fmt.Scanln(&choix)
	switch choix {
	case "1":
		if Personnage.CoteForce <= item01.CoteForce {
			clearScreen()
			item01.BuyLightSaber()
		}
	case "2":
		if Personnage.CoteForce <= item02.CoteForce {
			clearScreen()
			item02.BuyLightSaber()
		}

	case "3":
		if Personnage.CoteForce <= item03.CoteForce {
			clearScreen()
			item03.BuyWeapon()
		}
	case "4":
		if Personnage.CoteForce <= item04.CoteForce {
			clearScreen()
			item04.BuyWeapon()
		}
	case "5":
		ClearScreen()
		item05.BuyCristal()
	case "0":
		break
	default:
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

func (cristaux Cristal) BuyCristal() {
	if Personnage.Credit >= cristaux.Price {
		Personnage.Credit -= cristaux.Price
		Ajout_Crystal(cristaux)
	}
}
