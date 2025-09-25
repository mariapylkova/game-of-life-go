package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 80
	height = 15
)

type Universe [][]bool

func NewUniverse() Universe {
	field := make(Universe, height)
	for i := range field {
		field[i] = make([]bool, width)
	}
	return field
}

func (u Universe) Show() {
	for _, i := range u {
		for _, j := range i {
			if j == false {
				fmt.Printf(" ")
			} else {
				fmt.Printf("*")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func (u Universe) Seed() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 300; i++ {
		r, c := rand.Intn(height), rand.Intn(width)
		u[r][c] = true
	}
}

func (u Universe) Alive(x, y int) bool {
	y = (y + height) % height
	x = (x + width) % width
	return u[y][x]
}

func (u Universe) Neighbors(x, y int) int {
	count := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			if u.Alive(x+dx, y+dy) {
				count++
			}
		}
	}
	return count
}

func (u Universe) Next(x, y int) bool {
	n := u.Neighbors(x, y)
	alive := u.Alive(x, y)
	if alive && (n == 2 || n == 3) {
		return true
	}
	if !alive && n == 3 {
		return true
	}
	return false
}

func Step(a, b Universe) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			b[y][x] = a.Next(x, y)

		}
	}

}

func main() {
	field := NewUniverse()
	field2 := NewUniverse()
	field.Seed()
	for {
		fmt.Print("\x0c")
		field.Show()
		Step(field, field2)
		field, field2 = field2, field
		time.Sleep(300 * time.Millisecond)
	}
}
