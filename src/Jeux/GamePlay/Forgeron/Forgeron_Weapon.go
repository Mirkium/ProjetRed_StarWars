package forgeron

import (
	save "Game/Jeux/Sauvegarde"
	"fmt"
	//"time"
)

func MenuFWeapon() {
	ForgeronWeapon()
}

func ForgeronWeapon() {
	var choix string
	var Item1 save.Cristal
	var Item2 save.Item
	weapon1 := Forgeron_Weapon{
		"Ligth Saber Vador",
		500,
		100,
		666,
		1,
		save.Cristal{
			"Cristal Rouge",
			"\033[91m",
		},
		save.Item{
			"Manche sabre",
			200,
			"C'est un ùanche de sabre",
			1,
		},
		-100,
		"C'est le sabre de vador",
	}

	fmt.Println("Forgeron d'arme")
	fmt.Println("1- ", weapon1.Name)
	fmt.Println("   Pv Bonus : ", weapon1.PvBonus, " | Damage Bonus : ", weapon1.DamageBonus)
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
			AchatWeapon(weapon1, Item2, Item1)
		}
	}
	MenuFWeapon()
}

func AchatWeapon(BuyItem Forgeron_Weapon, Item2 save.Item, Item1 save.Cristal) {
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
		ForgeronWeapon()
	default:
		save.ClearScreen()
		AchatWeapon(BuyItem, Item2, Item1)
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
