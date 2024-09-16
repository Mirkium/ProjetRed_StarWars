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
	Armure          struct{}
	Inventaire      struct{}
	Classe          classe
}

type Weapon struct {
	Name        string
	DamageBonus int
	PvBonus     int
	Color       string
}

type Armure struct {
	Name        string
	DamageBonus int
	PvBonus     int
}

type classe struct {
	Name    string
	Energie int
	Abilite []Abilite
}

//=========================================Abilite<=================================

type Abilite struct {
	Name        string
	EnergieCost int
	Dammage     int
	Heal        int
}

//=========================================MOB=======================================

type Mob struct {
	Name        string
	PV_max      int
	PV_actuelle int
	Damage      int
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
