package animation

import (
    "fmt"
	"os"
	"time"

    "github.com/veandco/go-sdl2/sdl"
    "github.com/veandco/go-sdl2/img"
)

type Sprite struct {
	path string

	image *sdl.Surface

    texture *sdl.Texture

    frames []sdl.Rect
    frameW, frameH int32
    frameI, framesCount int

    src sdl.Rect
    dst sdl.Rect

    timeEnabled bool
    timeLast time.Time
    timeDuration time.Duration

    offsetX float64
}

func Load(fileName string) *sdl.Surface {
    image, err := img.Load(fileName)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to load PNG: %s\n", err)
        panic(err)
    }

    return image
}

func Texture(image *sdl.Surface, renderer *sdl.Renderer) *sdl.Texture {
    texture, err := renderer.CreateTextureFromSurface(image)
    if err != nil {
        panic(err)
    }

    return texture
}

func NewSprite(path string, framesCount int, duration int, renderer *sdl.Renderer) *Sprite {
	sprite := new(Sprite)
	sprite.path = path

	sprite.image = Load(path)
//     defer image.Free()

    sprite.texture = Texture(sprite.image, renderer)
//     defer texture.Destroy()

    sprite.frameW = sprite.image.W/int32(framesCount)
    sprite.frameH = sprite.image.H

    var x, y, w, h int32
    for i := 0; i < framesCount; i++ {

        x, y = sprite.frameW * int32(i), 0
        w, h = sprite.frameW, sprite.frameH

        sprite.frames = append(sprite.frames, sdl.Rect{x,y,w,h})
    }

    sprite.framesCount = framesCount
    sprite.frameI = 0
    sprite.src = sprite.frames[sprite.frameI]

    sprite.timeDuration = time.Millisecond * time.Duration(duration)
    sprite.timeLast = time.Now()
    sprite.timeEnabled = false

    sprite.offsetX = 0

	return sprite
}

func (sprite *Sprite) Free() {
    sprite.image.Free()
    sprite.texture.Destroy()
}

func (sprite *Sprite) Draw(renderer *sdl.Renderer) {
    renderer.Copy(sprite.texture, &sprite.frames[sprite.frameI], &sprite.dst)
}

func (sprite *Sprite) DrawFlip(renderer *sdl.Renderer) {
    renderer.CopyEx(sprite.texture, &sprite.frames[sprite.frameI], &sprite.dst, 0.0, nil, sdl.FLIP_HORIZONTAL)
}

func (sprite *Sprite) SetSrc(x, y, w, h int32) {
    sprite.src = sdl.Rect{x, y, w, h}
}

func (sprite *Sprite) SetDst(x, y, w, h int32) {
    sprite.dst = sdl.Rect{x, y, w, h}
}

func (sprite *Sprite) GetDstX() int32 {
    return sprite.dst.X
}

func (sprite *Sprite) GetDstY() int32 {
    return sprite.dst.Y
}

func (sprite *Sprite) DrawXYWH(x, y, w, h int32, renderer *sdl.Renderer) {
    sprite.SetDst(x, y, w, h)
    renderer.Copy(sprite.texture, &sprite.src, &sprite.dst)
}

func (sprite *Sprite) SetTimeEnabled(enabled bool) {
    sprite.timeLast = time.Now()
    sprite.timeEnabled = enabled
}

func (sprite *Sprite) NextFrame() {
    if (sprite.timeEnabled == false) {
        return
    }

    if (time.Since(sprite.timeLast) > sprite.timeDuration) {
        sprite.timeLast = time.Now()
        if sprite.frameI++; sprite.frameI >= sprite.framesCount {
            sprite.frameI = 0
        }
    }
}

func (sprite *Sprite) GetFrameW() int32 {
    return sprite.frameW
}

func (sprite *Sprite) GetFrameH() int32 {
    return sprite.frameH
}

func (sprite *Sprite) CheckCollisionSprite(sprite_being_compared *Sprite) bool {
    return checkCollisionRect(sprite.dst, sprite_being_compared.dst)
}

func checkCollisionRect(rec1, rec2 sdl.Rect) bool {
    return ((rec1.X <= (rec2.X + rec2.W) && (rec1.X + rec1.W) >= rec2.X) &&
            (rec1.Y <= (rec2.Y + rec2.H) && (rec1.Y + rec1.H) >= rec2.Y))
}

func (sprite *Sprite) CheckCollisionSpriteTop(sprite_being_compared *Sprite, v int32) bool {
    var rec1, rec2 sdl.Rect = sprite.dst, sprite_being_compared.dst

    rec1.H = v
    rec2.Y = rec2.Y + rec2.H - v
    rec2.H = v

    return checkCollisionRect(rec1, rec2)
}

func (sprite *Sprite) CheckCollisionSpriteBottom(sprite_being_compared *Sprite, v int32) bool {
    var rec1, rec2 sdl.Rect = sprite.dst, sprite_being_compared.dst

    rec1.Y = rec1.Y + v
    rec1.H = rec1.H - v
    rec2.Y = rec2.Y + v
    rec2.H = rec2.H - v

    return checkCollisionRect(rec1, rec2)
}

func (sprite *Sprite) OffsetNextX(x_offset, x_limit, x_for_reset float64) {
    if sprite.offsetX += x_offset; sprite.offsetX <= x_limit {
        sprite.offsetX = x_for_reset
    }
}

func (sprite *Sprite) OffsetLeftNextX(x_offset, x_limit, x_for_reset float64) {
    if sprite.offsetX += x_offset; sprite.offsetX >= x_limit {
        sprite.offsetX = x_for_reset
    }
}

func (sprite *Sprite) GetOffsetX() float64 {
    return sprite.offsetX
}
