package jeux

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

//music non fonctionnel
//music non fonctionnel
//music non fonctionnel
//music non fonctionnel
//music non fonctionnel
//music non fonctionnel
//music non fonctionnel
//music non fonctionnel
//music non fonctionnel
//music non fonctionnel
//music non fonctionnel
//music non fonctionnel
//music non fonctionnel
//music non fonctionnel
//music non fonctionnel
//music non fonctionnel
//music non fonctionnel
//music non fonctionnel
//music non fonctionnel
//music non fonctionnel
//music non fonctionnel
//music non fonctionnel
//music non fonctionnel
//music non fonctionnel
//music non fonctionnel

var cmd *exec.Cmd

func PlayMusic(track int) {
	var fileName string
	path, _ := os.Getwd()
	switch track {
	case 1: //batle of the heroes
		fileName = filepath.Join(path, "Jeux", "Music", "John Williams - Battle of the Heroes (Audio).mp3")
	case 2: // Duel of the Fates
		fileName = filepath.Join(path, "Jeux", "Music", "John Williams - Duel of the Fates (Star Wars Soundtrack) [HQ].mp3")
	case 3: //The Imperial Suite
		fileName = filepath.Join(path, "Jeux", "Music", "Michael Giacchino - The Imperial Suite (From Rogue One A Star Wars StoryAudio Only).mp3")
	case 4: //Dark maul theme
		fileName = filepath.Join(path, "Jeux", "Music", "Star Wars Battlefront II Soundtrack - Darth Maul Theme.mp3")
	case 5: //Intro
		fileName = filepath.Join(path, "Jeux", "Music", "Star Wars Intro HD 1080p.mp3")
	case 6: //Anakin vs Obi wan
		fileName = filepath.Join(path, "Jeux", "Music", "Star Wars Revenge of the Sith Soundtrack  Anakin vs Obi-Wan, the great duel.mp3")
	case 7: //Imperial march
		fileName = filepath.Join(path, "Jeux", "Music", "Star Wars- The Imperial March (Darth Vaders Theme).mp3")
	default:
		fmt.Println("Numéro de piste non valide")
		return
	}
	cmd := exec.Command("ffplay", "-nodisp", fileName)
	err := cmd.Start()
	if err != nil {
		return
	}
}

func StopMusic() error {
	if cmd != nil && cmd.Process != nil {
		err := cmd.Process.Kill()
		if err != nil {
			return err
		}
		fmt.Println("Musique arrêtée")
		cmd = nil
	}
	return nil
}
