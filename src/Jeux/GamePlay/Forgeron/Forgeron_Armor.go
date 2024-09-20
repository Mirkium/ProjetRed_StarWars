package forgeron

import (
	save "Game/Jeux/Sauvegarde"
	"fmt"
	//"time"
)

func MenuFArmor() {
	ForgeronArmor1()
	ForgeronArmor2()
	ForgeronArmor3()
}

func ForgeronArmor1() {
	var choix string
	var Item1 save.Item
	var Item2 save.Item

	armor1 := Forgeron_Armor{
		"Mandalorian armor",
		300,
		100,
		1000,
		save.Item{
			"Beskar alloy",
			0,
			"Beskar, also called Mandalorian iron, was an alloy used in particular in the making of Mandalorian armor.",
			1,
		},
		save.Item{
			"Durasteel Alloy",
			0,
			"Durasteel was a very resistant alloy created from different ores including zersium. Stronger than steel, it was used in the manufacture of armor and protection of ships.",
			1,
		},
		0,
		"Mandalorian armor referred to the traditional armor worn by human warriors of the clans of Mandalore.",
	}

	fmt.Println("Armorsmith")
	fmt.Println("1- ", armor1.Name)
	fmt.Println("   HP Bonus : ", armor1.PvBonus, " | Damage Bonus : ", armor1.DamageBonus)
	fmt.Println("Price :")
	fmt.Println("   credit : ", save.Personnage.Credit, " | ", armor1.ItemValeur_1.Name, " | ", armor1.ItemValeur_2.Name, " | way Force : ", save.Personnage.CoteForce)
	fmt.Println("____________________________________")
	fmt.Scanln(&choix) //input qui va prendre en considération l'objet voulu
	if choix == "1" {
		for _, Item := range save.Inventaire_Item {
			if Item.Name == armor1.ItemValeur_2.Name && save.Personnage.CoteForce <= armor1.CoteForce && save.Personnage.Credit >= armor1.Valeur {
				Item2 = armor1.ItemValeur_2
				break
			}
		}
		if Item1 == armor1.ItemValeur_1 && Item2 == armor1.ItemValeur_2 {
			AchatArmor1(armor1, Item2, Item1)
		}
	}
	MenuFArmor()
}

func AchatArmor1(BuyItem Forgeron_Armor, Item2 save.Item, Item1 save.Item) {
	var choix string
	NewArmor := save.Armor{
		BuyItem.Name,
		BuyItem.PvBonus,
		BuyItem.DamageBonus,
		BuyItem.Valeur,
		BuyItem.ItemValeur_1.Name,
		1,
		1,
		BuyItem.CoteForce,
		BuyItem.Description,
	}
	fmt.Println("Are you sure to buy ", BuyItem.Name)
	fmt.Println(" (1) yes   (2) no")
	fmt.Scanln(&choix)
	switch choix {
	case "1":
		save.Personnage.Credit -= BuyItem.Valeur
		save.Enlever_Item(Item2, 1)
		save.Ajout_Armur(NewArmor, 1)
	case "2":
		save.ClearScreen()
		ForgeronWeapon1()
	default:
		save.ClearScreen()
		AchatArmor1(BuyItem, Item2, Item1)
	}
}

func ForgeronArmor2() {
	var choix string
	var Item1 save.Item
	var Item2 save.Item

	armor2 := Forgeron_Armor{
		"StormTrooper Armor",
		300,
		100,
		1000,
		save.Item{
			"Chromium",
			0,
			"Chromium was an ore found on the mining planet of Burnin Konn where it was used as currency.",
			1,
		},
		save.Item{
			"Plastic composite",
			0,
			"A composite material is a material which is made up of at least two materials of different nature.",
			1,
		},
		-50,
		"Stormtrooper armor was standard protection worn by Stormtroopers of the Galactic Empire.",
	}

	fmt.Println("Armorsmith")
	fmt.Println("1- ", armor2.Name)
	fmt.Println("   HP Bonus : ", armor2.PvBonus, " | Damage Bonus : ", armor2.DamageBonus)
	fmt.Println("Price :")
	fmt.Println("   credit : ", save.Personnage.Credit, " | ", armor2.ItemValeur_1.Name, " | ", armor2.ItemValeur_2.Name, " | way Force : ", save.Personnage.CoteForce)
	fmt.Println("____________________________________")
	fmt.Scanln(&choix) //input qui va prendre en considération l'objet voulu
	if choix == "1" {
		for _, Item := range save.Inventaire_Item {
			if Item.Name == armor2.ItemValeur_2.Name && save.Personnage.CoteForce <= armor2.CoteForce && save.Personnage.Credit >= armor2.Valeur {
				Item2 = armor2.ItemValeur_2
				break
			}
		}
		if Item1 == armor2.ItemValeur_1 && Item2 == armor2.ItemValeur_2 {
			AchatArmor1(armor2, Item2, Item1)
		}
	}
	MenuFArmor()
}

