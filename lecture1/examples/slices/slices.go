package main

import "fmt"

// simply run `go run ./lecture1/examples/slices/slices.go` from the root of the repository
func main() {
	var x []int
	fmt.Printf("| var x[]int | : x == nil =%t; cap(x)=%d; len(x)=%d \n", x == nil, cap(x), len(x))

	x = make([]int, 3)
	fmt.Printf("| x := make([]int, 3) |; x=%v, cap(x)=%d; len(x)=%d \n", x, cap(x), len(x))

	x = make([]int, 0, 3)
	fmt.Printf("| x := make([]int, 0, 3) |; x=%v, cap(x)=%d; len(x)=%d \n", x, cap(x), len(x))

	x = append(x, 23)
	x = append(x, 42)
	x = append(x, 38)
	fmt.Printf("| appended 23,42,38 |; x=%v, cap(x)=%d; len(x)=%d \n", x, cap(x), len(x))

	x = append(x, 61)
	fmt.Printf("| appended 61 to full slice: cap*2|; x=%v, cap(x)=%d; len(x)=%d \n", x, cap(x), len(x))

	y := x[1:2]
	fmt.Printf("| y := x[1:2] |; y=%v, cap(y)=%d; len(y)=%d \n", y, cap(y), len(y))

	x = append(x, 196)
	fmt.Printf("| x := append(x, 196) |; x=%v, cap(x)=%d; len(x)=%d \n", x, cap(x), len(x))
	fmt.Printf("| y has 196? |; y=%v, cap(y)=%d; len(y)=%d \n", y, cap(y), len(y))

	y[0] = 111
	fmt.Printf("| y[0] := 111 |; y=%v, cap(y)=%d; len(y)=%d \n", y, cap(y), len(y))
	fmt.Printf("| x modified? |; x=%v, cap(x)=%d; len(x)=%d \n", x, cap(x), len(x))

	z := make([]int, 1, 2)
	copy(z, y)
	fmt.Printf("| copy(z, y) |; z=%v, cap(z)=%d; len(z)=%d \n", z, cap(z), len(z))

}
