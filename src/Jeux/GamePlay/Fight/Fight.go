package Fight

import (
	save "Game/Jeux/Sauvegarde"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var DcP int = 0 //DotCompteurPerso
var DDP int = 0 //DotDammagePerso
var DcM int = 0 //DotCompteurMob
var DDM int = 0 //DotDammageMob

func ChangeLevel(P *save.Perso, mobKilled save.Mob) {
	P.Xp_Actu += mobKilled.XpDrop
	if P.Xp_Actu > 1000*P.Level {
		P.Xp_Actu -= 1000 * P.Level
		P.Level++
	}
	fmt.Printf("You are level %d", P.Level)
}

func Ask() int {
	var choix string
	fmt.Scan(&choix)
	answer, verification := strconv.Atoi(choix)
	for verification != nil {
		save.ClearScreen()
		choix = ""
		//refaire l'affichage des abilité
		fmt.Println("Bad input")
		fmt.Println("Which ability do you want to choose ?")
		fmt.Scan(&choix)
		answer, verification = strconv.Atoi(choix)
	}
	return answer
}

func AskInt(min int, max int) int {
	a := Ask()
	for a < min || a > max {
		a = Ask()
	}
	return a
}

func DisplayAbilite(AbilitieList []save.Abilite, P save.Perso) {
	fmt.Println(" ============================================Abilite======================================")
	fmt.Println("|" + save.Formatage("Name", 20) + "|" + save.Formatage("Energie", 8) + "|" + save.Formatage("Dammage", 8) + "|" + save.Formatage("Dot duration", 14) + "|" + save.Formatage(" Dot dammage", 14) + "|  index   |")
	for in, element := range AbilitieList {
		d := element.Dammage + P.Weapon.DamageBonus - P.Armure.StatArmor
		if d <= 0 {
			d = 1
		}
		fmt.Println("|" + save.Formatage(element.Name, 20) + "|" + save.Formatage(strconv.Itoa(element.EnergieCost), 8) + "|" + save.Formatage(strconv.Itoa(d), 8) + "|" + save.Formatage(strconv.Itoa(element.DotCompteur), 14) + "|" + save.Formatage(strconv.Itoa(element.DotDammage), 14) + "|" + save.Formatage(strconv.Itoa(in+1), 7) + "|")
	}
	fmt.Println(" =========================================================================================")
	time.Sleep(3 * time.Second)
}

func Fight(P *save.Perso, mob *save.Mob, PlayerStart bool) bool {
	rand.Seed(time.Now().UnixNano())
	save.ClearScreen()
	//affichage du debut de combat
	fmt.Println("You decide to fight " + mob.Name)
	fmt.Println("The fight will start in 3")
	fmt.Printf("\nStats : \n\t-Pv : %d\n\t-Armor : %d\n", mob.PV_max, mob.Armor)
	time.Sleep(1 * time.Second)
	save.ClearScreen()
	fmt.Println("You decide to fight " + mob.Name)
	fmt.Println("The fight will start in 2")
	fmt.Printf("\nStats : \n\t-Pv : %d\n\t-Armor : %d\n", mob.PV_max, mob.Armor)
	time.Sleep(1 * time.Second)
	save.ClearScreen()
	fmt.Println("You decide to fight " + mob.Name)
	fmt.Println("The fight will start in 1")
	fmt.Printf("\nStats : \n\t-Pv : %d\n\t-Armor : %d\n", mob.PV_max, mob.Armor)
	time.Sleep(1 * time.Second)
	save.ClearScreen()
	for {
		if PlayerStart {
			//affichage
			save.ClearScreen()
			fmt.Println("You have " + strconv.Itoa(P.PV_actuelle) + " / " + strconv.Itoa(P.PV_max) + " and " + strconv.Itoa(P.Classe.Energie) + " energie.")
			fmt.Println(mob.Name + " have " + strconv.Itoa(mob.PV_actuelle) + " / " + strconv.Itoa(mob.PV_max))
			time.Sleep(1 * time.Second)
			AbilitieList := []save.Abilite{}
			AbilitieList = append(AbilitieList, P.AbilitieDefault...)
			AbilitieList = append(AbilitieList, P.Classe.Abilite...)
			DisplayAbilite(AbilitieList, *P)
			fmt.Println("Which ability do you want to choose ?")
			//choix des abilites
			answer := AskInt(1, len(AbilitieList)) // a tester une fois que le perso aura des abilites
			for AbilitieList[answer-1].EnergieCost > P.Classe.Energie {
				fmt.Println("Low energy, you can't us this abilite.")
				answer = AskInt(1, len(AbilitieList))
			}
			fmt.Println("you decide to use " + AbilitieList[answer-1].Name)
			P.Classe.Energie -= AbilitieList[answer-1].EnergieCost
			dam := AbilitieList[answer-1].Dammage + P.Weapon.DamageBonus - mob.Armor
			if dam <= 0 {
				dam = 1
			}
			mob.PV_actuelle -= dam
			if AbilitieList[answer-1].Heal > 0 {
				if P.PV_max > P.PV_actuelle+AbilitieList[answer-1].Heal {
					P.PV_actuelle += AbilitieList[answer-1].Heal
				} else {
					P.PV_actuelle = P.PV_max
				}
			}
			if AbilitieList[answer-1].DotCompteur > 0 {
				DcP = AbilitieList[answer-1].DotCompteur
				DDP = AbilitieList[answer-1].DotDammage
			}
			if DcP > 0 {
				DcP--
				mob.PV_actuelle -= DDP
			}
			if P.Classe.Energie+20 >= P.Classe.EnergieMax {
				P.Classe.Energie = P.Classe.EnergieMax
			} else {
				P.Classe.Energie += 20
			}
			if mob.PV_actuelle <= 0 {
				fmt.Println(mob.Name + " have " + strconv.Itoa(0) + " / " + strconv.Itoa(mob.PV_max))
				time.Sleep(2 * time.Second)
				fmt.Println("You win.")
				time.Sleep(3 * time.Second)
				ChangeLevel(P, *mob)
				return true
			} else {
				time.Sleep(2 * time.Second)
				fmt.Println("You have " + strconv.Itoa(P.PV_actuelle) + " / " + strconv.Itoa(P.PV_max))
				time.Sleep(2 * time.Second)
				fmt.Println(mob.Name + " have " + strconv.Itoa(mob.PV_actuelle) + " / " + strconv.Itoa(mob.PV_max))
				time.Sleep(3 * time.Second)
			}
			PlayerStart = false
		}
		if !PlayerStart {
			save.ClearScreen()
			fmt.Println("Ennemis : " + mob.Name)
			randomInt := rand.Intn(len(mob.Abilitie))
			fmt.Println(mob.Name + " decide to use " + mob.Abilitie[randomInt].Name)
			dam := mob.Abilitie[randomInt].Dammage - P.Armure.StatArmor
			if dam <= 0 {
				dam = 1
			}
			P.PV_actuelle -= dam
			if mob.Abilitie[randomInt].Heal > 0 {
				if mob.PV_max < mob.PV_actuelle+mob.Abilitie[randomInt].Heal {
					mob.PV_actuelle += mob.Abilitie[randomInt].Heal
				} else {
					mob.PV_actuelle = mob.PV_max
				}
			}
			if mob.Abilitie[randomInt].DotCompteur > 0 {
				DcM = mob.Abilitie[randomInt].DotCompteur
				DDM = mob.Abilitie[randomInt].DotDammage
			}
			if DcM > 0 {
				DcM--
				P.PV_actuelle -= DDM
			}
			if P.PV_actuelle < 0 {
				P.PV_actuelle = 0
			}
			fmt.Println("You have " + strconv.Itoa(P.PV_actuelle) + " / " + strconv.Itoa(P.PV_max))
			time.Sleep(2 * time.Second)
			if P.PV_actuelle <= 0 {
				save.ClearScreen()
				fmt.Println("you have been defeated.")
				return true
			} else {
				fmt.Println(mob.Name + " have " + strconv.Itoa(mob.PV_actuelle) + " / " + strconv.Itoa(mob.PV_max))
				time.Sleep(5 * time.Second)
			}
			PlayerStart = true
		}
	}
}
