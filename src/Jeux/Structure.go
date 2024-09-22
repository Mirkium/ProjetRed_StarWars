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
	AbilitieDefault   []Ability
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
	Abilite    []Ability
	EnergieMax int
}



//=========================================Abilite<=================================

type Ability struct {
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
	Abilitie    []Ability
	Drop        map[Item]int
	XpDrop      int
}

//========================================Marchand======================================



//var test ItemPacket
