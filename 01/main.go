package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseLists(filename string) ([]int, []int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	left := make([]int, 0)
	right := make([]int, 0)

	for scanner.Scan() {
		components := strings.SplitN(scanner.Text(), " ", 2)

		leftNum, leftErr := strconv.Atoi(strings.TrimSpace(components[0]))
		rightNum, rightErr := strconv.Atoi(strings.TrimSpace(components[1]))
		if leftErr != nil || rightErr != nil {
			log.Fatal("Error parsing numbers")
		}

		fmt.Printf("%d %d\n", leftNum, rightNum)

		left = append(left, leftNum)
		right = append(right, rightNum)
	}
	return left, right
}

func main() {
	left, right := parseLists("./data.txt")
	sort.Ints(left)
	sort.Ints(right)
	fmt.Printf("left: %v\n", left)
	fmt.Printf("right: %v\n", right)

	sum := 0
	for i := 0; i < len(left); i++ {
		sum += max(left[i], right[i]) - min(left[i], right[i])
	}

	fmt.Printf("Sum: %d\n", sum)

	similarityScore := 0
	for i := 0; i < len(left); i++ {

		leftnr := left[i]
		rightCnt := 0
		for _, num := range right {
			if num == leftnr {
				rightCnt++
			}
		}
		similarityScore += leftnr * rightCnt
	}

	fmt.Printf("Similarity: %d\n", similarityScore)
}
