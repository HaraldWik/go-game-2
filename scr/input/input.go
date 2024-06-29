package input

import (
	"github.com/veandco/go-sdl2/sdl"
)

func GetPressedKeys() []sdl.Keycode {
	keys := make([]sdl.Keycode, 0)

	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.KeyboardEvent:
			if t.Type == sdl.KEYDOWN {
				keys = append(keys, t.Keysym.Sym)
			} else if t.Type == sdl.KEYUP {
				// Remove key from keys slice if released
				for i, key := range keys {
					if key == t.Keysym.Sym {
						keys = append(keys[:i], keys[i+1:]...)
						break
					}
				}
			}
		}
	}
	return nil
}
