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

	SDL *sdl.Window

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
	// Lock the main thread for SDL2 and OpenGL
	runtime.LockOSThread()

	// Initialize SDL

	debug.Error(sdl.Init(sdl.INIT_VIDEO))

	// Create an SDL window
	var err error
	win.SDL, err = sdl.CreateWindow(win.Name, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(win.Size.X), int32(win.Size.Y), win.Flags|sdl.WINDOW_OPENGL)
	debug.Error(err)
	if win.MinSize != vec2.Zero() {
		win.SDL.SetMinimumSize(int32(win.MinSize.X), int32(win.MinSize.Y))
	}
	if win.MaxSize != vec2.Zero() {
		win.SDL.SetMaximumSize(int32(win.MinSize.X), int32(win.MinSize.Y))
	}

	// Set OpenGL attributes for version 2.1
	sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, 2)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_MINOR_VERSION, 1)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_PROFILE_MASK, sdl.GL_CONTEXT_PROFILE_CORE)

	// Create an OpenGL context
	_, err = win.SDL.GLCreateContext()
	debug.Error(err)

	// Initialize OpenGL
	debug.Error(gl.Init())

	// Enable depth test
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LEQUAL)
}

func (win *Win) BeginDraw(r, g, b float32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.ClearColor(r, g, b, 1.0)
}

func (win *Win) EndDraw(maxFps int32) {
	win.SDL.GLSwap()
	sdl.Delay(uint32(1000 / maxFps))
}

func (win *Win) Close() {
	win.SDL.Destroy()
}

func (win *Win) CloseEvent() bool {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			return true
		}
	}
	return false
}

func (win *Win) Update() {
	x, y := win.SDL.GetSize()
	win.SDL.SetSize(x, y)
}

func (win *Win) GetSize() vec2.Type {
	x, y := win.SDL.GetSize()
	return vec2.New(float32(x), float32(y))
}

func (win *Win) SetSize(size vec2.Type) {
	win.SDL.SetSize(int32(size.X), int32(size.Y))
}
