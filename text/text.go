package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var winTitle string = "Text"
var winWidth, winHeight int32 = 800, 600

func initWindow() *sdl.Window {
    window, err := sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, winWidth, winHeight, sdl.WINDOW_SHOWN)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
        panic(err)
    }

    return window
}

func initRenderer(window *sdl.Window) *sdl.Renderer {
    renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		panic(err)
	}

	return renderer
}

func initTTF() {
    err := ttf.Init()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to initialize TTF: %s\n", err)
        panic(err)
    }
}

func initFont(size int) *ttf.Font {
    font, err := ttf.OpenFont("font.ttf", size)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open font: %s\n", err)
		panic(err)
	}

	return font
}

func renderText(str string, font *ttf.Font, color sdl.Color, renderer *sdl.Renderer) (text *Text, err error) {
	surface, err := font.RenderUTF8Blended(str, color)
	if err != nil {
		return
	}

	defer surface.Free()

	text = &Text{}
	text.Texture, err = renderer.CreateTextureFromSurface(surface)
	_, _, text.Width, text.Height, _ = text.Texture.Query()

	return
}

type Text struct {
	Width int32
	Height int32
	Texture *sdl.Texture
}

func run() int {
	window := initWindow(); defer window.Destroy()
	renderer := initRenderer(window); defer renderer.Destroy()

	initTTF()
    font := initFont(32); defer font.Close()

    var i int = 0

	for {
        for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
            switch event.(type) {
                case *sdl.QuitEvent:
                    return 0
            }
        }

        renderer.Clear()
        renderer.SetDrawColor(0, 0, 0, 255)

        i++
        text, _ := renderText("count: " + strconv.Itoa(i), font, sdl.Color{255, 255, 255, 255}, renderer)

        renderer.Copy(text.Texture, nil, &sdl.Rect{100, 100, text.Width, text.Height})

        renderer.Present()
        sdl.Delay(150)
    }

	return 0
}

func main() {
	os.Exit(run())
}
