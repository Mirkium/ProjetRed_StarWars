package jeux

/*
import (
	
	"fmt"
	//"time"
)



func MenuFWeapon() {
	ForgeronWeapon()
	return
}

func ForgeronWeapon() {
	var choix string
	var Item1 Cristal
	var Item2 Item
	weapon1 := Forgeron_Weapon{
		"Light Saber Vador",
		500,
		100,
		666,
		1,
		Cristal{
			"Red Crystal",
			"\033[91m",
		},
		Item{
			"Saber handle",
			200,
			"It's a saber handle",
			1,
		},
		-100,
		"It's Vader's saber",
	}
	weapon2 := Forgeron_Weapon{
		"Light Saber Obiwan",
		500,
		100,
		444,
		1,
		Cristal{
			"Blue Crystal",
			"\033[91m",
		},
		Item{
			"Saer handle",
			200,
			"It's Vader's saber",
			1,
		},
		100,
		"This is Obiwan's saber",
	}
	weapon3 := Forgeron_Weapon{
		"Lightsaber count dooku",
		500,
		100,
		666,
		1,
		Cristal{
			"Red Crystal",
			"\033[91m",
		},
		Item{
			"Saber handle",
			200,
			"It's a curved saber handle",
			1,
		},
		-100,
		"This is Count Dooku's saber",
	}

	fmt.Println(Reset, "Weaponsmith")
	fmt.Println("1- ", weapon1.Name)
	fmt.Println("   HP Bonus : ", weapon1.PvBonus, " | Damage Bonus : ", weapon1.DamageBonus)
	fmt.Println("Price :")
	fmt.Println("   credit : ", Personnage.Credit, " | ", weapon1.ItemValeur_1.Name, " | ", weapon1.ItemValeur_2.Name, " | way Force : ", Personnage.CoteForce)
	fmt.Println("____________________________________")
	fmt.Println("Weaponsmith")
	fmt.Println("2- ", weapon2.Name)
	fmt.Println("   HP Bonus : ", weapon2.PvBonus, " | Damage Bonus : ", weapon2.DamageBonus)
	fmt.Println("Price :")
	fmt.Println("   credit : ", Personnage.Credit, " | ", weapon2.ItemValeur_1.Name, " | ", weapon2.ItemValeur_2.Name, " | way Force : ", Personnage.CoteForce)
	fmt.Println("____________________________________")
	fmt.Println("Weaponsmith")
	fmt.Println("3- ", weapon3.Name)
	fmt.Println("   HP Bonus : ", weapon3.PvBonus, " | Damage Bonus : ", weapon3.DamageBonus)
	fmt.Println("Price :")
	fmt.Println("   credit : ", Personnage.Credit, " | ", weapon3.ItemValeur_1.Name, " | ", weapon3.ItemValeur_2.Name, " | way Force : ", Personnage.CoteForce)
	fmt.Println("____________________________________")
	fmt.Scanln(&choix) //input qui va prendre en considération l'objet voulu
	if choix == "1" {
		for _, Item := range Inventaire_Item {
			if Item.Name == weapon1.ItemValeur_2.Name && Personnage.CoteForce <= weapon1.CoteForce && Personnage.Credit >= weapon1.Valeur {
				Item2 = weapon1.ItemValeur_2
				break
			}
		}
		for _, Item := range Inventaire_Crystal {
			if Item.Name == weapon1.ItemValeur_1.Name && Personnage.CoteForce <= weapon1.CoteForce && Personnage.Credit >= weapon1.Valeur {
				Item1 = weapon1.ItemValeur_1
				break
			}
		}
		if Item1 == weapon1.ItemValeur_1 && Item2 == weapon1.ItemValeur_2 && weapon1.Valeur <= Personnage.Credit {
			Craft(weapon1, weapon1.ItemValeur_1, weapon1.ItemValeur_2)
		}
	}

	if choix == "2" {
		for _, Item := range Inventaire_Item {
			if Item.Name == weapon2.ItemValeur_2.Name && Personnage.CoteForce <= weapon2.CoteForce && Personnage.Credit >= weapon2.Valeur {
				Item2 = weapon2.ItemValeur_2
				break
			}
		}
		for _, Item := range Inventaire_Crystal {
			if Item.Name == weapon2.ItemValeur_1.Name && Personnage.CoteForce <= weapon2.CoteForce && Personnage.Credit >= weapon2.Valeur {
				Item1 = weapon2.ItemValeur_1
				break
			}
		}
		if Item1 == weapon2.ItemValeur_1 && Item2 == weapon2.ItemValeur_2 && weapon2.Valeur <= Personnage.Credit {
			Craft(weapon2, weapon1.ItemValeur_1, weapon2.ItemValeur_2)
		}
	}

	if choix == "3" {
		for _, Item := range Inventaire_Item {
			if Item.Name == weapon3.ItemValeur_2.Name && Personnage.CoteForce <= weapon3.CoteForce && Personnage.Credit >= weapon3.Valeur {
				Item2 = weapon3.ItemValeur_2
				break
			}
		}
		for _, Item := range Inventaire_Crystal {
			if Item.Name == weapon3.ItemValeur_1.Name && Personnage.CoteForce <= weapon3.CoteForce && Personnage.Credit >= weapon3.Valeur {
				Item1 = weapon3.ItemValeur_1
				break
			}
		}
		if Item1 == weapon3.ItemValeur_1 && Item2 == weapon3.ItemValeur_2 && weapon3.Valeur <= Personnage.Credit {
			Craft(weapon1, weapon3.ItemValeur_1, weapon3.ItemValeur_2)
		}
	} else if choix == "0" {
		SmithchoiceAffichage()
	} else {
		ForgeronArmor()
		return
	}
}

func Craft(BuyItem Forgeron_Weapon, Item1 Cristal, Item2 Item) {
	var choix string
	NewWeapon := Lightsaber{
		BuyItem.Name,
		BuyItem.DamageBonus,
		BuyItem.PvBonus,
		BuyItem.Type,
		BuyItem.ItemValeur_1,
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
		Ajout_LightSaber(NewWeapon, 1)
	case "2":
		ClearScreen()
		ForgeronWeapon()
	default:
		ClearScreen()
		Craft(BuyItem, Item1, Item2)
	}
}
*/