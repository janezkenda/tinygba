package display

import (
	"unsafe"

	"runtime/volatile"

	"github.com/janezkenda/tinygba/registers"
)

const (
	Page0Addr = VRam
	Page1Addr = VRam + 0xA000
)

type Mode4 struct {
	addr  int
	page0 *[160][120]volatile.Register16
	page1 *[160][120]volatile.Register16
	vmem  *[160][120]volatile.Register16
}

func NewMode4() *Mode4 {
	registers.RegDispcnt.Set(registers.DcntMode4 | registers.DcntBg2)

	page0 := (*[160][120]volatile.Register16)(unsafe.Pointer(uintptr(Page0Addr)))
	page1 := (*[160][120]volatile.Register16)(unsafe.Pointer(uintptr(Page1Addr)))

	return &Mode4{
		addr:  Page0Addr,
		vmem:  page0,
		page0: page0,
		page1: page1,
	}
}

func (m4 *Mode4) FlipPage() {
	registers.RegDispcnt.Set(registers.RegDispcnt.Get() ^ registers.DcntPage)
	if m4.addr == Page1Addr {
		m4.addr = Page0Addr
		m4.vmem = m4.page0
	} else {
		m4.addr = Page1Addr
		m4.vmem = m4.page1
	}
}

func (m4 *Mode4) Plot(x, y int, c uint8) {
	current := m4.vmem[y][x/2].Get()
	var num uint16
	if x&1 == 0 {
		num = (current & 0xFF00) | uint16(c)
	} else {
		num = (current & 0xFF) | uint16(c)<<8
	}

	m4.vmem[y][x/2].Set(num)
}

func (m4 *Mode4) Fill(c uint8) {
	for y, l := range m4.vmem {
		for x := range l {
			m4.vmem[y][x].Set(uint16(c)<<8 | uint16(c))
		}
	}
}
