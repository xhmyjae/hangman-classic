package main

import (
	"bufio"
	"fmt"
	"os"
)


func (w *hangManData) inputLetter() {
	fmt.Print("Tentatives précédentes : ")
	for index, each := range w.Tried {
		if index == len(w.Tried)-1 {
			fmt.Println(each)
		} else {
			fmt.Print(each, " - ")
		}
	}
	fmt.Println("Tentatives restantes : ", 10-w.Attempts, "\n")
	w.killJose()
	for _, each := range w.HiddenWord {
		fmt.Print(each, " ")
	}
	fmt.Print("\n\nMettre une lettre : ")
	INletter := bufio.NewScanner(os.Stdin)
	INletter.Scan()
	letter := Capitalize(INletter.Text())

	if len(letter) > 1 {
		fmt.Println("Une seule lettre est requise.")
		fmt.Println("-------------")
		w.inputLetter()
	} else if !IsLetter(letter) {
		fmt.Println("Seules les lettres sont acceptées comme réponse.")
		fmt.Println("-------------")
		w.inputLetter()
	} else {
		w.checkTried(letter)
	}
}


func (w *hangManData) checkTried(letter string) {
	for _, each := range w.Tried {
		if letter == each {
			w.Attempts++
			fmt.Println("La lettre est déjà dans le mot!")
			w.death()
		}
	}
	w.Tried = append(w.Tried, letter)
	w.revealLetter(letter)
}


func (w *hangManData) revealLetter(letter string) {
	var anyFound = false
	for index, each := range w.ToFind {
		if letter == string(each) {
			w.HiddenWord[index] = letter
			anyFound = true
		}
	}
	if anyFound {
		w.endGame()
	}
	w.Attempts++
	fmt.Println("La lettre n'est pas dans le mot!")
	w.death()
}


func (w *hangManData) death() {
	fmt.Println("-------------")
	if w.Attempts >= 10 {
		fmt.Println("T'es mort bouffon.")
		w.startGame()
	} else {
		w.inputLetter()
	}
}


func (w *hangManData) killJose() {
	file := "hangman.txt"
	hangPos, _ := os.Open(file)
	scanner := bufio.NewScanner(hangPos)

	startLine := w.Attempts*7+1+1*w.Attempts
	endLine := (w.Attempts+1)*7+w.Attempts*1

	i:=1
	for scanner.Scan() {
		if i>=startLine && i<= endLine {
			fmt.Println(scanner.Text())
		}
		i++
	}
	fmt.Println()
}