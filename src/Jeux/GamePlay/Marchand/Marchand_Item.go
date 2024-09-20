package marchand

import (
	save "Game/Jeux/Sauvegarde"
	"fmt"
	"time"
)

func MenuItem() {
	MarchandItem()
}

func MarchandItem() {
	var choix string
	item01 := save.Item{
		Name:        "Force Detector",
		Quantite:    1,
		Valeur:      50000,
		Description: "A device used by the Galactic Empire to detect fugitive Jedi, the Force detector worked using Thaissen Crystals, collected on Mimban, which had the particularity of glowing in the presence of sentient beings in the Force."}
	item02 := save.Item{
		Name:        "Kitpack",
		Quantite:    1,
		Valeur:      50000,
		Description: "kit allowing its careful user a large quantity of HP"}
	item03 := save.Item{
		Name:        "hemorrhagic item",
		Quantite:    1,
		Valeur:      50000,
		Description: "allows the user to create a hemorrhaging effect on their opponent, removing a certain number of HP each turn"}
	item04 := save.Item{
		Name:        "Bandfill",
		Quantite:    1,
		Valeur:      50000,
		Description: "The Bandfill was a very complex and unique instrument popular during the era of the Old Republic and the Galactic Empire."}
	fmt.Println("")
	fmt.Println(item01.Name, " : ", item01.Description, item01.Quantite, item01.Valeur)
	fmt.Println(item02.Name, " : ", item01.Description, item01.Quantite, item01.Valeur)
	fmt.Println(item03.Name, " : ", item01.Description, item01.Quantite, item01.Valeur)
	fmt.Println(item04.Name, " : ", item01.Description, item01.Quantite, item01.Valeur)
	fmt.Println("")
	fmt.Println("credit : ", save.Personnage.Credit, " | way Force : ", save.Personnage.CoteForce)
	fmt.Println("")
	MenuItem()
	if choix == "1" {
		if save.Personnage.CoteForce <= -50 { //vérifie si on a assez de point Obscur
			save.ClearScreen()
			AchatItem(item01) // achète l'item01
		} else {
			fmt.Println("You don't believe enough in the Dark Side")
			time.Sleep(time.Second * 2)
			save.ClearScreen()
			MarchandItem()
		}
	} else if choix == "2" {
		if save.Personnage.CoteForce >= 50 { //vérifie si on a assez de point lumineux
			save.ClearScreen()
			AchatItem(item02)
		}
	} else if choix == "3" {
		save.ClearScreen()
		AchatItem(item03)
	} else if choix == "4" {
		save.ClearScreen()
		AchatItem(item04)
	} else if choix == "0" {
		fmt.Println("You go out")
	} else {
		save.ClearScreen()
		MarchandItem()
	}
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
		MarchandItem()
	default:
		save.ClearScreen()
		AchatItem(BuyItem)
	}
}
