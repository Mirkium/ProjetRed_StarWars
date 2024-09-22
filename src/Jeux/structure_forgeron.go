package jeux

import (

)

type Forgeron_Weapon struct {
	Name         string
	DamageBonus  int
	PvBonus      int
	Price       int
	Type         int
	Color        Cristal
	ItemValeur_1 Item
	ItemValeur_2 Lightsaber
	CoteForce    int
	Description  string
}

type Forgeron_Armor struct {
	Name         string
	PvBonus      int
	DamageBonus  int
	StatArmor    int
	Price       int
	ItemValeur_1 Item
	ItemValeur_2 Item
	CoteForce    int
	Description  string
}
