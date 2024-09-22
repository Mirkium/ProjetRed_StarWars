package jeux

import (
	"fmt"
)

var Armor_1 = Armor{
	Name:        "StormTrooper armor",
	PvBonus:     50,
	DamageBonus: 20,
	StatArmor:   5,
	Quantity:    10,
	Price:       100,
	CoteForce:   0,
	Description: "Standard StormTrooper armor",
}
var Armor_2 = Armor{
	Name:        "Mandalorian armor",
	PvBonus:     200,
	DamageBonus: 100,
	StatArmor:   100,
	Quantity:    10,
	Price:       1000,
	CoteForce:   0,
	Description: "Beskar Mandalorian armor",
}

func ArmorShop() {
	var choix string
	fmt.Println("Credit    : ", Green, Personnage.Credit)
	fmt.Println("Way Force : ", Personnage.CoteForce)
	fmt.Println("/=================ARMOR SHOP=================\\")
	fmt.Println("")
	fmt.Println(Yellow, "   1. ", Cyan, Armor_1.Name, "                     ", Cyan, Armor_1.Quantity, "/10")
	fmt.Println(Yellow, "     Damage Bonus : ", Cyan, Armor_1.DamageBonus, " | PV Bonus : ", Cyan, Armor_1.PvBonus, " | Armor : ", Cyan, Armor_1.StatArmor)
	fmt.Println(Yellow, "     Price : ", Cyan, Armor_1.Price, " | Way Force : ", Cyan, Armor_1.StatArmor)
	fmt.Println(" ____________________________________________")
	fmt.Println("")
	fmt.Println(Yellow, "   2. ", Cyan, Armor_2.Name, "                     ", Cyan, Armor_2.Quantity, "/10")
	fmt.Println(Yellow, "     Damage Bonus : ", Cyan, Armor_2.DamageBonus, " | PV Bonus : ", Cyan, Armor_2.PvBonus, " | Armor : ", Cyan, Armor_2.StatArmor)
	fmt.Println(Yellow, "     Price : ", Cyan, Armor_2.Price, " | Way Force : ", Cyan, Armor_2.StatArmor)
	fmt.Println(" ____________________________________________")
	fmt.Println("")
	fmt.Println(Yellow, "                   (0)", Cyan, " Exit", Reset)
	fmt.Println("")
	fmt.Println("\\============================================/")
	fmt.Println("")
	fmt.Print("Your choice : ")
	fmt.Scanln(&choix)
	switch choix {
	case "1":
		Armor_1.AchatArmor()
	case "2":
		Armor_2.AchatArmor()
	case "0":
		break
	default:
		clearScreen()
		ArmorShop()
	}
}

func (Armur Armor) AchatArmor() {
	if Personnage.Credit >= Armur.Price {
		Personnage.Credit -= Armur.Price
		Armur.Quantity -= 1
		Ajout_Armur(Armur, 1)
	}
}
