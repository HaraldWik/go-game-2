package app

import (
	"runtime"

	debug "github.com/HaraldWik/debug/scr"
	vec2 "github.com/HaraldWik/go-game-2/scr/vector/2"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/veandco/go-sdl2/sdl"
)

type Win struct {
	Name string

	Size    vec2.Type
	MinSize vec2.Type
	MaxSize vec2.Type

	Flags uint32

	sdl *sdl.Window

	FLAG_RESIZABLE     uint32 // Allowes window to be resized
	FLAG_FULLSCREEN    uint32 // Makes window fullscreen
	FLAG_MINIMIZED     uint32 // Makes window minimized
	FLAG_MAXIMIZED     uint32 // Makes window maximized
	FLAG_BORDERLESS    uint32 // Removes window borders (Not recomended)
	FLAG_SHOW          uint32 // Showes the window
	FLAG_HIDE          uint32 // Hides the window
	FLAG_FOREIGN       uint32 // In case the window is made by another framework then SDL2
	FLAG_ALLOW_HIGHDPI uint32 // (Window should be created in high-DPI mode if supported)
	FLAG_ALWAYS_ON_TOP uint32 // Makes so the window is allways above other windows
	FLAG_UTILITY       uint32 // Makes so window is treated as a utility
	FLAG_TOOLTIP       uint32 // Makes so window is treated as a tooltip
	FLAG_POPUP_MENU    uint32 // Makes so window is treated as a popup menu
}

func (app *App) NewWindow(name string, size vec2.Type) Win {
	win := Win{
		Name:               name,
		Size:               size,
		FLAG_RESIZABLE:     0x00000020,
		FLAG_FULLSCREEN:    0x00000001,
		FLAG_MINIMIZED:     0x00001000,
		FLAG_MAXIMIZED:     0x00000040,
		FLAG_BORDERLESS:    0x00000010,
		FLAG_SHOW:          0x00000004,
		FLAG_HIDE:          0x00000008,
		FLAG_FOREIGN:       0x00000800,
		FLAG_ALLOW_HIGHDPI: 0x00002000,
		FLAG_ALWAYS_ON_TOP: 0x00008000,
		FLAG_UTILITY:       0x00020000,
		FLAG_TOOLTIP:       0x00040000,
		FLAG_POPUP_MENU:    0x00080000,
	}
	app.WindowList = append(app.WindowList, win)
	return win
}

func (win *Win) Open() {
	// Window creation
	runtime.LockOSThread()
	debug.Error(sdl.Init(sdl.INIT_EVERYTHING))
	defer runtime.UnlockOSThread()

	var err error
	win.sdl, err = sdl.CreateWindow(win.Name, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(win.Size.X), int32(win.Size.Y), win.Flags|0x00000002)
	debug.Error(err)

	// OpenGL creation
	debug.Error(sdl.Init(uint32(sdl.INIT_VIDEO)))

	sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, 3)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_MINOR_VERSION, 3)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_PROFILE_MASK, sdl.GL_CONTEXT_PROFILE_CORE)

	_, err = win.sdl.GLCreateContext()
	debug.Error(err)

	// Initialize OpenGL
	debug.Error(gl.Init())

	// Enable depth test
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LEQUAL)
}

func (win Win) BeginDraw(r, g, b float32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.ClearColor(r, g, b, 1.0)
}

func (win Win) EndDrawOpenGL(maxFps int32) {
	win.sdl.GLSwap()
	sdl.Delay(uint32(1000 / maxFps))
}

func (win Win) Close() {
	win.sdl.Destroy()
}

func (win Win) CloseEvent() bool {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			return true
		}
	}
	return false
}

func (window Win) IsActive() bool {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			return true
		case *sdl.WindowEvent:
			return true
		case *sdl.KeyboardEvent:
			return true
		case *sdl.MouseButtonEvent:
			return true
		case *sdl.MouseMotionEvent:
			return true
		case *sdl.MouseWheelEvent:
			return true
		case *sdl.JoyButtonEvent:
			return true
		case *sdl.JoyAxisEvent:
			return true
		case *sdl.ControllerButtonEvent:
			return true
		case *sdl.ControllerAxisEvent:
			return true
		case *sdl.ControllerDeviceEvent:
			return true
		case *sdl.TouchFingerEvent:
			return true
		case *sdl.MultiGestureEvent:
			return true
		case *sdl.DollarGestureEvent:
			return true
		case *sdl.DropEvent:
			return true
		case *sdl.ClipboardEvent:
			return true
		default:
			return false
		}
	}
	return false
}
