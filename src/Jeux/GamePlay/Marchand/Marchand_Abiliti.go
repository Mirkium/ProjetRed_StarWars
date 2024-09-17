package marchand

import (
	save "Game/Jeux/Sauvegarde"
	"fmt"
)

type merc struct {
	name    string
	product []save.Abilite
}

type joue struct {
	credit    int
	coteForce int
	bag       []save.Abilite
}

var p1 joue

func (m merc) DisplayProduct() {
	fmt.Println("=== Abilitie Merchant ===")
	for index, product := range m.product {
		fmt.Printf("\t%d - %s price : %d, Remaining quantity : %d, Side of force required : %d,\n", (index + 1), product.Name, product.Price, product.Quantity, product.CoteForce)
	}
	fmt.Printf(" Remaining items: %d\n", m.MoneyRemaining())
}

func (m merc) MoneyRemaining() int {
	total := 0
	for _, obj := range m.product {
		total += obj.Quantity
	}
	return total
}

func (m *merc) DisplayMenu(u *joue) {
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
	MenuAbilitie(u, m)
}

func (u *joue) Buy(obj *save.Abilite) {
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
		u.bag = append(u.bag, save.Abilite{Name: obj.Name, Quantity: 1, Price: obj.Price})
		fmt.Printf("New item add in inventory : %s\n", obj.Name)
	}

	fmt.Printf("Purchase made : %s, remaining money : %d\n", obj.Name, u.credit)
}

func MenuAbilitie(u *joue, m *merc) {
	m.DisplayProduct()
	m.DisplayMenu(u)
}
func MarchandAbilitie(Credit int, CoteForce int) {
	item01 := save.Abilite{Name: "Absorption of life",
		EnergieCost: 200,
		Dammage:     100,
		Heal:        100,
		Quantity:    1,
		Price:       50000,
		CoteForce:   -10000,
		Description: "Life absorption is a dark power allowing you to vampirize an opponent's life force."}
	item02 := save.Abilite{Name: "Force Shield",
		EnergieCost: 100,
		Dammage:     0,
		Heal:        500,
		Quantity:    5,
		Price:       20000,
		CoteForce:   0,
		Description: "Power of the Force allowing its user to protect themselves from any type of aggression."}
	item03 := save.Abilite{Name: "Sith Alchemy",
		EnergieCost: 1000,
		Dammage:     1000,
		Heal:        0,
		Quantity:    1,
		Price:       30000,
		CoteForce:   -10000,
		Description: "Sith power allowing its user to manipulate very powerful lightning."}
	item04 := save.Abilite{Name: "Convection",
		EnergieCost: 200,
		Dammage:     200,
		Heal:        0,
		Quantity:    3,
		Price:       10000,
		CoteForce:   0,
		Description: "Convection allowed burning at a distance or on contact by concentrating the Force in the wrists."}
	fmt.Println("")
	fmt.Println(item01.Name, " : ", item01.Description, item01.EnergieCost, item01.Dammage, item01.Heal)
	fmt.Println(item02.Name, " : ", item02.Description, item02.EnergieCost, item02.Dammage, item02.Heal)
	fmt.Println(item03.Name, " : ", item03.Description, item03.EnergieCost, item03.Dammage, item03.Heal)
	fmt.Println(item04.Name, " : ", item04.Description, item04.EnergieCost, item04.Dammage, item04.Heal)
	fmt.Println("")
	fmt.Println("credit : ", Utilisateur.credit, " | way Force : ", Utilisateur.coteForce)
	fmt.Println("")
	marchand := merc{"Merchant", []save.Abilite{item01, item02, item03}}
	Utilisateur.credit = Credit
	Utilisateur.coteForce = CoteForce
	MenuAbilitie(&p1, &marchand)
}
