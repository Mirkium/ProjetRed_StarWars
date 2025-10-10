package createperso

import (
	"fmt"
	models "game/models"
	utils "game/utils"
)

var (
	listStyle  = []string{"Single LightSaber", "Two LightSaber", "Double LigthSaber"}
	indexStyle = 0
)

func ChooseWeapon() {
	for {
		models.ClearScreen()

		printChooseWeapon()

		key := utils.VerticalKey()
		switch key {
		case "up":
			if indexStyle > 0 {
				indexStyle--
			}
		case "down":
			if indexStyle < len(listStyle)-1 {
				indexStyle++
			}
		case "enter":
			models.ClearScreen()
			switch indexStyle {
			case 0:
				switch newPerso.Campagne.Way {
				case "Republic":
					newPerso.Weapon.Name = "Training Lightsaber"
					newPerso.Weapon.Cristal = models.Yellow
				case "Empire":
					newPerso.Weapon.Name = "Training Lightsaber"
					newPerso.Weapon.Cristal = models.Yellow
				}
				fmt.Println(newPerso, models.Reset)
				return
			case 1:
				switch newPerso.Campagne.Way {
				case "Republic":
					newPerso.Weapon.Name = "Training Lightsaber"
					newPerso.SecondWeapon.Name = "Training Lightsaber"
					newPerso.Weapon.Cristal = models.Yellow
					newPerso.SecondWeapon.Cristal = models.Yellow
				case "Empire":
					newPerso.Weapon.Name = "Training Lightsaber"
					newPerso.SecondWeapon.Name = "Training Lightsaber"
					newPerso.Weapon.Cristal = models.Yellow
					newPerso.SecondWeapon.Cristal = models.Yellow
				}
				fmt.Println(newPerso, models.Reset)
				return
			case 2:
				switch newPerso.Campagne.Way {
				case "Republic":
					newPerso.Weapon.Name = "Training Double-Bladed Lightsaber"
					newPerso.Weapon.Cristal = models.Yellow
				case "Empire":
					newPerso.Weapon.Name = "Training Double-Bladed Lightsaber"
					newPerso.Weapon.Cristal = models.Yellow
				}
				fmt.Println(newPerso, models.Reset)
				return
			}
		}
	}
}

func printChooseWeapon() {

	for i, opt := range listStyle {
		if i == indexStyle {
			switch i {
			case 0:
				fmt.Println(models.Red, opt, models.Reset, "\n")
				fmt.Println(models.Gray, "[▤▤▤▤▥{", models.Red, "■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■ \n", models.Reset)
			case 1:
				fmt.Println(models.Red, opt, models.Reset, "\n")
				fmt.Println(models.Gray, "[▤▤▤▤▥{", models.Red, "■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■ \n", models.Reset)
				fmt.Println(models.Red, "■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■", models.Gray, "}▥▤▤▤▤]", models.Reset)
			case 2:
				fmt.Println(models.Red, opt, models.Reset, "\n")
				fmt.Println(models.Red, "■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■", models.Gray, "}▥▤▤▤▤▤▤▤▥{", models.Red, "■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■ \n", models.Reset)
			}
		} else {
			switch i {
			case 0:
				fmt.Println(models.Yellow, opt, models.Reset, "\n")
				fmt.Println(models.Gray, "[▤▤▤▤▥{ \n", models.Reset)
			case 1:
				fmt.Println(models.Yellow, opt, models.Reset, "\n")
				fmt.Println(models.Gray, "[▤▤▤▤▥{ \n", models.Reset)
				fmt.Println(models.Red, "                                        ", models.Gray, "}▥▤▤▤▤]", models.Reset)
			case 2:
				fmt.Println(models.Yellow, opt, models.Reset, "\n")
				fmt.Println("                                         ", models.Gray, "}▥▤▤▤▤▤▤▤▥{ \n", models.Reset)
			}
		}
	}
	fmt.Print("   ")
}
