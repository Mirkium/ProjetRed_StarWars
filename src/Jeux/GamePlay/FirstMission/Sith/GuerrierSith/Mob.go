package guerriersith

import (
	save "Game/Jeux/Sauvegarde"
)

var AbilityLimaceKorRang1_1 save.Abilite = save.Abilite{"Crachat aside", 0, 40, 0, 1, 0, 0, "Lance un jet d'acide", 0, 0}
var AbilityLimaceKorRang1_2 save.Abilite = save.Abilite{"Coup basique", 0, 20, 0, 1, 0, 0, "Donne un coup basique", 0, 0}
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
