package registers

import (
	"unsafe"

	"runtime/volatile"
)

var IO = (*IORegisters)(unsafe.Pointer(uintptr(0x04000000)))

type IORegisters struct {
	LCD lcdRegisters
}

type lcdRegisters struct {
	DispCnt  volatile.Register16    // R/W  DISPCNT   LCD Control
	_        volatile.Register16    // R/W  -         Undocumented - Green Swap
	DispStat volatile.Register16    // R/W  DISPSTAT  General LCD Status (STAT,LYC)
	VCount   volatile.Register16    // R    VCOUNT    Vertical Counter (LY)
	Bg0Cnt   volatile.Register16    // R/W  BG0CNT    BG0 Control
	Bg1Cnt   volatile.Register16    // R/W  BG1CNT    BG1 Control
	Bg2Cnt   volatile.Register16    // R/W  BG2CNT    BG2 Control
	Bg3Cnt   volatile.Register16    // R/W  BG3CNT    BG3 Control
	Bg0HOfs  volatile.Register16    // W    BG0HOFS   BG0 X-Offset
	Bg0VOfs  volatile.Register16    // W    BG0VOFS   BG0 Y-Offset
	Bg1HOfs  volatile.Register16    // W    BG1HOFS   BG1 X-Offset
	Bg1VOfs  volatile.Register16    // W    BG1VOFS   BG1 Y-Offset
	Bg2HOfs  volatile.Register16    // W    BG2HOFS   BG2 X-Offset
	Bg2VOfs  volatile.Register16    // W    BG2VOFS   BG2 Y-Offset
	Bg3HOfs  volatile.Register16    // W    BG3HOFS   BG3 X-Offset
	Bg3VOfs  volatile.Register16    // W    BG3VOFS   BG3 Y-Offset
	Bg2PA    volatile.Register16    // W    BG2PA     BG2 Rotation/Scaling Parameter A (dx)
	Bg2PB    volatile.Register16    // W    BG2PB     BG2 Rotation/Scaling Parameter B (dmx)
	Bg2PC    volatile.Register16    // W    BG2PC     BG2 Rotation/Scaling Parameter C (dy)
	Bg2PD    volatile.Register16    // W    BG2PD     BG2 Rotation/Scaling Parameter D (dmy)
	Bg2X     volatile.Register16    // W    BG2X      BG2 Reference Point X-Coordinate
	Bg2Y     volatile.Register16    // W    BG2Y      BG2 Reference Point Y-Coordinate
	Bg3PA    volatile.Register16    // W    BG3PA     BG3 Rotation/Scaling Parameter A (dx)
	Bg3PB    volatile.Register16    // W    BG3PB     BG3 Rotation/Scaling Parameter B (dmx)
	Bg3PC    volatile.Register16    // W    BG3PC     BG3 Rotation/Scaling Parameter C (dy)
	Bg3PD    volatile.Register16    // W    BG3PD     BG3 Rotation/Scaling Parameter D (dmy)
	Bg3X     volatile.Register16    // W    BG3X      BG3 Reference Point X-Coordinate
	Bg3Y     volatile.Register16    // W    BG3Y      BG3 Reference Point Y-Coordinate
	Win0H    volatile.Register16    // W    WIN0H     Window 0 Horizontal Dimensions
	Win1H    volatile.Register16    // W    WIN1H     Window 1 Horizontal Dimensions
	Win0V    volatile.Register16    // W    WIN0V     Window 0 Vertical Dimensions
	Win1V    volatile.Register16    // W    WIN1V     Window 1 Vertical Dimensions
	WinIn    volatile.Register16    // R/W  WININ     Inside of Window 0 and 1
	WinOut   volatile.Register16    // R/W  WINOUT    Inside of OBJ Window & Outside of Windows
	Mosaic   volatile.Register16    // W    MOSAIC    Mosaic Size
	_        volatile.Register16    // -    -         Not used
	BldCnt   volatile.Register16    // R/W  BLDCNT    Color Special Effects Selection
	BldAlpha volatile.Register16    // R/W  BLDALPHA  Alpha Blending Coefficients
	BldY     volatile.Register16    // W    BLDY      Brightness (Fade-In/Out) Coefficient
	_        [4]volatile.Register16 // -    -         Not used
}

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

const (
	DStatVblIrq = 0x0008
)

var RegKeyInput = (*volatile.Register16)(unsafe.Pointer(uintptr(MemIo + 0x0130)))
