package main

import (
	"fmt"
	"math"
	"strings"
) 

func main() {

	directions := strings.Split("v p v v", " ")
	depth := len(directions) +1
	maxNodesAtBottom := int(math.Pow(2, float64(depth-1)))
	currentIndex := 0
	
	for _, direction := range directions {
		if direction == "v" {
			currentIndex = 2*currentIndex + 1
		} else {
			currentIndex = 2*currentIndex + 2
		}
	}
	exactValue := currentIndex - (maxNodesAtBottom - 1) + 1
	fmt.Println(exactValue)
	//Seems i was quite shit at solving perfect binary tree. Had to use ChatGPT. disappointment.
}

