package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func (w *hangManData) startGame() {
	clear()
	fmt.Print("  ______ _     __                                         _             _              _            __  \n")
	fmt.Print(" |  ____| |   / _|                                       (_)           | |            | |           \\_\\ \n")
	fmt.Print(" | |__  | |  | |_ __ _ _ __ ___   ___  ___  ___    __   ___  ___     __| | ___        | | ___  ___  ___ \n")
	fmt.Print(" |  __| | |  |  _/ _` | '_ ` _ \\ / _ \\/ __|/ _ \\   \\ \\ / / |/ _ \\   / _` |/ _ \\   _   | |/ _ \\/ __|/ _ \\\n")
	fmt.Print(" | |____| |  | || (_| | | | | | | (_) \\__ \\ (_) |   \\ V /| |  __/  | (_| |  __/  | |__| | (_) \\__ \\  __/\n")
	fmt.Println(" |______|_|  |_| \\__,_|_| |_| |_|\\___/|___/\\___/     \\_/ |_|\\___|   \\__,_|\\___|   \\____/ \\___/|___/\\___|\n")
	file := "words.txt"
	word := Capitalize(readLine(file, Counter(file)))
	w.Init(hideWord(word), word, 0, []string{}, []string{})
	reveal := len(w.ToFind)/2-1
	w.HiddenWord[reveal] = string(w.ToFind[reveal])
	//w.Tried = append(w.Tried, w.HiddenWord[reveal])
	w.inputLetter()
}


func (w *hangManData) endGame() {
	notHidden := strings.Join(w.HiddenWord[:], "")
	if notHidden == w.ToFind {
		fmt.Printf("\nTU AS SAUVE JOSE! Il retourne désormais à sa vie de pauvre.\nPS : Le mot était %v\n", w.ToFind)
		w.restart()
	} else {
		w.inputLetter()
	}
}


func (w *hangManData) restart() {
	fmt.Println("\nAppuie sur 1 pour recommencer, ou sur 2 pour quitter.")

	INtype := bufio.NewScanner(os.Stdin)
	INtype.Scan()
	Type := INtype.Text()

	switch Type {
	case "1":
		w.startGame()
	case "2":
		clear()
		os.Exit(0)
	default:
		clear()
		w.restart()
	}
}