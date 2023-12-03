package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const (
	MAX_RED   = 12
	MAX_GREEN = 13
	MAX_BLUE  = 14
)

func p(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// read file line by line
	// for each line, split by comma
	file, err := os.Open("input.txt")
	p(err)
	defer file.Close()
	// read file line by line
	sum_of_winning_game_numbers := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		game_num := parseGameNum(line)
		fmt.Println("Game", game_num)
		cubes := parseCubeMax(line)
		if IsGamePossible(cubes) {
			fmt.Println("Game", game_num, "is possible")
			sum_of_winning_game_numbers += game_num
		}
	}
	fmt.Println("Sum of winning game numbers:", sum_of_winning_game_numbers)
}

var maxValues = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func IsGamePossible(cubes []CubeGroup) bool {
	for _, cube := range cubes {
		maxValue, ok := maxValues[cube.Color]
		if !ok {
			continue
		}
		if cube.Value > maxValue {
			return false
		}
	}
	return true
}

type CubeGroup struct {
	Color string
	Value int
}

func parseCubeMax(line string) []CubeGroup {
	var cubes []CubeGroup
	r := regexp.MustCompile(`(\d+) (green|blue|red)`)
	matches := r.FindAllStringSubmatch(line, -1)
	for _, match := range matches {
		numOfBlocksStr := match[1]
		color := match[2]
		numOfBlocks, err := strconv.Atoi(numOfBlocksStr)
		p(err)
		cubes = append(cubes, CubeGroup{color, numOfBlocks})
	}
	return cubes
}

func parseGameNum(line string) int {
	r := regexp.MustCompile(`(\d+):`)
	match := r.FindStringSubmatch(line)
	gameNumStr := match[1]
	gameNum, err := strconv.Atoi(gameNumStr)
	p(err)
	return gameNum
}
