package main

import (
	"fmt"
	"strings"
)

const (
	estonianAlphabet  = "ABCDEFGHIJKLMNOPRSŠZŽTUVÕÄÖÜ"
	specialCharacters = " ,!"
	cypherMessage     = "Oöm te rimi iit im üefeshe, tmmt dläi omshpeä äddzm oežöh lerrötäeüeh sšfihepä tmsö äšmäö ne dot simtä üiip zimhef tö ommolšfötä äbztipä tmmt oöm te rekeh! Üeeäe iääi, üam römhö!"
)

var indexInAlphabet []int
var value int

func main() {
	fmt.Println("Day 4")
	toDecypher := strings.ToLower(cypherMessage)
	specialChars := strings.Split(specialCharacters, "")
	alphabet := strings.Split(strings.ToLower(estonianAlphabet), "")

	indexInAlphabet = decodeIntoArray(toDecypher, specialChars, alphabet)
	decypheredMessage := deCypher(alphabet)

	fmt.Println("The decyphered message:", strings.Join(decypheredMessage, ""))

}

func decodeIntoArray(toDecypher string, specialChars []string, alphabet []string) []int {
	for _, v := range toDecypher {
		for innerIndex, letter := range alphabet {
			if string(v) == letter {
				value = innerIndex
				indexInAlphabet = append(indexInAlphabet, value)
			}
		}
		for i, specialChar := range specialChars {
			if string(v) == specialChar {
				indexInAlphabet = append(indexInAlphabet, i+104)
			}
		}
	}
	//Offset cesaer by 4
	for i, v := range indexInAlphabet {
		indexInAlphabet[i] = v - 4

	}

	return indexInAlphabet

}

func deCypher(alphabet []string) (decypheredMessage []string) {
	
	for _, v := range indexInAlphabet {
		if v < len(alphabet) {
			if v < 0 {
				decypheredMessage = append(decypheredMessage, alphabet[len(alphabet)+v])
			} else {
				decypheredMessage = append(decypheredMessage, alphabet[v])
			}
		} else if v == 100 {
			decypheredMessage = append(decypheredMessage, " ")
		} else if v == 101 {
			decypheredMessage = append(decypheredMessage, ",")
		} else if v == 102 {
			decypheredMessage = append(decypheredMessage, "!")
		}
	}
	return decypheredMessage
}
