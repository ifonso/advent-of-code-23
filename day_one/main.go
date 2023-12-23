package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var numbers = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func getLeftmostNumber(s string) int {
	leftmost := -1
	leftmostIndex := -1

	// Find number
	for i, c := range s {
		if unicode.IsDigit(c) {
			num, err := strconv.Atoi(string(c))
			if err != nil {
				log.Fatal(err)
			}
			leftmost = num
			leftmostIndex = i
			break
		}
	}

	// Find number string
	for name, n := range numbers {
		contains, i := Contains(s, name)

		if contains {
			if leftmost == -1 {
				leftmost = n
				leftmostIndex = i
			}
			if leftmostIndex > i {
				leftmost = n
				leftmostIndex = i
			}
		}
	}

	return leftmost
}

func getRightmostNumber(s string) int {
	rightmost := -1
	rightmostIndex := -1
	rs := Reverse(s)

	// Find number
	for i, c := range rs {
		if unicode.IsDigit(c) {
			num, err := strconv.Atoi(string(c))
			if err != nil {
				log.Fatal(err)
			}
			rightmost = num
			rightmostIndex = i
			break
		}
	}

	// Find number string
	for name, n := range numbers {
		contains, i := Contains(rs, Reverse(name))

		if contains {
			if rightmost == -1 {
				rightmost = n
				rightmostIndex = i
			}
			if rightmostIndex > i {
				rightmost = n
				rightmostIndex = i
			}
		}
	}

	return rightmost
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		text := strings.ToLower(scanner.Text())
		total += getLeftmostNumber(text)*10 + getRightmostNumber(text)
	}

	fmt.Println("Result:", total)
}
