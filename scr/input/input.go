package input

import (
	vec2 "github.com/HaraldWik/go-game-2/scr/vector/2"
	"github.com/veandco/go-sdl2/sdl"
)

func GetPressedKeys() []string {
	keys := sdl.GetKeyboardState()

	var pressedKeys []string

	for keyCode := sdl.Scancode(0); keyCode < sdl.NUM_SCANCODES; keyCode++ {
		if keys[keyCode] != 0 {
			keyName := sdl.GetKeyName(sdl.GetKeyFromScancode(keyCode))
			pressedKeys = append(pressedKeys, keyName)
		}
	}

	mouseButtonMap := map[string]uint32{
		M_LMB: 1,
		M_MMB: 2,
		M_RMB: 4,
		M_B4:  8,
		M_B5:  16,
	}

	for button, value := range mouseButtonMap {
		if _, _, state := sdl.GetMouseState(); state == value {
			pressedKeys = append(pressedKeys, button)
		}
	}

	return pressedKeys
}

func IsPressed(keycode string) bool {
	keys := GetPressedKeys()
	for _, key := range keys {
		if key == keycode {
			return true
		}
	}

	return false
}

func IsReleased(keycode string) bool {
	return !IsPressed(keycode)
}

var previousPressed = map[string]bool{}

func IsJustPressed(keycode string) bool {
	current := IsPressed(keycode)
	pressed := current && !previousPressed[keycode]
	previousPressed[keycode] = current

	return pressed
}

var previousReleased = map[string]bool{}

func IsJustReleased(keycode string) bool {

	current := IsPressed(keycode) && !IsJustPressed(keycode)
	released := !current && previousReleased[keycode]
	previousReleased[keycode] = current

	return released
}

func MousePosition() vec2.Type {
	x, y, _ := sdl.GetMouseState()

	return vec2.New(float32(x), float32(y))
}
