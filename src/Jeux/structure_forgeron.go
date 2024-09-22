package jeux

import (

)

type Forgeron_Weapon struct {
	Name         string
	PvBonus      int
	DamageBonus  int
	Valeur       int
	Type         int
	ItemValeur_1 Cristal
	ItemValeur_2 Item
	CoteForce    int
	Description  string
}

type Forgeron_Armor struct {
	Name         string
	PvBonus      int
	DamageBonus  int
	StatArmor    int
	Valeur       int
	ItemValeur_1 Item
	ItemValeur_2 Item
	CoteForce    int
	Description  string
}
