package in3d

import (
	"errors"
	"image"
	"image/draw"
	"os"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/seemywingz/go-toolbox"
)

// NewTexture : greate GL reference to provided texture
func NewTexture(file string) uint32 {
	imgFile, err := os.Open(file)
	toolbox.EoE(err, "Error Loading Texture")

	img, _, err := image.Decode(imgFile)
	toolbox.EoE(err, "Error Decoding Image")

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		toolbox.EoE(errors.New("unsupported stride"), "Error Getting RGB Strride")
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	var texture uint32
	gl.Enable(gl.TEXTURE_2D)
	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))
	gl.BindTexture(gl.TEXTURE_2D, 0)
	gl.Disable(gl.TEXTURE_2D)
	return texture
}
