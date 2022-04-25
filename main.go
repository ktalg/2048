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

func init() {
	rand.Seed(time.Now().UnixNano())
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	randFillOne()
	randFillOne()
	printGrid()
}

var applyKey = map[termbox.Key]func(f func([4]int) [4]int){
	termbox.KeyArrowUp: func(f func([4]int) [4]int) {
		for i := 0; i < 4; i++ {
			var col [4]int
			for j := 0; j < 4; j++ {
				col[j] = grid[j][i]
			}
			col = f(col)
			for j, newVal := range col {
				grid[j][i] = newVal
			}
		}
	},
	termbox.KeyArrowDown: func(f func([4]int) [4]int) {
		for i := 0; i < 4; i++ {
			var col [4]int
			for j := 3; j >= 0; j-- {
				col[j] = grid[j][i]
			}
			col = f(col)
			for j, newVal := range col {
				grid[3-j][i] = newVal
			}
		}
	},
	termbox.KeyArrowLeft: func(f func([4]int) [4]int) {
		for i := 0; i < 4; i++ {
			grid[i] = f(grid[i])
		}
	},
	termbox.KeyArrowRight: func(f func([4]int) [4]int) {
		for i := 0; i < 4; i++ {
			row := grid[i]
			for i := 0; i < 4/2; i++ {
				row[i], row[3-i] = row[3-i], row[i]
			}
			grid[3-i] = f(row)
		}
	},
}

func merge(arr [4]int) [4]int {
	var newArr [4]int
	var arrI int
	for i := 1; i < 4; i++ {
		if arr[i] == 0 {
			continue
		}
		if arr[i] == arr[i-1] {
			add := arr[i] << 1
			newArr[arrI] = add
			i++

			score += add
			empSize++
		} else {
			newArr[arrI] = arr[i]
		}
		arrI++
	}
	return newArr
}
func printGrid() {
	termbox.Sync()
	for _, row := range grid {
		fmt.Println(row)
	}
}

func main() {
	for {
		key := input()
		apply := applyKey[key]
		if apply != nil {
			apply(merge)
			randFillOne()
			printGrid()
			if empSize == 0 && death() {
				fmt.Println("end!")
				os.Exit(0)
			}
		}
	}
}

func input() termbox.Key {
	for {
		event := termbox.PollEvent()
		switch event.Type {
		case termbox.EventKey:
			if event.Key == termbox.KeyCtrlC {
				os.Exit(0)
			}
			return event.Key
		}
		fmt.Println("xx", event)
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
		row := grid[i]
		for j := 0; j < 3; j++ {
			if row[j] == row[j+1] {
				return false
			}
		}
	}
	return true
}
