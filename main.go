package main

import ("fmt"
		"os"
		"strings")

func main() {
	fmt.Println("hello world")

	rawFile, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println("Error opening file:")
	}
	file := string(rawFile)
	splittedFileByRows := strings.Split(file, "\n")
	fmt.Println("Amount of rows", len(splittedFileByRows))
	var firstXHeight int 
	var lastXHeight int
	var firstXLeftness int = 999
	var lastXRightness int = 0
	var mostXrightness int = 0
	var foundXHeight bool = false

	
	for i, v := range splittedFileByRows {
		xExists := strings.Contains(v, "X")

		if xExists && foundXHeight != true {
			fmt.Println("found X at row:", i)
			foundXHeight = true
			firstXHeight = i
		} else if xExists {
			lastXHeight = i
		}

		if firstXLeftness <= 999 && xExists {
			firstXLeftness = strings.IndexAny(v, "X")
		} 

		if lastXRightness >= 0 && xExists  {
			lastXRightness = strings.LastIndexAny(v, "X")
			if lastXRightness > mostXrightness {
				mostXrightness = lastXRightness
			}
		}

		xExists = false
		
	}

	foundXHeight = false
	// var lowestXHeight int
	for i := len(splittedFileByRows)-1; i >= 0; i-- {
		if strings.Contains(splittedFileByRows[i], "X") && foundXHeight != true {
			fmt.Println("found X at row:", i)
			foundXHeight = true
			// lowestXHeight = i
		}
	}
	fmt.Println("MostRightX:", mostXrightness)
	fmt.Println("MostLeftX:", firstXLeftness)
	fmt.Println("firstXheight:", firstXHeight)
	fmt.Println("lowest X:", lastXHeight)
	fullCircumference := (firstXHeight+mostXrightness+firstXLeftness+lastXHeight+4)
	fmt.Println("FullCirc is :", fullCircumference)
	//fmt.Println(string(file)) //prints correctly to terminal.
	fmt.Println(strings.Count(file,"\n")) // we have 999 new lines. 



	// we have to find the first X and mark its height for top
	
	// We have to find the most right X and mark its X-axis for right
	// We have to find the most left X and mark its X-axis for left
	// We have to find the very last X to mark the bottom
	// Calc the circumfernce(Top+Bottom+Left+Right) height and width + 1


}