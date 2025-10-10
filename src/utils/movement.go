package utils

import (
	"os"
	"golang.org/x/term"
)

func HorizontaleKey() string {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	var buf [3]byte
	os.Stdin.Read(buf[:])

	if buf[0] == 27 && buf[1] == 91 {
		switch buf[2] {
		case 68:
			return "left" // flèche gauche
		case 67:
			return "right" // flèche droite
		}
	}
	if buf[0] == 13 {
		return "enter"
	}
	return ""
}


func VerticalKey() string {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	var buf [3]byte
	os.Stdin.Read(buf[:])

	if buf[0] == 27 && buf[1] == 91 {
		switch buf[2] {
		case 65:
			return "up"
		case 66:
			return "down"
		}
	}
	if buf[0] == 13 {
		return "enter"
	}
	return ""
}