package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	z := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		x_vals := make([]uint8, dx)
		for j := 0; j < dx; j++ {
			x_vals[j] = uint8(i*j)
		}
		z[i] = x_vals
	}
	return z
}

func main() {
	pic.Show(Pic)
}
