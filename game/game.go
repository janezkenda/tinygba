package game

import (
	"fmt"
	"math/rand"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Point struct {
	X, Y int
}

type Entry struct {
	Value   int
	Blocked bool
}

type Game struct {
	Board map[Point]*Entry
	Score int
	Moved bool
	End   bool
	Win   bool
}

func NewGame() Game {
	g := Game{Board: map[Point]*Entry{}}
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			g.Board[Point{x, y}] = &Entry{}
		}
	}

	g.AddNewNumber()
	g.AddNewNumber()
	return g
}

func (g *Game) Print() {
	fmt.Println("┌──────┬──────┬──────┬──────┐")
	for x := 0; x < 4; x++ {
		fmt.Print("│")
		for y := 0; y < 4; y++ {
			el := g.Board[Point{x, y}]

			if el.Value == 0 {
				fmt.Print("      ")
			} else {
				fmt.Printf(" %4d ", el.Value)
			}
			fmt.Print("│")
		}
		if x < 3 {
			fmt.Println("\n├──────┼──────┼──────┼──────┤")
		}
	}
	fmt.Println("\n└──────┴──────┴──────┴──────┘")
}

func (g *Game) AddNewNumber() {
	var zeroes []Point

	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			pt := Point{x, y}
			el := g.Board[pt]
			if el.Value == 0 {
				zeroes = append(zeroes, pt)
			}
		}
	}

	if len(zeroes) == 0 {
		return
	}

	num := 2
	if rand.Intn(10) == 0 {
		num = 4
	}

	pt := zeroes[rand.Intn(len(zeroes))]

	val := g.Board[pt]
	val.Value = num
	g.Board[pt] = val

	if g.canMove() {
		return
	}

	g.End = true
}

func (g *Game) canMove() bool {
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			if g.Board[Point{x, y}].Value == 0 {
				return true
			}
		}
	}

	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			el := g.Board[Point{x, y}]
			if g.testAdd(x+1, y, el.Value) ||
				g.testAdd(x-1, y, el.Value) ||
				g.testAdd(x, y+1, el.Value) ||
				g.testAdd(x, y-1, el.Value) {
				return true
			}
		}
	}

	return false
}

func (g *Game) testAdd(x, y, value int) bool {
	if x < 0 || x > 3 || y < 0 || y > 3 {
		return false
	}

	return g.Board[Point{x, y}].Value == value
}

func (g *Game) mv(pt, other *Entry) {
	if pt.Value != 0 && other.Value == pt.Value &&
		!pt.Blocked && !other.Blocked {
		pt.Value = 0

		other.Value *= 2
		other.Blocked = true

		if other.Value == 2048 {
			g.Win = true
		}

		g.Score += other.Value
		g.Moved = true
	} else if other.Value == 0 && pt.Value != 0 {
		other.Value = pt.Value
		pt.Value = 0

		g.Moved = true
	}
}

func (g *Game) Move(d Direction) {
	switch d {
	case Up:
		for x := 0; x < 4; x++ {
			for y := 1; y < 4; y++ {
				if g.Board[Point{x, y}].Value == 0 {
					continue
				}

				for i := y; i > 0; i-- {
					pt := g.Board[Point{x, i}]
					other := g.Board[Point{x, i - 1}]

					g.mv(pt, other)
				}
			}
		}
	case Down:
		for x := 0; x < 4; x++ {
			for y := 2; y >= 0; y-- {
				if g.Board[Point{x, y}].Value == 0 {
					continue
				}

				for i := y; i < 3; i++ {
					pt := g.Board[Point{x, i}]
					other := g.Board[Point{x, i + 1}]

					g.mv(pt, other)
				}
			}
		}
	case Left:
		for y := 0; y < 4; y++ {
			for x := 1; x < 4; x++ {
				if g.Board[Point{x, y}].Value == 0 {
					continue
				}

				for i := x; i > 0; i-- {
					pt := g.Board[Point{i, y}]
					other := g.Board[Point{i - 1, y}]

					g.mv(pt, other)
				}
			}
		}
	case Right:
		for y := 0; y < 4; y++ {
			for x := 2; x >= 0; x-- {
				if g.Board[Point{x, y}].Value == 0 {
					continue
				}

				for i := x; i < 3; i++ {
					pt := g.Board[Point{i, y}]
					other := g.Board[Point{i + 1, y}]

					g.mv(pt, other)
				}
			}
		}
	}

	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			g.Board[Point{x, y}].Blocked = false
		}
	}
}

func (g *Game) GameOver() {
	fmt.Println("Game Over!")
	fmt.Println("Score: ", g.Score)
}

func (g *Game) GameWon() {
	fmt.Println("You Won!")
	fmt.Println("Score: ", g.Score)
}
