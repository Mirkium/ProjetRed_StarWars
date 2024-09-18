package forgeron

import (
	save "Game/Jeux/Sauvegarde"
	"fmt"
)

/*func ForgeronTest() {
	Sword := save.Item{"Sword", 100, "sedfgujhch,;u,nyubgkjvfj", 1}
	fmt.Println("épée : lingot de fer + baton de bois")
	fmt.Println("100€")
	//Tu choisis l'épée
	price := 100
	var Item1 save.Item
	var Item2 save.Item
	for _, Item := range save.Inventaire {
		if Item.Name == "lingot de fer" {
			Item1 = Item
		}
		if Item.Name == "baton de bois" {
			Item2 = Item
		}
	}
	if save.Personnage.Credit >= price && Item1.Name == "lingot de fer" && Item2.Name == "baton de bois" {
		save.Enlever_Item(Item1, 1)
		save.Enlever_Item(Item2, 1)
		save.Personnage.Credit -= price
		save.Ajout_Item(Sword, 1)
	}
}*/

func ForgeronTest() {
	LaserSaber := save.Item{"laser saber", 100, "sedfgujhch,;u,nyubgkjvfj", 1}
	fmt.Println("épée : iron ingot + ")
	fmt.Println("100€")
	//Tu choisis l'épée
	price := 100
	var Item1 save.Item
	var Item2 save.Item
	for _, Item := range save.Inventaire {
		if Item.Name == "lingot de fer" {
			Item1 = Item
		}
		if Item.Name == "baton de bois" {
			Item2 = Item
		}
	}
	if save.Personnage.Credit >= price && Item1.Name == "lingot de fer" && Item2.Name == "baton de bois" {
		save.Enlever_Item(Item1, 1)
		save.Enlever_Item(Item2, 1)
		save.Personnage.Credit -= price
		save.Ajout_Item(LaserSaber, 1)
	}
}
func (f Forgeron) DisplayProduct() {
	fmt.Println("=== Armor Forgeron ===")
	for index, product := range f.Product {
		fmt.Printf("\t%d - %s price : %d, Remaining quantity : %d", (index + 1), product.Name, product.Price, product.Quantity)
	}
	fmt.Printf(" Remaining items: %d\n", f.MoneyRemaining())
}

func (f Forgeron) MoneyRemaining() int {
	total := 0
	for _, obj := range f.Product {
		total += obj.Quantity
	}
	return total
}

func (f *Forgeron) DisplayMenu(u *User) {
	fmt.Println("List of choices : ")
	for index, product := range f.Product {
		if product.Quantity > 0 {
			fmt.Printf("\t%d - Buy %s\n", (index + 1), product.Name)
		}
	}
	fmt.Println("Your choice?")
	var choix int
	fmt.Scan(&choix)

	if choix < 1 || choix > len(f.Product) {
		fmt.Println("Invalid choice...")
		f.DisplayMenu(u)
		return

	}
}
func MenuWeapon(u *User, f *Forgeron) {
	f.DisplayProduct()
	f.DisplayMenu(u)
}
