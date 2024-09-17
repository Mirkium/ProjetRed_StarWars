package marchand

import (
	save "Game/Jeux/Sauvegarde"
	"fmt"
)



var Utilisateur User

func (m March) DisplayProduct() {
	fmt.Println("=== Merchant ===")
	for index, product := range m.product {
		fmt.Printf("\t%d - %s price : %d, Remaining quantity : %d, Side of force required : %d,\n", (index + 1), product.Name, product.Price, product.Quantity, product.CoteForce)
	}
	fmt.Printf(" Remaining items: %d\n", m.MoneyRemaining())
}

func (m March) MoneyRemaining() int {
	total := 0
	for _, obj := range m.product {
		total += obj.Quantity
	}
	return total
}

func (m *March) DisplayMenu(u *User) {
	fmt.Println("List of choices : ")
	for index, product := range m.product {
		if product.Quantity > 0 {
			fmt.Printf("\t%d - Buy %s\n", (index + 1), product.Name)
		}
	}
	fmt.Println("Your choice?")
	var choix int
	fmt.Scan(&choix)

	if choix < 1 || choix > len(m.product) {
		fmt.Println("Invalid choice...")
		m.DisplayMenu(u)
		return
	}

	selectedProduct := &m.product[choix-1]

	if selectedProduct.Quantity <= 0 {
		fmt.Printf("Item %s is no longer available from the merchant.\n", selectedProduct.Name)
	} else {
		u.Buy(selectedProduct)
		selectedProduct.Quantity--
	}
	Menu(u, m)
}

func (u *User) Buy(obj *save.Objects) {
	if u.credit < obj.Price {
		fmt.Println("Insufficient credit...")
		return
	}

	if obj.CoteForce < 0 && obj.CoteForce > u.coteForce {
		fmt.Println(("Missing good point of force... "))
		return
	} else {
		u.coteForce -= obj.CoteForce
	}

	if obj.CoteForce > 0 && obj.CoteForce > u.coteForce {
		fmt.Println(("Missing dark point of force... "))
		return
	} else {
		u.coteForce -= obj.CoteForce
	}

	u.credit -= obj.Price
	isFind := false
	for index, objBag := range u.bag {
		if objBag.Name == obj.Name {
			isFind = true
			u.bag[index].Quantity++
			fmt.Printf("Item add : %s, quantity in inventory: %d\n", objBag.Name, u.bag[index].Quantity)
			break
		}
	}

	if !isFind {
		u.bag = append(u.bag, save.Objects{Name: obj.Name, Quantity: 1, Price: obj.Price})
		fmt.Printf("New item add in inventory : %s\n", obj.Name)
	}

	fmt.Printf("Purchase made : %s, remaining money : %d\n", obj.Name, u.credit)
}

func Menu(u *User, m *March) {
	m.DisplayProduct()
	m.DisplayMenu(u)
}
func Marchand(Credit int, CoteForce int) {
	item01 := save.Objects{Name: "kitpack",
		PvBonus:     20,
		DamageBonus: 0,
		Color:       "Green",
		Quantity:    0,
		Price:       45,
		CoteForce:   0,
		Description: "Allows you to heal some of the player's missing HP."}
	item02 := save.Objects{Name: "Bandfill",
		PvBonus:     23,
		DamageBonus: 10,
		Color:       "Brown",
		Quantity:    1,
		Price:       45,
		CoteForce:   0,
		Description: " both wind and percussion instrument that requires great skills to master."}
	item03 := save.Objects{Name: "Hemorrhagic item",
		PvBonus:     0,
		DamageBonus: 20,
		Color:       "Green and Purple",
		Quantity:    1,
		Price:       20,
		CoteForce:   0,
		Description: "allows you to create hemorrhage in the enemy who will receive the next attack, gradually removing HP"}
	fmt.Println("")
	fmt.Println(item01.Name, " : ", item01.Description, item01.PvBonus, item01.DamageBonus, item01.Color)
	fmt.Println(item02.Name, " : ", item02.Description, item02.PvBonus, item02.DamageBonus, item02.Color)
	fmt.Println(item03.Name, " : ", item03.Description, item03.PvBonus, item03.DamageBonus, item03.Color)
	fmt.Println("")
	fmt.Println("credit : ", Utilisateur.credit, " | way Force : ", Utilisateur.coteForce)
	fmt.Println("")
	Marchand := March{"Merchant", []save.Objects{item01, item02, item03}}
	Utilisateur.credit = Credit
	Utilisateur.coteForce = CoteForce
	Menu(&Utilisateur, &Marchand)
}
