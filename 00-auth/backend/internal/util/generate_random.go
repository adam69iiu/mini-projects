package util

import (
	"fmt"
	"math/rand/v2"

	"github.com/google/uuid"
)


var nouns = []string {
 "adam", "hana", "san", "anas", "dana", "mama","ahmad", "naza",
 "kaza", "jhon", "dano", "fara", "hela", "deza", 
}

var father = []string {
	"qashmar", "rambo", "jambo", "vana", "kafo", "gdo","farxo",
}
func GenerateUsername() string {
  noun := nouns[rand.IntN(len(nouns))]
  second := father[rand.IntN(len(father))]
	num := rand.IntN(9000) + 1000
	return  fmt.Sprintf("%s-%s-%d",noun,second, num)
}


func GenerateID() string {
	return uuid.New().String()
}
