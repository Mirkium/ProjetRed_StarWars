package marchand

import (
	save "Game/Jeux/Sauvegarde"
	"fmt"
)



var joueur Player

func (m Marcht) DisplayProduct() {
	fmt.Println("=== Weapon Merchant ===")
	for index, product := range m.product {
		fmt.Printf("\t%d - %s price : %d, Remaining quantity : %d, Side of force required : %d,\n", (index + 1), product.Name, product.Price, product.Quantity, product.CoteForce)
	}
	fmt.Printf(" Remaining items: %d\n", m.MoneyRemaining())
}

func (m Marcht) MoneyRemaining() int {
	total := 0
	for _, obj := range m.product {
		total += obj.Quantity
	}
	return total
}

func (m *Marcht) DisplayMenu(u *Player) {
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
	MenuWeapon(u, m)
}

func (u *Player) Buy(obj *save.Weapon) {
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
		u.bag = append(u.bag, save.Weapon{Name: obj.Name, Quantity: 1, Price: obj.Price})
		fmt.Printf("New item add in inventory : %s\n", obj.Name)
	}

	fmt.Printf("Purchase made : %s, remaining money : %d\n", obj.Name, u.credit)
}

func MenuWeapon(u *Player, m *Marcht) {
	m.DisplayProduct()
	m.DisplayMenu(u)
}
func MarchandWeapon(Credit int, CoteForce int) {
	item01 := save.Weapon{Name: "LaserSaber Sith",
		PvBonus:     50,
		DamageBonus: 0,
		Color:       "Red",
		Quantity:    5,
		Price:       45,
		CoteForce:   -50,
		Description: "melee weapon, This weapon requires training and its use is greatly improved when its user uses the Force."}
	item02 := save.Weapon{Name: "LaserSaber Jedi",
		PvBonus:     50,
		DamageBonus: 10,
		Color:       "Blue",
		Quantity:    1,
		Price:       45,
		CoteForce:   50,
		Description: " melee weapon, this weapon requires training and its use is greatly improved when its user uses the Force."}
	item03 := save.Weapon{Name: "E-11 Blaster Rifle",
		PvBonus:     30,
		DamageBonus: 10,
		Color:       "Black and White",
		Quantity:    1,
		Price:       30,
		CoteForce:   0,
		Description: "A powerful, light and compact weapon, the E-11 was used widely throughout the galaxy for nearly a century and a half."}
	item04 := save.Weapon{Name: "Electro-Proton Bomb",
		PvBonus:     100,
		DamageBonus: 0,
		Color:       "Purle in explosion",
		Quantity:    3,
		Price:       20000,
		CoteForce:   0,
		Description: "Invented on behalf of the Grand Army of the Republic by Doctor Sionver Boll, the Electro-Proton Bomb released upon its explosion an electromagnetic pulse capable of destroying hundreds of Battle Droids in a few seconds.."}
	fmt.Println("")
	fmt.Println(item01.Name, " : ", item01.Description, item01.PvBonus, item01.DamageBonus, item01.Color)
	fmt.Println(item02.Name, " : ", item02.Description, item02.PvBonus, item02.DamageBonus, item02.Color)
	fmt.Println(item03.Name, " : ", item03.Description, item03.PvBonus, item03.DamageBonus, item03.Color)
	fmt.Println(item04.Name, " : ", item04.Description, item04.PvBonus, item04.DamageBonus, item04.Color)
	fmt.Println("")
	fmt.Println("credit : ", Utilisateur.credit, " | way Force : ", Utilisateur.coteForce)
	fmt.Println("")
	marchand := Marcht{"Merchant", []save.Weapon{item01, item02, item03}}
	Utilisateur.credit = Credit
	Utilisateur.coteForce = CoteForce
	MenuWeapon(&joueur, &marchand)

}
