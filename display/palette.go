package display

import (
	"unsafe"

	"runtime/volatile"

	"github.com/janezkenda/tinygba/registers"
)

const (
	PaletteBG     = registers.MemPalette
	PaletteSprite = registers.MemPalette + 0x200
)

type Palette struct {
	mem *[256]volatile.Register16
}

func newPalette(addr int) *Palette {
	return &Palette{
		mem: (*[256]volatile.Register16)(unsafe.Pointer(uintptr(addr))),
	}
}

func NewBGPalette() *Palette {
	return newPalette(PaletteBG)
}

func NewSpritePalette() *Palette {
	return newPalette(PaletteSprite)
}

func (p *Palette) SetColor(i int, c Color) {
	p.mem[i].Set(uint16(c))
}
