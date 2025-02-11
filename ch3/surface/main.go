// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
	"time"
	"sync"
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

type Point struct {
	x float64
	y float64
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

	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	
	// ach, bch, cch, dch := make(chan Point), make(chan Point), make(chan Point), make(chan Point)
	wg := &sync.WaitGroup{}
	pts := make(chan Point, 4)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			wg.Add(1)
			go func (i, j int) {
				defer wg.Done()
				
			}
		}
	}

	fmt.Printf("<polygon points=")
	go func() {
		i := 0
		for p := range pts {
			if (i > 0) {
				fmt.Printf(" ")
			}
			fmt.Printf("%g,%g", p.x, p.y);
			i++
		}
		fmt.Printf("/>\n")
		fmt.Println("</svg>")
	} ()

	wg.Wait()
	close(pts)

	// go func() {
	// 	wg.Wait()
	// 	close(pts)
	// } ()

	// i := 0
	// for p := range pts {
	// 	if (i > 0) {
	// 		fmt.Printf(" ")
	// 	}
	// 	fmt.Printf("%g,%g", p.x, p.y);
	// 	i++
	// }
	// fmt.Printf("/>\n")
	// fmt.Println("</svg>")
}

func corner(i, j int, pt chan <- Point, wg *sync.WaitGroup) {
	defer wg.Done()
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	pt <- Point{sx, sy}	
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

//!-
