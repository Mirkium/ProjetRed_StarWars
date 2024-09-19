package marchand

import (
	save "Game/Jeux/Sauvegarde"
	"fmt"
)

func MenuItem(u *save.Perso, m *MarchantItem) {
	MarchandItem(u.Credit, u.CoteForce)
}

func MarchandItem(Credit int, CoteForce int) {
	item01 := save.Item{
		Name:        "darth vader helmet",
		Quantite:    1,
		Valeur:      50000,
		Description: "Having been badly burned on Moustafar, Darth Vader had no choice but to wear a breathing helmet to survive."}
	item02 := save.Item{
		Name:        "darth vader helmet",
		Quantite:    1,
		Valeur:      50000,
		Description: "Having been badly burned on Moustafar, Darth Vader had no choice but to wear a breathing helmet to survive."}
	item03 := save.Item{
		Name:        "darth vader helmet",
		Quantite:    1,
		Valeur:      50000,
		Description: "Having been badly burned on Moustafar, Darth Vader had no choice but to wear a breathing helmet to survive."}
	item04 := save.Item{
		Name:        "darth vader helmet",
		Quantite:    1,
		Valeur:      50000,
		Description: "Having been badly burned on Moustafar, Darth Vader had no choice but to wear a breathing helmet to survive."}
	fmt.Println("")
	fmt.Println(item01.Name, " : ", item01.Description, item01.Quantite, item01.Valeur)
	fmt.Println(item02.Name, " : ", item01.Description, item01.Quantite, item01.Valeur)
	fmt.Println(item03.Name, " : ", item01.Description, item01.Quantite, item01.Valeur)
	fmt.Println(item04.Name, " : ", item01.Description, item01.Quantite, item01.Valeur)
	fmt.Println("")
	fmt.Println("credit : ", save.Personnage.Credit, " | way Force : ", save.Personnage.CoteForce)
	fmt.Println("")
	marchand := MarchantItem{"Merchant", []save.Item{item01, item02, item03}}
	save.Personnage.Credit = Credit
	save.Personnage.CoteForce = CoteForce
	MenuItem(&save.Personnage, &marchand)
}

func AchatItem(BuyItem save.Item) {
	var choix string
	fmt.Println("Are you sure to buy ", BuyItem.Name)
	fmt.Println(" (1) yes   (2) no")
	fmt.Scanln(&choix)
	switch choix {
	case "1":
		save.Personnage.Credit -= BuyItem.Valeur
		save.Ajout_Item(BuyItem, 1)
	case "2":
		save.ClearScreen()
		MarchandWeapon(save.Personnage.Credit, save.Personnage.CoteForce)
	default:
		save.ClearScreen()
		AchatItem(BuyItem)
	}
}
