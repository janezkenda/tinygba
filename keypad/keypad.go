package keypad

import "github.com/janezkenda/tinygba/registers"

type Key uint16

var current, previous Key

func Poll() {
	previous = current
	current = Key(^registers.RegKeyInput.Get()) & 0x03FF
}

func (k Key) IsDown() bool {
	return current&k > 0
}

func (k Key) IsUp() bool {
	return ^current&k > 0
}

func (k Key) WasDown() bool {
	return previous&k > 0
}

func (k Key) WasUp() bool {
	return ^previous&k > 0
}

func (k Key) Transit() bool {
	return (current^previous)&k > 0
}

func (k Key) Held() bool {
	return (current&previous)&k > 0
}

func (k Key) Hit() bool {
	return (current & ^previous)&k > 0
}

func (k Key) Released() bool {
	return (^current&previous)&k > 0
}

const (
	A Key = 1 << iota
	B
	Select
	Start
	Right
	Left
	Up
	Down
	R
	L
)
