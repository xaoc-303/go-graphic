package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/veandco/go-sdl2/sdl"

	"./lib/animation"
	"./lib/audio"
	"./lib/text"
	. "./lib/random"
)

var winTitle string = "Go-SDL2 Video Game"
var winWidth, winHeight int32 = 800, 450

const (
    directionRight = iota
    directionLeft
)

var direction = directionRight

const (
	stateStay = iota
	stateWalk
	stateJump
	stateJumpWalk
)

var state int = stateStay

var jumpCount, jumpCountMax int = 0, 2;

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

func run() int {
    var score int = 0;

    window := initWindow(); defer window.Destroy()
	renderer := initRenderer(window); defer renderer.Destroy()

	text.InitTTF()
    font := text.InitFont(14); defer font.Close()

    bg := animation.NewSprite("./resources/img/cyberpunk_street_background.png", 1, 0, renderer); defer bg.Free()
    mg := animation.NewSprite("./resources/img/cyberpunk_street_midground.png", 1, 0, renderer); defer mg.Free()
    fg := animation.NewSprite("./resources/img/cyberpunk_street_foreground.png", 1, 0, renderer); defer fg.Free()

    sc := animation.NewSprite("./resources/img/scarfy.png", 6, 123, renderer); defer sc.Free()
    sc.SetDst(100, 290, sc.GetFrameW(), sc.GetFrameH())
    sc.SetTimeEnabled(true)

    var positionX float64 = 100.0
    var positionY float64 = 290.0
    var velocityX float64 = 4.0
    var velocityY float64 = 0.0
    var gravity float64 = 0.5

    enemy := animation.NewSprite("./resources/img/enemy.png", 10, 50, renderer); defer enemy.Free()
    enemy.SetTimeEnabled(true)
    enemyBig := animation.NewSprite("./resources/img/enemy.png", 10, 50, renderer); defer enemyBig.Free()
    enemyBig.SetTimeEnabled(true)

    star := animation.NewSprite("./resources/img/coin-star.png", 6, 120, renderer); defer star.Free()
    star.SetTimeEnabled(true)

    // render

    var enemyOffsetX, enemyBigOffsetX float64 = 0.0, 0.0
    var starOffsetX float64 = 0.0
    var starY int32 = 300

    music := audio.LoadMusic("./resources/sound/track.wav"); defer music.Free()
    music.Play()

    sfxCoin := audio.LoadSound("./resources/sound/coin.wav"); defer sfxCoin.Sound.Free()
    sfxCoin.Sound.Volume(90)

    sfxKick := audio.LoadSound("./resources/sound/kick.wav"); defer sfxKick.Sound.Free()
    sfxScoreDown := audio.LoadSound("./resources/sound/score-down.wav"); defer sfxScoreDown.Sound.Free()

	for {
        for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
            switch t := event.(type) {
                case *sdl.KeyboardEvent:
                    if t.Type == 768 { // press key
                        if (t.Repeat == 0 && sdl.GetKeyName(t.Keysym.Sym) == "Up") {
                            if state == stateWalk || state == stateJumpWalk {
                                state = stateJumpWalk
                            } else {
                                state = stateJump
                            }
                            sc.SetTimeEnabled(false)
                            if jumpCount++; jumpCount <= jumpCountMax {
                                velocityY = -12.0
                            }
                        }

                        switch sdl.GetKeyName(t.Keysym.Sym) {
                            case "Right":
                                direction = directionRight
                                if state != stateJump && state != stateJumpWalk {
                                    state = stateWalk
                                }
                            case "Left":
                                direction = directionLeft
                                if state != stateJump && state != stateJumpWalk {
                                    state = stateWalk
                                }
                        }
                    }
                    if t.Type == 769 { // unpressed key
                        switch sdl.GetKeyName(t.Keysym.Sym) {
                            case "Right":
                                direction = directionRight
                                if state != stateJump && state != stateJumpWalk {
                                    state = stateStay
                                }
                            case "Left":
                                direction = directionLeft
                                if state != stateJump && state != stateJumpWalk {
                                    state = stateStay
                                }
                        }
                    }
                case *sdl.QuitEvent:
                    return 0
            }
        }

        renderer.Clear()
        renderer.SetDrawColor(0, 0, 0, 255)

        if state == stateWalk || state == stateJumpWalk {
            if direction == directionRight {
                bg.OffsetNextX(-0.1, float64(-winWidth), 0)
                mg.OffsetNextX(-0.5, float64(-winWidth), 0)
                fg.OffsetNextX(-2.0, float64(-winWidth), 0)
                if starOffsetX -= 2.0; starOffsetX <= 0 {
                    starOffsetX = float64(Random(0, int(winWidth)) + float64(winWidth/2))
                    starY = int32(Random(100, 250))
                }
                if enemyOffsetX -= 4.0; enemyOffsetX <= 0 {
                    enemyOffsetX = float64(Random(0, int(winWidth)) + float64(winWidth))
                }
                if enemyBigOffsetX -= 4.0; enemyBigOffsetX <= 0 {
                    enemyBigOffsetX = float64(Random(0, int(winWidth)) + float64(winWidth*2))
                }
            } else {
                bg.OffsetLeftNextX(0.1, 0, float64(-winWidth))
                mg.OffsetLeftNextX(0.5, 0, float64(-winWidth))
                fg.OffsetLeftNextX(2.0, 0, float64(-winWidth))
                starOffsetX += 2.0
                if enemyOffsetX -= 2.0; enemyOffsetX <= 0 {
                    enemyOffsetX = float64(Random(0, int(winWidth)) + float64(winWidth))
                }
                if enemyBigOffsetX -= 2.0; enemyBigOffsetX <= 0 {
                    enemyBigOffsetX = float64(Random(0, int(winWidth)) + float64(winWidth*2))
                }
            }
        } else {
            if enemyOffsetX -= 3.0; enemyOffsetX <= 0 {
                enemyOffsetX = float64(Random(0, int(winWidth)) + float64(winWidth))
            }
            if enemyBigOffsetX -= 3.0; enemyBigOffsetX <= 0 {
                enemyBigOffsetX = float64(Random(0, int(winWidth)) + float64(winWidth*2))
            }
        }

        bg.SetDst(int32(bg.GetOffsetX()), 0, winWidth * 2, winHeight)
        mg.SetDst(int32(mg.GetOffsetX()), 0, winWidth * 2, winHeight)
        fg.SetDst(int32(fg.GetOffsetX()), 0, winWidth * 2, winHeight)
        star.SetDst(int32(starOffsetX), starY, 30, 30)
        enemy.SetDst(int32(enemyOffsetX), 320, 100, 100)
        enemyBig.SetDst(int32(enemyBigOffsetX), 270, 130, 150)

        if state == stateJump || state == stateJumpWalk {
            velocityY += gravity
            positionY += velocityY
            positionX += velocityX

            if positionY > 290.0 {
                positionY = 290.0
                velocityY = 0.0
                state = stateWalk
                sc.SetTimeEnabled(true)
                jumpCount = 0
            }

            sc.SetDst(100, int32(positionY), sc.GetFrameW(), sc.GetFrameH())
        }

        if star.CheckCollisionSprite(sc) {
            sfxCoin.Sound.Play(1, 0)
            score++;
            starOffsetX = float64(Random(0, int(winWidth)) + float64(winWidth))
            starY = int32(Random(100, 250))
        }

        if enemy.CheckCollisionSpriteTop(sc, 15) || enemyBig.CheckCollisionSpriteTop(sc, 15) {
            sfxKick.Sound.Play(1, 0)
            velocityY = -12.0
        }

        if enemy.CheckCollisionSpriteBottom(sc, 15) || enemyBig.CheckCollisionSpriteBottom(sc, 15) {
            sfxScoreDown.Sound.Play(1, 0)
            score = 0;
        }

        bg.Draw(renderer)
        mg.Draw(renderer)
        fg.Draw(renderer)

        scoreText, _ := text.RenderText("SCORE: " + strconv.Itoa(score), font, sdl.Color{255, 255, 255, 255}, renderer)
        scoreText.DrawXY(700, 10, renderer)

        star.NextFrame()
        star.Draw(renderer)

        enemy.NextFrame()
        enemy.Draw(renderer)
        enemyBig.NextFrame()
        enemyBig.Draw(renderer)

        sc.NextFrame()
        if direction == directionRight {
            sc.Draw(renderer)
        } else {
            sc.DrawFlip(renderer)
        }

        renderer.Present()
        sdl.Delay(15)
    }

	return 0
}

func main() {
	os.Exit(run())
}
