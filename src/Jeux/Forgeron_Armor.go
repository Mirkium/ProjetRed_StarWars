package jeux

/*
import (
	"fmt"
	//"time"
)

func MenuFArmor() {
	ForgeronArmor()
	return
}

func ForgeronArmor() {
	var choix string
	var Item1 Item
	var Item2 Item
	armor1 := Forgeron_Armor{
		"Mandalorian armor",
		300,
		100,
		1000,
		50,
		Item{
			"Beskar alloy",
			0,
			"Beskar, also called Mandalorian iron, was an alloy used in particular in the making of Mandalorian armor and its component parts.",
			1,
		},
		Item{
			"Durasteel Alloy",
			0,
			"Durasteel was a very resistant alloy created from different ores including zersium. Stronger than steel, it was used in the manufacture of armor and protection of ships.",
			1,
		},
		0,
		"Mandalorian armor referred to the traditional armor worn by human warriors of the clans of Mandalore.",
	}
	armor2 := Forgeron_Armor{
		"Stormtrooper armor",
		500,
		100,
		20,
		444,
		Item{
			"Plastic composite",
			0,
			"plastics reinforced with fibers, fillers, particles, powders and other matrix reinforcements to provide improved strength and/or rigidity.",
			1,
		},
		Item{
			"iron",
			0,
			"a strong, hard magnetic silvery-grey metal, the chemical element of atomic number 26, much used as a material for construction and manufacturing, especially in the form of steel.",
			1,
		},
		-50,
		"Stormtrooper armor was standard protection worn by Stormtroopers of the Galactic Empire.",
	}
	armor3 := Forgeron_Armor{
		"Ubese armor",
		500,
		100,
		40,
		400,
		Item{
			"leather",
			0,
			"Animal skin separated from the flesh, tanned and prepared.",
			1,
		},
		Item{
			"iron",
			0,
			"a strong, hard magnetic silvery-grey metal, the chemical element of atomic number 26, much used as a material for construction and manufacturing, especially in the form of steel.",
			1,
		},
		-50,
		"The Ubese armor is a distinctive armor thanks to its strange helmet that can resemble a dog's mouth, as well as the voice modulator which gives the wearer's voice a cold electronic tone.",
	}

	fmt.Println("Armorsmith")
	fmt.Println("1- ", armor1.Name)
	fmt.Println("   HP Bonus : ", armor1.PvBonus, " | Damage Bonus : ", armor1.DamageBonus)
	fmt.Println("Price :")
	fmt.Println("   credit : ", Personnage.Credit, " | ", armor1.ItemValeur_1.Name, " | ", armor1.ItemValeur_2.Name, " | way Force : ", Personnage.CoteForce)
	fmt.Println("____________________________________")
	fmt.Println("Armorsmith")
	fmt.Println("2- ", armor2.Name)
	fmt.Println("   HP Bonus : ", armor2.PvBonus, " | Damage Bonus : ", armor2.DamageBonus)
	fmt.Println("Price :")
	fmt.Println("   credit : ", Personnage.Credit, " | ", armor2.ItemValeur_1.Name, " | ", armor2.ItemValeur_2.Name, " | way Force : ", Personnage.CoteForce)
	fmt.Println("____________________________________")
	fmt.Println("Armorsmith")
	fmt.Println("3- ", armor3.Name)
	fmt.Println("   HP Bonus : ", armor3.PvBonus, " | Damage Bonus : ", armor3.DamageBonus)
	fmt.Println("Price :")
	fmt.Println("   credit : ", Personnage.Credit, " | ", armor3.ItemValeur_1.Name, " | ", armor3.ItemValeur_2.Name, " | way Force : ", Personnage.CoteForce)
	fmt.Println("____________________________________")
	fmt.Scanln(&choix) //input qui va prendre en considération l'objet voulu
	if choix == "1" {
		for _, Item := range Inventaire_Item {
			if Item.Name == armor1.ItemValeur_2.Name && Personnage.CoteForce <= armor1.CoteForce && Personnage.Credit >= armor1.Valeur {
				Item2 = armor1.ItemValeur_2
				break
			}
		}
		for _, Item := range Inventaire_Item {
			if Item.Name == armor1.ItemValeur_1.Name && Personnage.CoteForce <= armor1.CoteForce && Personnage.Credit >= armor1.Valeur {
				Item1 = armor1.ItemValeur_1
				break
			}
		}
		if Item1 == armor1.ItemValeur_1 && Item2 == armor1.ItemValeur_2 && armor1.Valeur <= Personnage.Credit {
			CraftArmor(armor1, armor1.ItemValeur_1, armor1.ItemValeur_2)
		}
	}

	if choix == "2" {
		for _, Item := range Inventaire_Item {
			if Item.Name == armor2.ItemValeur_2.Name && Personnage.CoteForce <= armor2.CoteForce && Personnage.Credit >= armor2.Valeur {
				Item2 = armor1.ItemValeur_2
				break
			}
		}
		for _, Item := range Inventaire_Item {
			if Item.Name == armor2.ItemValeur_1.Name && Personnage.CoteForce <= armor2.CoteForce && Personnage.Credit >= armor2.Valeur {
				Item1 = armor2.ItemValeur_1
				break
			}
		}
		if Item1 == armor2.ItemValeur_1 && Item2 == armor2.ItemValeur_2 && armor2.Valeur <= Personnage.Credit {
			CraftArmor(armor2, armor2.ItemValeur_1, armor2.ItemValeur_2)
		}
	}

	if choix == "3" {
		for _, Item := range Inventaire_Item {
			if Item.Name == armor3.ItemValeur_2.Name && Personnage.CoteForce <= armor3.CoteForce && Personnage.Credit >= armor3.Valeur {
				Item2 = armor3.ItemValeur_2
				break
			}
		}
		for _, Item := range Inventaire_Item {
			if Item.Name == armor3.ItemValeur_1.Name && Personnage.CoteForce <= armor3.CoteForce && Personnage.Credit >= armor3.Valeur {
				Item1 = armor3.ItemValeur_1
				break
			}
		}
		if Item1 == armor3.ItemValeur_1 && Item2 == armor3.ItemValeur_2 && armor3.Valeur <= Personnage.Credit {
			CraftArmor(armor2, armor3.ItemValeur_1, armor3.ItemValeur_2)
		}
	} else if choix == "0" {
		SmithchoiceAffichage()
	} else {
		ForgeronArmor()
		return
	}
}

/*
	func AchatWeapon1(BuyItem Forgeron_Weapon, Item2 Item, Item1 Cristal) {
		var choix string
		NewWeapon := Weapon{
			BuyItem.Name,
			BuyItem.PvBonus,
			BuyItem.DamageBonus,
			BuyItem.Type,
			BuyItem.ItemValeur_1.Color,
			1,
			BuyItem.Valeur,
			BuyItem.CoteForce,
			BuyItem.Description,
		}
		fmt.Println("Are you sure to buy ", BuyItem.Name)
		fmt.Println(" (1) yes   (2) no")
		fmt.Scanln(&choix)
		switch choix {
		case "1":
			Personnage.Credit -= BuyItem.Valeur
			Enlever_Item(Item2, 1)
			Ajout_Weapon(NewWeapon, 1)
		case "2":
			ClearScreen()
			ForgeronWeapon1()
		default:
			ClearScreen()
			AchatWeapon1(BuyItem, Item2, Item1)
		}
	}

	func ForgeronWeapon3() {
		var choix string
		var Item1 Cristal
		var Item2 Item


		fmt.Scanln(&choix) //input qui va prendre en considération l'objet voulu
		if choix == "1" {
			for _, Item := range Inventaire_Item {
				if Item.Name == weapon3.ItemValeur_2.Name && Personnage.CoteForce <= weapon3.CoteForce && Personnage.Credit >= weapon3.Valeur {
					Item2 = weapon3.ItemValeur_2
					break
				}
			}
			if Item1 == weapon3.ItemValeur_1 && Item2 == weapon3.ItemValeur_2 {
				AchatWeapon3(weapon3, Item2, Item1)
			}
		}
		MenuFWeapon()
	}

	func AchatWeapon3(BuyItem Forgeron_Weapon, Item2 Item, Item1 Cristal) {
		var choix string
		NewWeapon := Weapon{
			BuyItem.Name,
			BuyItem.DamageBonus,
			BuyItem.PvBonus,
			BuyItem.Type,
			BuyItem.ItemValeur_1.Color,
			1,
			BuyItem.Valeur,
			BuyItem.CoteForce,
			BuyItem.Description,
		}
		fmt.Println("Are you sure to buy ", BuyItem.Name)
		fmt.Println(" (1) yes   (2) no")
		fmt.Scanln(&choix)
		switch choix {
		case "1":
			Personnage.Credit -= BuyItem.Valeur
			Enlever_Item(Item2, 1)
			Ajout_Weapon(NewWeapon, 1)
		case "2":
			ClearScreen()
			ForgeronWeapon2()
		default:
			ClearScreen()
			AchatWeapon3(BuyItem, Item2, Item1)
		}
	}

func CraftArmor(BuyItem Forgeron_Armor, Item1 Item, Item2 Item) {
	var choix string
	NewArmor := Armor{
		BuyItem.Name,
		BuyItem.DamageBonus,
		BuyItem.PvBonus,
		BuyItem.StatArmor,
		1,
		BuyItem.Valeur,
		BuyItem.CoteForce,
		BuyItem.Description,
	}
	fmt.Println("Are you sure to buy ", BuyItem.Name)
	fmt.Println(" (1) yes   (2) no")
	fmt.Scanln(&choix)
	switch choix {
	case "1":
		Personnage.Credit -= BuyItem.Valeur
		Enlever_Item(Item2, 1)
		Ajout_Armur(NewArmor, 1)
	case "2":
		ClearScreen()
		ForgeronArmor()
	default:
		ClearScreen()
		CraftArmor(BuyItem, Item1, Item2)
	}
}

func ForgeronTest() {
	Sword := Item{"Sword", 100, "sedfgujhch,;u,nyubgkjvfj", 1}
	fmt.Println("épée : lingot de fer + baton de bois")
	fmt.Println("100€")
	//Tu choisis l'épée
	price := 100
	var Item1 Item
	var Item2 Item
	for _, Item := range Inventaire_Weapon {
		if Item.Name == "lingot de fer" {
			Item1 =
		}
		if Item.Name == "baton de bois" {
			Item2 =
		}
	}
	if Personnage.Credit >= price && Item1.Name == "lingot de fer" && Item2.Name == "baton de bois" {
		Enlever_Item(Item1, 1)
		Enlever_Item(Item2, 1)
		Personnage.Credit -= price
		Ajout_Item(Sword, 1)
	}
}*/
