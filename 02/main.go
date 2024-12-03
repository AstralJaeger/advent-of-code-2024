package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseLists(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	data := make([][]int, 0)

	for scanner.Scan() {
		dataLine := make([]int, 0)
		components := strings.Split(scanner.Text(), " ")
		for _, component := range components {
			num, err := strconv.Atoi(strings.TrimSpace(component))
			if err != nil {
				log.Fatal("Error parsing numbers")
			}
			dataLine = append(dataLine, num)
		}
		data = append(data, dataLine)
	}
	return data
}

func isMonotonicInc(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] >= arr[i+1] {
			return false
		}
	}
	return true
}

func isMonotonicDec(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] <= arr[i+1] {
			return false
		}
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func isSafe(dataLine []int) bool {
	if isMonotonicInc(dataLine) || isMonotonicDec(dataLine) {
		for j := 0; j < len(dataLine)-1; j++ {
			if max(dataLine[j], dataLine[j+1])-min(dataLine[j], dataLine[j+1]) > 3 {
				return false
			}
		}
		return true
	}
	return false
}

func isSafeWithDampener(dataLine []int) bool {
	if isSafe(dataLine) {
		return true
	}

	for i := 0; i < len(dataLine); i++ {
		// Create a new array with the i-th element removed
		modified := append([]int{}, dataLine[:i]...)
		modified = append(modified, dataLine[i+1:]...)

		if isSafe(modified) {
			return true
		}
	}

	return false
}

func main() {
	data := parseLists("./data.txt")
	safeCnt := 0

	for i, dataLine := range data {
		isSafeReport := isSafeWithDampener(dataLine)

		if isSafeReport {
			fmt.Printf("DataLine[%d]: %v is safe\n", i, dataLine)
			safeCnt++
		} else {
			fmt.Printf("DataLine[%d]: %v is unsafe\n", i, dataLine)
		}
	}

	fmt.Printf("Safe count with dampener: %d\n", safeCnt)
}
