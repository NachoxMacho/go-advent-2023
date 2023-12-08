package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil { log.Fatal(err) }
	sum := 0
	for _, line := range strings.Split(string(content), "\n") {
		if len(line) == 0 { continue }
		lineSplit := strings.Split(line, "")
		fmt.Print("Line => ", lineSplit)
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
		fmt.Print(" | Left => ", lineSplit[left])
		fmt.Print(" | Right => ", lineSplit[right])
		fmt.Println()
		num, err := strconv.Atoi(lineSplit[left] + lineSplit[right])
		if err != nil { log.Fatal(err) }
		sum += num
	}
	fmt.Println("Sum =>", sum)
}
