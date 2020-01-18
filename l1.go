package main

import (
	"fmt"
	"math"
)

func main() {
	for r := 2.8; r < 4.0; r += 0.0002 {
		x := 0.5
		for i := 0; i < 100; i++ {
			x = r * x * (1.0 - x)
		}
		past := make([]float64, 1)
		past[0] = x
	DONE:
		for i := 0; i < 100; i++ {
			x = r * x * (1.0 - x)
			for _, n := range past {
				eps := math.Abs(x - n)
				if eps < .00001 {
					break DONE
				}
			}
			past = append(past, x)
		}
		for _, point := range past {
			fmt.Printf("%f\t%f\n", r, point)
		}
	}
}
