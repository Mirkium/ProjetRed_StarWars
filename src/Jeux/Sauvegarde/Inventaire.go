package Sauvegarde

import (
	"fmt"
	"strconv"
)

var Inventaire []Item

func Ajout_Item(Objet Item, quantite int) {
	/*
		Rajoute un item si il n'existe pas déja
		sinon rajoute la quantite.
		Verifie si l'inventaire est complet egalement
	*/
	AlreadyExiste := false
	for _, ele := range Inventaire {
		if Objet.Name == ele.Name {
			AlreadyExiste = true
			ele.Quantite += Objet.Quantite
			return
		}
	}
	if len(Inventaire) >= 9 && !AlreadyExiste {
		fmt.Println("L'inventaire est complet")
		return
	} else {
		Inventaire = append(Inventaire, Objet)
	}
}

func Enlever_Item(Objet Item, Quantite int) {
	/*
		Retire la quantite d'item demande et le supprime si <= 0
		Ne fais rien si l'item n'existe pas.
	*/
	for index, element := range Inventaire {
		if element.Name == Objet.Name {
			if element.Quantite-Objet.Quantite <= 0 {
				Inventaire = append(Inventaire[:index], Inventaire[index+1:]...)
				return
			} else {
				element.Quantite -= Objet.Quantite
				return
			}
		}
	}
	fmt.Println("L'item n'a pas été trouvé")
}

func Stats(Objet Item) string {
	/*
		renvoie la description de l'item
	*/
	return Objet.Description
}

func DisplayInventaire() {
	var Exit string
	/*
		Affiche les diiferents item
	*/
	ClearScreen()
	i := 1
	fmt.Println(" ========================Inventaire=========================")
	for _, element := range Inventaire {
		fmt.Printf("|Index : %d | Item :  %s | Quantité : %s |\n", i, Formatage(element.Name, 16), Formatage(strconv.Itoa(element.Quantite), 4))
		i++
	}
	fmt.Println("                         (0)  Exit")
	fmt.Println(" ===========================================================")
	fmt.Scanln(&Exit)
	switch Exit {
	case "0":
		break
	default:
		ClearScreen()
		DisplayInventaire()
	}
}

func Formatage(c string, z int) string {
	res := ""
	l := len(c)
	n := (z - l) / 2
	if l%2 == 1 {
		res += " "
	}
	for i := 0; i < n; i++ {
		res += " "
	}
	res += c
	for i := 0; i <= n+1; i++ {
		res += " "
	}
	return res
}

func formatagePrix(prix int) string {
	res := strconv.Itoa(prix)
	return Formatage(res+" credits", 16)
}

func DisplayInventaireIndex(index1 int) {
	if index1 > len(Inventaire) {
		fmt.Println("Index trop grand")
	} else {
		i := 1
		for _, element := range Inventaire {
			if index1 == i {
				ClearScreen()
				fmt.Println(" ========================Inventaire=========================")
				fmt.Println("|--------------------------Name-----------------------------|")
				fmt.Println("|                                                           |")
				fmt.Println("|                    " + Formatage(element.Name, 16) + "                     |")
				fmt.Println("|                                                           |")
				fmt.Println("|-----------------------Description-------------------------|")
				fmt.Println("|                                                           |")
				fmt.Println("|   " + element.Description)
				fmt.Println("|                                                           |")
				fmt.Println("---------------------------Prix-----------------------------|")
				fmt.Println("|                                                           |")
				fmt.Println("|                     " + formatagePrix(element.Valeur) + "                    |")
				fmt.Println("|                                                           |")
				fmt.Println("|------------------------Quantites--------------------------|")
				fmt.Println("|                                                           |")
				fmt.Println("|                          " + Formatage(strconv.Itoa(element.Quantite), 3) + "                           |")
				fmt.Println("|                                                           |")
				fmt.Println(" ===========================================================")
				//fmt.Printf("Item :  %s, Quantité : %d\n", index.Objet.Name, element)
			}
			i++
		}
	}
}

func Vendre(index int, quantite int) {
	var choix string
	DisplayInventaire()
	fmt.Println("Quel item veut tu vendre ?")
	fmt.Scanln(&choix)
	if len(choix) != 1 {
		fmt.Println("error 404 : pas le bon input")
	}
	for _, let := range choix {
		if let >= '0' && let <= '9' {
			//DisplayInventaireIndex(int(let - 48))
			i := 1
			for _, key := range Inventaire {
				if index == i {
					Enlever_Item(key, quantite)
					//return value * quantite
				}
			}
		} else {
			fmt.Println("on t'a demandé un chiffre pas autre chose.")
		}
	}
}

func RegardeStat() {
	DisplayInventaire()
	var choix string
	fmt.Println("Veux tu voir plus de détails sur un item (si non rentre juste 0)")
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
