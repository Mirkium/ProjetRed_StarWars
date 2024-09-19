package marchand

import (
	save "Game/Jeux/Sauvegarde"
	"fmt"
)

func MenuItem() {
	MarchandItem()
}

func MarchandItem() {
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
	MenuItem()
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
		MarchandWeapon()
	default:
		save.ClearScreen()
		AchatItem(BuyItem)
	}
}
