package main

import (
	"bufio"
	"fmt"
	"os"
)


func (w *hangManData) inputLetter() {
	fmt.Println("\n***********")
	w.killJose()
	fmt.Print("Tentatives précédentes : ")
	for index, each := range w.Tried {
		if index == len(w.Tried)-1 {
			fmt.Print(each)
		} else {
			fmt.Print(each, " - ")
		}
	}
	fmt.Println("\nTentatives restantes : ", 10-w.Attempts, "\n")
	for _, each := range w.HiddenWord {
		fmt.Print(each, " ")
	}
	fmt.Print("\n\nMettre une lettre : ")
	INletter := bufio.NewScanner(os.Stdin)
	INletter.Scan()
	letter := Capitalize(INletter.Text())

	if len(letter) > 1 {
		clear()
		fmt.Println("Une seule lettre est requise.")
		w.inputLetter()
	} else if !IsLetter(letter) {
		clear()
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
			clear()
			fmt.Printf("La lettre %v est déjà dans le mot!\n", letter)
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
		clear()
		w.endGame()
	}
	w.Attempts++
	clear()
	fmt.Printf("La lettre %v n'est pas dans le mot!\n", letter)
	w.death()
}


func (w *hangManData) death() {
	if w.Attempts >= 10 {
		fmt.Printf("\nTu as tué Josè!!! Sa famille et son chien attendront son retour pour toujours...\nPS : Le mot était %v\n", w.ToFind)
		w.restart()
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