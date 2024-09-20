package guerriersith

import (
	save "Game/Jeux/Sauvegarde"
)

var ClassTank save.Classe = save.Classe{"Sith Warrior Absorption", 150, []save.Abilite{save.Abilite{"Damage Heal", 10, 0, 40,1,0,0,"Use the force for healing",3,0}, save.Abilite{"Vital theif", 30, 20+save.Personnage.Force, 80+save.Personnage.Force,1,0,0,"stealing life",0,0},save.Abilite{"Force wave", 10, 5+save.Personnage.Force, 5+save.Personnage.Force, 1, 0, 0,"Launch a Force wave",0,0}}}
var ClassBurst save.Classe = save.Classe{"Sith Warrior Fureur", 120, []save.Abilite{save.Abilite{"Force jump assault", 50, 30+save.Personnage.Force, 0, 1, 0, 0, "Use the Force for jump to enemy", 0, 0}, save.Abilite{"surgical strike",50, 80+save.Personnage.Force, 0, 1, 0, 0, "hits the enemy's vital points", 0, 0}, save.Abilite{"force strangulation", 20, 10+save.Personnage.Force, save.Personnage.Force,1, 0,0, "Use the Force for enemy's strangle", 0, 0}}}
var ClassDot save.Classe = save.Classe{"Sith Warrior Extermination", 100, []save.Abilite{save.Abilite{"Bleeding", 40, 20+save.Personnage.Force, 0, 1, 0, 0, "inflicts bleeding on the enemy", 3, 20+save.Personnage.Force}, save.Abilite{"force infusion", 50, 30+save.Personnage.Force, 0, 1, 0, 0, "infusion of force into the enemy's body", 3, 30+save.Personnage.Force}}}

