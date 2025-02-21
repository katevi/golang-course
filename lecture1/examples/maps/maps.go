package main

import "fmt"

// simply run `go run ./lecture1/examples/maps/maps.go` from the root of the repository
func main() {
	var m map[string]int
	fmt.Printf("| m is not initialized |: nil?=%t; len(m)=%d \n", m == nil, len(m))

	x := make(map[string]int, 5)
	fmt.Printf("| x := make(map[string]int, 5) |: empty?=%t; len(x)=%d \n", x == nil, len(x))

	x["a"] = 42
	x["c"] = 99
	fmt.Println(x)

	valB, okB := x["b"]
	valA, okA := x["a"]
	fmt.Printf("| x[b] |: valB=%v, okB=%t \n", valB, okB)
	fmt.Printf("| x[a] |: valA=%v, okA=%t \n", valA, okA)

	fmt.Println("len(x) = ", len(x))
	fmt.Println("traversing map x")
	for key, value := range x {
		fmt.Printf("x[%s] = %d; ", key, value)
	}
	fmt.Println()

	y := map[string]int{"1": 1, "2": 2, "46": 46}
	fmt.Println("y = ", y, ", len(y) = ", len(y))

	delete(y, "2")
	fmt.Println("y = ", y, ", len(y) = ", len(y))
}
