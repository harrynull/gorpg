package main

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
)

type GameMap struct {
	gameMap [][]Tile
	mobs []Mob
}

func MakeMap(mapFile string) *GameMap {
	// read mobs
	jsonFile, err := os.Open("mobs.json")

	if err != nil {
		return nil
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var mobs []Mob
	json.Unmarshal(byteValue, &mobs)

	var mobsMap = make(map[string]Mob)
	for _, mob := range mobs {
		mobsMap[mob.Symbol] = mob
	}

	// read map
	file, err := os.Open(mapFile)

	if err != nil {
		return nil
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	const Width = 10
	const Height = 10
	gameMap := make([][]Tile, Height)
	for i := 0; i < Height; i++ {
		gameMap[i] = make([]Tile, Width)
		for j := 0; j < Width; j++ {
			gameMap[i][j] = MakeTileFromChar(string(lines[i][j]), &mobsMap)
		}
	}
	return &GameMap{gameMap, mobs}
}

func (m *GameMap) collisionCheck(x int, y int) bool {
	if x < 0 || x > 10 || y < 0 || y > 10 {
		return true
	}
	return !m.gameMap[y][x].passable()
}
