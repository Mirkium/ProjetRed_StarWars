package createperso

import (
	"fmt"
	models "game/models"
	utils "game/utils"
)

var newPerso models.Perso

var (
	choixCampagne = []string{"RePUBLIC", "EMPIRE", "QUIT"}
	indexCampagne = 0
)

func printCampagne() {
	fmt.Println(models.Cyan, "                                      .........", models.Red, " ..... ...")
	fmt.Println(models.Cyan, "                              ....-=*#%%%%%%%%%", models.Red, " %%%%%##*+=:....")
	fmt.Println(models.Cyan, "                            ..+%%%%%%%%%%%%%%%#", models.Red, " %%%%%%%%%%%%##=...")
	fmt.Println(models.Cyan, "                         .-*%%%%%%%#*+=:.... ..", models.Red, " .....:-=+*#%%%%%#*-.. ")
	fmt.Println(models.Cyan, "                      .:#%%%%%%#=......       .", models.Red, " ##=..    ...:+#%%%%#*:.")
	fmt.Println(models.Cyan, "                   ..-#%%%%%*-...-=-...       .", models.Red, " #%%%#=..       .=#%%%%#-..")
	fmt.Println(models.Cyan, "                  .-%%%%%#-...*#+.....        .", models.Red, " #%%%#%%#=..     ..:+%%###-.")
	fmt.Println(models.Cyan, "                .-#%%%%*:..-#%+....           .", models.Red, " #%%%=.-*%%=.       ..=#####-.")
	fmt.Println(models.Cyan, "               .*%%%%#:..-#%#:..              .", models.Red, " #%%#-......          ..=####+.")
	fmt.Println(models.Cyan, "             .:#%%%%=...*%%#..         ...   ..", models.Red, " #%%#:           .::.   .:*###*:.")
	fmt.Println(models.Cyan, "            .-%%%%#....#%%%..          .-.   =:", models.Red, " #%%#..          :###*-....=###*-.")
	fmt.Println(models.Cyan, "           .=%%%%+.  .*%%%+          .:#:-==+#:", models.Red, " #%%+..           ..=####=:.-***#=.")
	fmt.Println(models.Cyan, "          .-%%%%+.. .+%%%%=          :##:-+#%#:", models.Red, " ###=.             ..=*#####*=#***-.")
	fmt.Println(models.Cyan, "          .#%%%+..=.:#%%%%=         .+%%=....+:", models.Red, " ###:.          ..:+##******#-:#***:")
	fmt.Println(models.Cyan, "         .#%%%*..==.=%%%%%#.        .#%%%*.. ..", models.Red, " ###.         ..=*##********#:.=****.")
	fmt.Println(models.Cyan, "        .=%%%%-.:%-.+%%%%%%*.       .:*#%%%#-..", models.Red, " ####*:..  ..:+#####********#: .****=.")
	fmt.Println(models.Cyan, "        .*%%%*..*%:.*%%%%%%%*.         .+%%%%-.", models.Red, " #######*-.=*######****#*-=##: .-#***: ")
	fmt.Println(models.Cyan, "        -%%%%:.-%%:.*%%%%%%%%%:..        -%%%-.", models.Red, " ##################**-....-#*:  .****-.")
	fmt.Println(models.Cyan, "       .+%%%#..+%%=.+%%%%%%%%###-.        =%%-.", models.Red, " ###############+-.. ..  .-#*:   -#**=.")
	fmt.Println(models.Cyan, "       .*%%%+..#%%#.-%%%%%%%######-.      :*%-.", models.Red, " ############:...         ....   :#**+.")
	fmt.Println(models.Cyan, "       .#%%%=..#%%%-.*%%%%%%#######=      .*%-.", models.Red, " ############:.                  :****.")
	fmt.Println(models.Cyan, "       .#%%%=..*%%%#:.#%%%%%%#######      .+%-.", models.Red, " ############:.                  .*##*.")
	fmt.Println(models.Cyan, "       .#%%%=..=%%%%#..#%%%%%%######      .+%-.", models.Red, " ############:.                  :*##*.")
	fmt.Println(models.Cyan, "        *%%%+. .%%%%%%:.*%%%%%%#####.     .+%-.", models.Red, " ############=..          .::.   :###+.")
	fmt.Println(models.Cyan, "        =%%%#.  -%%%%%%+.-#%%%%%####      .+%-.", models.Red, " ###############*=:...    -#*:   =###=.")
	fmt.Println(models.Cyan, "        -%%%%-  .+%%%%%%%*..#%%%%###      .+#-.", models.Red, " ####################+..  :##:  .####: ")
	fmt.Println(models.Cyan, "        .*%%%#. ..*%%%%%%%%%*=-=+*##.     .*%-.", models.Red, " #######=:.:*############==##: .=###*. ")
	fmt.Println(models.Cyan, "        .-%%%%=....=%%%%%%%%%%%%%#=..     .*%-.", models.Red, " ####=...   ..=##############: :####-..")
	fmt.Println(models.Cyan, "         .*%%%#:.:+::#%%%%%%%%%%%%%#.     :*#-.", models.Red, " ###.           :*######%%%%#:.+%%%+.. ")
	fmt.Println(models.Cyan, "          .#%%%#...#*.:*%%%%%%%%%%##*.    :##-.", models.Red, " ###:           ...=#%%%%%%%%-=%%%*... ")
	fmt.Println(models.Cyan, "          .:#%%%#:..*%#=.:+*#%%%%%#*+.    =%%-.", models.Red, " ###=              ..:*%%%%#++%%%#:.   ")
	fmt.Println(models.Cyan, "           .-%%%%#:..+%%%##**++**###*:.   *#%-.", models.Red, " ###+            ..:*#%%*-..+%%%#:.    ")
	fmt.Println(models.Cyan, "            .:#%%%%=..:#%%%%%%%#######+. -%#%-.", models.Red, " ####.           :#%#+:...:#%%%#:.     ")
	fmt.Println(models.Cyan, "             ..#%%%%#...:#%%######%####*.###%-.", models.Red, " ####:           ..... ..=#%%%*..      ")
	fmt.Println(models.Cyan, "               .=%%%%%+....:=++==#######%###%-.", models.Red, " ####- ...-..         .-#%%%#=...      ")
	fmt.Println(models.Cyan, "                .:#%####*.  .*#############%%-.", models.Red, " ####=.+###=.       .-#%%%%*:.         ")
	fmt.Println(models.Cyan, "                  .:#######:...-*##########%%-.", models.Red, " #######*-...    .:+#%%%%*:..          ")
	fmt.Println(models.Cyan, "                   ..:*%#####*=....-+#%##%%%%-.", models.Red, " ####*-...    .:+#%%%%#+:..            ")
	fmt.Println(models.Cyan, "                      ..+#######%*-:.......:-..", models.Red, " #*-.......:=#%%%%%%#=....             ")
	fmt.Println(models.Cyan, "                         .:=###########*++=-:::", models.Red, " ::-==+*##%%%%%%%#=:. ..               ")
	fmt.Println(models.Cyan, "                            ..:*#######%%%%%%##", models.Red, " ########%%%##+:...                    ")
	fmt.Println(models.Cyan, "                               ....-=+#%%%%%%##", models.Red, " ######*+-:....                        ")
	fmt.Println(models.Cyan, "                                        .......", models.Red, " ....", models.Reset)
}
func printHorizontalMenu() {
	fmt.Print("\n\n                                ")

	for i, opt := range choixCampagne {
		if i == indexCampagne {

			fmt.Print(models.Yellow, opt, models.Reset)
		} else {

			fmt.Print(models.Cyan, opt, models.Reset)
		}
		fmt.Print("   ")
	}
}

