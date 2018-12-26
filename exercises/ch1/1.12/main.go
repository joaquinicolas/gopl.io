package main

import (
	"strconv"
	"fmt"
	"log"
	"math"
	"image"
	"image/gif"
	"io"
	"math/rand"
	"net/http"
	"image/color"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIdx = 0
	blackIdx = 1
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request)  {
		values := r.URL.Query()
		cycles, err := strconv.Atoi(values.Get("cycles"))
		if err != nil {
			 fmt.Fprintf(w, "An has happened: %v", err)
			 return
		}
		lissajous(w, cycles)
	}

	http.HandleFunc("/", handler)
	
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
	return
}


func lissajous(out io.Writer, c int)  {
	var cycles = 5.0
	const (
		
		res = 0.001
		size = 100
		nframes = 64
		delay = 8
	)

	if c > 0 {
		cycles = float64(c)
	}

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size + int(x*size+0.5), size+int(y*size+0.5),
						blackIdx)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}