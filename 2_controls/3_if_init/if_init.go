package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomNumber() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	// works on switch too
	if i := randomNumber(); i > 5 {
		fmt.Println(i, "Win!")
	} else {
		fmt.Println(i, "Lose.")
	}
}
