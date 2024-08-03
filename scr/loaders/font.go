package load

import (
	"image"
	"image/color"
	"log"
	"os"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/golang/freetype"
)

type Font struct {
	Path      string
	Context   *freetype.Context
	Size, DPI float32
}

func NewFont(path string, size, dpi float32) *Font {
	fontData, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read font file: %v", err)
	}

	font, err := freetype.ParseFont(fontData)
	if err != nil {
		log.Fatalf("failed to parse font: %v", err)
	}

	context := freetype.NewContext()
	context.SetDPI(float64(dpi))
	context.SetFont(font)
	context.SetFontSize(float64(size))
	context.SetClip(image.Rect(0, 0, 1000000, 1000000))
	context.SetDst(image.NewRGBA(image.Rect(0, 0, 10000, 10000)))
	context.SetSrc(image.NewUniform(color.White))

	return &Font{
		Path:    path,
		Context: context,
		Size:    size,
		DPI:     dpi,
	}
}

func (f *Font) RenderTextToTexture(text string) Texture {
	img := image.NewRGBA(image.Rect(0, 0, 10000, 10000))
	f.Context.SetDst(img)

	// Calculate the starting point for the text
	pt := freetype.Pt(10, 10+int(f.Context.PointToFixed(float64(f.Size))>>6)) // Position to start drawing
	_, err := f.Context.DrawString(text, pt)
	if err != nil {
		log.Fatalf("failed to draw string: %v", err)
	}

	var image uint32
	gl.GenTextures(1, &image)
	gl.BindTexture(gl.TEXTURE_2D, image)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(img.Bounds().Dx()), int32(img.Bounds().Dy()), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(img.Pix))

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)

	return Texture{
		Image: image,
	}
}
