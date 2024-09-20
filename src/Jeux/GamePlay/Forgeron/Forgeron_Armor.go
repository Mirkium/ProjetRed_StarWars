package forgeron

import (
	save "Game/Jeux/Sauvegarde"
	"fmt"
	//"time"
)

func MenuFArmor() {
	ForgeronArmor()
	return
}

func ForgeronArmor() {
	var choix string
	var Item1 save.Item
	var Item2 save.Item
	armor1 := Forgeron_Armor{
		"Mandalorian armor",
		300,
		100,
		1000,
		50,
		save.Item{
			"Beskar alloy",
			0,
			"Beskar, also called Mandalorian iron, was an alloy used in particular in the making of Mandalorian armor and its component parts.",
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
	armor2 := Forgeron_Armor{
		"Stormtrooper armor",
		500,
		100,
		20,
		444,
		save.Item{
			"Plastic composite",
			0,
			"plastics reinforced with fibers, fillers, particles, powders and other matrix reinforcements to provide improved strength and/or rigidity.",
			1,
		},
		save.Item{
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
		save.Item{
			"leather",
			0,
			"Animal skin separated from the flesh, tanned and prepared.",
			1,
		},
		save.Item{
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
	fmt.Println("   credit : ", save.Personnage.Credit, " | ", armor1.ItemValeur_1.Name, " | ", armor1.ItemValeur_2.Name, " | way Force : ", save.Personnage.CoteForce)
	fmt.Println("____________________________________")
	fmt.Println("Armorsmith")
	fmt.Println("2- ", armor2.Name)
	fmt.Println("   HP Bonus : ", armor2.PvBonus, " | Damage Bonus : ", armor2.DamageBonus)
	fmt.Println("Price :")
	fmt.Println("   credit : ", save.Personnage.Credit, " | ", armor2.ItemValeur_1.Name, " | ", armor2.ItemValeur_2.Name, " | way Force : ", save.Personnage.CoteForce)
	fmt.Println("____________________________________")
	fmt.Println("Armorsmith")
	fmt.Println("3- ", armor3.Name)
	fmt.Println("   HP Bonus : ", armor3.PvBonus, " | Damage Bonus : ", armor3.DamageBonus)
	fmt.Println("Price :")
	fmt.Println("   credit : ", save.Personnage.Credit, " | ", armor3.ItemValeur_1.Name, " | ", armor3.ItemValeur_2.Name, " | way Force : ", save.Personnage.CoteForce)
	fmt.Println("____________________________________")
	fmt.Scanln(&choix) //input qui va prendre en considération l'objet voulu
	if choix == "1" {
		for _, Item := range save.Inventaire_Item {
			if Item.Name == armor1.ItemValeur_2.Name && save.Personnage.CoteForce <= armor1.CoteForce && save.Personnage.Credit >= armor1.Valeur {
				Item2 = armor1.ItemValeur_2
				break
			}
		}
		for _, Item := range save.Inventaire_Item {
			if Item.Name == armor1.ItemValeur_1.Name && save.Personnage.CoteForce <= armor1.CoteForce && save.Personnage.Credit >= armor1.Valeur {
				Item1 = armor1.ItemValeur_1
				break
			}
		}
		if Item1 == armor1.ItemValeur_1 && Item2 == armor1.ItemValeur_2 && armor1.Valeur <= save.Personnage.Credit {
			CraftArmor(armor1, armor1.ItemValeur_1, armor1.ItemValeur_2)
		}
	}

	if choix == "2" {
		for _, Item := range save.Inventaire_Item {
			if Item.Name == armor2.ItemValeur_2.Name && save.Personnage.CoteForce <= armor2.CoteForce && save.Personnage.Credit >= armor2.Valeur {
				Item2 = armor1.ItemValeur_2
				break
			}
		}
		for _, Item := range save.Inventaire_Item {
			if Item.Name == armor2.ItemValeur_1.Name && save.Personnage.CoteForce <= armor2.CoteForce && save.Personnage.Credit >= armor2.Valeur {
				Item1 = armor2.ItemValeur_1
				break
			}
		}
		if Item1 == armor2.ItemValeur_1 && Item2 == armor2.ItemValeur_2 && armor2.Valeur <= save.Personnage.Credit {
			CraftArmor(armor2, armor2.ItemValeur_1, armor2.ItemValeur_2)
		}
	}

	if choix == "3" {
		for _, Item := range save.Inventaire_Item {
			if Item.Name == armor3.ItemValeur_2.Name && save.Personnage.CoteForce <= armor3.CoteForce && save.Personnage.Credit >= armor3.Valeur {
				Item2 = armor3.ItemValeur_2
				break
			}
		}
		for _, Item := range save.Inventaire_Item {
			if Item.Name == armor3.ItemValeur_1.Name && save.Personnage.CoteForce <= armor3.CoteForce && save.Personnage.Credit >= armor3.Valeur {
				Item1 = armor3.ItemValeur_1
				break
			}
		}
		if Item1 == armor3.ItemValeur_1 && Item2 == armor3.ItemValeur_2 && armor3.Valeur <= save.Personnage.Credit {
			CraftArmor(armor2, armor3.ItemValeur_1, armor3.ItemValeur_2)
		}

	} else {
		ForgeronArmor()
		return
	}
}

/*
	func AchatWeapon1(BuyItem Forgeron_Weapon, Item2 save.Item, Item1 save.Cristal) {
		var choix string
		NewWeapon := save.Weapon{
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
			save.Personnage.Credit -= BuyItem.Valeur
			save.Enlever_Item(Item2, 1)
			save.Ajout_Weapon(NewWeapon, 1)
		case "2":
			save.ClearScreen()
			ForgeronWeapon1()
		default:
			save.ClearScreen()
			AchatWeapon1(BuyItem, Item2, Item1)
		}
	}

	func ForgeronWeapon3() {
		var choix string
		var Item1 save.Cristal
		var Item2 save.Item


		fmt.Scanln(&choix) //input qui va prendre en considération l'objet voulu
		if choix == "1" {
			for _, Item := range save.Inventaire_Item {
				if Item.Name == weapon3.ItemValeur_2.Name && save.Personnage.CoteForce <= weapon3.CoteForce && save.Personnage.Credit >= weapon3.Valeur {
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

	func AchatWeapon3(BuyItem Forgeron_Weapon, Item2 save.Item, Item1 save.Cristal) {
		var choix string
		NewWeapon := save.Weapon{
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
			save.Personnage.Credit -= BuyItem.Valeur
			save.Enlever_Item(Item2, 1)
			save.Ajout_Weapon(NewWeapon, 1)
		case "2":
			save.ClearScreen()
			ForgeronWeapon2()
		default:
			save.ClearScreen()
			AchatWeapon3(BuyItem, Item2, Item1)
		}
	}
*/
func CraftArmor(BuyItem Forgeron_Armor, Item1 save.Item, Item2 save.Item) {
	var choix string
	NewArmor := save.Armor{
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
		save.Personnage.Credit -= BuyItem.Valeur
		save.Enlever_Item(Item2, 1)
		save.Ajout_Armur(NewArmor, 1)
	case "2":
		save.ClearScreen()
		ForgeronArmor()
	default:
		save.ClearScreen()
		CraftArmor(BuyItem, Item1, Item2)
	}
}

/*func ForgeronTest() {
	Sword := save.Item{"Sword", 100, "sedfgujhch,;u,nyubgkjvfj", 1}
	fmt.Println("épée : lingot de fer + baton de bois")
	fmt.Println("100€")
	//Tu choisis l'épée
	price := 100
	var Item1 save.Item
	var Item2 save.Item
	for _, Item := range save.Inventaire_Weapon {
		if Item.Name == "lingot de fer" {
			Item1 =
		}
		if Item.Name == "baton de bois" {
			Item2 =
		}
	}
	if save.Personnage.Credit >= price && Item1.Name == "lingot de fer" && Item2.Name == "baton de bois" {
		save.Enlever_Item(Item1, 1)
		save.Enlever_Item(Item2, 1)
		save.Personnage.Credit -= price
		save.Ajout_Item(Sword, 1)
	}
}*/
