package marchand

import (
	save "Game/Jeux/Sauvegarde"
)

type MarchantWeapon struct {
	name    string
	product []save.Weapon
}

type MarchantArmor struct {
	name    string
	product []save.Armor
}

type MarchandItem struct {
	name    string
	product []save.Item
}

type MarchandAbiliti struct {
	name    string
	product []save.Abilite
}
