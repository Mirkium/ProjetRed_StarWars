package Fight

/*import (
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
		fmt.Println("Quelle abilité souhaitez vous choisir ?")
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
	fmt.Println(" =====================Abilités======================")
	fmt.Println("|" + save.Formatage("Name", 16) + "|" + save.Formatage("Dammage", 8) + "|" + save.Formatage("Energie", 8) + "|  index   |")
	for in, element := range AbilitieList {
		d := element.Dammage + P.Weapon[0].DamageBonus - P.Armure[0].StatArmor
		if d <= 0 {
			d = 1
		}
		fmt.Println("|" + save.Formatage(element.Name, 16) + "|" + save.Formatage(strconv.Itoa(d), 8) + "|" + save.Formatage(strconv.Itoa(element.EnergieCost), 8) + "|" + save.Formatage(strconv.Itoa(in+2), 7) + "|")
	}
	fmt.Println(" ===================================================")
	time.Sleep(3 * time.Second)
}

func Fight(P *save.Perso, mob *save.Mob, PlayerStart bool) bool {
	rand.Seed(time.Now().UnixNano())
	save.ClearScreen()
	//affichage du debut de combat
	fmt.Println("Vous décidez d'affronter " + mob.Name)
	fmt.Println("Le combat will start in 3")
	fmt.Printf("\nStats : \n\t-Pv : %d\n\t-Armor : %d\n", mob.PV_max, mob.Armor)
	time.Sleep(1 * time.Second)
	save.ClearScreen()
	fmt.Println("Vous décidez d'affronter " + mob.Name)
	fmt.Println("Le combat will start in 2")
	fmt.Printf("\nStats : \n\t-Pv : %d\n\t-Armor : %d\n", mob.PV_max, mob.Armor)
	time.Sleep(1 * time.Second)
	save.ClearScreen()
	fmt.Println("Vous décidez d'affronter " + mob.Name)
	fmt.Println("Le combat will start in 1")
	fmt.Printf("\nStats : \n\t-Pv : %d\n\t-Armor : %d\n", mob.PV_max, mob.Armor)
	time.Sleep(1 * time.Second)
	save.ClearScreen()
	for {
		if PlayerStart {
			//affichage
			save.ClearScreen()
			fmt.Println("Il reste " + strconv.Itoa(mob.PV_actuelle) + " / " + strconv.Itoa(mob.PV_max) + " à " + mob.Name)
			time.Sleep(1 * time.Second)
			AbilitieList := []save.Abilite{}
			AbilitieList = append(AbilitieList, P.AbilitieDefault...)
			AbilitieList = append(AbilitieList, P.Classe.Abilite...)
			DisplayAbilite(AbilitieList, *P)
			fmt.Println("Quelle abilité souhaitez vous choisir ?")
			//choix des abilites
			answer := AskInt(1, len(AbilitieList)) // a tester une fois que le perso aura des abilites
			for AbilitieList[answer-1].EnergieCost > P.Classe.Energie {
				fmt.Println("Manque d'energie")
				answer = AskInt(1, len(AbilitieList))
			}
			fmt.Println("Vous décidez d'utliser " + AbilitieList[answer-1].Name)
			P.Classe.Energie -= AbilitieList[answer-1].EnergieCost
			dam := AbilitieList[answer-1].Dammage + P.Weapon[0].DamageBonus - mob.Armor
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
			if mob.PV_actuelle <= 0 {
				fmt.Println("Il reste " + strconv.Itoa(0) + " / " + strconv.Itoa(mob.PV_max) + " à " + mob.Name)
				time.Sleep(2 * time.Second)
				fmt.Println("You win")
				time.Sleep(3 * time.Second)
				ChangeLevel(P, *mob)
				return true
			} else {
				time.Sleep(2 * time.Second)
				fmt.Println("Il vous reste " + strconv.Itoa(P.PV_actuelle) + " / " + strconv.Itoa(P.PV_max))
				time.Sleep(2 * time.Second)
				fmt.Println("Il reste " + strconv.Itoa(mob.PV_actuelle) + " / " + strconv.Itoa(mob.PV_max) + " à " + mob.Name)
				time.Sleep(3 * time.Second)
			}
			PlayerStart = false
		}
		if !PlayerStart {
			save.ClearScreen()
			fmt.Println("Ennemis : " + mob.Name)
			randomInt := rand.Intn(len(mob.Abilitie))
			fmt.Println(mob.Name + " a décidé d'utiliser " + mob.Abilitie[randomInt].Name)
			dam := mob.Abilitie[randomInt].Dammage - P.Armure[0].StatArmor
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
			time.Sleep(2 * time.Second)
			if P.PV_actuelle < 0 {
				P.PV_actuelle = 0
			}
			fmt.Println("Il vous reste " + strconv.Itoa(P.PV_actuelle) + " / " + strconv.Itoa(P.PV_max))
			time.Sleep(2 * time.Second)
			if P.PV_actuelle <= 0 {
				save.ClearScreen()
				fmt.Println("Vous avez été vaincu.")
				return true
			} else {
				fmt.Println("\nIl reste " + strconv.Itoa(mob.PV_actuelle) + " / " + strconv.Itoa(mob.PV_max) + " à " + mob.Name)
				time.Sleep(5 * time.Second)
			}
			PlayerStart = true
		}
	}
}
*/