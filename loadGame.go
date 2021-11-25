package main

import (
	"fmt"
	"strings"
)


func (w *hangManData) startGame() {
	fmt.Println(("-----------------------"))
	fmt.Println("ESSAYONS DE SAUVER JOSE")
	fmt.Println(("-----------------------"))
	file := "words.txt"
	word := Capitalize(readLine(file, Counter(file)))
	w.Init(hideWord(word), word, 0, []string{}, []string{})
	reveal := len(w.ToFind)/2-1
	w.HiddenWord[reveal] = string(w.ToFind[reveal])
	w.Tried = append(w.Tried, w.HiddenWord[reveal])
	w.inputLetter()
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
