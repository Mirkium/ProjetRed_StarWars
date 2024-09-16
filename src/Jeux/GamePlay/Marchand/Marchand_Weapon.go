package marchand

import (
	"fmt"
)

type Weapon struct {
	name        string
	quantity    int
	price       int
	CoteForce   int
	Description string
}

type marcht struct {
	name    string
	product []Weapon
}

type player struct {
	credit    int
	coteForce int
	bag       []Weapon
}

var joueur player

func (m marcht) DisplayProduct() {
	fmt.Println("=== Weapon Merchant ===")
	for index, product := range m.product {
		fmt.Printf("\t%d - %s price : %d, Remaining quantity : %d, Side of force required : %d,\n", (index + 1), product.name, product.price, product.quantity, product.CoteForce)
	}
	fmt.Printf(" Remaining items: %d\n", m.MoneyRemaining())
}

func (m marcht) MoneyRemaining() int {
	total := 0
	for _, obj := range m.product {
		total += obj.quantity
	}
	return total
}

func (m *marcht) DisplayMenu(u *player) {
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
	MenuWeapon(u, m)
}

func (u *player) Buy(obj *Weapon) {
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
		u.bag = append(u.bag, Weapon{name: obj.name, quantity: 1, price: obj.price})
		fmt.Printf("New item add in inventory : %s\n", obj.name)
	}

	fmt.Printf("Purchase made : %s, remaining money : %d\n", obj.name, u.credit)
}

func MenuWeapon(u *player, m *marcht) {
	m.DisplayProduct()
	m.DisplayMenu(u)
}
func MarchandWeapon(Credit int, CoteForce int) {
	item01 := Weapon{"LaserSaber Sith",
		45,
		5,
		-50,
		"melee weapon, This weapon requires training and its use is greatly improved when its user uses the Force."}
	item02 := Weapon{"LaserSaber Jedi",
		20,
		1,
		50,
		" melee weapon, this weapon requires training and its use is greatly improved when its user uses the Force."}
	item03 := Weapon{"E-11 Blaster Rifle",
		20,
		1,
		0,
		"A powerful, light and compact weapon, the E-11 was used widely throughout the galaxy for nearly a century and a half."}
	item04 := Weapon{"Electro-Proton Bomb",
		2000,
		1,
		0,
		"Invented on behalf of the Grand Army of the Republic by Doctor Sionver Boll, the Electro-Proton Bomb released upon its explosion an electromagnetic pulse capable of destroying hundreds of Battle Droids in a few seconds.."}
	fmt.Println("")
	fmt.Println(item01.name, " : ", item01.Description)
	fmt.Println(item02.name, " : ", item02.Description)
	fmt.Println(item03.name, " : ", item03.Description)
	fmt.Println(item04.name, " : ", item04.Description)
	fmt.Println("")
	fmt.Println("credit : ", Utilisateur.credit, " | way Force : ", Utilisateur.coteForce)
	fmt.Println("")
	marchand := marcht{"Merchant", []Weapon{item01, item02, item03}}
	Utilisateur.credit = Credit
	Utilisateur.coteForce = CoteForce
	MenuWeapon(&joueur, &marchand)
}
