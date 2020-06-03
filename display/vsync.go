package display

import "github.com/janezkenda/tinygba/registers"

func VSync() {
	for registers.IO.LCD.VCount.Get() >= 160 {
	}
	for registers.IO.LCD.VCount.Get() < 160 {
	}
}
