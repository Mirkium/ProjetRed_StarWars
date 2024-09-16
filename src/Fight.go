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
	fmt.Println("Vous décidez d'affronter " + mob.Name)
	FightIsOver := false
	for !FightIsOver {
		if PlayerStart {
			if P.PV_actuelle <= 0 {
				fmt.Println("Vous avez été vaincu.")
				FightIsOver = true
			} else {
				//affichage
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
						if len(AbilitieList) > int(e+48) {
							fmt.Println("Mauvais input, trop grand")
						} else {
							if AbilitieList[int(e+48)].EnergieCost > P.Classe.Energie {
								fmt.Println("Manque d'energie")
							} else {
								P.Classe.Energie -= AbilitieList[int(e+48)].EnergieCost
								mob.PV_actuelle -= AbilitieList[int(e)+48].Dammage //ne pas oubliez l'application de l'armure
								if P.PV_max > P.PV_actuelle+AbilitieList[int(e+48)].Heal {
									P.PV_actuelle += AbilitieList[int(e+48)].Heal
								} else {
									P.PV_actuelle = P.PV_max
								}
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
				randomInt := rand.Intn(len(mob.Abilitie))
				P.PV_actuelle -= mob.Abilitie[randomInt].Dammage
				if mob.PV_max < mob.PV_actuelle+mob.Abilitie[randomInt].Heal {
					mob.PV_actuelle += mob.Abilitie[randomInt].Heal
				} else {
					mob.PV_actuelle = mob.PV_max
				}
			}
			if P.PV_actuelle <= 0 {
				fmt.Println("Vous avez été vaincu.")
				FightIsOver = true
			}
			PlayerStart = true
		}
	}
}
