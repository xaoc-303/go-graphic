package animation

import (
    "time"
    "math/rand"
)

func Random(min int, max int) float64 {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    return float64(r.Intn(max - min +1) + min)
}
