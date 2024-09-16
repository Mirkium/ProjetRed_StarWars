package marchand

import (
	"fmt"
)

type Objects struct {
	name        string
	quantity    int
	price       int
	CoteForce   int
	Description string
}

type march struct {
	name    string
	product []Objects
}

type User struct {
	credit    int
	coteForce int
	bag       []Objects
}

var Utilisateur User

func (m march) DisplayProduct() {
	fmt.Println("=== Merchant ===")
	for index, product := range m.product {
		fmt.Printf("\t%d - %s price : %d, Remaining quantity : %d, Side of force required : %d,\n", (index + 1), product.name, product.price, product.quantity, product.CoteForce)
	}
	fmt.Printf(" Remaining items: %d\n", m.MoneyRemaining())
}

func (m march) MoneyRemaining() int {
	total := 0
	for _, obj := range m.product {
		total += obj.quantity
	}
	return total
}

func (m *march) DisplayMenu(u *User) {
	fmt.Println("List of choices : ")
	for index, product := range m.product {
		if product.quantity > 0 {
			fmt.Printf("\t%d - Buy %s\n", (index + 1), product.name)
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

	if selectedProduct.quantity <= 0 {
		fmt.Printf("Item %s is no longer available from the merchant.\n", selectedProduct.name)
	} else {
		u.Buy(selectedProduct)
		selectedProduct.quantity--
	}
	Menu(u, m)
}

func (u *User) Buy(obj *Objects) {
	if u.credit < obj.price {
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

	u.credit -= obj.price
	isFind := false
	for index, objBag := range u.bag {
		if objBag.name == obj.name {
			isFind = true
			u.bag[index].quantity++
			fmt.Printf("Item add : %s, quantity in inventory: %d\n", objBag.name, u.bag[index].quantity)
			break
		}
	}

	if !isFind {
		u.bag = append(u.bag, Objects{name: obj.name, quantity: 1, price: obj.price})
		fmt.Printf("New item add in inventory : %s\n", obj.name)
	}

	fmt.Printf("Purchase made : %s, remaining money : %d\n", obj.name, u.credit)
}

func Menu(u *User, m *march) {
	m.DisplayProduct()
	m.DisplayMenu(u)
}
func Marchand(Credit int, CoteForce int) {
	item01 := Objects{"kitpack",
		5,
		45,
		0,
		"Allows you to heal some of the player's missing HP."}
	item02 := Objects{"Bandfill",
		1,
		20,
		0,
		" both wind and percussion instrument that requires great skills to master."}
	item03 := Objects{"Hemorrhagic item",
		20,
		1,
		0,
		"allows you to create hemorrhage in the enemy who will receive the next attack, gradually removing HP"}
	fmt.Println("")
	fmt.Println(item01.name, " : ", item01.Description)
	fmt.Println(item02.name, " : ", item02.Description)
	fmt.Println(item03.name, " : ", item03.Description)
	fmt.Println("")
	fmt.Println("credit : ", Utilisateur.credit, " | way Force : ", Utilisateur.coteForce)
	fmt.Println("")
	marchand := march{"Merchant", []Objects{item01, item02, item03}}
	Utilisateur.credit = Credit
	Utilisateur.coteForce = CoteForce
	Menu(&Utilisateur, &marchand)
}
