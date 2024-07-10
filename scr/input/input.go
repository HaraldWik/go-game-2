package input

import "github.com/veandco/go-sdl2/sdl"

func GetPressedKeys() []string {
	// Get the current state of the keyboard
	keys := sdl.GetKeyboardState()

	// Create a slice to hold the names of the pressed keys
	var pressedKeys []string

	// Iterate through all possible key codes
	for keyCode := sdl.Scancode(0); keyCode < sdl.NUM_SCANCODES; keyCode++ {
		if keys[keyCode] != 0 {
			keyName := sdl.GetKeyName(sdl.GetKeyFromScancode(keyCode))
			pressedKeys = append(pressedKeys, keyName)
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
