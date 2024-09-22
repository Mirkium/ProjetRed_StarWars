package jeux

import (
	"fmt"
)

var Weapon_1 = Forgeron_Weapon{
	Name:         "Vador Lightsaber",
	DamageBonus:  500,
	PvBonus:      500,
	Price:        500,
	Type:         1,
	Color:        Cristal{"Red Crystal", Red, 200},
	ItemValeur_1: Item{"Sith Holocron", 1, 300, "A sith holocron"},
	ItemValeur_2: Lightsaber{Name: "LightSaber Jedi", PvBonus: 50, DamageBonus: 10, Type: 1, Color: Cristal{"Cristal Bleue", Cyan, 66}, Quantity: 1, Price: 66, CoteForce: 50, Description: " melee weapon, this weapon requires training and its use is greatly improved when its user uses the Force."},
	CoteForce:    -100,
	Description:  "Vador lightsaber",
}
var Weapon_2 = Forgeron_Weapon{
	Name:         "Vador Lightsaber",
	DamageBonus:  500,
	PvBonus:      500,
	Price:        500,
	Type:         1,
	Color:        Cristal{"Red Crystal", Red, 200},
	ItemValeur_1: Item{"Sith Holocron", 1, 300, "A sith holocron"},
	ItemValeur_2: Lightsaber{
		Name:        "LightSaber Sith",
		PvBonus:     50,
		DamageBonus: 0,
		Type:        1,
		Color:       Cristal{"Cristal rouge", Red, 666},
		Quantity:    5,
		Price:       666,
		CoteForce:   -50,
		Description: "melee weapon, This weapon requires training and its use is greatly improved when its user uses the Force.",
	},
	CoteForce:   -100,
	Description: "Vador lightsaber",
}

func WeaponForge() {
	var choix string
	var Item1 Item
	var Item2 Lightsaber
	fmt.Println("Credit    : ", Personnage.Credit)
	fmt.Println("Way Force : ", Personnage.CoteForce)
	fmt.Println("/=================WEAPON FORGE=================\\")
	fmt.Println("")
	fmt.Println(Yellow, " 1. ", Cyan, Weapon_1.Name)
	fmt.Println(gray, "   []####[", Weapon_1.Color.Color, "===========================================>", Reset)
	fmt.Println(Yellow, "   Damage Bonus :", Weapon_1.DamageBonus, " |", Yellow, " Pv Bonus : ", Cyan, Weapon_1.PvBonus)
	fmt.Println(Yellow, "   Price : ", Green, Weapon_1.Price, "  |", Yellow, "  Way Force : ", Cyan, Weapon_1.CoteForce, Reset, " | ", Cyan, Weapon_1.ItemValeur_1.Name, Reset, " | ", Cyan, Weapon_2.ItemValeur_2.Name, Reset)
	fmt.Println(" ____________________________________________")
	fmt.Println("")
	fmt.Println(Yellow, " 1. ", Cyan, Weapon_2.Name)
	fmt.Println(Weapon_2.Color.Color, "<===========================================", Reset, gray, "]#####[]#####[", Weapon_2.Color.Color, "===========================================>", Reset)
	fmt.Println(Yellow, "   Damage Bonus :", Weapon_2.DamageBonus, " |", Yellow, " Pv Bonus : ", Cyan, Weapon_2.PvBonus)
	fmt.Println(Yellow, "   Price : ", Green, Weapon_2.Price, "  |", Yellow, "  Way Force : ", Cyan, Weapon_2.CoteForce, Reset, " | ", Cyan, Weapon_2.ItemValeur_1.Name, Reset, " | ", Cyan, Weapon_2.ItemValeur_2.Name, Reset)
	fmt.Println(" ____________________________________________")
	fmt.Println("")
	fmt.Println(Yellow, "                 (0)", Cyan, " Exit", Reset)
	fmt.Println("")
	fmt.Println("\\=============================================/")
	fmt.Println("")
	fmt.Print("Your choice : ")
	fmt.Scanln(&choix)
	switch choix {
	case "1":
		for _, K := range Inventaire_Item {
			if K == Weapon_1.ItemValeur_1 {
				Item1 = K
			}
		}
		for _, L := range Inventaire_LightSaber {
			if L == Weapon_1.ItemValeur_2 {
				Item2 = L
			}
		}
		if Item1 == Weapon_1.ItemValeur_1 && Item2 == Weapon_1.ItemValeur_2 {
			Weapon_1.CraftWeapon(Item1, Item2)
		}
	case "2":
		for _, K := range Inventaire_Item {
			if K == Weapon_2.ItemValeur_1 {
				Item1 = K
			}
		}
		for _, L := range Inventaire_LightSaber {
			if L == Weapon_2.ItemValeur_2 {
				Item2 = L
			}
		}
		if Item1 == Weapon_2.ItemValeur_1 && Item2 == Weapon_2.ItemValeur_2 {
			Weapon_1.CraftWeapon(Item1, Item2)
		}
	case "0":
		break
	default:
		clearScreen()
		WeaponForge()
	}
}

func (arme Forgeron_Weapon) CraftWeapon(Item1 Item, Item2 Lightsaber) {
	NewSaber := Lightsaber{
		Name:        arme.Name,
		DamageBonus: arme.DamageBonus,
		PvBonus:     arme.PvBonus,
		Type:        arme.Type,
		Color:       arme.Color,
		Quantity:    1,
		Price:       arme.Price,
		CoteForce:   arme.CoteForce,
		Description: arme.Description,
	}
	if Personnage.Credit >= arme.Price {
		Enlever_Item(Item1, 1)
		Enlever_LightSaber(Item2, 1)
		Ajout_LightSaber(NewSaber, 1)
		Personnage.Credit -= arme.Price
	} else {
		clearScreen()
		WeaponForge()
	}
}
