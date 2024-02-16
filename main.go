package main

import (
	"fmt"
	"math/rand"
)

type Tile rune

const (
	WALL_TILE  = '#'
	FLOOR_TILE = '.'
)

type World [][]Tile

func main() {
	w := 50
	h := 20

	world := make(World, h)
	for i := range world {
		world[i] = make([]Tile, w)
		for j := range world[i] {
			world[i][j] = WALL_TILE
		}
	}

	printWorld(world)
	addNoise(world, 60, int64(rand.Intn(100)))
    addBorders(world)
	printWorld(world)
    newWorld := cellularAutomata(world,3)
    printWorld(newWorld)
}

func cellularAutomata(world World, iterations int) World {
	temp := world
	for i := 0; i < iterations; i++ {
		temp = cellularAutomataRun(temp)
	}
	return temp
}
func cellularAutomataRun(world World) World {
	result := make(World, len(world))
	for i := range result {
		result[i] = make([]Tile, len(world[0]))
		for j := range world[i] {
			result[i][j] = world[i][j]
		}
	}

	for i := 1; i <= len(world)-2; i++ {
		for j := 1; j <= len(world[0])-2; j++ {

			counter := 0
			for y := i - 1; y <= i+1; y++ {
				for x := j - 1; x <= j+1; x++ {
					if y == i && x == j {
						continue
					}

					tile := world[y][x]
					if tile == WALL_TILE {
						counter++
					}
				}
			}

			if counter > 4 {
				result[i][j] = WALL_TILE
			} else {
				result[i][j] = FLOOR_TILE
			}
		}
	}
	return result
}

func addBorders(world World) {
	for i, row := range world {
		if i == 0 || i == len(world)-1 {
			for j := range row {
				world[i][j] = WALL_TILE
			}
		} else {
			world[i][0] = WALL_TILE
			world[i][len(world[i])-1] = WALL_TILE
		}
	}
}

func addNoise(world World, density int, seed int64) {
	dice := rand.New(rand.NewSource(seed))

	for i, row := range world {
		for j := range row {
            roll := dice.Intn(100)
			if roll > density {
				world[i][j] = FLOOR_TILE
			}
		}
	}
}

func printWorld(world World) {
	for _, row := range world {
		fmt.Printf("%s\n", string(row))
	}
	fmt.Println("----")
}
