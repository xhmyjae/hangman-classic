package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func IsLetter(s string) bool {
	if len(s) == 0 {
		return true
	} else {
		for _, cara := range s {
			if !(rune('a') <= cara && cara <= rune('z') || rune('A') <= cara && cara <= rune('Z')) {
				return false
			}
		}
		return true
	}
}

func Capitalize(s string) string {
	var res string
	for _, cara := range s {
		if 'a' <= cara && cara <= 'z' {
			cara -= 32
			res += string(cara)
		} else {
			res += string(cara)
		}
	}
	return res
}

func lineCounter(file string) int {
	scanner := bufio.NewScanner(strings.NewReader(file))

	scanner.Split(bufio.ScanLines)

	count := 0
	for scanner.Scan() {
		count++
	}
	return count
}

func readLine(onefile string, ranNum int) string {
	var word string
	file, _ := os.Open(onefile)
	scanner := bufio.NewScanner(file)
	randomNumber := rand.Intn(ranNum)
	i := 0
	for scanner.Scan() {
		if i == randomNumber {
			word = scanner.Text()
		}
		i++
	}
	return word
}

func hideWord(word string) []string {
	wordLen := len(word)
	var hiddenWord []string
	for i:=0; i<wordLen; i++ {
		hiddenWord = append(hiddenWord, "_")
	}
	return hiddenWord
}

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

func (w *hangManData) death() {
	if w.Attempts >= 10 {
		fmt.Println("T'es mort bouffon.")
	} else {
		fmt.Println("La lettre n'est pas dans le mot! Il te reste", 10-w.Attempts)
		w.inputLetter()
	}
}

func (w *hangManData) revealLetter(letter string) {
	for index, each := range w.ToFind {
		if letter == string(each) {
			w.HiddenWord[index] = letter
			w.inputLetter()
		}
	}
	w.Attempts++
	w.death()
}