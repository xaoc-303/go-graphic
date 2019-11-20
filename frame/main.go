package main

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/img"
)

var winTitle string = "Go-SDL2 Frame"
var winWidth, winHeight int32 = 800, 600

var fileName string = "./resources/balloon.png"

func run() int {
    // window
    var window *sdl.Window
    window, err := sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
    winWidth, winHeight, sdl.WINDOW_SHOWN)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
        panic(err)
    }
    defer window.Destroy()

	// renderer
	var renderer *sdl.Renderer
	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
        panic(err)
    }
    defer renderer.Destroy()

	// image
	var image *sdl.Surface
	image, err = img.Load(fileName)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to load PNG: %s\n", err)
        panic(err)
    }
    defer image.Free()

    // texture
    var texture *sdl.Texture
    texture, err = renderer.CreateTextureFromSurface(image)
    if err != nil {
        panic(err)
    }
    defer texture.Destroy()

    // set area image
    var src sdl.Rect
    src = sdl.Rect{0, 0, image.W, image.H}

    //rendering
    for {
        for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
            switch event.(type) {
                case *sdl.QuitEvent:
                    return 0
            }
        }

        renderer.Clear()
        renderer.SetDrawColor(0, 0, 188, 255)

        renderer.Copy(texture, &src, &sdl.Rect{100, 20, image.W/4, image.H/4})
        renderer.Copy(texture, &src, &sdl.Rect{300, 20, image.W/4, image.H/4})

        renderer.CopyEx(texture, &src, &sdl.Rect{100, 220, image.W/4, image.H/4}, 0.0, nil, sdl.FLIP_HORIZONTAL)
        renderer.CopyEx(texture, &src, &sdl.Rect{300, 220, image.W/4, image.H/4}, 0.0, nil, sdl.FLIP_VERTICAL)

        renderer.CopyEx(texture, &src, &sdl.Rect{500, 420, image.W/4, image.H/4}, 45.0, nil, sdl.FLIP_NONE)

        renderer.Present()

        sdl.Delay(15)
    }

	return 0
}

func main() {
	os.Exit(run())
}
