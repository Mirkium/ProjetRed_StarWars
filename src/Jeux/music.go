package jeux
/*
import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	//"github.com/faiface/beep"
	//"github.com/faiface/beep/mp3"
	//"github.com/faiface/beep/speaker"
)

// Variable globale pour contrôler la musique
//var control *beep.Ctrl
var done chan bool

// Fonction pour lancer la musique
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
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier:", err)
		return
	}
	defer f.Close()
	//streamer, format, err := mp3.Decode(f)
	if err != nil {
		fmt.Println("Erreur lors du décodage:", err)
		return
	}
	defer streamer.Close()
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	control = &beep.Ctrl{Streamer: streamer, Paused: false}
	done = make(chan bool)
	speaker.Play(beep.Seq(control, beep.Callback(func() {
		done <- true
	})))
	CreatePerso()
}
func StopMusic() {
	if control != nil {
		speaker.Lock()
		control.Paused = true
		speaker.Unlock()
	}
}
*/
