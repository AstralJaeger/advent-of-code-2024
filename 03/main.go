package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readFile(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := ""

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		data += scanner.Text()
	}
	return data
}

var numberr = regexp.MustCompile("[\\d]{1,3}")

func parseNumbersFromMul(instruction string) (int, int) {
	numbers := numberr.FindAllString(instruction, -1)
	nrs := make([]int, 2)
	for j, number := range numbers {
		n, err := strconv.Atoi(number)
		if err != nil {
			log.Fatal(err)
		}
		nrs[j] = n
	}
	return nrs[0], nrs[1]
}

func main() {

	data := readFile("./mydata.txt")
	//fmt.Printf("Raw Data: %n\n", data)
	r := regexp.MustCompile("(do(n't)?\\(\\))|(mul\\(([\\d]{1,3}),([\\d]{1,3})\\))")

	results := r.FindAllString(data, -1)
	sum := 0
	skip := false
	for i, result := range results {
		if strings.HasPrefix(result, "do") {
			fmt.Printf("Match %3d: do()\n", i)
			skip = false
		}
		if strings.HasPrefix(result, "don't") {
			fmt.Printf("Match %3d: don't()\n", i)
			skip = true
		}

		if skip {
			a, b := parseNumbersFromMul(result)
			fmt.Printf("Match %3d: mul(%3d,%3d): %s\n", i, a, b, "SKIPPED")
		}

		if !skip && strings.HasPrefix(result, "mul") {
			a, b := parseNumbersFromMul(result)
			product := a * b
			sum += product
			fmt.Printf("Match %3d: mul(%3d,%3d): %6d\n", i, a, b, product)
		}
	}

	fmt.Printf("Sum: %d\n", sum)
}
