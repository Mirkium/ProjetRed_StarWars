package guerriersith

import (
	save "Game/Jeux/Sauvegarde"
)

var classeJedi1 save.Classe = save.Classe{"Way of the spirit", 100, []save.Abilite{save.Abilite{"Meditation", 10, 0, 30, 1, 0, 0, "An Abilite for rest a litle bit during a fight.", 0, 0}, save.Abilite{"Force Attack", 15, 20, 0, 1, 0, 0, "A Force Attack who project the ennemie on the wall", 0, 0}}}
var classeJedi2 save.Classe = save.Classe{"Way of the training", 150, []save.Abilite{save.Abilite{"Hissatsu Majishirīzu", 20, 30, 0, 1, 0, 0, "A serie of punch", 0, 0}, save.Abilite{"right hook", 40, 40, 0, 1, 0, 0, "A right hook on the face of the ennemie", 0, 0}}}
var classeJedi3 save.Classe = save.Classe{"Way of the instinct", 120, []save.Abilite{save.Abilite{"Charge of the Jedi", 20, 35, 0, 1, 0, 0, "Charge on the enemie", 0, 0}, save.Abilite{"Dot of the jedi", 40, 10, 0, 1, 0, 0, "A dot build in Jedi laboratory", 3, 10}}}

var ClassTank save.Classe = save.Classe{"Absorption", 150, []save.Abilite{save.Abilite{"Damage Heal", 10, 0, 40,1,0,0,"Use the force for healing",3,0}, save.Abilite{"Vital theif", 30, 40, 80,1,0,0,"stealing life",0,0},save.Abilite{"Force wave", 10, 20, 10, 1, 0, 0}}}