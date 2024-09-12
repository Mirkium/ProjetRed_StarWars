package main

import (
	save "Game/Jeux/Sauvegarde"
	"fmt"
	"math/rand"
	"time"
)

func ChangeLevel(P *save.Perso, mobKilled save.Mob) {
	P.Xp_Actu += mobKilled.XpDrop
	if P.Xp_Actu > 1000*P.Level {
		P.Xp_Actu -= 1000 * P.Level
		P.Level++
	}
}

func Fight(P *save.Perso, mob *save.Mob) {
	rand.Seed(time.Now().UnixNano())
	randomInt := rand.Intn(len(mob.Abilitie))
	fmt.Println(randomInt)
}
