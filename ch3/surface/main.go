// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

type zFunc func(x, y float64) float64

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

//Exercise 3.2 - allow the image to display exxbog, moguls, and a saddle.
func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	switch r.URL.Path {
	case "/corner":
		displaySVG(w, fWave)
	case "/saddle":
		displaySVG(w, fSaddle)
	case "/egg":
		displaySVG(w, fEgg)
	}

}

func displaySVG(output io.Writer, f zFunc) {
	fmt.Fprintf(output, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	zMin, zMax := findMinMax(f)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, f)
			bx, by := corner(i, j, f)
			cx, cy := corner(i, j+1, f)
			dx, dy := corner(i+1, j+1, f)
			//Exercise 1.1 prevent the code from printing non numeric float64 vals.
			if !(math.IsNaN(ax) || math.IsNaN(ay) || math.IsNaN(bx) || math.IsNaN(by) || math.IsNaN(cx) || math.IsNaN(cy) || math.IsNaN(dx) || math.IsNaN(dy)) {
				fmt.Fprintf(output, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='stroke:%s;fill:#333333'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, color(i, j, zMin, zMax, f))
			}
		}
	}
	fmt.Fprintln(output, "</svg>")
}

//For exercise 3.3 - adding color grading to
func color(i, j int, min, max float64, f zFunc) string {
	localMin := math.NaN()
	localMax := math.NaN()
	for xOffset := 0; xOffset < 2; xOffset++ {
		for yOffset := 0; yOffset < 2; yOffset++ {
			x := xyrange * (float64(i+xOffset)/cells - 0.5)
			y := xyrange * (float64(j+yOffset)/cells - 0.5)
			z := f(x, y)
			if math.IsNaN(localMax) || z > localMax {
				localMax = z
			}
			if math.IsNaN(localMin) || z < localMin {
				localMin = z
			}

		}
	}

	fmt.Printf("Local Max: %f\t Local Min: %f \t Max: %f \t Min: %f\n", localMax, localMin, max, min)

	if math.Abs(localMax) > math.Abs(localMin) {
		redIndex := math.Exp(math.Abs(localMax)) / math.Exp(math.Abs(max)) * 255
		if redIndex > 255 {
			redIndex = 255
		}
		return fmt.Sprintf("#%02x0000", int(redIndex))
	} else {
		blueIndex := math.Exp(math.Abs(localMin)) / math.Exp(math.Abs(min)) * 255
		if blueIndex > 255 {
			blueIndex = 255
		}
		return fmt.Sprintf("#0000%02x", int(blueIndex))
	}
}

func findMinMax(f zFunc) (float64, float64) {
	min := math.MaxFloat64
	max := math.SmallestNonzeroFloat64
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			for xOffset := 0; xOffset < 2; xOffset++ {
				for yOffset := 0; yOffset < 2; yOffset++ {
					x := xyrange * (float64(i+xOffset)/cells - 0.5)
					y := xyrange * (float64(j+yOffset)/cells - 0.5)
					z := f(x, y)
					if !math.IsNaN(z) && !math.IsInf(z, 0) && z > max {
						max = z
					} else if !math.IsNaN(z) && !math.IsInf(z, 0) && z < min {
						min = z
					}

				}
			}
		}
	}

	return min, max
}

func corner(i, j int, f zFunc) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func fWave(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

//Exercise 1.2 functions
func fEgg(x, y float64) float64 {
	return (math.Sin(x) + math.Cos(y)) / 15
}

func fSaddle(x, y float64) float64 {
	h := math.Pow(x, 2)/math.Pow(22, 2) - math.Pow(y, 2)/math.Pow(15, 2)
	return h
}

//!-
