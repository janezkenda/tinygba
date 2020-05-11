package display

import "github.com/janezkenda/tinygba/registers"

func VSync() {
	for registers.RegVCount.Get() >= 160 {
	}
	for registers.RegVCount.Get() < 160 {
	}
}
