package audio

import (
	"github.com/veandco/go-sdl2/mix"
)

type Music struct {
	path string
	Music *mix.Music
}

type Sound struct {
	path string
    Sound *mix.Chunk
}

func Init() {
    err := mix.OpenAudio(mix.DEFAULT_FREQUENCY, mix.DEFAULT_FORMAT, mix.DEFAULT_CHANNELS, mix.DEFAULT_CHUNKSIZE * 2)
    if err != nil {
        println("OpenAudio: %s\n", err)
        panic(err)
    }
//     defer mix.CloseAudio()
}

func LoadMusic(path string) *Music {
    music := new(Music)
    music.path = path

    Init()

    var err error
    music.Music, err = mix.LoadMUS(music.path)
    if err != nil {
        println("LoadMUS: %s\n", err)
        panic(err)
    }
//     music.Free()

    return music
}

func (music *Music) Free() {
    mix.CloseAudio()
    music.Music.Free()
}

func (music *Music) Play() {
    music.Music.Play(-1)
}

func LoadSound(path string) *Sound {
    sound := new(Sound)
    sound.path = path

    Init()

    var err error
    sound.Sound, err = mix.LoadWAV(sound.path)
    if err != nil {
        println("LoadWAV: %s\n", err)
        panic(err)
    }
//     defer Sound.Free()

    return sound
}
