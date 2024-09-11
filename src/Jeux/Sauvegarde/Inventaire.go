package Sauvegarde

import "fmt"

var Inventaire map[ItemPacket]int = make(map[ItemPacket]int)

func Ajout_Item(Objet ItemPacket) {
	/*
		Rajoute un item si il n'existe pas déja
		sinon rajoute la quantite.
		Verifie si l'inventaire est complet egalement
	*/
	if len(Inventaire) >= 9 {
		fmt.Println("L'inventaire est complets")
	} else {
		IsFind := false
		for key := range Inventaire {
			if key.Objet.Name == Objet.Objet.Name {
				Inventaire[Objet] += Objet.Quantite
				IsFind = true
				break
			}
		}
		if !IsFind {
			Inventaire[Objet] = Objet.Quantite
		}
	}
}

func Enlever_Item(Objet Item, Quantite int) {
	/*
		Retire la quantite d'item demande et le supprime si <= 0
		Ne fais rien si l'item n'existe pas.
	*/
	for key, value := range Inventaire {
		if key.Objet == Objet {
			if value-Quantite <= 0 {
				delete(Inventaire, key)
			} else {
				Inventaire[key] -= Quantite
			}
		}
	}
}

func Stats(Objet Item) string {
	/*
		renvoie la description de l'item
	*/
	return Objet.Description
}

func DisplayInventaire() {
	/*
		Affiche les diiferents item
	*/
	i := 1
	for index, element := range Inventaire {
		fmt.Printf("Item :  %s, Quantité : %d, Index : %d\n", index.Objet.Name, element, i)
		i++
	}
}

func DisplayInventaireIndex(index1 int) {
	if index1 > len(Inventaire) {
		fmt.Println("Index trop grand")
	} else {
		i := 1
		for index, element := range Inventaire {
			if index1 == i {
				fmt.Printf("Item :  %s, Quantité : %d\n", index.Objet.Name, element)
			}
			i++
		}
	}
}

func Selectionner() {
	/*
		Montre les inforamtion de l'item selectionner.
		Gere les autres cas d'ereur.
	*/
	var choix string
	fmt.Scanln(&choix)
	if len(choix) != 1 {
		fmt.Println("error 404 : pas le bon input")
	}
	for _, let := range choix {
		if let >= '0' && let <= '9' {
			DisplayInventaireIndex(int(let - 48))
		} else {
			fmt.Println("on t'a demandé un chiffre pas autre chose.")
		}
	}
}
