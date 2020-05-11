package display

type Color uint16

const (
	ClrBlack  Color = 0x0000
	ClrRed    Color = 0x001F
	ClrLime   Color = 0x03E0
	ClrYellow Color = 0x03FF
	ClrBlue   Color = 0x7C00
	ClrMag    Color = 0x7C1F
	ClrCyan   Color = 0x7FE0
	ClrWhite  Color = 0x7FFF
)

func RGB15(red, green, blue uint32) Color {
	return Color(red | green<<5 | blue<<10)
}
