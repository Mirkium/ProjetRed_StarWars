package jeux

import (
	"fmt"
)

var Ability_1 = Ability{
	Name:        "Force lightning",
	EnergieCost: 20,
	Dammage:     50,
	Heal:        0,
	Quantity:    1,
	Price:       200,
	CoteForce:   -100,
	Description: "Throw lightning force on the enemy.",
	DotCompteur: 2,
	DotDammage:  50,
}

func AbilityShop() {
	var choix string
	fmt.Println("Credit    : ", Personnage.Credit)
	fmt.Println("Way Force : ", Personnage.CoteForce)
	fmt.Println("/======================ABILITY SHOP======================\\")
	fmt.Println("")
	fmt.Println("   1. ", Ability_1.Name)
	fmt.Println("     Damage : ", Ability_1.Dammage, " | Dot Counter : ", Ability_1.DotCompteur, " | Dot Damage : ", Ability_1.DotDammage, " | Energie Cost : ", Ability_1.EnergieCost)
	fmt.Println("     Price : ", Ability_1.Price, " | Way Force : ", Ability_1.CoteForce)
	fmt.Println("   ____________________________________________________")
	fmt.Println("")
	fmt.Println("                        (0) Exit")
	fmt.Println("")
	fmt.Println("\\========================================================/")
	fmt.Println("")
	fmt.Print("Your choice : ")
	fmt.Scanln(&choix)
	switch choix {
	case "1":

	case "0":

	default:

	}
}

func (Competence Ability) AchatAbility() {
	if Personnage.Credit >= Competence.Price {
		Personnage.Credit -= Competence.Price
		Personnage.AbilitieDefault = append(Personnage.AbilitieDefault, Competence)
	} else {
		clearScreen()
		AbilityShop()
	}
}
