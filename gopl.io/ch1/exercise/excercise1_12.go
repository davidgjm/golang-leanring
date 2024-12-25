package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 /// first color in palette
	blackIndex = 1 //next color in palette
)

//var cycles = 5

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Printf("%s: %s\n", k, v)
	}
	cycles := 5
	if v, err := strconv.Atoi(r.Form.Get("cycles")); err == nil {
		cycles = v
	}

	size := 100
	if s, err := strconv.Atoi(r.Form.Get("size")); err == nil {
		size = s
	}
	lissajous(w, cycles, size)
}

func lissajous(out io.Writer, cycles int, size int) {
	if cycles == 0 {
		cycles = 5
	}
	if size == 0 {
		size = 100
	}
	const (
		res     = 0.001
		nframes = 64
		delay   = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < math.Pi*(float64(cycles))*2; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*(float64(size))+0.5), size+int(y*(float64(size))+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
