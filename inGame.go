package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func (w *hangManData) inputLetter() {
	fmt.Println(w.HiddenWord)
	fmt.Println("Mettre une lettre :")
	INletter := bufio.NewScanner(os.Stdin)
	INletter.Scan()
	letter := INletter.Text()

	if len(letter) > 1 {
		fmt.Println("Une seule lettre est requise.")
		w.inputLetter()
	} else if !IsLetter(letter) {
		fmt.Println("Seules les lettres sont acceptées comme réponse.")
		w.inputLetter()
	} else {
		w.checkTried(letter)
	}
}


func (w *hangManData) checkTried(letter string) {
	for _, each := range w.Tried {
		if letter == each {
			w.Attempts++
			w.death()
		}
	}
	w.Tried = append(w.Tried, letter)
	w.revealLetter(letter)
}


func (w *hangManData) revealLetter(letter string) {
	for index, each := range w.ToFind {
		if letter == string(each) {
			w.HiddenWord[index] = letter
			w.endGame()
		}
	}
	w.Attempts++
	w.death()
}


func (w *hangManData) death() {
	if w.Attempts >= 10 {
		fmt.Println("T'es mort bouffon.")
		w.startGame()
	} else {
		fmt.Println("La lettre n'est pas dans le mot! Il te reste", 10-w.Attempts)
		w.inputLetter()
	}
}


func (w *hangManData) endGame() {
	notHidden := strings.Join(w.HiddenWord[:], "")
	if notHidden == w.ToFind {
		fmt.Println("GG, t'as trouvé le mot!")
		w.startGame()
	} else {
		w.inputLetter()
	}
}

func (w *hangManData) startGame() {
	file := "words.txt"
	word := readLine(file, 37)
	w.Init(hideWord(word), word, 0, []string{}, []string{})
	fmt.Println(word)
	reveal := len(w.ToFind)/2-1
	w.HiddenWord[reveal] = string(w.ToFind[reveal])
	w.inputLetter()

}