package chevalierjedi

import (
	Fight "Game/Jeux/GamePlay/Fight"
	save "Game/Jeux/Sauvegarde"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

const (
	Reset   = "\033[0m"
	Red     = "\033[91m"
	Green   = "\033[32m"
	Yellow  = "\033[93m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[96m"
	White   = "\033[37m"
	gray    = "\033[90m"
)

func JediKnight() {
	ClearScreen()
	fmt.Println("                      ", Cyan, "LONG TIME AGO IN A GALAXY FAR,")
	fmt.Println("                      FAR AWAY...", Reset)
	time.Sleep(time.Second * 2)
	ClearScreen()
	fmt.Println(Yellow, "      ▄▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀█     █▀▀▀▀▀█        █▀▀▀▀▀▀▀▀▀▄        ")
	fmt.Println("     █                            █    █   ▄   █       █   █▀▀▄   █        ")
	fmt.Println("      ▀▄    ▀█▀▀▀▀▀▀▀▀▀█   █▀▀▀▀▀▀▀   █   █ █   █      █   ▀▀▀  ▄▀         ")
	fmt.Println("        ▀▄    ▀▄       █   █         █   █▄▄▄█   █     █   ▄     ▀▄        ")
	fmt.Println(" ▄▄▄▄▄▄▄▄▄█▄    ▀▄     █   █        █    ▄▄▄▄▄    █    █   █▀▄     ▀▄▄▄▄▄▄▄▄▄▄")
	fmt.Println(" █               █     █   █       █    █     █    █   █   █  ▀▄             █")
	fmt.Println(" █▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▀     █▄▄▄█      █▄▄▄▄█       █▄▄▄▄█  █▄▄▄█    ▀▄▄▄▄▄▄▄▄▄▄▄▄█")
	fmt.Println("")
	fmt.Println("█▀▀▀▀█     █▀▀▀█     █▀▀▀▀█   █▀▀▀▀▀█        █▀▀▀▀▀▀▀▀▀▄      ▄▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀█")
	fmt.Println(" █    █   █     █   █    █   █   ▄   █       █   █▀▀▄   █    █                █")
	fmt.Println("  █    █ █       █ █    █   █   █ █   █      █   ▀▀▀  ▄▀      ▀▄    ▀█▀▀▀▀▀▀▀▀▀")
	fmt.Println("   █    ▀    ▄    ▀    █   █   █▄▄▄█   █     █   ▄     ▀▄       ▀▄    ▀▄")
	fmt.Println("    █       █ █       █   █    ▄▄▄▄▄    █    █   █▀▄     ▀▄▄▄▄▄▄▄▄█▄    ▀▄")
	fmt.Println("     █     █   █     █   █    █     █    █   █   █  ▀▄                    █")
	fmt.Println("      █▄▄▄█     █▄▄▄█   █▄▄▄▄█       █▄▄▄▄█  █▄▄▄█    ▀▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▀")
	time.Sleep(time.Second * 5)
	ClearScreen()
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("                                  PROJET RED ")
	fmt.Println("")
	fmt.Println("                                 SITH WARRIOR")
	time.Sleep(time.Second * 2)
	ClearScreen()
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("                                  PROJET RED ")
	fmt.Println("")
	fmt.Println("                                 SITH WARRIOR")
	fmt.Println("")
	time.Sleep(time.Second * 2)
	ClearScreen()
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("                                  PROJET RED ")
	fmt.Println("")
	fmt.Println("                                 SITH WARRIOR")
	fmt.Println("")
	fmt.Println("                           It is a dark era for the")
	time.Sleep(time.Second * 2)
	ClearScreen()
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("                                  PROJET RED ")
	fmt.Println("")
	fmt.Println("                                 SITH WARRIOR")
	fmt.Println("")
	fmt.Println("                           It is a dark  era  for  the")
	fmt.Println("                           JEDI ORDER. The Sith Empire")
	time.Sleep(time.Second * 2)
	ClearScreen()
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("                                  PROJET RED ")
	fmt.Println("")
	fmt.Println("                                 SITH WARRIOR")
	fmt.Println("")
	fmt.Println("                           It is a dark  era  for  the")
	fmt.Println("                           JEDI ORDER. The Sith Empire")
	fmt.Println("                           obliterared the Jedi Temple")
	time.Sleep(time.Second * 2)
	ClearScreen()
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("                                  PROJET RED ")
	fmt.Println("")
	fmt.Println("                                 SITH WARRIOR")
	fmt.Println("")
	fmt.Println("                           It is a dark  era  for  the")
	fmt.Println("                           JEDI ORDER. The Sith Empire")
	fmt.Println("                           obliterared the Jedi Temple")
	fmt.Println("                           on Coruscant and slaughtered")
	time.Sleep(time.Second * 2)
	ClearScreen()
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("                                  PROJET RED ")
	fmt.Println("")
	fmt.Println("                                 SITH WARRIOR")
	fmt.Println("")
	fmt.Println("                           It is a dark  era  for  the")
	fmt.Println("                           JEDI ORDER. The Sith Empire")
	fmt.Println("                           obliterared the Jedi Temple")
	fmt.Println("                           on Coruscant and slaughtered")
	fmt.Println("                           many of the Republic's brave")
	time.Sleep(time.Second * 2)
	ClearScreen()
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("                                  PROJET RED ")
	fmt.Println("")
	fmt.Println("                                 SITH WARRIOR")
	fmt.Println("")
	fmt.Println("                          It is a dark era for the")
	fmt.Println("                          JEDI ORDER. The Sith Empire")
	fmt.Println("                          obliterared the Jedi Temple")
	fmt.Println("                          on Coruscant and slaughtered")
	fmt.Println("                          many of the Republic's brave")
	fmt.Println("                          defenders during the last war.")
	time.Sleep(time.Second * 2)
	ClearScreen()
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("                                  PROJET RED ")
	fmt.Println("")
	fmt.Println("                                 SITH WARRIOR")
	fmt.Println("")
	fmt.Println("                          It is a dark era for the")
	fmt.Println("                          JEDI ORDER. The Sith Empire")
	fmt.Println("                          obliterared the Jedi Temple")
	fmt.Println("                          on Coruscant and slaughtered")
	fmt.Println("                          many of the Republic's brave")
	fmt.Println("                          defenders during the last war.")
	fmt.Println("")
	time.Sleep(time.Second * 2)
	ClearScreen()
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("                                  PROJET RED ")
	fmt.Println("")
	fmt.Println("                                 SITH WARRIOR")
	fmt.Println("")
	fmt.Println("                          It is a dark era for the")
	fmt.Println("                          JEDI ORDER. The Sith Empire")
	fmt.Println("                          obliterared the Jedi Temple")
	fmt.Println("                          on Coruscant and slaughtered")
	fmt.Println("                          many of the Republic's brave")
	fmt.Println("                          defenders during the last war.")
	fmt.Println("")
	fmt.Println("                          The surviving Jedi have")
	time.Sleep(time.Second * 2)
	ClearScreen()
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("                                  PROJET RED ")
	fmt.Println("")
	fmt.Println("                                 SITH WARRIOR")
	fmt.Println("")
	fmt.Println("                          It is a dark era for the")
	fmt.Println("                          JEDI ORDER. The Sith Empire")
	fmt.Println("                          obliterared the Jedi Temple")
	fmt.Println("                          on Coruscant and slaughtered")
	fmt.Println("                          many of the Republic's brave")
	fmt.Println("                          defenders during the last war.")
	fmt.Println("")
	fmt.Println("                          The surviving Jedi have")
	fmt.Println("                          withdrawn to their ancient")
	time.Sleep(time.Second * 2)
	ClearScreen()
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("                                  PROJET RED ")
	fmt.Println("")
	fmt.Println("                                 SITH WARRIOR")
	fmt.Println("")
	fmt.Println("                          It is a dark era for the")
	fmt.Println("                          JEDI ORDER. The Sith Empire")
	fmt.Println("                          obliterared the Jedi Temple")
	fmt.Println("                          on Coruscant and slaughtered")
	fmt.Println("                          many of the Republic's brave")
	fmt.Println("                          defenders during the last war.")
	fmt.Println("")
	fmt.Println("                          The surviving Jedi have")
	fmt.Println("                          withdrawn to their ancient")
	fmt.Println("                          homeworld of TYTHON, where")
	time.Sleep(time.Second * 2)
	ClearScreen()
	fmt.Println("")
	fmt.Println("")
	fmt.Println("                                  PROJET RED ")
	fmt.Println("")
	fmt.Println("                                 SITH WARRIOR")
	fmt.Println("")
	fmt.Println("                          It is a dark era for the")
	fmt.Println("                          JEDI ORDER. The Sith Empire")
	fmt.Println("                          obliterared the Jedi Temple")
	fmt.Println("                          on Coruscant and slaughtered")
	fmt.Println("                          many of the Republic's brave")
	fmt.Println("                          defenders during the last war.")
	fmt.Println("")
	fmt.Println("                          The surviving Jedi have")
	fmt.Println("                          withdrawn to their ancient")
	fmt.Println("                          homeworld of TYTHON, where")
	fmt.Println("                          they take advantage of a")
	time.Sleep(time.Second * 2)
	ClearScreen()
	fmt.Println("")
	fmt.Println("                                  PROJET RED ")
	fmt.Println("")
	fmt.Println("                                 SITH WARRIOR")
	fmt.Println("")
	fmt.Println("                          It is a dark era for the")
	fmt.Println("                          JEDI ORDER. The Sith Empire")
	fmt.Println("                          obliterared the Jedi Temple")
	fmt.Println("                          on Coruscant and slaughtered")
	fmt.Println("                          many of the Republic's brave")
	fmt.Println("                          defenders during the last war.")
	fmt.Println("")
	fmt.Println("                          The surviving Jedi have")
	fmt.Println("                          withdrawn to their ancient")
	fmt.Println("                          homeworld of TYTHON, where")
	fmt.Println("                          they take advantage of a")
	fmt.Println("                          fragile peace to train a new")
	time.Sleep(time.Second * 2)
	ClearScreen()
	fmt.Println("                                  PROJET RED ")
	fmt.Println("")
	fmt.Println("                                 SITH WARRIOR")
	fmt.Println("")
	fmt.Println("                          It is a dark era for the")
	fmt.Println("                          JEDI ORDER. The Sith Empire")
	fmt.Println("                          obliterared the Jedi Temple")
	fmt.Println("                          on Coruscant and slaughtered")
	fmt.Println("                          many of the Republic's brave")
	fmt.Println("                          defenders during the last war.")
	fmt.Println("")
	fmt.Println("                          The surviving Jedi have")
	fmt.Println("                          withdrawn to their ancient")
	fmt.Println("                          homeworld of TYTHON, where")
	fmt.Println("                          they take advantage of a")
	fmt.Println("                          fragile peace to train a new")
	fmt.Println("                          generation of gardians for the")
	time.Sleep(time.Second * 2)
	ClearScreen()
	fmt.Println(" ")
	fmt.Println("                                 SITH WARRIOR")
	fmt.Println("")
	fmt.Println("                          It is a dark era for the")
	fmt.Println("                          JEDI ORDER. The Sith Empire")
	fmt.Println("                          obliterared the Jedi Temple")
	fmt.Println("                          on Coruscant and slaughtered")
	fmt.Println("                          many of the Republic's brave")
	fmt.Println("                          defenders during the last war.")
	fmt.Println("")
	fmt.Println("                          The surviving Jedi have")
	fmt.Println("                          withdrawn to their ancient")
	fmt.Println("                          homeworld of TYTHON, where")
	fmt.Println("                          they take advantage of a")
	fmt.Println("                          fragile peace to train a new")
	fmt.Println("                          generation of gardians for the")
	fmt.Println("                          galaxy.")
	time.Sleep(time.Second * 2)
	ClearScreen()
	fmt.Println("                                 SITH WARRIOR")
	fmt.Println("")
	fmt.Println("                          It is a dark era for the")
	fmt.Println("                          JEDI ORDER. The Sith Empire")
	fmt.Println("                          obliterared the Jedi Temple")
	fmt.Println("                          on Coruscant and slaughtered")
	fmt.Println("                          many of the Republic's brave")
	fmt.Println("                          defenders during the last war.")
	fmt.Println("")
	fmt.Println("                          The surviving Jedi have")
	fmt.Println("                          withdrawn to their ancient")
	fmt.Println("                          homeworld of TYTHON, where")
	fmt.Println("                          they take advantage of a")
	fmt.Println("                          fragile peace to train a new")
	fmt.Println("                          generation of gardians for the")
	fmt.Println("                          galaxy.")
	fmt.Println(" ")
	time.Sleep(time.Second * 2)
	ClearScreen()
	fmt.Println("   ")
	fmt.Println("                          It is a dark era for the")
	fmt.Println("                          JEDI ORDER. The Sith Empire")
	fmt.Println("                          obliterared the Jedi Temple")
	fmt.Println("                          on Coruscant and slaughtered")
	fmt.Println("                          many of the Republic's brave")
	fmt.Println("                          defenders during the last war.")
	fmt.Println("")
	fmt.Println("                          The surviving Jedi have")
	fmt.Println("                          withdrawn to their ancient")
	fmt.Println("                          homeworld of TYTHON, where")
	fmt.Println("                          they take advantage of a")
	fmt.Println("                          fragile peace to train a new")
	fmt.Println("                          generation of gardians for the")
	fmt.Println("                          galaxy.")
	fmt.Println(" ")
	fmt.Println("                          Now a new hope emerges. A")
	time.Sleep(time.Second * 2)
	ClearScreen()
	fmt.Println("                          It is a dark era for the")
	fmt.Println("                          JEDI ORDER. The Sith Empire")
	fmt.Println("                          obliterared the Jedi Temple")
	fmt.Println("                          on Coruscant and slaughtered")
	fmt.Println("                          many of the Republic's brave")
	fmt.Println("                          defenders during the last war.")
	fmt.Println("")
	fmt.Println("                          The surviving Jedi have")
	fmt.Println("                          withdrawn to their ancient")
	fmt.Println("                          homeworld of TYTHON, where")
	fmt.Println("                          they take advantage of a")
	fmt.Println("                          fragile peace to train a new")
	fmt.Println("                          generation of gardians for the")
	fmt.Println("                          galaxy.")
	fmt.Println(" ")
	fmt.Println("                          Now a new hope emerges. A")
	fmt.Println("                          young Padawan strong in")
	time.Sleep(time.Second * 2)
	ClearScreen()
	fmt.Println("                          JEDI ORDER. The Sith Empire")
	fmt.Println("                          obliterared the Jedi Temple")
	fmt.Println("                          on Coruscant and slaughtered")
	fmt.Println("                          many of the Republic's brave")
	fmt.Println("                          defenders during the last war.")
	fmt.Println("")
	fmt.Println("                          The surviving Jedi have")
	fmt.Println("                          withdrawn to their ancient")
	fmt.Println("                          homeworld of TYTHON, where")
	fmt.Println("                          they take advantage of a")
	fmt.Println("                          fragile peace to train a new")
	fmt.Println("                          generation of gardians for the")
	fmt.Println("                          galaxy.")
	fmt.Println(" ")
	fmt.Println("                          Now a new hope emerges. A")
	fmt.Println("                          young Padawan strong in")
	fmt.Println("                          the Force journeys to Tython's")
	time.Sleep(time.Second * 2)
	ClearScreen()
	fmt.Println("                          obliterared the Jedi Temple")
	fmt.Println("                          on Coruscant and slaughtered")
	fmt.Println("                          many of the Republic's brave")
	fmt.Println("                          defenders during the last war.")
	fmt.Println("")
	fmt.Println("                          The surviving Jedi have")
	fmt.Println("                          withdrawn to their ancient")
	fmt.Println("                          homeworld of TYTHON, where")
	fmt.Println("                          they take advantage of a")
	fmt.Println("                          fragile peace to train a new")
	fmt.Println("                          generation of gardians for the")
	fmt.Println("                          galaxy.")
	fmt.Println(" ")
	fmt.Println("                          Now a new hope emerges. A")
	fmt.Println("                          young Padawan strong in")
	fmt.Println("                          the Force journeys to Tython's")
	fmt.Println("                          dangerous wilderness to")
	time.Sleep(time.Second * 2)
	ClearScreen()
	fmt.Println("                          on Coruscant and slaughtered")
	fmt.Println("                          many of the Republic's brave")
	fmt.Println("                          defenders during the last war.")
	fmt.Println("")
	fmt.Println("                          The surviving Jedi have")
	fmt.Println("                          withdrawn to their ancient")
	fmt.Println("                          homeworld of TYTHON, where")
	fmt.Println("                          they take advantage of a")
	fmt.Println("                          fragile peace to train a new")
	fmt.Println("                          generation of gardians for the")
	fmt.Println("                          galaxy.")
	fmt.Println(" ")
	fmt.Println("                          Now a new hope emerges. A")
	fmt.Println("                          young Padawan strong in")
	fmt.Println("                          the Force journeys to Tython's")
	fmt.Println("                          dangerous wilderness to")
	fmt.Println("                          complete the final Jedi trials")
	time.Sleep(time.Second * 2)
	ClearScreen()
	fmt.Println("                          many of the Republic's brave")
	fmt.Println("                          defenders during the last war.")
	fmt.Println("")
	fmt.Println("                          The surviving Jedi have")
	fmt.Println("                          withdrawn to their ancient")
	fmt.Println("                          homeworld of TYTHON, where")
	fmt.Println("                          they take advantage of a")
	fmt.Println("                          fragile peace to train a new")
	fmt.Println("                          generation of gardians for the")
	fmt.Println("                          galaxy.")
	fmt.Println(" ")
	fmt.Println("                          Now a new hope emerges. A")
	fmt.Println("                          young Padawan strong in")
	fmt.Println("                          the Force journeys to Tython's")
	fmt.Println("                          dangerous wilderness to")
	fmt.Println("                          complete the final Jedi trials")
	fmt.Println("                          and become a Knight of the")
	time.Sleep(time.Second * 2)
	ClearScreen()
	fmt.Println("                          defenders during the last war.")
	fmt.Println("")
	fmt.Println("                          The surviving Jedi have")
	fmt.Println("                          withdrawn to their ancient")
	fmt.Println("                          homeworld of TYTHON, where")
	fmt.Println("                          they take advantage of a")
	fmt.Println("                          fragile peace to train a new")
	fmt.Println("                          generation of gardians for the")
	fmt.Println("                          galaxy.")
	fmt.Println(" ")
	fmt.Println("                          Now a new hope emerges. A")
	fmt.Println("                          young Padawan strong in")
	fmt.Println("                          the Force journeys to Tython's")
	fmt.Println("                          dangerous wilderness to")
	fmt.Println("                          complete the final Jedi trials")
	fmt.Println("                          and become a Knight of the")
	fmt.Println("                          Republic....")
	time.Sleep(time.Second * 2)
	ClearScreen()
	fmt.Println("")
	fmt.Println("                          The surviving Jedi have")
	fmt.Println("                          withdrawn to their ancient")
	fmt.Println("                          homeworld of TYTHON, where")
	fmt.Println("                          they take advantage of a")
	fmt.Println("                          fragile peace to train a new")
	fmt.Println("                          generation of gardians for the")
	fmt.Println("                          galaxy.")
	fmt.Println(" ")
	fmt.Println("                          Now a new hope emerges. A")
	fmt.Println("                          young Padawan strong in")
	fmt.Println("                          the Force journeys to Tython's")
	fmt.Println("                          dangerous wilderness to")
	fmt.Println("                          complete the final Jedi trials")
	fmt.Println("                          and become a Knight of the")
	fmt.Println("                          Republic....")
	fmt.Println(" ")
	time.Sleep(time.Second * 2)
	ClearScreen()
	fmt.Println("                          The surviving Jedi have")
	fmt.Println("                          withdrawn to their ancient")
	fmt.Println("                          homeworld of TYTHON, where")
	fmt.Println("                          they take advantage of a")
	fmt.Println("                          fragile peace to train a new")
	fmt.Println("                          generation of gardians for the")
	fmt.Println("                          galaxy.")
	fmt.Println(" ")
	fmt.Println("                          Now a new hope emerges. A")
	fmt.Println("                          young Padawan strong in")
	fmt.Println("                          the Force journeys to Tython's")
	fmt.Println("                          dangerous wilderness to")
	fmt.Println("                          complete the final Jedi trials")
	fmt.Println("                          and become a Knight of the")
	fmt.Println("                          Republic....")
	fmt.Println("")
	fmt.Println(" ")
	time.Sleep(time.Second * 2)
	ClearScreen()
	fmt.Println("                          withdrawn to their ancient")
	fmt.Println("                          homeworld of TYTHON, where")
	fmt.Println("                          they take advantage of a")
	fmt.Println("                          fragile peace to train a new")
	fmt.Println("                          generation of gardians for the")
	fmt.Println("                          galaxy.")
	fmt.Println(" ")
	fmt.Println("                          Now a new hope emerges. A")
	fmt.Println("                          young Padawan strong in")
	fmt.Println("                          the Force journeys to Tython's")
	fmt.Println("                          dangerous wilderness to")
	fmt.Println("                          complete the final Jedi trials")
	fmt.Println("                          and become a Knight of the")
	fmt.Println("                          Republic....")
	fmt.Println("     ")
	fmt.Println("")
	fmt.Println(" ")
	time.Sleep(time.Second * 2)
	ClearScreen()
	fmt.Println("                          homeworld of TYTHON, where")
	fmt.Println("                          they take advantage of a")
	fmt.Println("                          fragile peace to train a new")
	fmt.Println("                          generation of gardians for the")
	fmt.Println("                          galaxy.")
	fmt.Println(" ")
	fmt.Println("                          Now a new hope emerges. A")
	fmt.Println("                          young Padawan strong in")
	fmt.Println("                          the Force journeys to Tython's")
	fmt.Println("                          dangerous wilderness to")
	fmt.Println("                          complete the final Jedi trials")
	fmt.Println("                          and become a Knight of the")
	fmt.Println("                          Republic....")
	fmt.Println("  ")
	fmt.Println("     ")
	fmt.Println("")
	fmt.Println(" ")
	time.Sleep(time.Second * 2)
	ClearScreen()
	fmt.Println("                          they take advantage of a")
	fmt.Println("                          fragile peace to train a new")
	fmt.Println("                          generation of gardians for the")
	fmt.Println("                          galaxy.")
	fmt.Println(" ")
	fmt.Println("                          Now a new hope emerges. A")
	fmt.Println("                          young Padawan strong in")
	fmt.Println("                          the Force journeys to Tython's")
	fmt.Println("                          dangerous wilderness to")
	fmt.Println("                          complete the final Jedi trials")
	fmt.Println("                          and become a Knight of the")
	fmt.Println("                          Republic....")
	fmt.Println(" ")
	fmt.Println("  ")
	fmt.Println("     ")
	fmt.Println("")
	fmt.Println(" ")
	time.Sleep(time.Second * 2)
	ClearScreen()
	fmt.Println("                          fragile peace to train a new")
	fmt.Println("                          generation of gardians for the")
	fmt.Println("                          galaxy.")
	fmt.Println(" ")
	fmt.Println("                          Now a new hope emerges. A")
	fmt.Println("                          young Padawan strong in")
	fmt.Println("                          the Force journeys to Tython's")
	fmt.Println("                          dangerous wilderness to")
	fmt.Println("                          complete the final Jedi trials")
	fmt.Println("                          and become a Knight of the")
	fmt.Println("                          Republic....")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("  ")
	fmt.Println("     ")
	fmt.Println("")
	fmt.Println(" ")
	time.Sleep(time.Second * 2)
	ClearScreen()
	fmt.Println("                          generation of gardians for the")
	fmt.Println("                          galaxy.")
	fmt.Println(" ")
	fmt.Println("                          Now a new hope emerges. A")
	fmt.Println("                          young Padawan strong in")
	fmt.Println("                          the Force journeys to Tython's")
	fmt.Println("                          dangerous wilderness to")
	fmt.Println("                          complete the final Jedi trials")
	fmt.Println("                          and become a Knight of the")
	fmt.Println("                          Republic....")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("  ")
	fmt.Println("     ")
	fmt.Println("")
	fmt.Println(" ")
	time.Sleep(time.Second * 2)
	ClearScreen()
	fmt.Println("                          galaxy.")
	fmt.Println(" ")
	fmt.Println("                          Now a new hope emerges. A")
	fmt.Println("                          young Padawan strong in")
	fmt.Println("                          the Force journeys to Tython's")
	fmt.Println("                          dangerous wilderness to")
	fmt.Println("                          complete the final Jedi trials")
	fmt.Println("                          and become a Knight of the")
	fmt.Println("                          Republic....")
	fmt.Println("  ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("  ")
	fmt.Println("     ")
	fmt.Println("")
	fmt.Println(" ")
	time.Sleep(time.Second * 2)
	ClearScreen()
	fmt.Println(" ")
	fmt.Println("                          Now a new hope emerges. A")
	fmt.Println("                          young Padawan strong in")
	fmt.Println("                          the Force journeys to Tython's")
	fmt.Println("                          dangerous wilderness to")
	fmt.Println("                          complete the final Jedi trials")
	fmt.Println("                          and become a Knight of the")
	fmt.Println("                          Republic....")
	fmt.Println(" ")
	fmt.Println("  ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("  ")
	fmt.Println("     ")
	fmt.Println("")
	fmt.Println(" ")
	time.Sleep(time.Second * 2)
	ClearScreen()
	fmt.Println("                          Now a new hope emerges. A")
	fmt.Println("                          young Padawan strong in")
	fmt.Println("                          the Force journeys to Tython's")
	fmt.Println("                          dangerous wilderness to")
	fmt.Println("                          complete the final Jedi trials")
	fmt.Println("                          and become a Knight of the")
	fmt.Println("                          Republic....")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("  ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("  ")
	fmt.Println("     ")
	fmt.Println("")
	fmt.Println(" ")
	time.Sleep(time.Second * 2)
	ClearScreen()
	fmt.Println("                          young Padawan strong in")
	fmt.Println("                          the Force journeys to Tython's")
	fmt.Println("                          dangerous wilderness to")
	fmt.Println("                          complete the final Jedi trials")
	fmt.Println("                          and become a Knight of the")
	fmt.Println("                          Republic....")
	fmt.Println("  ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("  ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("  ")
	fmt.Println("     ")
	fmt.Println("")
	fmt.Println(" ")
	time.Sleep(time.Second * 2)
	ClearScreen()
	fmt.Println("                          the Force journeys to Tython's")
	fmt.Println("                          dangerous wilderness to")
	fmt.Println("                          complete the final Jedi trials")
	fmt.Println("                          and become a Knight of the")
	fmt.Println("                          Republic....")
	fmt.Println(" ")
	fmt.Println("  ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("  ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("  ")
	fmt.Println("     ")
	fmt.Println("")
	fmt.Println(" ")
	time.Sleep(time.Second * 2)
	ClearScreen()
	fmt.Println("                          dangerous wilderness to")
	fmt.Println("                          complete the final Jedi trials")
	fmt.Println("                          and become a Knight of the")
	fmt.Println("                          Republic....")
	fmt.Println("   ")
	fmt.Println(" ")
	fmt.Println("  ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("  ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("  ")
	fmt.Println("     ")
	fmt.Println("")
	fmt.Println(" ")
	time.Sleep(time.Second * 2)
	ClearScreen()
	fmt.Println("                          complete the final Jedi trials")
	fmt.Println("                          and become a Knight of the")
	fmt.Println("                          Republic....")
	fmt.Println("  ")
	fmt.Println("   ")
	fmt.Println(" ")
	fmt.Println("  ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("  ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("  ")
	fmt.Println("     ")
	fmt.Println("")
	fmt.Println(" ")
	time.Sleep(time.Second * 2)
	ClearScreen()
	fmt.Println("                          and become a Knight of the")
	fmt.Println("                          Republic....")
	fmt.Println("  ")
	fmt.Println("  ")
	fmt.Println("   ")
	fmt.Println(" ")
	fmt.Println("  ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("  ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("  ")
	fmt.Println("     ")
	fmt.Println("")
	fmt.Println(" ")
	time.Sleep(time.Second * 2)
	ClearScreen()
	fmt.Println("                          Republic....", Reset)
	fmt.Println(" ")
	fmt.Println("  ")
	fmt.Println("  ")
	fmt.Println("   ")
	fmt.Println(" ")
	fmt.Println("  ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("  ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("  ")
	fmt.Println("     ")
	fmt.Println("")
	fmt.Println(" ")
	time.Sleep(time.Second * 2)
	ClearScreen()
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

func CampaingJediKnight(MC *save.Perso) {
	//fonction a appeller apres l'intro
	//Prologue
	fmt.Println("Clone : Are you okay general?")
	time.Sleep(3 * time.Second)
	fmt.Println("\nYou : Yes, I fell asleep a bit")
	time.Sleep(3 * time.Second)
	fmt.Println("\nClone : Need I remind you of the current mission?")
	time.Sleep(3 * time.Second)
	fmt.Println("\nYou : Yes please")
	time.Sleep(3 * time.Second)
	fmt.Println("\n Clone : Okay")
	save.ClearScreen()
	time.Sleep(3 * time.Second)
	fmt.Println("Clone : We are on Coruscant, on the floor minus 32.")
	time.Sleep(3 * time.Second)
	save.ClearScreen()
	fmt.Println("Clone : Our mission is to bring criminals who are in relation with the head of black march.")
	time.Sleep(3 * time.Second)
	save.ClearScreen()
	fmt.Println("Clone : WARNING")
	time.Sleep(3 * time.Second)
	save.ClearScreen()
	fmt.Println("Unknown : You are lucky " + MC.Name + ".")
	time.Sleep(2 * time.Second)
	fmt.Println("\nYou : Oh, you think, so we will see.")
	time.Sleep(4 * time.Second)
	save.ClearScreen()
	fmt.Println("You decide to fight this unknown enemy.")
	unknown := save.Mob{"brigands", 100, 100, 0, []save.Abilite{save.Abilite{"punch", 0, 10, 0, 0, 0, 0, "Just a punch"}}, map[save.Item]int{}, 100}
	Fight.Fight(MC, &unknown, true)
}
