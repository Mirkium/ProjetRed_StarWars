package Sauvegarde

import (
	"fmt"
	"strconv"
)

var Inventaire_Item []Item
var Inventaire_Weapon []Weapon
var Inventaire_Armor []Armor

//Inventaire des Item

func Ajout_Item(Objet Item, quantite int) {
	/*
		Rajoute un item si il n'existe pas déja
		sinon rajoute la quantite.
		Verifie si l'Inventaire_Item est complet egalement
	*/
	AlreadyExiste := false
	for _, ele := range Inventaire_Item {
		if Objet.Name == ele.Name {
			AlreadyExiste = true
			ele.Quantite += Objet.Quantite
			return
		}
	}
	if len(Inventaire_Item) >= 9 && !AlreadyExiste {
		fmt.Println("L'Inventaire_Item est complet")
		return
	} else {
		Inventaire_Item = append(Inventaire_Item, Objet)
	}
}

func Enlever_Item(Objet Item, Quantite int) {
	/*
		Retire la quantite d'item demande et le supprime si <= 0
		Ne fais rien si l'item n'existe pas.
	*/
	for index, element := range Inventaire_Item {
		if element.Name == Objet.Name {
			if element.Quantite-Objet.Quantite <= 0 {
				Inventaire_Item = append(Inventaire_Item[:index], Inventaire_Item[index+1:]...)
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
	fmt.Println(" ========================Inventaire_Item=========================")
	for _, element := range Inventaire_Item {
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
	if index1 > len(Inventaire_Item) {
		fmt.Println("Index trop grand")
	} else {
		i := 1
		for _, element := range Inventaire_Item {
			if index1 == i {
				ClearScreen()
				fmt.Println(" ========================Inventaire_Item=========================")
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
			//DisplayInventaire_ItemIndex(int(let - 48))
			i := 1
			for _, key := range Inventaire_Item {
				if index == i {
					Enlever_Item(key, quantite)
					//return value * quantite
				}
				i++
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

//Inventaire des Weapons==========================================================================================

func Ajout_Weapon(Objet Weapon, quantite int) {
	/*
		Rajoute un item si il n'existe pas déja
		sinon rajoute la quantite.
		Verifie si l'Inventaire_Item est complet egalement
	*/
	AlreadyExiste := false
	for _, ele := range Inventaire_Weapon {
		if Objet.Name == ele.Name {
			AlreadyExiste = true
			ele.Quantity += Objet.Quantity
			return
		}
	}
	if len(Inventaire_Weapon) >= 9 && !AlreadyExiste {
		fmt.Println("L'Inventaire_Item est complet")
		return
	} else {
		Inventaire_Weapon = append(Inventaire_Weapon, Objet)
	}
}

func Enlever_Weapon(Objet Weapon, Quantite int) {
	/*
		Retire la quantite d'item demande et le supprime si <= 0
		Ne fais rien si l'item n'existe pas.
	*/
	for index, element := range Inventaire_Weapon {
		if element.Name == Objet.Name {
			if element.Quantity-Objet.Quantity <= 0 {
				Inventaire_Weapon = append(Inventaire_Weapon[:index], Inventaire_Weapon[index+1:]...)
				return
			} else {
				element.Quantity -= Objet.Quantity
				return
			}
		}
	}
	fmt.Println("L'item n'a pas été trouvé")
}

func Stats_Weapon(Objet Weapon) string {
	/*
		renvoie la description de l'item
	*/
	return Objet.Description
}

func DisplayInventaire_Weapon() {
	/*
		Affiche les diiferents item
	*/
	ClearScreen()
	i := 1
	fmt.Println(" ========================     Weapon    =========================")
	for _, element := range Inventaire_Weapon {
		fmt.Printf("|Index : %d | Item :  %s | Quantité : %s |\n", i, Formatage(element.Name, 16), Formatage(strconv.Itoa(element.Quantity), 4))
		i++
	}
	fmt.Println(" ===========================================================")
}

func DisplayInventaireIndexWeapon(index1 int) {
	if index1 > len(Inventaire_Weapon) {
		fmt.Println("Index trop grand")
	} else {
		i := 1
		for _, element := range Inventaire_Weapon {
			if index1 == i {
				ClearScreen()
				fmt.Println(" ========================     Weapon    =========================")
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
				fmt.Println("|                     " + formatagePrix(element.Price) + "                    |")
				fmt.Println("|                                                           |")
				fmt.Println("|------------------------Quantites--------------------------|")
				fmt.Println("|                                                           |")
				fmt.Println("|                          " + Formatage(strconv.Itoa(element.Quantity), 3) + "                           |")
				fmt.Println("|                                                           |")
				fmt.Println(" ===========================================================")
				//fmt.Printf("Item :  %s, Quantité : %d\n", index.Objet.Name, element)
			}
			i++
		}
	}
}

func VendreWeapon(index int, quantite int) {
	var choix string
	DisplayInventaire_Weapon()
	fmt.Println("Quel item veut tu vendre ?")
	fmt.Scanln(&choix)
	if len(choix) != 1 {
		fmt.Println("error 404 : pas le bon input")
	}
	for _, let := range choix {
		if let >= '0' && let <= '9' {
			//DisplayInventaire_ItemIndex(int(let - 48))
			i := 1
			for _, key := range Inventaire_Weapon {
				if index == i {
					Enlever_Weapon(key, quantite)
					//return value * quantite
				}
				i++
			}
		} else {
			fmt.Println("on t'a demandé un chiffre pas autre chose.")
		}
	}
}

func RegardeStatWeapon() {
	DisplayInventaire_Weapon()
	var choix string
	fmt.Println("Veux tu voir plus de détails sur un item (si non rentre juste 0)")
	fmt.Scanln(&choix)
	if len(choix) != 1 {
		fmt.Println("error 404 : pas le bon input")
	}
	for _, let := range choix {
		if let >= '0' && let <= '9' {
			DisplayInventaireIndexWeapon(int(let - '0'))
		} else {
			fmt.Println("on t'a demandé un chiffre pas autre chose.")
		}
	}
}

//Inventaire Armur

func Ajout_Armur(Objet Armor, quantite int) {
	/*
		Rajoute un item si il n'existe pas déja
		sinon rajoute la quantite.
		Verifie si l'Inventaire_Item est complet egalement
	*/
	AlreadyExiste := false
	for _, ele := range Inventaire_Armor {
		if Objet.Name == ele.Name {
			AlreadyExiste = true
			ele.Quantity += Objet.Quantity
			return
		}
	}
	if len(Inventaire_Armor) >= 9 && !AlreadyExiste {
		fmt.Println("L'Inventaire_Item est complet")
		return
	} else {
		Inventaire_Armor = append(Inventaire_Armor, Objet)
	}
}

func Enlever_Armor(Objet Armor, Quantite int) {
	/*
		Retire la quantite d'item demande et le supprime si <= 0
		Ne fais rien si l'item n'existe pas.
	*/
	for index, element := range Inventaire_Armor {
		if element.Name == Objet.Name {
			if element.Quantity-Objet.Quantity <= 0 {
				Inventaire_Armor = append(Inventaire_Armor[:index], Inventaire_Armor[index+1:]...)
				return
			} else {
				element.Quantity -= Objet.Quantity
				return
			}
		}
	}
	fmt.Println("L'item n'a pas été trouvé")
}

func Stats_Armor(Objet Armor) string {
	/*
		renvoie la description de l'item
	*/
	return Objet.Description
}

func DisplayInventaire_Armor() {
	/*
		Affiche les diiferents item
	*/
	ClearScreen()
	i := 1
	fmt.Println(" ========================     Weapon    =========================")
	for _, element := range Inventaire_Armor {
		fmt.Printf("|Index : %d | Item :  %s | Quantité : %s |\n", i, Formatage(element.Name, 16), Formatage(strconv.Itoa(element.Quantity), 4))
		i++
	}
	fmt.Println(" ===========================================================")
}

func DisplayInventaireIndexArmor(index1 int) {
	if index1 > len(Inventaire_Armor) {
		fmt.Println("Index trop grand")
	} else {
		i := 1
		for _, element := range Inventaire_Armor {
			if index1 == i {
				ClearScreen()
				fmt.Println(" ========================     Weapon    =========================")
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
				fmt.Println("|                     " + formatagePrix(element.Price) + "                    |")
				fmt.Println("|                                                           |")
				fmt.Println("|------------------------Quantites--------------------------|")
				fmt.Println("|                                                           |")
				fmt.Println("|                          " + Formatage(strconv.Itoa(element.Quantity), 3) + "                           |")
				fmt.Println("|                                                           |")
				fmt.Println(" ===========================================================")
				//fmt.Printf("Item :  %s, Quantité : %d\n", index.Objet.Name, element)
			}
			i++
		}
	}
}

func VendreArmor(index int, quantite int) {
	var choix string
	DisplayInventaire_Armor()
	fmt.Println("Quel item veut tu vendre ?")
	fmt.Scanln(&choix)
	if len(choix) != 1 {
		fmt.Println("error 404 : pas le bon input")
	}
	for _, let := range choix {
		if let >= '0' && let <= '9' {
			//DisplayInventaire_ItemIndex(int(let - 48))
			i := 1
			for _, key := range Inventaire_Armor {
				if index == i {
					Enlever_Armor(key, quantite)
					//return value * quantite
				}
				i++
			}
		} else {
			fmt.Println("on t'a demandé un chiffre pas autre chose.")
		}
	}
}

func RegardeStaArmor() {
	DisplayInventaire_Armor()
	var choix string
	fmt.Println("Veux tu voir plus de détails sur un item (si non rentre juste 0)")
	fmt.Scanln(&choix)
	if len(choix) != 1 {
		fmt.Println("error 404 : pas le bon input")
	}
	for _, let := range choix {
		if let >= '0' && let <= '9' {
			DisplayInventaireIndexWeapon(int(let - '0'))
		} else {
			fmt.Println("on t'a demandé un chiffre pas autre chose.")
		}
	}
}
