package marchand

import (
	save "Game/Jeux/Sauvegarde"
	"fmt"
)

type merchant struct {
	name    string
	product []save.Armor
}

type plyer struct {
	credit    int
	coteForce int
	bag       []save.Armor
}

var j1 plyer

func (m merchant) DisplayProduct() {
	fmt.Println("=== Armor Merchant ===")
	for index, product := range m.product {
		fmt.Printf("\t%d - %s price : %d, Remaining quantity : %d, Side of force required : %d,\n", (index + 1), product.Name, product.Price, product.Quantity, product.CoteForce)
	}
	fmt.Printf(" Remaining items: %d\n", m.MoneyRemaining())
}

func (m merchant) MoneyRemaining() int {
	total := 0
	for _, obj := range m.product {
		total += obj.Quantity
	}
	return total
}

func (m *merchant) DisplayMenu(u *plyer) {
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
	MenuArmor(u, m)
}

func (u *plyer) Buy(obj *save.Armor) {
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
		u.bag = append(u.bag, save.Armor{Name: obj.Name, Quantity: 1, Price: obj.Price})
		fmt.Printf("New item add in inventory : %s\n", obj.Name)
	}

	fmt.Printf("Purchase made : %s, remaining money : %d\n", obj.Name, u.credit)
}

func MenuArmor(u *plyer, m *merchant) {
	m.DisplayProduct()
	m.DisplayMenu(u)
}
func MarchandArmor(Credit int, CoteForce int) {
	item01 := save.Armor{Name: "darth vader helmet",
		PvBonus:     200,
		DamageBonus: 100,
		Color:       "Black",
		Quantity:    1,
		Price:       50000,
		CoteForce:   -10000,
		Description: "Having been badly burned on Moustafar, Darth Vader had no choice but to wear a breathing helmet to survive."}
	item02 := save.Armor{Name: "stormtrooper armor",
		PvBonus:     40,
		DamageBonus: 10,
		Color:       "White and Black",
		Quantity:    5,
		Price:       1000,
		CoteForce:   0,
		Description: "Developed by the Imperial Department of Military Research, the Stormtrooper's armor was made from simple, inexpensive materials. So the entire armor was made of 18 removable white plastic composite plates that the Stormtroopers slipped over a special black tunic."}
	item03 := save.Armor{Name: "Jedi Battle Armor",
		PvBonus:     100,
		DamageBonus: 50,
		Color:       "Brown and White",
		Quantity:    1,
		Price:       4000,
		CoteForce:   10000,
		Description: "During the Sith Wars, the Jedi were forced to take on greater protection. These were unique armors, fashioned for each Jedi."}
	item04 := save.Armor{Name: "Gungan Personal Energy Shield",
		PvBonus:     0,
		DamageBonus: 10,
		Color:       "Purple",
		Quantity:    3,
		Price:       10000,
		CoteForce:   0,
		Description: "Produced during the Galactic Republic era by the Otoh Gunga Defense League on Naboo, the Gungan Personal Energy Shield was a means of protection used by the soldiers and scouts of General Tobler Ceel's Grand Army."}
	fmt.Println("")
	fmt.Println(item01.Name, " : ", item01.Description, item01.PvBonus, item01.DamageBonus, item01.Color)
	fmt.Println(item02.Name, " : ", item02.Description, item02.PvBonus, item02.DamageBonus, item02.Color)
	fmt.Println(item03.Name, " : ", item03.Description, item03.PvBonus, item03.DamageBonus, item03.Color)
	fmt.Println(item04.Name, " : ", item04.Description, item04.PvBonus, item04.DamageBonus, item04.Color)
	fmt.Println("")
	fmt.Println("credit : ", Utilisateur.credit, " | way Force : ", Utilisateur.coteForce)
	fmt.Println("")
	marchand := merchant{"Merchant", []save.Armor{item01, item02, item03}}
	Utilisateur.credit = Credit
	Utilisateur.coteForce = CoteForce
	MenuArmor(&j1, &marchand)
}
