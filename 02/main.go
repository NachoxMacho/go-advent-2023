package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Pull struct {
	num   int
	color string
}

type Game struct {
	id    int
	pulls []Pull
	valid bool
}

func getCubePower(g Game) int {
	minRed, minBlue, minGreen := 0, 0, 0
	for _, p := range g.pulls {
		if p.color == "red" && p.num > minRed { minRed = p.num }
		if p.color == "blue" && p.num > minBlue { minBlue = p.num }
		if p.color == "green" && p.num > minGreen { minGreen = p.num }
	}
	return minRed * minBlue * minGreen	
}

func isGameValid(g Game) bool {
	for _, pull := range g.pulls {
		if (pull.color == "red" && pull.num > 12) {
			return false
		}
		if (pull.color == "green" && pull.num > 13) {
			return false
		}
		if (pull.color == "blue" && pull.num > 14) {
			return false
		}
	} 
	return true
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	games := []Game{}
	
	for _, line := range strings.Split(string(content), "\n") {
		if len(line) == 0 { continue }
		fmt.Println("line =>", line)
		id, err := strconv.Atoi(strings.ReplaceAll(strings.Split(line, ":")[0], "Game ", ""))
		if err != nil { log.Fatal(err) }
		allGamesString := strings.Split(line, ":")[1]
		game := Game{
			id: id,
			valid: true,
		}
		for _, gameString := range strings.Split(allGamesString, ";") {
			fmt.Println(" | Game =>", gameString)
			
			for _, pullString := range strings.Split(gameString, ",") {
				num, err := strconv.Atoi(strings.Split(strings.Trim(pullString, " "), " ")[0])
				if err != nil { log.Fatal(err) }
				game.pulls = append(game.pulls, Pull{
					num: num,
					color : strings.Split(strings.Trim(pullString, " "), " ")[1],
				})
			}
		}
		fmt.Println("id =>", game.id)
		games = append(games, game)
		fmt.Println()
	}
	sum := 0
	for _, game := range games {
		// if (isGameValid(game)) { sum += game.id}
		sum += getCubePower(game)
	}
	fmt.Println("Final =>", sum)
}
