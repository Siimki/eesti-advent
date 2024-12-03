package main

import (
	"fmt"
	"os"
	"strings"
)

func day2() {
	rawFile, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println("Error opening file:")
		return
	}
	file := string(rawFile)
	splittedFileByRows := strings.Split(file, "\n")

	var topMost int = -1
	var bottomMost int
	var mostLeft int = 999
	var mostRight int = 0
	var mostXrightness int = 0


	for i, row := range splittedFileByRows {
		if strings.Contains(row, "X") {
			if topMost == -1 {
				topMost = i
			}
			bottomMost = i
			leftIndex := strings.IndexAny(row, "X")

			if leftIndex < mostLeft {
				mostLeft = leftIndex
			}

			if mostRight >= 0 {
				mostRight = strings.LastIndexAny(row, "X")
				if mostRight > mostXrightness {
					mostXrightness = mostRight
				}
			}
		}
	}

	width := mostXrightness - mostLeft
	height := bottomMost - topMost
	realFullCirc := width + 1 + width + 1 + height + 1 + height + 1 + 4
	fmt.Println("We need that amount of fence!", realFullCirc)

	//ideally change the var names of the mostRight and mostXrightness. But i Have to go now. it is 00:22 and I wake up at 06:45 tomorrow to drive my girlfriend to the bus....
}
