package main

import (
	"fmt"
	"os"
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

var winTitle string = "Go-SDL2 Through Stars"
var winWidth, winHeight int32 = 800, 600
var cx float64 = float64(winWidth / 2)
var cy float64 = float64(winHeight / 2)
var lastTime time.Time

var num_stars int = 512
var max_depth float64 = 32
var stars [][3]float64

var rect sdl.Rect

const coolDownTime = time.Millisecond * 5

func random(min int, max int) float64 {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    return float64(r.Intn(max - min +1) + min)
}

func initStars() {
    for i := 0; i < num_stars; i++ {
        star := [3]float64{random(-25, 25), random(-25, 25), random(1, int(max_depth))}
        stars = append(stars, star)
    }
}

func moveStars(renderer *sdl.Renderer) {
    for i := 0; i < len(stars); i++ {
        star := &stars[i]

        if star[2] -= 0.2; star[2] <= 0 {
            star[0] = random(-25, 25)
            star[1] = random(-25, 25)
            star[2] = max_depth
        }

        // coords convert 3D to 2D for perspective projection
        k := float64(128.0 / star[2])
        x := int32(star[0] * k + cx)
        y := int32(star[1] * k + cy)

        if (0 <= x && x < winWidth) && (0 <= y && y < winHeight) {
            size := int32((1 - star[2] / max_depth) * 3)
            shade := uint8((1 - star[2] / max_depth) * 255)

            rect = sdl.Rect{x, y, size, size}
            renderer.SetDrawColor(shade, shade, shade, 255)
            renderer.FillRect(&rect)
        }
    }
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

	initStars()

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

    renderer.SetDrawColor(255, 255, 255, 255)
    moveStars(renderer)

    renderer.Present()
}

func main() {
	os.Exit(run())
}
