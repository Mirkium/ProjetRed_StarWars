package jeux

import (
	
)

var AbilityLimaceKorRang1_1 Abilite = Abilite{"Crachat aside", 0, 20, 0, 1, 0, 0, "Lance un jet d'acide", 0, 0}
var AbilityLimaceKorRang1_2 Abilite = Abilite{"Coup basique", 0, 10, 0, 1, 0, 0, "Donne un coup basique", 0, 0}
var AbilityLimaceKorRang2_1 Abilite = Abilite{"Crachat aside", 0, 40, 0, 1, 0, 0, "Lance un jet d'acide", 0, 0}
var AbilityLimaceKorRang2_2 Abilite = Abilite{"Coup basique", 0, 20, 0, 1, 0, 0, "Donne un coup basique", 0, 0}
var AbilityLimaceKorRang3_1 Abilite = Abilite{"Crachat aside", 0, 60, 0, 1, 0, 0, "Lance un jet d'acide", 0, 0}
var AbilityLimaceKorRang3_2 Abilite = Abilite{"Coup basique", 0, 30, 0, 1, 0, 0, "Donne un coup basique", 0, 0}
var ItemLimace Item = Item{"Dent de Limace Kor", 30, 100, "Dent tombée d'une Limace Kor"}

var LimaceKor_Rang1 Mob = Mob{
	"Limace Kor Rang 1",
	400,
	400,
	0,
	[]Abilite{AbilityLimaceKorRang1_1, AbilityLimaceKorRang1_2},
	map[Item]int{ItemLimace: ItemLimace.Quantity},
	20,
}

var LimaceKor_Rang2 Mob = Mob{
	"Limace Kor Rang 2",
	800,
	800,
	10,
	[]Abilite{AbilityLimaceKorRang2_1, AbilityLimaceKorRang2_2},
	map[Item]int{ItemLimace: ItemLimace.Quantity},
	50,
}

var LimaceKor_Rang3 Mob = Mob{
	"Limace Kor Rang 3",
	1250,
	1250,
	20,
	[]Abilite{AbilityLimaceKorRang3_1, AbilityLimaceKorRang3_2},
	map[Item]int{ItemLimace: ItemLimace.Quantity},
	100,
}

