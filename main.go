package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a, b := NewUniverse(), NewUniverse()
	a.Seed()

	for i := 0; i < 300; i++ {

		fmt.Print("\033[H\033[2J") //очищение экрана в VisualStudio - на разных API или ОС может отличаться механизм очищения. "\x0c" в Go Playground, "\033[H" на macOS. 

		Step(a, b)
		a.Show()
		time.Sleep(time.Second / 30)
		a, b = b, a

	}
}

const (
	width  = 80
	height = 15
)

// Universe является двухмерным полем клеток.
type Universe [][]bool

// NewUniverse возвращает пустую вселенную.
func NewUniverse() Universe {
	Uni := make(Universe, height)
	for height := range Uni {
		Uni[height] = make([]bool, width)
	}
	return Uni
}

// Seed заполняет вселенную случайными живыми клетками.
func (u Universe) Seed() Universe {
	for _, line := range u {
		for i := range line {
			if rand.Intn(4) == 3 {
				line[i] = true
			}
		}
	}
	return u
}

// Alive сообщает, является ли клетка живой.
// Если координаты за пределами вселенной, возвращаемся к началу.
func (u Universe) Alive(x, y int) bool {
	var result bool

	for y < 0 || y >= height {
		if y < 0 {
			y += height
		} else {
			y %= height
		}
	}
	for x < 0 || x >= width {
		if x < 0 {
			x += width
		} else {
			x %= width
		}
	}
	for column, line := range u {
		if column == y {
			for indexLife := range line {
				if indexLife == x {
					result = line[indexLife]
					break
				}
			}
		}
	}
	return result
}

// Neighbors подсчитывает прилегающие живые клетки.
func (u Universe) Neighbors(x, y int) int {
	count := 0

	if u.Alive(x-1, y-1) {
		count++
	}
	if u.Alive(x, y-1) {
		count++
	}
	if u.Alive(x+1, y-1) {
		count++
	}
	if u.Alive(x-1, y) {
		count++
	}
	if u.Alive(x+1, y) {
		count++
	}
	if u.Alive(x-1, y+1) {
		count++
	}
	if u.Alive(x, y+1) {
		count++
	}
	if u.Alive(x+1, y+1) {
		count++
	}
	return count
}

// Next возвращает состояние определенной клетки на следующем шаге.
func (u Universe) Next(x, y int) bool {
	var result bool
	if u.Alive(x, y) {
		if u.Neighbors(x, y) < 2 {
			result = false
		} else if u.Neighbors(x, y) == 2 || u.Neighbors(x, y) == 3 {
			result = true
		} else if u.Neighbors(x, y) > 3 {
			result = false
		}
	} else {
		if u.Neighbors(x, y) == 3 {
			result = true
		}
	}
	return result
}

// Step обновляет состояние следующей вселенной (b) из
// текущей вселенной (a).
func Step(a, b Universe) {

	for column, line := range b {
		for index := range line {
			line[index] = a.Next(index, column)
		}
	}
}

// Show возвращает обновлённую вселенную.
func (u Universe) Show() {
	for _, line := range u {
		for _, life := range line {
			if life {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}

}
