package main

import (
	"fmt"
	"math"
	"sync"
    "time"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

type point struct {
	i, j  int
	ax, ay float64
}

func traceDuring() func() {
	start := time.Now()
	return func() {
		interval := time.Since(start)
		fmt.Printf("spend time: %d ms\n", interval.Milliseconds())
	}
}

func main() {
    defer traceDuring()()

	results := make(chan point)
	var wg sync.WaitGroup

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			wg.Add(1)
			go func(i, j int) {
				defer wg.Done()
				ax, ay := corner(i+1, j)
				results <- point{i, j, ax, ay}
			}(i, j)
		}
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for p := range results {
		bx, by := corner(p.i, p.j)
		cx, cy := corner(p.i, p.j+1)
		dx, dy := corner(p.i+1, p.j+1)
		fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
			p.ax, p.ay, bx, by, cx, cy, dx, dy)
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
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

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
