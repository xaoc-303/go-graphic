package text

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Text struct {
	width int32
	height int32
	texture *sdl.Texture
}

func InitTTF() {
    err := ttf.Init()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to initialize TTF: %s\n", err)
        panic(err)
    }
}

func InitFont(size int) *ttf.Font {
    font, err := ttf.OpenFont("./resources/font.ttf", size)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open font: %s\n", err)
		panic(err)
	}

	return font
}

func RenderText(str string, font *ttf.Font, color sdl.Color, renderer *sdl.Renderer) (text *Text, err error) {
	surface, err := font.RenderUTF8Blended(str, color)
	if err != nil {
		return
	}

	defer surface.Free()

	text = &Text{}
	text.texture, err = renderer.CreateTextureFromSurface(surface)
	_, _, text.width, text.height, _ = text.texture.Query()

	return
}

func (text *Text) DrawXY(x, y int32, renderer *sdl.Renderer) {
    renderer.Copy(text.texture, nil, &sdl.Rect{x, y, text.width, text.height})
}
