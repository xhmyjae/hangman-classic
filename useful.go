package main

import (
	"bufio"
	"math/rand"
	"os"
)


// IsLetter will check if a letter is a special character or not
func IsLetter(s string) bool {
	if len(s) == 0 {
		return false
	} else {
		for _, cara := range s {
			if !(rune('a') <= cara && cara <= rune('z') || rune('A') <= cara && cara <= rune('Z')) {
				return false
			}
		}
		return true
	}
}


// Capitalize will put the input letter in upper-case
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


// Counter will count the amount of lines in a file
func Counter(fil string) int {
	file, _ := os.Open(fil)
	scanner := bufio.NewScanner(file)
	var count int
	for scanner.Scan() { //while \n available
		count++
	}
	return count
}


// readLine will read the text at a certain line specified in parameters
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


// hideWord will create a array with the letters of the "hidden" word in it
func hideWord(word string) []string {
	wordLen := len(word)
	var hiddenWord []string
	for i:=0; i<wordLen; i++ {
		hiddenWord = append(hiddenWord, "_")
	}
	return hiddenWord
}


// clear will delete all the text written in the terminal
func clear() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
}