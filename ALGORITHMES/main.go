package main

import "fmt"

func main() {
	var x, y, z int

	fmt.Print("Enter two integer values X and Y: ")
	fmt.Scanf("%d %d", &x, &y)
	z = x
	x = y
	y = z
	fmt.Printf("After swapping X and Y, we have X=%d and Y=%d\n", x, y)
}
