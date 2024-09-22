package jeux

type Perso struct {
	Name              string
	Level             int
	PV_max            int
	PV_actuelle       int
	Force             int
	CoteForce         int
	Credit            int
	Xp_Actu           int
	AbilitieDefault   []Abilite
	Weapon            Weapon
	IsSecondaryWeapon bool
	IsDoubleBlade     bool
	SecondaryWeapon   Weapon
	Armure            Armor
	Classe            Classe
}


//============================================================================================================

type Armor struct {
	Name        string
	PvBonus     int
	DamageBonus int
	StatArmor   int
	Quantity    int
	Price       int
	CoteForce   int
	Description string
}

type Classe struct {
	Name       string
	Energie    int
	Abilite    []Abilite
	EnergieMax int
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
	DotCompteur int
	DotDammage  int
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


//var test ItemPacket
