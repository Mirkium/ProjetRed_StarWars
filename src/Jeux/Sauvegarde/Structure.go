package Sauvegarde

type Perso struct {
	Name        string
	Level       int
	PV_max      int
	PV_actuelle int
	Force       int
	CoteForce   int
	Credit      int
	Weapon      struct{}
	Armure      struct{}
	Inventaire  struct{}
	classe      struct{}
}

type Weapon struct {
	Name string
	DamageBonus int
	PvBonus int
	Color string
}

type Armure struct {
	Name string
	DamageBonus int
	PvBonus int
}

//========================================INVENTAIRE======================================

type Item struct {
	Name        string
	Valeur      int
	Description string
}

type ItemPacket struct {
	Objet    Item
	Quantite int
}

//var test ItemPacket
