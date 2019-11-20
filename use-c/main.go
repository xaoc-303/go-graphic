package main

// https://golang.org/cmd/cgo/

/*
#include <SDL2/SDL_mixer.h>
#include <stdio.h>
int getNumber() {
    return 777;
}
*/
import "C"

import (
	"fmt"
)

func main() {
    fmt.Println(C.getNumber())
    fmt.Println(C.AUDIO_S16SYS)
}
