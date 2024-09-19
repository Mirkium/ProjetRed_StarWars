package marchand

import (
	save "Game/Jeux/Sauvegarde"
	"fmt"
)

func MenuArmor() {
	MarchandArmor()
}

func MarchandArmor() {
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
	fmt.Println("credit : ", save.Personnage.Credit, " | way Force : ", save.Personnage.CoteForce)
	fmt.Println("")
	MenuArmor()
}
