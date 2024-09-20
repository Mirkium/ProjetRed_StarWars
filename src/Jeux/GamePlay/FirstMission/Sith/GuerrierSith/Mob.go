package guerriersith

import (
	save "Game/Jeux/Sauvegarde"
)

var AbilityLimaceKorRang1_1 save.Abilite = save.Abilite{"Crachat aside", 0, 20, 0, 1, 0, 0, "Lance un jet d'acide", 0, 0}
var AbilityLimaceKorRang1_2 save.Abilite = save.Abilite{"Coup basique", 0, 10, 0, 1, 0, 0, "Donne un coup basique", 0, 0}
var AbilityLimaceKorRang2_1 save.Abilite = save.Abilite{"Crachat aside", 0, 40, 0, 1, 0, 0, "Lance un jet d'acide", 0, 0}
var AbilityLimaceKorRang2_2 save.Abilite = save.Abilite{"Coup basique", 0, 20, 0, 1, 0, 0, "Donne un coup basique", 0, 0}
var AbilityLimaceKorRang3_1 save.Abilite = save.Abilite{"Crachat aside", 0, 60, 0, 1, 0, 0, "Lance un jet d'acide", 0, 0}
var AbilityLimaceKorRang3_2 save.Abilite = save.Abilite{"Coup basique", 0, 30, 0, 1, 0, 0, "Donne un coup basique", 0, 0}
var ItemLimace save.Item = save.Item{"Dent de Limace Kor", 30, "Dent tombée d'une Limace Kor", 2}

var LimaceKor_Rang1 save.Mob = save.Mob{
	"Limace Kor Rang 1",
	400,
	400,
	0,
	[]save.Abilite{AbilityLimaceKorRang1_1, AbilityLimaceKorRang1_2},
	map[save.Item]int{ItemLimace: ItemLimace.Quantite},
	20,
}

var LimaceKor_Rang2 save.Mob = save.Mob{
	"Limace Kor Rang 2",
	800,
	800,
	10,
	[]save.Abilite{AbilityLimaceKorRang2_1, AbilityLimaceKorRang2_2},
	map[save.Item]int{ItemLimace: ItemLimace.Quantite},
	50,
}

var LimaceKor_Rang3 save.Mob = save.Mob{
	"Limace Kor Rang 3",
	1250,
	1250,
	20,
	[]save.Abilite{AbilityLimaceKorRang3_1, AbilityLimaceKorRang3_2},
	map[save.Item]int{ItemLimace: ItemLimace.Quantite},
	100,
}

