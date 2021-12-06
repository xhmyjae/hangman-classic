package main

import (
	"bufio"
	"fmt"
	"os"
)


/* inputLetter display information such as the amount of attempts left and the letters already input, and it allows the user to input text (letters/word).
It will also load the function linked with the input */
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
		w.checkWord(letter)
	} else if !IsLetter(letter) {
		clear()
		fmt.Println("Les caractères spéciaux ne sont pas acceptés comme réponse.")
		w.inputLetter()
	} else {
		w.checkTried(letter)
	}
}


/* checkTried will check if the input letter has already been said by the user */
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


/* revealLetter will display the letter said if it's in the word, if not it will increase the attempts and load the death function */
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


 /* death will check if the attempts are higher or equal to 10, if so it'll end the game and load the restart function, if not the game continue */
func (w *hangManData) death() {
	if w.Attempts >= 10 {
		fmt.Printf("\nTu as tué Josè!!! Sa famille et son chien attendront son retour pour toujours...\nPS : Le mot était %v\n", w.ToFind)
		w.restart()
	} else {
		w.inputLetter()
	}
}


/* killJose will display the hangman, it changes depending on the attempts number thanks to the startLine and endLine variables */
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


/* checkWord will load if the users input is a word. It will check if it's the expected word or not */
func (w *hangManData) checkWord(word string) {
	if word == w.ToFind {
		for index, each := range w.ToFind {
			w.HiddenWord[index] = string(each)
		}
		w.endGame()
	} else {
		onlyLetters := true
		for _, each := range word {
			if !IsLetter(string(each)) {
				onlyLetters = false
			}
		}
		if onlyLetters {
			w.Attempts += 2
			fmt.Printf("Le mot n'était pas %v!", word)
		} else {
			fmt.Println("Les caractères spéciaux ne sont pas acceptés comme réponse.")
		}
		w.death()
	}
}