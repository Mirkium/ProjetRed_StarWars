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
<<<<<<< HEAD
	//Weapon          struct{}
	//Armure          struct{}
	//Inventaire      struct{}
	Classe classe
=======
	Weapon          []Weapon
	Armure          struct{}
	Inventaire      struct{}
	Classe          classe
>>>>>>> 303821bd3eee723783fb6890c44f1fa131a8beb3
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
