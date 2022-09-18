package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Level struct {
	Width  int
	Height int
	Layout []Tiles
}

type Tiles struct {
	Value int
	X     int
	Y     int
}

func main() {
	level, err := readLevel("./level.json")
	if err != nil {
		panic(err)
	}

	levelAsArray := level.toArray()
	ProcessLevel(levelAsArray)
}

func readLevel(levelPath string) (*Level, error) {
	content, err := ioutil.ReadFile(levelPath)
	if err != nil {
		log.Fatal("Error reading file:", err)
		return nil, err
	}

	var level *Level
	err = json.Unmarshal(content, &level)
	if err != nil {
		log.Fatal("Error unmarshalling file:", err)
		return nil, err
	}

	return level, nil
}

func (l *Level) toArray() []int {
	levelAsArray := make([]int, l.Width*l.Height)

	for _, tile := range l.Layout {
		levelAsArray[tile.X*l.Width+tile.Y] = tile.Value
	}

	fmt.Println("levelAsArray:", levelAsArray)
	return levelAsArray
}

func ProcessLevel(level []int) {
	goal := findGoal(level)
	fmt.Println("goal", goal)
}

func findGoal(level []int) int {
	for count, tile := range level {
		if tile == -10 {
			return count
		}
	}
	return -1
}
