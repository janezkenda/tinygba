package input

import "github.com/janezkenda/tinygba/registers"

type Key uint16

const KeyMask Key = 0x03FF

var current, previous Key

func Poll() {
	previous = current
	current = Key(^registers.RegKeyInput.Get()) & KeyMask
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

type Keypad struct {
	current, prev Key
}

func NewKeyPad() *Keypad {
	return &Keypad{}
}

func (k *Keypad) Poll() {
	k.prev = k.current
	k.current = Key(^registers.RegKeyInput.Get()) & KeyMask
}

func (k *Keypad) Current() Key {
	return k.current
}

func (k *Keypad) Previous() Key {
	return k.prev
}

func (k *Keypad) IsDown(key Key) bool {
	return k.current&key > 0
}

func (k *Keypad) IsUp(key Key) bool {
	return ^k.current&key > 0
}

func (k *Keypad) WasDown(key Key) bool {
	return k.prev&key > 0
}

func (k *Keypad) WasUp(key Key) bool {
	return ^k.prev&key > 0
}

func (k *Keypad) Transit(key Key) bool {
	return (k.current^k.prev)&key > 0
}

func (k *Keypad) Held(key Key) bool {
	return (k.current&k.prev)&key > 0
}

func (k *Keypad) Hit(key Key) bool {
	return (k.current & ^k.prev)&key > 0
}

func (k *Keypad) Released(key Key) bool {
	return (^k.current&k.prev)&key > 0
}
