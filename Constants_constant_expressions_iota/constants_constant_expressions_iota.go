package main

import "fmt"

func main() {
	const a = 42
	const b float32 = 3
	var f32 float32 = b
	var f64 float64 = float64(b)

	fmt.Println(f32, f64)

	const c = iota
	fmt.Println(c)

	const (
		d = 2 * 5
		e // 2 * 5
		f = iota // 2 related to position in the current group
		g // 3
		h = 10 * iota
	)

	fmt.Println(d, e, f, g, h)
}