func CreatePerso() {
	var pseudo string
	for {
		models.ClearScreen()
		printCampagne()

		printHorizontalMenu()

		key := utils.HorizontaleKey()
		switch key {
		case "left":
			if indexCampagne > 0 {
				indexCampagne--
			}
		case "right":
			if indexCampagne < len(choixCampagne)-1 {
				indexCampagne++
			}
		case "enter":
			models.ClearScreen()
			switch indexCampagne {
			case 0:
				fmt.Println(models.Cyan, "You have choice Repulic Galactic")
				fmt.Print("Enter your Name : ", models.Yellow)
				fmt.Scanln(&pseudo)
				newPerso.Name = pseudo
				newPerso.Campagne.Name = "Republic Galactic - prologue"
				newPerso.Campagne.Nb_maxMission = 10
				newPerso.Campagne.ActualMission = 1
				newPerso.Campagne.Way = "Republic"
				ChooseWeapon()
				return
			case 1:
				fmt.Println(models.Cyan, "You have choice Empire")
				fmt.Print("Enter your Name : ", models.Yellow)
				fmt.Scanln(&pseudo)
				newPerso.Name = pseudo
				newPerso.Campagne.Name = "Empire - prologue"
				newPerso.Campagne.Nb_maxMission = 10
				newPerso.Campagne.ActualMission = 1
				newPerso.Campagne.Way = "Empire"
				ChooseWeapon()
				return
			case 2:
				fmt.Println(models.Cyan, "Vous quittez le jeu", models.Reset)
				return
			}
		}
	}
}
