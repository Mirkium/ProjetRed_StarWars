package jeux

/*
import (
	"fmt"
	"time"
)



func MarchandAbilitie() {
	var choix string
	item01 := Abilite{
		Name:        "Absorption of life",
		EnergieCost: 200,
		Dammage:     100,
		Heal:        100,
		Quantity:    1,
		Price:       50000,
		CoteForce:   -10000,
		Description: "Life absoprion is a dark power allowing you to wampirize an opponent's life force."}
	item02 := Abilite{
		Name:        "Force Shield",
		EnergieCost: 100,
		Dammage:     0,
		Heal:        500,
		Quantity:    1,
		Price:       20000,
		CoteForce:   0,
		Description: "Power of the Force allowing its user to protect themselves from any type of aggression"}
	item03 := Abilite{
		Name:        "Sith Alchemy",
		EnergieCost: 1000,
		Dammage:     1000,
		Heal:        0,
		Quantity:    1,
		Price:       30000,
		CoteForce:   -10000,
		Description: "Sither power allowing its user to manipulate very powerful lightning."}
	item04 := Abilite{
		Name:        "Convection",
		EnergieCost: 200,
		Dammage:     200,
		Heal:        0,
		Quantity:    3,
		Price:       10000,
		CoteForce:   0,
		Description: "Convection allowed burning at a distance or on contact by concentrating the Force in the wrists"}
	fmt.Println("Merchant")
	fmt.Println("")
	fmt.Println("1- ", item01.Name, " : ", item01.Description, item01.EnergieCost, item01.Dammage, item01.Heal)
	fmt.Println("2- ", item02.Name, " : ", item02.Description, item02.EnergieCost, item02.Dammage, item02.Heal)
	fmt.Println("3- ", item03.Name, " : ", item03.Description, item03.EnergieCost, item03.Dammage, item03.Heal)
	fmt.Println("4- ", item04.Name, " : ", item04.Description, item04.EnergieCost, item04.Dammage, item04.Heal)
	fmt.Println("")
	fmt.Println("credit : ", Personnage.Credit, " | way Force : ", Personnage.CoteForce)
	fmt.Println("")
	fmt.Scanln(&choix) //input qui va prendre en considération l'objet voulu
	if choix == "1" {
		if Personnage.CoteForce <= -50 { //vérifie si on a assez de point Obscur
			ClearScreen()
			AchatAbilitie(item01) // achète l'item01
		} else {
			fmt.Println("You don't believe enough in the Dark Side")
			time.Sleep(time.Second * 2)
			ClearScreen()
			MarchandAbilitie()
		}
	} else if choix == "2" {
		ClearScreen()
		AchatAbilitie(item02)
	} else if choix == "3" {
		if Personnage.CoteForce >= 50 { //vérifie si on a assez de point lumineux
			ClearScreen()
			AchatAbilitie(item03)
		} else {

		}
	} else if choix == "4" {
		ClearScreen()
		AchatAbilitie(item04)
	} else if choix == "0" {
		Merchantchoice()
	} else {
		ClearScreen()
		MarchandAbilitie()
	}
	MenuAbilitie()
}

func AchatAbilitie(BuyItem Abilite) {
	var choix string
	fmt.Println("Are you sure to buy ", BuyItem.Name)
	fmt.Println(" (1) yes   (2) no")
	fmt.Scanln(&choix)
	switch choix {
	case "1":
		Personnage.Credit -= BuyItem.Price
		//Ajout_Abilitie(BuyItem, 1)
	case "2":
		ClearScreen()
		MarchandAbilitie()
	default:
		ClearScreen()
		AchatAbilitie(BuyItem)
	}
}
*/