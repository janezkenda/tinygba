package display

import "image/color"

type Color uint16

func (c Color) ToRGB() color.RGBA {
	return color.RGBA{
		R: uint8(c & 0x1F),
		G: uint8(c >> 5 & 0x1F),
		B: uint8(c >> 10 & 0x1F),
	}
}

func ColorFromRGB(c color.RGBA) Color {
	return Color(c.B)<<10 | Color(c.G)<<5 | Color(c.R)
}

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

func Gradient(min, max Color, size int) []Color {
	minC := min.ToRGB()
	maxC := max.ToRGB()

	rDiff := int(maxC.R)-int(minC.R)
	gDiff := int(maxC.G)-int(minC.G)
	bDiff := int(maxC.B)-int(minC.B)

	var gradient []Color
	for i := 0; i < size; i++ {

		gradient = append(gradient, ColorFromRGB(color.RGBA{
			R: uint8(int(minC.R) + rDiff*i/size),
			G: uint8(int(minC.G) + gDiff*i/size),
			B: uint8(int(minC.B) + bDiff*i/size),
		}))
	}
	return gradient
}
