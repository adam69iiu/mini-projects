package main


import (
	"fmt"
	"math/rand/v2"
)


var nouns = []string {
 "adam", "hana", "san", "anas", "dana", "mama","ahmad", "naza",
}

var father = []string {
	"qashmar", "rambo", "jambo", "vana", "kafo", "gdo","farxo",
}
func main() {
  noun := nouns[rand.IntN(len(nouns))]
  second := father[rand.IntN(len(father))]
	num := rand.IntN(9000) + 1000
	fmt.Printf("%s-%s-%d",noun,second, num)
}
