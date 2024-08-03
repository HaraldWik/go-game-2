package app

import (
	"image"
	"log"
	"runtime"

	gfx "github.com/HaraldWik/go-game-2/scr/graphics"
	vec2 "github.com/HaraldWik/go-game-2/scr/vector/2"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/veandco/go-sdl2/sdl"
)

type Window struct {
	Name string

	size    vec2.Type
	minSize vec2.Type
	maxSize vec2.Type

	flags uint32

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

	maxFPS uint32
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
		log.Fatalf("Failed to init window videoe:%v", err)
	}

	// Create an SDL window
	var err error
	w.SDL, err = sdl.CreateWindow(w.Name, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(w.size.X), int32(w.size.Y), w.flags|sdl.WINDOW_OPENGL)
	if err != nil {
		log.Fatalf("Failed to create window %s:%v", w.Name, err)
	}

	if w.minSize != vec2.Zero() {
		w.SDL.SetMinimumSize(int32(w.minSize.X), int32(w.minSize.Y))
	}
	if w.maxSize != vec2.Zero() {
		w.SDL.SetMaximumSize(int32(w.minSize.X), int32(w.minSize.Y))
	}

	// Set OpenGL attributes for version 2.1
	sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, 2)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_MINOR_VERSION, 1)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_PROFILE_MASK, sdl.GL_CONTEXT_PROFILE_CORE)
	sdl.GLSetAttribute(sdl.GL_DOUBLEBUFFER, 1)

	// Create an OpenGL context
	if _, err = w.SDL.GLCreateContext(); err != nil {
		log.Fatalf("Failed to create OpenGL Context on window %s:%v", w.Name, err)
	}

	// Initialize OpenGL
	if err := gl.Init(); err != nil {
		log.Fatalf("Failed to init OpenGL on window %s: %v", w.Name, err)
	}
}

func (w *Window) BeginDraw() {
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LEQUAL)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (w *Window) EndDraw() {
	gfx.GFX2D.DrawCycle()
	w.SDL.GLSwap()
	sdl.Delay(uint32(1000 / w.maxFPS))
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

func (w *Window) SetSize(size vec2.Type) {
	w.SDL.SetSize(int32(size.X), int32(size.Y))
}

func (w *Window) GetSize() vec2.Type {
	x, y := w.SDL.GetSize()
	return vec2.New(float32(x), float32(y))
}

func (w *Window) SetMinSize(size vec2.Type) {
	w.SDL.SetMinimumSize(int32(size.X), int32(size.Y))
}

func (w *Window) GetMinSize() vec2.Type {
	return w.minSize
}

func (w *Window) SetMaxSize(size vec2.Type) {
	w.SDL.SetMaximumSize(int32(size.X), int32(size.Y))
}

func (w *Window) GetMaxSize() vec2.Type {
	return w.maxSize
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

func (w *Window) SetIcon(image image.Image) {
	bounds := image.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	// Create a new SDL surface with the same dimensions as the image
	surface, err := sdl.CreateRGBSurface(0, int32(width), int32(height), 32,
		0x0000FF, 0x00FF00, 0xFF0000, 0xFF000000)
	if err != nil {
		log.Fatalf("%v", err)
	}

	// Lock the surface to access its pixels
	surface.Lock()
	defer surface.Unlock()

	// Get the surface pixels
	pixels := surface.Pixels()

	// Convert the Go image to SDL format
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Get the color of the pixel from the Go image
			r, g, b, a := image.At(x, y).RGBA()
			// Convert to 8-bit color values
			r8 := uint8(r >> 8)
			g8 := uint8(g >> 8)
			b8 := uint8(b >> 8)
			a8 := uint8(a >> 8)

			// Calculate the index in the surface pixel array
			index := (y*int(surface.Pitch) + x*4)

			// Set the pixel color in the SDL surface
			pixels[index] = r8
			pixels[index+1] = g8
			pixels[index+2] = b8
			pixels[index+3] = a8
		}
	}

	w.SDL.SetIcon(surface)
}

func (w *Window) SetMaxFPS(value uint32) {
	w.maxFPS = value
}

func (w *Window) GetMaxFPS() uint32 {
	return w.maxFPS
}

func (w *Window) SetFlags(flags uint32) {
	w.flags = flags
}

func (w *Window) GetFlags() uint32 {
	return w.flags
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
