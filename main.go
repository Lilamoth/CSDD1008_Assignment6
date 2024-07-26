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

func handler(w http.ResponseWriter, r *http.Request) {
	dc := gg.NewContext(800, 800)
	hostname, _ := os.Hostname()
	seed := time.Now().UnixNano() + int64(hostname[0])
	drawRandomArt(dc, seed)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
