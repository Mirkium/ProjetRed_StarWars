package models

import (
	"os"
	"os/exec"
	"runtime"
)

const (
	Reset   = "\033[0m"
	Red     = "\033[91m"
	Green   = "\033[32m"
	Yellow  = "\033[93m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	Gray    = "\033[90m"
	White   = "\033[37m"
)

func ClearScreen() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
