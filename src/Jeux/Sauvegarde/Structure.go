package Sauvegarde

type Perso struct {
	Name            string
	Level           int
	PV_max          int
	PV_actuelle     int
	Force           int
	CoteForce       int
	Credit          int
	Xp_Actu         int
	AbilitieDefault Abilite
	Weapon          []Weapon
	SecondaryWeapon []Weapon
	Armure          []Armor
	Classe          classe
}

type Weapon struct {
	Name        string
	DamageBonus int
	PvBonus     int
	Color       string
	Quantity    int
	Price       int
	CoteForce   int
	Description string
}

type SecondaryWeapon struct {
	Name        string
	DamageBonus int
	PvBonus     int
	Color       string
}

type Armure struct {
	Name        string
	PvBonus     int
	DamageBonus int
	Color       string
	Quantity    int
	Price       int
	CoteForce   int
	Description string
}

type Armor struct {
	Name        string
	PvBonus     int
	DamageBonus int
	Color       string
	Quantity    int
	Price       int
	CoteForce   int
	Description string
}

type classe struct {
	Name    string
	Energie int
	Abilite []Abilite
}

type Objects struct {
	Name        string
	PvBonus     int
	DamageBonus int
	Color       string
	Quantity    int
	Price       int
	CoteForce   int
	Description string
}

//=========================================Abilite<=================================

type Abilite struct {
	Name        string
	EnergieCost int
	Dammage     int
	Heal        int
	Quantity    int
	Price       int
	CoteForce   int
	Description string
}

//=========================================MOB=======================================

type Mob struct {
	Name        string
	PV_max      int
	PV_actuelle int
	Armor       int
	Abilitie    []Abilite
	Drop        map[Item]int
	XpDrop      int
}

//========================================INVENTAIRE======================================

type Item struct {
	Name        string
	Valeur      int
	Description string
	Quantite    int
}

//var test ItemPacket
