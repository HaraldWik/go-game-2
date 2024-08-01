package load

import (
	"image"
	"image/draw"
	"log"
	"os"

	vec2 "github.com/HaraldWik/go-game-2/scr/vector/2"
	"github.com/go-gl/gl/v2.1/gl"
)

type Texture struct {
	Image uint32
	Size  vec2.Type
}

func EmptyTexture() Texture {
	return Texture{
		Image: 0,
		Size:  vec2.Zero(),
	}
}

func NewTexture(filePath string) Texture {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("could not open file: %v\n", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatalf("could not decode image: %v\n", err)
	}

	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{}, draw.Src)

	var image uint32
	gl.GenTextures(1, &image)
	gl.BindTexture(gl.TEXTURE_2D, image)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(rgba.Bounds().Dx()), int32(rgba.Bounds().Dy()), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(rgba.Pix))

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)

	return Texture{
		Image: image,
		Size: vec2.New(
			float32(rgba.Bounds().Dx()),
			float32(rgba.Bounds().Dy()),
		),
	}
}
