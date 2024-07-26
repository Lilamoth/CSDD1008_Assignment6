package main

import (
	"math/rand"

	"github.com/fogleman/gg"
)

func drawRandomArt(dc *gg.Context, seed int64) {
	rand.Seed(seed)
	for i := 0; i < 10; i++ {
		x := rand.Float64() * float64(dc.Width())
		y := rand.Float64() * float64(dc.Height())
		r := rand.Float64() * 0.5 * float64(dc.Width())
		dc.DrawCircle(x, y, r)
		dc.SetRGB(rand.Float64(), rand.Float64(), rand.Float64())
		dc.Fill()

	}
}

func main() {

}
