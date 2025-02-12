package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	image := make([][]uint8, dy)
	for s := range image {
		image[s] = make([]uint8, dx)
		for e := range image[s] {
			image[s][e] = uint8(s * e) // Mathematical function: s * e
		}
	}
	return image
}

func main() {
	pic.Show(Pic)
}
