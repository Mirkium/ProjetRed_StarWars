package models

type Perso struct {
	Name           string    `json:"Name"`
	Campagne       Chapitre  `json:"Campagne"`
	WayForce       float32   `json:"WayForc"`
	Armor          []Armor   `json:"Armor"`
	Weapon         Weapon    `json:"Weapon"`
	SecondWeapon   Weapon    `json:"SecondWeapon"`
	HP_max         int       `json:"HP_max"`
	HP_actual      int       `json:"HP_actual"`
	Attack         int       `json:"Attack"`
	DefaultAbility []Ability `json:"DefaultAbility"`
	Class          Class     `json:"Class"`
}



type Class struct {
	Name    string    `json:"Name"`
	Ability []Ability `json:"Ability"`
}

type Ability struct {
	Name         string `json:"Name"`
	Attack       int    `json:"Attack"`
	Consomation  int    `json:"Consomation"`
	Recuperation int    `json:"Recuperation"`
}

type Chapitre struct {
	Name          string `json:"Name"`
	Nb_maxMission int    `json:"Nb_maxMission"`
	ActualMission int    `json:"ActualMission"`
	Way           string `json:"Way"`
}
