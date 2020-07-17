package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	script := os.Args[1]
	fmt.Println(fmt.Sprintf("********* script: %s ***********", script))
	givens := strings.Split(os.Args[2], "")
	boardArea := 81
	for i := len(givens); i < boardArea; i++ {
		givens = append(givens, ".")
	}
	var choices []string
	for i := 0; i < 81; i++ {
		choices = append(choices, "123456789")
	}
	subsets := [][]int{
		// rows
		{0, 1, 2, 9, 10, 11, 18, 19, 20},
		{3, 4, 5, 12, 13, 14, 21, 22, 23},
		{6, 7, 8, 15, 16, 17, 24, 25, 26},
		{27, 28, 29, 36, 37, 38, 45, 46, 47},
		{30, 31, 32, 39, 40, 41, 48, 49, 50},
		{33, 34, 35, 42, 43, 44, 51, 52, 53},
		{54, 55, 56, 63, 64, 65, 72, 73, 74},
		{57, 58, 59, 66, 67, 68, 75, 76, 77},
		{60, 61, 62, 69, 70, 71, 78, 79, 80},

		// columns
		{0, 3, 6, 27, 30, 33, 54, 57, 60},
		{1, 4, 7, 28, 31, 34, 55, 58, 61},
		{2, 5, 8, 29, 32, 35, 56, 59, 62},
		{9, 12, 15, 36, 39, 42, 63, 66, 69},
		{10, 13, 16, 37, 40, 43, 64, 67, 70},
		{11, 14, 17, 38, 41, 44, 65, 68, 71},
		{18, 21, 24, 45, 48, 51, 72, 75, 78},
		{19, 22, 25, 46, 49, 52, 73, 76, 79},
		{20, 23, 26, 47, 50, 53, 74, 77, 80},

		// squares
		{0, 1, 2, 3, 4, 5, 6, 7, 8},
		{9, 10, 11, 12, 13, 14, 15, 16, 17},
		{18, 19, 20, 21, 22, 23, 24, 25, 26},
		{27, 28, 29, 30, 31, 32, 33, 34, 35},
		{36, 37, 38, 39, 40, 41, 42, 43, 44},
		{45, 46, 47, 48, 49, 50, 51, 52, 53},
		{54, 55, 56, 57, 58, 59, 60, 61, 62},
		{63, 64, 65, 66, 67, 68, 69, 70, 71},
		{72, 73, 74, 75, 76, 77, 78, 79, 80},
	}

	// remove choices eliminated by the at-most-one rule
	// fmt.Println("--- iterate over subsets to build matrix of options ---")
	for _, subset := range subsets {
		for _, i := range subset {
			// fmt.Println(fmt.Sprintf("%d, givens = %v", i, givens[i]))
			digit := givens[i]
			if digit == "." {
				continue
			}
			for _, j := range subset {
				// fmt.Println(fmt.Sprintf("%d, %d - digit = %v", i, j, digit))
				// remove the specified digit from the set of advertised choices for this spot
				choices[j] = strings.Replace(choices[j], digit, "", -1)
			}
		}
	}

	// identify choices mandated by the at-least-one-rule
	// var unique []string
	for _, subset := range subsets {
		var counts = make([]int, boardArea)
		var w = make([]int, boardArea)
		// fmt.Printf("counts = %#v\n", counts)
		for _, i := range subset {
			if givens[i] == "." {
				continue
			}
			r := regexp.MustCompile(`(.)`)
			// fmt.Printf("c = %#v\n", c)
			// fmt.Printf("findallsubmatch = %#v\n", r.FindAllStringSubmatch(c, -1))
			nums := r.FindAllStringSubmatch(choices[i], -1)
			for _, num := range nums {
				d, _ := strconv.Atoi(num[0])
				counts[d] = counts[d] + 1
				w[d] = i
			}
			// fmt.Println(fmt.Sprintf("%v", c))
		}
	}

	fmt.Println("--          --")
	fmt.Println("-- sudokant --")
	fmt.Println("--          --")
	fmt.Println(fmt.Sprintf("choices: %v", choices))
}
