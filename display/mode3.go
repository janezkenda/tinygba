package display

import (
	"image/color"
	"unsafe"

	"runtime/volatile"

	"github.com/janezkenda/tinygba/registers"
)

const VRam = 0x06000000

type Mode3 struct {
	vmem *[160][240]volatile.Register16
}

func NewMode3() *Mode3 {
	registers.RegDispcnt.Set(registers.DcntMode3 | registers.DcntBg2)
	return &Mode3{
		vmem: (*[160][240]volatile.Register16)(unsafe.Pointer(uintptr(VRam))),
	}
}

func (m3 *Mode3) Plot(x, y int, c Color) {
	m3.vmem[y][x].Set(uint16(c))
}

func (m3 *Mode3) Line(x1, y1, x2, y2 int, c Color) {
	bmp16Line(x1, y1, x2, y2, c, m3.vmem)
}

func (m3 *Mode3) Rect(left, top, right, bottom int, c Color) {
	bmp16Rect(left, top, right, bottom, c, m3.vmem)
}

func (m3 *Mode3) Frame(left, top, right, bottom int, c Color) {
	bmp16Frame(left, top, right, bottom, c, m3.vmem)
}

func (m3 *Mode3) Fill(c Color) {
	for y, l := range m3.vmem {
		for x := range l {
			m3.vmem[y][x].Set(uint16(c))
		}
	}
}

// Methods for implementing the Displayer interface, so it can be used with TinyFont.
func (m3 *Mode3) Size() (int16, int16) {
	return 240, 160
}

func (m3 *Mode3) SetPixel(x, y int16, c color.RGBA) {
	m3.Plot(int(x), int(y), ColorFromRGB(c))
}

func (m3 *Mode3) Display() error {
	return nil
}
