package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/mix"
)

var Music *mix.Music
var Sound *mix.Chunk

func main() {
    var err error

    // init
    err = mix.OpenAudio(mix.DEFAULT_FREQUENCY, mix.DEFAULT_FORMAT, mix.DEFAULT_CHANNELS, mix.DEFAULT_CHUNKSIZE * 2)
    if err != nil {
        return
    }
    defer mix.CloseAudio()

    // music
    Music, err = mix.LoadMUS("track.wav")
    if err != nil {
        println("LoadMUS: %s\n", err)
        panic(err)
    }
    defer Music.Free()

    Music.Play(-1)

    // sound
    Sound, err = mix.LoadWAV("test.wav")
    if err != nil {
        println("LoadWAV: %s\n", err)
        panic(err)
    }
    defer Sound.Free()

    Sound.Play(1, 0)

    // delay for playing
    for mix.Playing(-1) == 1 {
        sdl.Delay(16)
    }

    sdl.Delay(8888)
}
