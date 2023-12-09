package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Digit struct {
	str string
	val string
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil { log.Fatal(err) }
	sum := 0
	
	digits := []Digit {
		{ str: "nine", val: "n9e", },
		{ str: "eight", val: "e8t", },
		{ str: "seven", val: "s7n", },
		{ str: "six", val: "s6x", },
		{ str: "five", val: "f5e", },
		{ str: "four", val: "f4r", },
		{ str: "three", val: "t3e", },
		{ str: "two", val: "t2o", },
		{ str: "one", val: "o1e", },
	}
	
	for _, line := range strings.Split(string(content), "\n") {
		if len(line) == 0 { continue }
		originalLine := line

		for _, digit := range digits {
			line = strings.ReplaceAll(line, digit.str, digit.str + digit.val)
		}

		lineSplit := strings.Split(line, "")
		left, right := 0, len(lineSplit) - 1
		for left < right {
			_, err := strconv.Atoi(lineSplit[left])
			if err == nil { break }
			left++
		}
		for right > left {
			_, err := strconv.Atoi(lineSplit[right])
			if err == nil { break }
			right--
		}
		num, err := strconv.Atoi(lineSplit[left] + lineSplit[right])
		if err != nil { log.Fatal(err) }
		sum += num
		fmt.Print("Line => ", originalLine)
		fmt.Print(" | Replaced => ", line)
		fmt.Print(" | Left => ", lineSplit[left])
		fmt.Print(" | Right => ", lineSplit[right])
		fmt.Print(" | Sum => ", sum)
		fmt.Println()
	}
	fmt.Println("Sum =>", sum)
}
