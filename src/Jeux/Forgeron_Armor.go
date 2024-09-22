package jeux

import (
	"fmt"
)

var ArmorForge_1 = Forgeron_Armor{
	Name:         "Anakin Skywalker outfit",
	DamageBonus:  66,
	PvBonus:      66,
	StatArmor:    10,
	Price:        200,
	ItemValeur_1: Item{"Jedi outfi", 1, 100, "Jedi outfi"},
	ItemValeur_2: Item{"Padme necklace", 1, 100, "Padme necklace"},
	CoteForce:    0,
	Description:  "Anakin Skywalker outfit",
}

func ArmorForge() {
	var choix string
	var Item1 Item
	var Item2 Item
	fmt.Println("Credit    : ", Green, Personnage.Credit)
	fmt.Println("Way Force : ", Personnage.CoteForce)
	fmt.Println("/=================ARMOR FORGE=================\\")
	fmt.Println("")
	fmt.Println(Yellow, "   1. ", Cyan, ArmorForge_1.Name)
	fmt.Println(Yellow, "     Damage Bonus : ", Cyan, ArmorForge_1.DamageBonus, " | PV Bonus : ", Cyan, ArmorForge_1.PvBonus, " | Armor : ", Cyan, ArmorForge_1.StatArmor)
	fmt.Println(Yellow, "     Price : ", Cyan, ArmorForge_1.Price, " | Way Force : ", Cyan, ArmorForge_1.StatArmor, " | ", ArmorForge_1.ItemValeur_1.Name, " | ", ArmorForge_1.ItemValeur_2.Name)
	fmt.Println(" ____________________________________________")
	fmt.Println("")
	fmt.Println(Yellow,"                   (0)", Cyan, " Exit",Reset)
	fmt.Println("")
	fmt.Println("\\=============================================")
	fmt.Println("")
	fmt.Print("Your choice : ")
	fmt.Scanln(&choix)
	switch choix {
	case "1":
		for _, k := range Inventaire_Item {
			if k == ArmorForge_1.ItemValeur_1 {
				Item1 = k
			} else if k == ArmorForge_1.ItemValeur_2 {
				Item2 = k
			}
		}
		if Item1 == ArmorForge_1.ItemValeur_1 && Item2 == ArmorForge_1.ItemValeur_2 {
			ArmorForge_1.CraftArmor(Item1, Item2)
		}
	case "0":

	default:

	}
}

func (armur Forgeron_Armor) CraftArmor(Item1 Item, Item2 Item) {
	NewArmor := Armor{
		Name:        armur.Name,
		PvBonus:     armur.PvBonus,
		DamageBonus: armur.DamageBonus,
		StatArmor:   armur.StatArmor,
		Quantity:    1,
		Price:       armur.Price,
		CoteForce:   armur.CoteForce,
		Description: armur.Description,
	}
	if Personnage.Credit >= armur.Price {
		Enlever_Item(Item1, 1)
		Enlever_Item(Item2, 1)
		Ajout_Armur(NewArmor, 1)
		Personnage.Credit -= armur.Price
	}
}
