package display

import "runtime/volatile"

func bmp16Line(x1, y1, x2, y2 int, c Color, vmem *[160][240]volatile.Register16) {
	clr := uint16(c)

	var dx, xstep int
	if x1 > x2 {
		xstep = -1
		dx = x1 - x2
	} else {
		xstep = 1
		dx = x2 - x1
	}

	var dy, ystep int
	if y1 > y2 {
		ystep = -1
		dy = y1 - y2
	} else {
		ystep = 1
		dy = y2 - y1
	}

	if dy == 0 {
		for i := 0; i <= dx; i++ {
			vmem[y1][x1+i].Set(clr)
		}
	} else if dy == 0 {
		for i := 0; i <= dy; i++ {
			vmem[y1+i][x1].Set(clr)
		}
	} else if dx >= dy {
		dd := 2*dy - dx

		for i := 0; i <= dx; i++ {
			vmem[y1][x1].Set(clr)
			if dd >= 0 {
				dd -= 2 * dx
				y1 += ystep
			}
			dd += 2 * dy
			x1 += xstep
		}
	} else {
		dd := 2*dx - dy

		for i := 0; i <= dy; i++ {
			vmem[y1][x1].Set(clr)
			if dd >= 0 {
				dd -= 2 * dy
				x1 += xstep
			}
			dd += 2 * dx
			y1 += ystep
		}
	}
}

func bmp16Rect(left, top, right, bottom int, c Color, vmem *[160][240]volatile.Register16) {
	width := right - left
	height := bottom - top

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			vmem[top+y][left+x].Set(uint16(c))
		}
	}
}

func bmp16Frame(left, top, right, bottom int, c Color, vmem *[160][240]volatile.Register16) {
	right--
	bottom--

	bmp16Line(left, top, right, top, c, vmem)
	bmp16Line(left, bottom, right, bottom, c, vmem)

	bmp16Line(left, top, left, bottom, c, vmem)
	bmp16Line(right, top, right, bottom, c, vmem)
}
