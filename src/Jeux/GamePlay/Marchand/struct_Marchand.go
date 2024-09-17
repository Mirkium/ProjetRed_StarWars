package marchand

import (
	save "Game/Jeux/Sauvegarde"
)

type Mar struct {
	name    string
	product []save.Weapon
}

type Play struct {
	credit    int
	coteForce int
	bag       []save.Weapon
}

type Marcht struct {
	name    string
	product []save.Weapon
}

type Player struct {
	credit    int
	coteForce int
	bag       []save.Weapon
}

type March struct {
	name    string
	product []save.Objects
}

type User struct {
	credit    int
	coteForce int
	bag       []save.Objects
}

type Merchant struct {
	name    string
	product []save.Armor
}

type Plyer struct {
	credit    int
	coteForce int
	bag       []save.Armor
}

type Merc struct {
	name    string
	product []save.Abilite
}

type Joue struct {
	credit    int
	coteForce int
	bag       []save.Abilite
}