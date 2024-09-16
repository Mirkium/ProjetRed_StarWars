package Sauvegarde

import (
	"fmt"
	"strconv"
)

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
	ClearScreen()
	i := 1
	fmt.Println(" ========================Inventaire=========================")
	for index, element := range Inventaire {
		fmt.Printf("|Index : %d | Item :  %s | Quantité : %s |\n", i, formatage(index.Objet.Name, 16), formatage(string(element+48), 4))
		i++
	}
	fmt.Println(" ===========================================================")
}

func formatage(c string, z int) string {
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
	return formatage(res+" credits", 16)
}

func DisplayInventaireIndex(index1 int) {
	if index1 > len(Inventaire) {
		fmt.Println("Index trop grand")
	} else {
		i := 1
		for index, element := range Inventaire {
			if index1 == i {
				ClearScreen()
				fmt.Println(" ========================Inventaire=========================")
				fmt.Println("|--------------------------Name-----------------------------|")
				fmt.Println("|                                                           |")
				fmt.Println("|                    " + formatage(index.Objet.Name, 16) + "                     |")
				fmt.Println("|                                                           |")
				fmt.Println("|-----------------------Description-------------------------|")
				fmt.Println("|                                                           |")
				fmt.Println("|   " + index.Objet.Description)
				fmt.Println("|                                                           |")
				fmt.Println("---------------------------Prix-----------------------------|")
				fmt.Println("|                                                           |")
				fmt.Println("|                     " + formatagePrix(index.Objet.Valeur) + "                    |")
				fmt.Println("|                                                           |")
				fmt.Println("|------------------------Quantites--------------------------|")
				fmt.Println("|                                                           |")
				fmt.Println("|                          " + formatage(strconv.Itoa(element), 3) + "                           |")
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
			for key := range Inventaire {
				if index == i {
					Enlever_Item(key.Objet, quantite)
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
