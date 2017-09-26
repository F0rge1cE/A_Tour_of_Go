package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	a := [][]uint8{}

	for i := 0; i < dy; i = i + 1 {
		a = append(a, []uint8{})

		for j := 0; j < dx; j = j + 1 {
			a[i] = append(a[i], uint8(j+i))
		}
		//println(len(a),cap(a))
	}

	return a
}

func main() {
	pic.Show(Pic)
}
