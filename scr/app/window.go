package app

import (
	"log"
	"runtime"

	vec2 "github.com/HaraldWik/go-game-2/scr/vector/2"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/veandco/go-sdl2/sdl"
)

type Window struct {
	Name string

	size    vec2.Type
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

	MaxFPS uint32
}

func (app *App) NewWindow(name string, size vec2.Type) Window {
	window := Window{
		Name:               name,
		size:               size,
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
	app.WindowList = append(app.WindowList, window)
	return window
}

func (w *Window) Open() {
	// Lock the main thread for SDL2 and OpenGL
	runtime.LockOSThread()

	// Initialize SDL
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		log.Fatalf("Failed to init window videoe:\n%v\n", err)
	}

	// Create an SDL window
	var err error
	w.SDL, err = sdl.CreateWindow(w.Name, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(w.size.X), int32(w.size.Y), w.Flags|sdl.WINDOW_OPENGL)
	if err != nil {
		log.Fatalf("Failed to create window %s:\n%v\n", w.Name, err)
	}

	if w.MinSize != vec2.Zero() {
		w.SDL.SetMinimumSize(int32(w.MinSize.X), int32(w.MinSize.Y))
	}
	if w.MaxSize != vec2.Zero() {
		w.SDL.SetMaximumSize(int32(w.MinSize.X), int32(w.MinSize.Y))
	}

	// Set OpenGL attributes for version 2.1
	sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, 2)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_MINOR_VERSION, 1)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_PROFILE_MASK, sdl.GL_CONTEXT_PROFILE_CORE)
	sdl.GLSetAttribute(sdl.GL_DOUBLEBUFFER, 1)

	// Create an OpenGL context
	if _, err = w.SDL.GLCreateContext(); err != nil {
		log.Fatalf("Failed to create OpenGL Context on window %s:\n%v\n", w.Name, err)
	}

	// Initialize OpenGL
	if err := gl.Init(); err != nil {
		log.Fatalf("Failed to init OpenGL on window %s:\n%v\n", w.Name, err)
	}
}

func (w *Window) BeginDraw() {
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LEQUAL)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (w *Window) EndDraw() {
	w.SDL.GLSwap()
	sdl.Delay(uint32(1000 / w.MaxFPS))
}

func (w *Window) Close() {
	w.SDL.Destroy()
}

func (w *Window) CloseEvent() bool {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			return true
		}
	}
	return false
}

func (w *Window) GetSize() vec2.Type {
	x, y := w.SDL.GetSize()
	return vec2.New(float32(x), float32(y))
}

func (w *Window) SetSize(size vec2.Type) {
	w.SDL.SetSize(int32(size.X), int32(size.Y))
}

func (w *Window) Minimize() {
	w.SDL.Minimize()
}

func (w *Window) Maximize() {
	w.SDL.Maximize()
}

func (w *Window) SetAlwaysOnTop(onTop bool) {
	w.SDL.SetAlwaysOnTop(onTop)
}

func (w *Window) Hide() {
	w.SDL.Hide()
}

func (w *Window) Show() {
	w.SDL.Show()
}

var lastFrameTime uint32

func (w *Window) GetDeltaTime() float32 {
	// Initialize timing if it's the first call
	if lastFrameTime == 0 {
		lastFrameTime = sdl.GetTicks()
		return 0.0 // Return 0 for the first frame as there's no previous delta time
	}

	currentTime := sdl.GetTicks()
	deltaTime := float32(currentTime-lastFrameTime) / 1000.0
	lastFrameTime = currentTime
	return deltaTime
}
