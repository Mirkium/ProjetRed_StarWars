package main

import (
	save "Game/Jeux/Sauvegarde"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func ChangeLevel(P *save.Perso, mobKilled save.Mob) {
	P.Xp_Actu += mobKilled.XpDrop
	if P.Xp_Actu > 1000*P.Level {
		P.Xp_Actu -= 1000 * P.Level
		P.Level++
	}
}

func Fight(P *save.Perso, mob *save.Mob, PlayerStart bool) {
	rand.Seed(time.Now().UnixNano())
	save.ClearScreen()
	//affichage du debut de combat
	fmt.Println("Vous décidez d'affronter " + mob.Name)
	fmt.Println("Le combat will start in 3")
	time.Sleep(1 * time.Second)
	save.ClearScreen()
	fmt.Println("Vous décidez d'affronter " + mob.Name)
	fmt.Println("Le combat will start in 2")
	time.Sleep(1 * time.Second)
	save.ClearScreen()
	fmt.Println("Vous décidez d'affronter " + mob.Name)
	fmt.Println("Le combat will start in 1")
	time.Sleep(1 * time.Second)
	save.ClearScreen()
	FightIsOver := false
	for !FightIsOver {
		if PlayerStart {
			if P.PV_actuelle <= 0 {
				save.ClearScreen() //combat termine
				FightIsOver = true
			} else {
				//affichage
				save.ClearScreen()
				fmt.Println("Il reste " + strconv.Itoa(mob.PV_actuelle) + " / " + strconv.Itoa(mob.PV_max) + " à " + mob.Name)
				time.Sleep(2 * time.Second)
				AbilitieList := []save.Abilite{}
				fmt.Println(" =================Abilités=================")
				fmt.Println("|" + save.Formatage("Name", 16) + "|" + save.Formatage("Dammage", 8) + "|" + save.Formatage("Energie", 8) + "|")
				fmt.Println("|" + save.Formatage(P.AbilitieDefault.Name, 16) + "|" + save.Formatage(strconv.Itoa(P.AbilitieDefault.Dammage), 8) + "|" + save.Formatage(strconv.Itoa(P.AbilitieDefault.EnergieCost), 8) + "|")
				AbilitieList = append(AbilitieList, P.AbilitieDefault)
				if len(P.Classe.Abilite) != 0 {
					//affichage des abilite de la classe
					for _, element := range P.Classe.Abilite {
						fmt.Println("|" + save.Formatage(element.Name, 16) + "|" + save.Formatage(strconv.Itoa(element.Dammage), 8) + "|" + save.Formatage(strconv.Itoa(element.EnergieCost), 8) + "|")
						AbilitieList = append(AbilitieList, element)
					}
				}
				fmt.Println(" ==========================================")
				time.Sleep(5 * time.Second)
				fmt.Println("Quelle abilité souhaitez vous choisir ?")
				//choix des abilites
				var choix string
				fmt.Scanln(&choix)
				if len(choix) != 1 {
					fmt.Println("Mauvais input")
				}
				for _, e := range choix {
					if !(e >= '0' && e < '9') {
						fmt.Print("mauvais input")
					} else {
						//quelle abilite veux tu choisir
						save.ClearScreen()
						c, _ := strconv.Atoi(choix)
						if len(AbilitieList) < c { //verifier le type rune
							fmt.Println("Mauvais input, trop grand")
						} else {
							if AbilitieList[c].EnergieCost > P.Classe.Energie {
								fmt.Println("Manque d'energie")
							} else {
								fmt.Println("Vous décidez d'utliser " + AbilitieList[c].Name)
								P.Classe.Energie -= AbilitieList[c].EnergieCost
								mob.PV_actuelle -= AbilitieList[c].Dammage //ne pas oubliez l'application de l'armure
								if P.PV_max > P.PV_actuelle+AbilitieList[c].Heal {
									P.PV_actuelle += AbilitieList[c].Heal
								} else {
									P.PV_actuelle = P.PV_max
								}
								time.Sleep(2 * time.Second)
								fmt.Println("Il vous reste " + strconv.Itoa(P.PV_actuelle) + " / " + strconv.Itoa(P.PV_max))
								time.Sleep(2 * time.Second)
								fmt.Println("Il reste " + strconv.Itoa(mob.PV_actuelle) + " / " + strconv.Itoa(mob.PV_max) + " à " + mob.Name)
								time.Sleep(3 * time.Second)
							}
						}
					}
				}
			}
			PlayerStart = false
		}
		if !PlayerStart {
			if mob.PV_actuelle <= 0 {
				fmt.Println("Le mob adverse a été batue")
				FightIsOver = true
			} else {
				save.ClearScreen()
				fmt.Println("Ennemis : " + mob.Name)
				randomInt := rand.Intn(len(mob.Abilitie))
				fmt.Println(mob.Name + " a décidé d'utiliser " + mob.Abilitie[randomInt].Name)
				P.PV_actuelle -= mob.Abilitie[randomInt].Dammage
				if mob.Abilitie[randomInt].Heal > 0 {
					if mob.PV_max < mob.PV_actuelle+mob.Abilitie[randomInt].Heal {
						mob.PV_actuelle += mob.Abilitie[randomInt].Heal
					} else {
						mob.PV_actuelle = mob.PV_max
					}
				}
			}
			time.Sleep(2 * time.Second)
			fmt.Println("Il vous reste " + strconv.Itoa(P.PV_actuelle) + " / " + strconv.Itoa(P.PV_max))
			time.Sleep(2 * time.Second)
			if P.PV_actuelle <= 0 {
				save.ClearScreen()
				fmt.Println("Vous avez été vaincu.")
				FightIsOver = true
			} else {
				fmt.Println("Il reste " + strconv.Itoa(mob.PV_actuelle) + " / " + strconv.Itoa(mob.PV_max) + " à " + mob.Name)
				time.Sleep(5 * time.Second)
			}
			PlayerStart = true
		}
	}
	//afichage de fin de combat
	save.ClearScreen()
	if P.PV_actuelle <= 0 {
		fmt.Println(mob.Name + " vous a vaincue, il lui reste " + strconv.Itoa(mob.PV_actuelle) + "pv.")
	} else {
		fmt.Println("Vous avez vaincue " + mob.Name + ", il vous reste " + strconv.Itoa(P.PV_actuelle) + "pv.")
	}
}
