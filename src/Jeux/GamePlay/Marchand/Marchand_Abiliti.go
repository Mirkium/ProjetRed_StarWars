package marchand

import (
	save "Game/Jeux/Sauvegarde"
	"fmt"
)

func MenuAbilitie() {
	MarchandAbiliti()
}

<<<<<<< HEAD
func MarchandAbiliti() {
=======
func (m MarchandAbiliti) DisplayProduct() {
	fmt.Println("=== Armor MarchantArmor ===")
	for index, product := range m.product {
		fmt.Printf("\t%d - %s price : %d, Remaining quantity : %d, Side of force required : %d,\n", (index + 1), product.Name, product.Price, product.Quantity, product.CoteForce)
	}
	fmt.Printf(" Remaining items: %d\n", m.MoneyRemaining())
}

func (m MarchandAbiliti) MoneyRemaining() int {
	total := 0
	for _, obj := range m.product {
		total += obj.Quantity
	}
	return total
}

func (m *MarchantArmor) DisplayMenu(u *save.Perso) {
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
		fmt.Printf("Item %s is no longer available from the MarchantArmor.\n", selectedProduct.Name)
	} else {
		selectedProduct.Quantity--
	}
	MenuAbilitie(u, m)
}

func Marchand_Abiliti(Credit int, CoteForce int) {
>>>>>>> eaeb2da9ec69c6fe6d6e0cb520c7b602a872a6a0
	item01 := save.Abilite{Name: "Jedi Battle Armor",
		EnergieCost: 100,
		Dammage:     50,
		Heal:        0,
		Quantity:    1,
		Price:       4000,
		CoteForce:   10000,
		Description: "During the Sith Wars, the Jedi were forced to take on greater protection. These were unique armors, fashioned for each Jedi."}
	item02 := save.Abilite{Name: "Jedi Battle Armor",
		EnergieCost: 100,
		Dammage:     50,
		Heal:        0,
		Quantity:    1,
		Price:       4000,
		CoteForce:   10000,
		Description: "During the Sith Wars, the Jedi were forced to take on greater protection. These were unique armors, fashioned for each Jedi."}
	item03 := save.Abilite{Name: "Jedi Battle Armor",
		EnergieCost: 100,
		Dammage:     50,
		Heal:        0,
		Quantity:    1,
		Price:       4000,
		CoteForce:   10000,
		Description: "During the Sith Wars, the Jedi were forced to take on greater protection. These were unique armors, fashioned for each Jedi."}
	item04 := save.Abilite{
		Name:        "Gungan Personal Energy Shield",
		EnergieCost: 0,
		Dammage:     10,
		Heal:        0,
		Quantity:    3,
		Price:       10000,
		CoteForce:   0,
		Description: "Produced during the Galactic Republic era by the Otoh Gunga Defense League on Naboo, the Gungan Personal Energy Shield was a means of protection used by the soldiers and scouts of General Tobler Ceel's Grand Army."}
	fmt.Println("")
	fmt.Println(item01.Name, " : ", item01.Description, item01.EnergieCost, item01.Dammage, item01.Heal, item01.Quantity)
	fmt.Println(item02.Name, " : ", item02.Description, item02.EnergieCost, item02.Dammage, item02.Heal, item02.Quantity)
	fmt.Println(item03.Name, " : ", item03.Description, item03.EnergieCost, item03.Dammage, item03.Heal, item03.Quantity)
	fmt.Println(item01.Name, " : ", item04.Description, item04.EnergieCost, item04.Dammage, item04.Heal, item04.Quantity)
	fmt.Println("")
	fmt.Println("credit : ", save.Personnage.Credit, " | way Force : ", save.Personnage.CoteForce)
	fmt.Println("")
	MenuAbilitie()
}
