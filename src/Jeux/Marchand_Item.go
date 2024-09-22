package jeux

import (
	"fmt"
)

var Item_1 = Item{
	Name: "Saber handle",
	Quantity: 5,
	Price: 200,
	Description: "Lightsaber handle used in its design.",
}
var Item_2 = Item{
	Name: "rancor skin",
	Quantity: 5,
	Price: 200,
	Description: "Products used to craft armor.",
}
var Item_3 = Kit{
	Name: "KitPack rang 0",
	Heal: 50,
	Quantity: 10,
	Price: 0,
	Description: "Makeshift KitPack that regenerates you, but is not very effective.",
}
var Item_4 = Kit{
	Name: "KitPack rang 1",
	Heal: 100,
	Quantity: 10,
	Price: 50,
	Description: "Makeshift KitPack that regenerates you, but is only moderately effective.",
}

func ItemShop() {
	var choix string
	
	fmt.Println("/=================ITEM SHOP=================\\")
	fmt.Println("")
	fmt.Println(Yellow, " 1. ", Cyan, Item_1.Name)
	fmt.Println(Yellow, "   Price : ",Green, Item_1.Price,Reset)
	fmt.Println(" ____________________________________________")
	fmt.Println("")
	fmt.Println(Yellow, " 1. ", Cyan, Item_2.Name)
	fmt.Println(Yellow, "   Price : ",Green, Item_2.Price,Reset)
	fmt.Println(" ____________________________________________")
	fmt.Println("")
	fmt.Println(Yellow, " 1. ", Cyan, Item_3.Name)
	fmt.Println(Yellow,"   Heal : ",Cyan, Item_3.Heal)
	fmt.Println(Yellow, "   Price : ",Green, Item_3.Price,Reset)
	fmt.Println(" ____________________________________________")
	fmt.Println("")
	fmt.Println(Yellow, " 1. ", Cyan, Item_4.Name)
	fmt.Println(Yellow,"   Heal : ",Cyan, Item_4.Heal)
	fmt.Println(Yellow, "   Price : ",Green, Item_4.Price,Reset)
	fmt.Println(" ____________________________________________")
	fmt.Println("")
	fmt.Println(Yellow, "                 (0)", Cyan, " Exit",Reset)
	fmt.Println("")
	fmt.Println("\\=============================================/")
	fmt.Println("")
	fmt.Print("Your choice : ")
	fmt.Scanln(&choix)
	switch choix {
	case "1" :
		Item_1.AchatItem()
	case "2" :
		Item_2.AchatItem()
	case "3" :

	case "4" :

	case "0" :
		break
	default :
		clearScreen()
		ItemShop()
	}
}

func (Object Item) AchatItem() {
	if Personnage.Credit >= Object.Price {
		Personnage.Credit -= Object.Price
		Object.Quantity -= 1
		Ajout_Item(Object, 1)
	}
}

