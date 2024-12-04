package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	// "regexp"
	"bufio"
	"strconv"
)
	
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main () {
	input, err := os.Open("./input.txt")
	check(err)
	// re := regexp.MustCompile("[0-9]+")
    fileScan := bufio.NewScanner(input)
    fileScan.Split(bufio.ScanLines)
    var leftCoordinateArray []int
	var leftCoordinate int
	var rightCoordinateArray []int
	var rightCoordinate int
	var distanceArray []int
	var finalDistance int

    for fileScan.Scan() {
		lineText := strings.Split(fileScan.Text(), "   ")
		leftCoordinate, err = strconv.Atoi(lineText[0])
		leftCoordinateArray = append(leftCoordinateArray, leftCoordinate)
		rightCoordinate, err = strconv.Atoi(lineText[1])
		rightCoordinateArray = append(rightCoordinateArray, rightCoordinate)
        // coordinatesLines = append(coordinatesLines, fileScan.Text())
    }

	sort.Ints(leftCoordinateArray)
	sort.Ints(rightCoordinateArray)
	fmt.Println(leftCoordinateArray)
	fmt.Println(rightCoordinateArray)
	
	for index, element := range leftCoordinateArray {
		var distanceValue int
		if element > rightCoordinateArray[index] {
			distanceValue = element - rightCoordinateArray[index]
			distanceArray = append(distanceArray, distanceValue)
		} else {
			distanceValue = rightCoordinateArray[index] - element
			distanceArray = append(distanceArray, distanceValue)
		}
	}

	fmt.Println("Distance Array: ", distanceArray)
	for _, distance := range distanceArray {
		finalDistance = finalDistance + distance
	}

	fmt.Println("Final Distance: ", finalDistance)
}