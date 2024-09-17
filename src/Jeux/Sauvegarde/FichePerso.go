package Sauvegarde

import (
	"fmt"
)

func DisplayCharacter() {
	fmt.Println("/======================================\\")
	fmt.Println("           ",Personnage.Name," Lv :", Personnage.Level)
	fmt.Println("\\======================================/")
}