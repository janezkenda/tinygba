package main

import "github.com/janezkenda/tinygba/display"

func main() {
	m3 := display.NewMode3()

	m3.Fill(display.RGB15(12, 12, 14))

	m3.Rect(12, 8, 108, 72, display.ClrRed)
	m3.Rect(108, 72, 132, 88, display.ClrLime)
	m3.Rect(132, 88, 228, 152, display.ClrBlue)

	for i := 0; i <= 8; i++ {
		j := uint32(3*i + 7)
		m3.Line(132+11*i, 9, 226, 12+7*i, display.RGB15(j, 0, j))
		m3.Line(226-11*i, 70, 133, 69-7*i, display.RGB15(j, 0, j))

		m3.Line(15+11*i, 88, 104-11*i, 150, display.RGB15(0, j, j))
	}

	m3.Frame(132, 8, 228, 72, display.ClrCyan)
	m3.Frame(109, 73, 131, 87, display.ClrBlack)
	m3.Frame(12, 88, 108, 152, display.ClrYellow)

	for {
	}
}
