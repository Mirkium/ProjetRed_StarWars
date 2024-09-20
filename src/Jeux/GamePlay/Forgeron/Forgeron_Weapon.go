package forgeron

import (
	save "Game/Jeux/Sauvegarde"
	"fmt"
	//"time"
)

func MenuFWeapon() {
	ForgeronWeapon1()
	ForgeronWeapon2()
	ForgeronWeapon3()
}

func ForgeronWeapon1() {
	var choix string
	var Item1 save.Cristal
	var Item2 save.Item

	weapon1 := Forgeron_Weapon{
		"Light Saber Vador",
		500,
		100,
		666,
		1,
		save.Cristal{
			"Red Crystal",
			"\033[91m",
		},
		save.Item{
			"Saber handle",
			200,
			"It's a saber handle",
			1,
		},
		-100,
		"It's Vader's saber",
	}

	fmt.Println("Weaponsmith")
	fmt.Println("1- ", weapon1.Name)
	fmt.Println("   HP Bonus : ", weapon1.PvBonus, " | Damage Bonus : ", weapon1.DamageBonus)
	fmt.Println("Price :")
	fmt.Println("   credit : ", save.Personnage.Credit, " | ", weapon1.ItemValeur_1.Name, " | ", weapon1.ItemValeur_2.Name, " | way Force : ", save.Personnage.CoteForce)
	fmt.Println("____________________________________")
	fmt.Scanln(&choix) //input qui va prendre en considération l'objet voulu
	if choix == "1" {
		for _, Item := range save.Inventaire_Item {
			if Item.Name == weapon1.ItemValeur_2.Name && save.Personnage.CoteForce <= weapon1.CoteForce && save.Personnage.Credit >= weapon1.Valeur {
				Item2 = weapon1.ItemValeur_2
				break
			}
		}
		if Item1 == weapon1.ItemValeur_1 && Item2 == weapon1.ItemValeur_2 {
			AchatWeapon1(weapon1, Item2, Item1)
		}
	}
	MenuFWeapon()
}

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

func ForgeronWeapon2() {
	var choix string
	var Item1 save.Cristal
	var Item2 save.Item

	weapon2 := Forgeron_Weapon{
		"Light Saber Obiwan",
		500,
		100,
		444,
		1,
		save.Cristal{
			"Blue Crystal",
			"\033[91m",
		},
		save.Item{
			"Saer handle",
			200,
			"It's Vader's saber",
			1,
		},
		100,
		"This is Obiwan's saber",
	}

	fmt.Println("Weaponsmith")
	fmt.Println("1- ", weapon2.Name)
	fmt.Println("   HP Bonus : ", weapon2.PvBonus, " | Damage Bonus : ", weapon2.DamageBonus)
	fmt.Println("Price :")
	fmt.Println("   credit : ", save.Personnage.Credit, " | ", weapon2.ItemValeur_1.Name, " | ", weapon2.ItemValeur_2.Name, " | way Force : ", save.Personnage.CoteForce)
	fmt.Println("____________________________________")
	fmt.Scanln(&choix) //input qui va prendre en considération l'objet voulu
	if choix == "1" {
		for _, Item := range save.Inventaire_Item {
			if Item.Name == weapon2.ItemValeur_2.Name && save.Personnage.CoteForce <= weapon2.CoteForce && save.Personnage.Credit >= weapon2.Valeur {
				Item2 = weapon2.ItemValeur_2
				break
			}
		}
		if Item1 == weapon2.ItemValeur_1 && Item2 == weapon2.ItemValeur_2 {
			AchatWeapon2(weapon2, Item2, Item1)
		}
	}
	MenuFWeapon()
}

func AchatWeapon2(BuyItem Forgeron_Weapon, Item2 save.Item, Item1 save.Cristal) {
	var choix string
	NewWeapon := save.Weapon{
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
		save.Personnage.Credit -= BuyItem.Valeur
		save.Enlever_Item(Item2, 1)
		save.Ajout_Weapon(NewWeapon, 1)
	case "2":
		save.ClearScreen()
		ForgeronWeapon2()
	default:
		save.ClearScreen()
		AchatWeapon2(BuyItem, Item2, Item1)
	}
}

func ForgeronWeapon3() {
	var choix string
	var Item1 save.Cristal
	var Item2 save.Item

	weapon3 := Forgeron_Weapon{
		"Lightsaber count dooku",
		500,
		100,
		666,
		1,
		save.Cristal{
			"Red Crystal",
			"\033[91m",
		},
		save.Item{
			"Saber handle",
			200,
			"It's a curved saber handle",
			1,
		},
		-100,
		"This is Count Dooku's saber",
	}

	fmt.Println("Weaponsmith")
	fmt.Println("1- ", weapon3.Name)
	fmt.Println("   HP Bonus : ", weapon3.PvBonus, " | Damage Bonus : ", weapon3.DamageBonus)
	fmt.Println("Price :")
	fmt.Println("   credit : ", save.Personnage.Credit, " | ", weapon3.ItemValeur_1.Name, " | ", weapon3.ItemValeur_2.Name, " | way Force : ", save.Personnage.CoteForce)
	fmt.Println("____________________________________")
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
func Weaponchoice() {
	fmt.Println("Which Weapon do you want to buy ?")
	fmt.Println("1 - Light Saber Vador")
	fmt.Println("2 - Light Saber Obiwan")
	fmt.Println("3 - Lightsaber count dooku")
	fmt.Println("4 - Exit")
	var choix string
	fmt.Print("Enter the number corresponding to your choice: ")
	_, err := fmt.Scan(&choix)

	if err != nil {
		fmt.Println("Error during entry. Please enter a number.")
		return
	}

	switch choix {
	case "1":
		fmt.Println("You have chosen to buy the Light Saber Vador")
		MenuFWeapon()
	case "2":

		fmt.Println("You have chosen to buy the Light Saber Obiwan")
		MenuFArmor()
	case "3":

		fmt.Println("You have chosen to buy the Light Saber count dooku")
		MenuFWeapon()
	case "4":
		break
	default:
		fmt.Println("Invalid choice. Please choose a number between 1 and 4.")
		save.ClearScreen()
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