func AchatArmor2(BuyItem Forgeron_Armor, Item2 save.Item, Item1 save.Item) {
	var choix string
	NewArmor := save.Armor{
		BuyItem.Name,
		BuyItem.PvBonus,
		BuyItem.DamageBonus,
		BuyItem.Valeur,
		BuyItem.ItemValeur_1.Name,
		1,
		1,
		BuyItem.CoteForce,
		BuyItem.Description,
	}
	fmt.Println("Are you sure to buy ", BuyItem.Name)
	fmt.Println(" (1) yes   (2) no")
	fmt.Scanln(&choix)
	switch choix {
	case "1":
		save.Personnage.Credit -= BuyItem.Valeur
		save.Enlever_Item(Item2, 1)
		save.Ajout_Armur(NewArmor, 1)
	case "2":
		save.ClearScreen()
		ForgeronWeapon1()
	default:
		save.ClearScreen()
		AchatArmor1(BuyItem, Item2, Item1)
	}
}

func ForgeronArmor3() {
	var choix string
	var Item1 save.Item
	var Item2 save.Item

	armor3 := Forgeron_Armor{
		"Ubese Armor",
		100,
		50,
		300,
		save.Item{
			"leather",
			0,
			"Animal skin separated from the flesh, tanned and prepared.",
			1,
		},
		save.Item{
			"Iron",
			0,
			"a strong, hard magnetic silvery-grey metal",
			1,
		},
		-50,
		"Stormtrooper armor was standard protection worn by Stormtroopers of the Galactic Empire.",
	}

	fmt.Println("Armorsmith")
	fmt.Println("1- ", armor3.Name)
	fmt.Println("   HP Bonus : ", armor3.PvBonus, " | Damage Bonus : ", armor3.DamageBonus)
	fmt.Println("Price :")
	fmt.Println("   credit : ", save.Personnage.Credit, " | ", armor3.ItemValeur_1.Name, " | ", armor3.ItemValeur_2.Name, " | way Force : ", save.Personnage.CoteForce)
	fmt.Println("____________________________________")
	fmt.Scanln(&choix) //input qui va prendre en considération l'objet voulu
	if choix == "1" {
		for _, Item := range save.Inventaire_Item {
			if Item.Name == armor3.ItemValeur_2.Name && save.Personnage.CoteForce <= armor3.CoteForce && save.Personnage.Credit >= armor3.Valeur {
				Item2 = armor3.ItemValeur_2
				break
			}
		}
		if Item1 == armor3.ItemValeur_1 && Item2 == armor3.ItemValeur_2 {
			AchatArmor1(armor3, Item2, Item1)
		}
	}
	MenuFArmor()
}

func AchatArmor3(BuyItem Forgeron_Armor, Item2 save.Item, Item1 save.Item) {
	var choix string
	NewArmor := save.Armor{
		BuyItem.Name,
		BuyItem.PvBonus,
		BuyItem.DamageBonus,
		BuyItem.Valeur,
		BuyItem.ItemValeur_1.Name,
		1,
		1,
		BuyItem.CoteForce,
		BuyItem.Description,
	}
	fmt.Println("Are you sure to buy ", BuyItem.Name)
	fmt.Println(" (1) yes   (2) no")
	fmt.Scanln(&choix)
	switch choix {
	case "1":
		save.Personnage.Credit -= BuyItem.Valeur
		save.Enlever_Item(Item2, 1)
		save.Ajout_Armur(NewArmor, 1)
	case "2":
		save.ClearScreen()
		ForgeronWeapon1()
	default:
		save.ClearScreen()
		AchatArmor1(BuyItem, Item2, Item1)
	}
}
