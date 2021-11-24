package main

import (
	"math/rand"
	"time"
)

var Elword hangManData

func main() {
	rand.Seed(time.Now().UnixNano())
	//numLines := lineCounter(file)
	//randNum := rand.Intn(numLines)
	Elword.startGame()
}
