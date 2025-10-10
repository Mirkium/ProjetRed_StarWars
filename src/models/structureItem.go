package models

type Armor struct {
	Name     string `json:"Name"`
	BonusDMG int    `json:"BonusDMG"`
	BonusHP  int    `json:"BonusHP`
}

type Weapon struct {
	Name     string `json:"Name"`
	Cristal  string `json:"Cristal"`
}