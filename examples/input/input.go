package main

import (
	"unsafe"

	"runtime/volatile"

	"github.com/janezkenda/tinygba/display"
	"github.com/janezkenda/tinygba/input"
)

func main() {
	vmem := (*[0xA000]volatile.Register32)(unsafe.Pointer(uintptr(display.VRam)))

	for i, c := range gbaPicBitmap {
		(*vmem)[i].Set(c)
	}

	bp := display.NewBGPalette()
	for i, c := range gbaPicPalette {
		bp.SetColor(i, c)
	}

	display.NewMode4()

	kp := input.NewKeyPad()

	var frame int
	for {
		display.VSync()
		if frame == 900 {
			frame = 0
			kp.Poll()
		}

		for i := 0; i <= 9; i++ {
			var c display.Color
			k := input.Key(1 << i)
			if kp.Hit(k) {
				c = display.ClrRed
			} else if kp.Released(k) {
				c = display.ClrYellow
			} else if kp.Held(k) {
				c = display.ClrLime
			} else {
				c = display.RGB15(27, 27, 29)
			}
			bp.SetColor(5+i, c)
		}
		frame++
	}
}
