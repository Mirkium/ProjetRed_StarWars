package forgeron

import (
	save "Game/Jeux/Sauvegarde"
)

type User struct {
	credit    int
	coteForce int
	bag       []save.Objects
}

var Utilisateur User

type Item struct {
	Name        string
	Valeur      int
	Description string
	Quantite    int
}
type Forgeron struct {
	Name    string
	Product []save.Weapon
}
