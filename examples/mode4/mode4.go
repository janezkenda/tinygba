package main

import (
	"unsafe"

	"runtime/volatile"

	"github.com/janezkenda/tinygba/display"
)

func main() {
	page0 := (*[0xA000]volatile.Register32)(unsafe.Pointer(uintptr(display.Page0Addr)))
	page1 := (*[0xA000]volatile.Register32)(unsafe.Pointer(uintptr(display.Page1Addr)))

	for i := 0; i < 16; i++ {
		i1 := i * 144 / 4
		for j := 0; j < 144/4; j++ {
			page0[i*240/4+j].Set(frontBitmap[i1+j])
			page1[i*240/4+j].Set(backBitmap[i1+j])
		}
	}

	bgPal := display.NewBGPalette()
	for i, c := range palette {
		bgPal.SetColor(i, c)
	}

	m4 := display.NewMode4()

	var i int
	for {
		display.VSync()
		i++
		if i >= 9000 {
			i = 0
			m4.FlipPage()
		}
	}
}
