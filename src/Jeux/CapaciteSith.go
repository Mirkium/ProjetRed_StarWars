package jeux

import (
	
)

var ClassTank Classe = Classe{"Sith Warrior Absorption", 150, []Ability{Ability{"Damage Heal", 10, 50, 40, 1, 0, 0, "Use the force for healing", 3, 0}, Ability{"Vital theif", 30, 80, 80, 1, 0, 0, "stealing life", 0, 0}, Ability{"Force wave", 10, 20, 20, 1, 0, 0, "Launch a Force wave", 0, 0}}, 150}
var ClassBurst Classe = Classe{"Sith Warrior Fureur", 120, []Ability{Ability{"Force jump assault", 50, 70, 0, 1, 0, 0, "Use the Force for jump to enemy", 0, 0}, Ability{"surgical strike", 50, 200, 0, 1, 0, 0, "hits the enemy's vital points", 0, 0}, Ability{"force strangulation", 20, 40, Personnage.Force, 1, 0, 0, "Use the Force for enemy's strangle", 0, 0}}, 120}
var ClassDot Classe = Classe{"Sith Warrior Extermination", 100, []Ability{Ability{"Bleeding", 40, 30, 0, 1, 0, 0, "inflicts bleeding on the enemy", 3, 30}, Ability{"force infusion", 50, 50, 0, 1, 0, 0, "infusion of force into the enemy's body", 3, 50}}, 100}
