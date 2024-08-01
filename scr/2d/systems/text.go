package s2d

import (
	"log"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func CreateTextureFromText(fontPath string, text string, fontSize int, color sdl.Color) uint32 {
	// Load the font
	font, err := ttf.OpenFont(fontPath, fontSize)
	if err != nil {
		log.Fatalf("Error loading font: %v", err)
	}
	defer font.Close()

	// Render text to SDL surface
	textSurface, err := font.RenderUTF8Solid(text, color)
	if err != nil {
		log.Fatalf("Error rendering text: %v", err)
	}
	defer textSurface.Free()

	// Create OpenGL texture from SDL surface
	var texture uint32
	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	width, height := textSurface.W, textSurface.H
	pixels := textSurface.Pixels()

	// Upload the texture data
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(width), int32(height), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(pixels))

	return texture
}
