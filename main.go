package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"math/rand"
	"os"
	"time"
)

var grid = [4][4]int{}
var empSize = 16
var score int
var step int

var test bool

func init() {
	if test {
		return
	}
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	rand.Seed(time.Now().UnixNano())
	randFillOne()
	randFillOne()
	printStatus()
}

var applyKey = map[termbox.Key]func(f func([4]int) [4]int) bool{
	termbox.KeyArrowUp: func(f func([4]int) [4]int) (changed bool) {
		for i := 0; i < 4; i++ {
			var col [4]int
			for j := 0; j < 4; j++ {
				col[j] = grid[j][i]
			}
			newCol := f(col)
			if newCol == col {
				continue
			}
			changed = true
			for j, newVal := range newCol {
				grid[j][i] = newVal
			}
		}
		return
	},
	termbox.KeyArrowDown: func(f func([4]int) [4]int) (changed bool) {
		for i := 0; i < 4; i++ {
			var col [4]int
			for j := 0; j < 4; j++ {
				col[j] = grid[3-j][i]
			}
			newCol := f(col)
			if newCol == col {
				continue
			}
			changed = true
			for j, newVal := range newCol {
				grid[3-j][i] = newVal
			}
		}
		return
	},
	termbox.KeyArrowLeft: func(f func([4]int) [4]int) (changed bool) {
		for i := 0; i < 4; i++ {
			newRow := f(grid[i])
			if grid[i] == newRow {
				continue
			}
			changed = true
			grid[i] = newRow
		}
		return
	},
	termbox.KeyArrowRight: func(f func([4]int) [4]int) (changed bool) {
		for i := 0; i < 4; i++ {
			var row [4]int
			for j := 0; j < 4; j++ {
				row[j] = grid[i][3-j]
			}
			newRow := f(row)
			if row == newRow {
				continue
			}
			changed = true
			for j, newVal := range newRow {
				grid[i][3-j] = newVal
			}
		}
		return
	},
}

func merge(arr [4]int) [4]int {
	var newArr [4]int
	var ni int
	var lastNum int
	for i := 0; i < 4; i++ {
		if arr[i] == 0 {
			continue
		}
		if arr[i] == lastNum {
			db := lastNum << 1
			lastNum = 0
			newArr[ni-1] = db

			score += db
			empSize++
		} else {
			newArr[ni] = arr[i]
			lastNum = arr[i]
			ni++
		}
	}
	return newArr
}
func printStatus() {
	if test {
		return
	}
	termbox.Sync()
	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Println("score: ", score)
	fmt.Println("step: ", step)
}

func main() {
	for {
		key := input()
		apply := applyKey[key]
		if apply == nil {
			continue
		}
		changed := apply(merge)
		if !changed {
			continue
		}
		step++
		randFillOne()
		printStatus()
		if empSize == 0 && death() {
			fmt.Println("end!")
			exit()
		}
	}
}

func input() termbox.Key {
	for {
		event := termbox.PollEvent()
		switch event.Type {
		case termbox.EventKey:
			if event.Key == termbox.KeyCtrlC {
				exit()
			}
			return event.Key
		}
	}
}

func randFillOne() {
	r := rand.Intn(3)
	for {
		var emps []int
		for i, v := range grid[r] {
			if v == 0 {
				emps = append(emps, i)
			}
		}
		if emps != nil {
			grid[r][emps[rand.Intn(len(emps))]] = 2 << (rand.Int() & 1)
			empSize--
			return
		}
		r++
		if r == 4 {
			r = 0
		}
	}
}

func death() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			cell := grid[i][j]
			if cell == grid[i][j+1] || cell == grid[i+1][j] {
				return false
			}
		}
	}
	// check last row and col
	for i := 0; i < 3; i++ {
		if grid[i][3] == grid[i+1][3] {
			return false
		}
		if grid[3][i] == grid[3][i+1] {
			return false
		}
	}
	return true
}

func exit() {
	termbox.Close()
	os.Exit(0)
}
