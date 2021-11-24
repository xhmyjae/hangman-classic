package main

import (
	"fmt"
	"math/rand"
	"time"
)

var Elword hangManData

func main() {
	rand.Seed(time.Now().UnixNano())
	file := "words.txt"
	//numLines := lineCounter(file)
	//randNum := rand.Intn(numLines)
	word := readLine(file, 37)
	Elword.Init(hideWord(word), word, 0, []string{}, []string{})
	fmt.Println(word)
	Elword.inputLetter()


}
