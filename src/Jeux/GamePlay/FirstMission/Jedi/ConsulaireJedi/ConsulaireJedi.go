package consulairejedi

import (
	fight "Game/Jeux/GamePlay/Fight"
	save "Game/Jeux/Sauvegarde"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

var ClasseConsulaire1 save.Classe = save.Classe{"Guardian", 110, []save.Abilite{save.Abilite{"Meditation", 10, 0, 30, 1, 0, 0, "An Abilite for rest a litle bit during a fight.", 0, 0}, save.Abilite{"Force wave", 30, 25, 0, 1, 0, 0, "A projection of the force who atack the ennemie", 0, 0}}}
var ClasseConsulaire2 save.Classe = save.Classe{"Sentinel", 100, []save.Abilite{save.Abilite{"Toxic gaz", 40, 15, 0, 1, 0, 0, "A toxic Gaz for the enemie but not for you", 5, 15}, save.Abilite{"right hook", 40, 40, 0, 1, 0, 0, "A right hook on the face of the ennemie", 0, 0}}}
var ClasseConsulaire3 save.Classe = save.Classe{"Nurse", 120, []save.Abilite{save.Abilite{"Heal", 20, 0, 20, 1, 0, 0, "Heal yourself", 0, 0}, save.Abilite{"Toxic gaz", 40, 15, 0, 1, 0, 0, "A toxic Gaz for the enemie but not for you", 5, 15}}}

var ClasseConsulaireList [3]save.Classe = [3]save.Classe{ClasseConsulaire1, ClasseConsulaire2, ClasseConsulaire3}

func JediConsular() {

}

func ClearScreen() {
	var cmd *exec.Cmd
	// Détecter le système d'exploitation
	switch runtime.GOOS {
	case "windows":
		// Commande pour Windows
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		// Commande pour les systèmes Unix-like
		cmd = exec.Command("clear")
	}
	// Définir la sortie de la commande sur Stdout
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Quete1(MC *save.Perso) {
	fmt.Printf("Librarian : So you want the simulate the fight betwen you and Obi-wan Kenobi?")
	time.Sleep(1 * time.Second)
	fmt.Printf("(1)Yes     (2)I just want to tain my technic.I think he is a good test for that.")
	var choix string
	fmt.Scan(&choix)
	_, verification := strconv.Atoi(choix)
	for verification != nil {
		fmt.Println("for ", verification)
		ClearScreen()
		choix = ""
		fmt.Println("(1)Yes     (2)I just want to train my technics.I think he is a good test for that.")
		fmt.Println("\nWe ask you to enter 1 or 2")
		fmt.Scan(&choix)
		_, verification = strconv.Atoi(choix)
	}
	ClearScreen()
	fmt.Printf("Librarian : Okay, I start the simulation.")
	ClearScreen()
	Obi_Wan := save.Mob{"Obi-wan", 100, 100, 20, []save.Abilite{save.Abilite{"Light saber", 0, 10, 0, 1, 0, 0, "Light Saber of obi wan", 0, 0}}, map[save.Item]int{}, 1000}
	fight.Fight(MC, &Obi_Wan, true)
	fmt.Println("Librarian : Congratulations, you win the level 4 / 10 of this simulation.")
}

func Quete2(MC *save.Perso) {
	fmt.Printf("Librarian : So you want the simulate the fight betwen you and Yoda ?")
	time.Sleep(1 * time.Second)
	fmt.Printf("(1)Yes     (2)Yes I want to see his experiece")
	var choix string
	fmt.Scan(&choix)
	_, verification := strconv.Atoi(choix)
	for verification != nil {
		fmt.Println("for ", verification)
		ClearScreen()
		choix = ""
		fmt.Println("(1)Yes     (2)I just want to train my technics.I think he is a good test for that.")
		fmt.Println("\nWe ask you to enter 1 or 2")
		fmt.Scan(&choix)
		_, verification = strconv.Atoi(choix)
	}
	ClearScreen()
	fmt.Printf("Librarian : Okay, I start the simulation.")
	ClearScreen()
	Obi_Wan := save.Mob{"Yoda", 300, 300, 60, []save.Abilite{save.Abilite{"Light saber", 0, 10, 0, 1, 0, 0, "Light Saber of yoda", 0, 0}, save.Abilite{"Self-Healing", 0, 0, 70, 1, 0, 0, "Self healing", 0, 0}}, map[save.Item]int{}, 1000}
	fight.Fight(MC, &Obi_Wan, true)
	fmt.Println("Librarian : Congratulations, you win the level 8 / 10 of this simulation.")
}
