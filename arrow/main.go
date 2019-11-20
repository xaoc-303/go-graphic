package main

import (
	"fmt"
	"os"
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

var winTitle string = "Go-SDL2 Arrow"
var winWidth, winHeight int32 = 800, 600
var cx, cy int
var lastTime time.Time

var angle float64 = 0

const coolDownTime = time.Millisecond * 5

func msCoord(val float64, hlen float64) [2]int {
    var coord [2]int
    val *= 0.5
    var valPi = math.Pi * val / 180

    if val >= 0 && val <= 180 {
        coord[0] = cx + int(hlen * math.Sin(valPi))
        coord[1] = cy - int(hlen * math.Cos(valPi))
    } else {
        coord[0] = cx - int(hlen * -math.Sin(valPi))
        coord[1] = cy - int(hlen * math.Cos(valPi))
    }

    return coord
}

func run() int {
	var window *sdl.Window
	var renderer *sdl.Renderer

	window, err := sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return 1
	}
	defer window.Destroy()

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		return 2
	}
	defer renderer.Destroy()

	for {
        for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
            switch event.(type) {
                case *sdl.QuitEvent:
                    return 0
            }
        }

        tick(window, renderer)

        sdl.Delay(15)
    }

	return 0
}

func tick(window *sdl.Window, renderer *sdl.Renderer) {
    if (time.Since(lastTime) < coolDownTime) {
        return
    }

    renderer.SetDrawColor(0, 0, 0, 255)
    renderer.Clear()

    lastTime = time.Now()

    cx = int(winWidth) / 2
    cy = int(winHeight) / 2

    angle++
    var handCoord [2]int = msCoord(angle, 140)
    renderer.SetDrawColor(255, 255, 255, 255)
    renderer.DrawLine(int32(cx), int32(cy), int32(handCoord[0]), int32(handCoord[1]))

    renderer.Present()
}

func main() {
	os.Exit(run())
}
