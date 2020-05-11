package input

import "github.com/janezkenda/tinygba/registers"

type Key uint16

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

const KeyMask Key = 0x03FF

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
