package registers

import (
	"unsafe"

	"runtime/volatile"
)

const (
	MemIo      = 0x04000000
	MemPalette = 0x05000000
	MemVram    = 0x06000000
)

var RegDispcnt = (*volatile.Register16)(unsafe.Pointer(uintptr(MemIo)))

const (
	DcntMode0 = 0x0000
	DcntMode1 = 0x0001
	DcntMode2 = 0x0002
	DcntMode3 = 0x0003
	DcntMode4 = 0x0004
	DcntMode5 = 0x0005

	DcntPage = 0x0010

	DcntBg0 = 0x0100
	DcntBg1 = 0x0200
	DcntBg2 = 0x0400
	DcntBg3 = 0x0800
	DcntObj = 0x1000

	ScreenWidth  = 240
	ScreenHeight = 160
)

var RegDispStat = (*volatile.Register16)(unsafe.Pointer(uintptr(MemIo + 0x0004)))

const (
	DStatVblIrq = 0x0008
)

var RegVCount = (*volatile.Register16)(unsafe.Pointer(uintptr(MemIo + 0x0006)))

var RegKeyInput = (*volatile.Register16)(unsafe.Pointer(uintptr(MemIo + 0x0130)))
