package Sauvegarde

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Partie struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Joueur Perso  `json:"Joueur"`
}

func Verif() bool {
	// Chemin vers le fichier JSON
	filePath := "./Save.json"

	// Vérifier si le fichier existe
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("Le fichier Sauvegarde.json n'existe pas.")
		return false
	}

	// Lire le fichier
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
		return false
	}

	// Décoder le fichier JSON
	var partie []Partie
	err = json.Unmarshal(file, &partie)
	if err != nil {
		fmt.Println("Erreur lors du décodage JSON:", err)
		return false
	}

	// Vérifier si des données de joueur existent
	if len(partie) == 0 {
		fmt.Println("Aucune Partie Sauvegarder")
		return false
	} else {
		fmt.Println("Partie trouver")
		ClearScreen()
		return true
	}
}

func ChargeGame() {
	filePath := "Save.json"

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
		return
	}

	var partie []Partie
	err = json.Unmarshal(file, &partie)
	if err != nil {
		fmt.Println("Erreur lors du décodage JSON:", err)
		return
	}

	ClearScreen()
	fmt.Println(Blue, "      SAUVEGARDE", Reset)
	fmt.Println("_________________________")
	var table []int
	var choixSauvegarde int
	for k := 1; k < len(partie)+1; k++ {
		fmt.Println(k, "- ", partie[k].Name)
		table = append(table, k)
	}
	fmt.Scanln(&choixSauvegarde)
	for L := 1; L < len(table)+1; L++ {
		if choixSauvegarde == table[L] {

		}
	}
}
