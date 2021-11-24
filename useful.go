package main

import (
	"bufio"
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
