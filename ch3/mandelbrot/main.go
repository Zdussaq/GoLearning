// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

type ImageParams struct {
	xmin, ymin, xmax, ymax, width, height float64
}

func main() {
	imageParams := ImageParams{-2, -2, 2, 2, 1024, 1024}

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		serveImage(rw, r, imageParams)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func serveImage(rw http.ResponseWriter, r *http.Request, imageParams ImageParams) {

	r.ParseForm()
	for k, v := range r.Form {
		switch k {
		case "xmin":
			imageParams.xmin, _ = strconv.ParseFloat(v[0], 64)
		case "xmax":
			imageParams.xmax, _ = strconv.ParseFloat(v[0], 64)
		case "ymin":
			imageParams.ymin, _ = strconv.ParseFloat(v[0], 64)
		case "ymax":
			imageParams.ymax, _ = strconv.ParseFloat(v[0], 64)
		case "height":
			imageParams.height, _ = strconv.ParseFloat(v[0], 64)
		}
	}

	if imageParams.ymin >= imageParams.xmax || imageParams.xmin >= imageParams.xmax {
		http.Error(rw, "min coordinate greater than max", http.StatusBadRequest)
		return
	}

	img := image.NewRGBA(image.Rect(0, 0, int(imageParams.width), int(imageParams.height)))
	for py := 0; py < int(imageParams.height); py++ {
		y := float64(py)/imageParams.height*(imageParams.ymax-imageParams.ymin) + imageParams.ymin
		for px := 0; px < int(imageParams.width); px++ {
			x := float64(px)/imageParams.width*(imageParams.xmax-imageParams.xmin) + imageParams.xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(rw, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

//!-

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
