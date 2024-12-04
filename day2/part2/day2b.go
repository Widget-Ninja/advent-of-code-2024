package main

import (
	"fmt"
	"os"
	// "sort"
	"strings"
	"bufio"
	"strconv"
)
	
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func RemoveIndex(s []int, index int) []int {
    return append(s[:index], s[index+1:]...)
}

func main () {
	input, err := os.Open("./input-simple.txt")
	check(err)
    fileScan := bufio.NewScanner(input)
    fileScan.Split(bufio.ScanLines)

	var safeUnsafeArray []string
	safeCount := 0
	fileScanIndex := 0

    for fileScan.Scan() {
		lineText := strings.Split(fileScan.Text(), " ")
		var reportArray []int
		var reportValue int
		for _, value := range lineText {
			reportValue, err = strconv.Atoi(value)
			reportArray = append(reportArray, reportValue)
		}
		fmt.Println(reportArray)
		safeUnsafeArray = append(safeUnsafeArray, "null")

		unsafeBuffer := 0

		if reportArray[0] == reportArray[1] {
			safeUnsafeArray[fileScanIndex] = "unsafe"
		} else if reportArray[0] > reportArray[1] {
			for index := 0; index < len(reportArray) - 1; index++ {
				if reportArray[index] == reportArray[index + 1] {
					safeUnsafeArray[fileScanIndex] = "unsafe"
					break
				} else if reportArray[index] - reportArray[index + 1] > 3 {
					safeUnsafeArray[fileScanIndex] = "unsafe"
					break
				} else if reportArray[index] - reportArray[index + 1] < 0 {
					safeUnsafeArray[fileScanIndex] = "unsafe"
					break
				} else {
					safeUnsafeArray[fileScanIndex] = "safe"
				}
			}
		} else if reportArray[0] < reportArray[1] {
			for index := 0; index < len(reportArray) - 1; index++ {
				if reportArray[index + 1] == reportArray[index] {
					safeUnsafeArray[fileScanIndex] = "unsafe"
					break
				} else if reportArray[index + 1] - reportArray[index] > 3 {
					safeUnsafeArray[fileScanIndex] = "unsafe"
					break
				} else if reportArray[index + 1] - reportArray[index] < 0 {
					safeUnsafeArray[fileScanIndex] = "unsafe"
					break
				} else {
					safeUnsafeArray[fileScanIndex] = "safe"
				}
			}
		}

		fileScanIndex = fileScanIndex + 1
    }

	fmt.Println("Safe Unsafe Array: ", safeUnsafeArray)
	for _, safeUnsafe := range safeUnsafeArray {
		if safeUnsafe == "safe" {
		safeCount = safeCount + 1
		}
	}

	fmt.Println("Number of Safe Reports: ", safeCount)
}