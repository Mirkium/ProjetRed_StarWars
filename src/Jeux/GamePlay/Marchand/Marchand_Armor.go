package marchand

import (
	"fmt"
)

type Armor struct {
	name        string
	quantity    int
	price       int
	CoteForce   int
	Description string
}

type merchant struct {
	name    string
	product []Armor
}

type plyer struct {
	credit    int
	coteForce int
	bag       []Armor
}

var j1 plyer

func (m merchant) DisplayProduct() {
	fmt.Println("=== Armor Merchant ===")
	for index, product := range m.product {
		fmt.Printf("\t%d - %s price : %d, Remaining quantity : %d, Side of force required : %d,\n", (index + 1), product.name, product.price, product.quantity, product.CoteForce)
	}
	fmt.Printf(" Remaining items: %d\n", m.MoneyRemaining())
}

func (m merchant) MoneyRemaining() int {
	total := 0
	for _, obj := range m.product {
		total += obj.quantity
	}
	return total
}

func (m *merchant) DisplayMenu(u *plyer) {
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
	MenuArmor(u, m)
}

func (u *plyer) Buy(obj *Armor) {
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
		u.bag = append(u.bag, Armor{name: obj.name, quantity: 1, price: obj.price})
		fmt.Printf("New item add in inventory : %s\n", obj.name)
	}

	fmt.Printf("Purchase made : %s, remaining money : %d\n", obj.name, u.credit)
}

func MenuArmor(u *plyer, m *merchant) {
	m.DisplayProduct()
	m.DisplayMenu(u)
}
func MarchandArmor(Credit int, CoteForce int) {
	item01 := Armor{"darth vader helmet",
		500,
		1,
		-10000,
		"Having been badly burned on Moustafar, Darth Vader had no choice but to wear a breathing helmet to survive."}
	item02 := Armor{"stormtrooper armor",
		100,
		5,
		0,
		"Developed by the Imperial Department of Military Research, the Stormtrooper's armor was made from simple, inexpensive materials. So the entire armor was made of 18 removable white plastic composite plates that the Stormtroopers slipped over a special black tunic."}
	item03 := Armor{"Jedi Battle Armor",
		20,
		1,
		0,
		"During the Sith Wars, the Jedi were forced to take on greater protection. These were unique armors, fashioned for each Jedi."}
	item04 := Armor{"Gungan Personal Energy Shield",
		2000,
		1,
		0,
		"roduit à l'époque de la République Galactique par la Ligue de Défense d'Otoh Gunga sur Naboo, le Bouclier Energétique Personnel Gungan était un moyen de protection employé par les soldats et éclaireurs de la Grande Armée du Général Tobler Ceel."}
	fmt.Println("")
	fmt.Println(item01.name, " : ", item01.Description)
	fmt.Println(item02.name, " : ", item02.Description)
	fmt.Println(item03.name, " : ", item03.Description)
	fmt.Println(item04.name, " : ", item04.Description)
	fmt.Println("")
	fmt.Println("credit : ", Utilisateur.credit, " | way Force : ", Utilisateur.coteForce)
	fmt.Println("")
	marchand := merchant{"Merchant", []Armor{item01, item02, item03}}
	Utilisateur.credit = Credit
	Utilisateur.coteForce = CoteForce
	MenuArmor(&j1, &marchand)
}
