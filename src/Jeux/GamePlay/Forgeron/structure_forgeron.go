package forgeron

import (
	save "Game/Jeux/Sauvegarde"
)

type Forgeron_Weapon struct {
	Name         string
	PvBonus      int
	DamageBonus  int
	Valeur       int
	Type         int
	ItemValeur_1 save.Cristal
	ItemValeur_2 save.Item
	CoteForce    int
	Description  string
}

type Forgeron_Armor struct {
	Name         string
	PvBonus      int
	DamageBonus  int
	Valeur       int
	ItemValeur_1 save.Item
	ItemValeur_2 save.Item
	CoteForce    int
	Description  string
}
