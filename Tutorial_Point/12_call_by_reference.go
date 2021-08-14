package main

import "fmt"

func swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}
func main() {
	var a int = 20
	var b int = 30

	fmt.Printf("Before swap, value of a : %d\n", a)
	fmt.Printf("Before swap, value of b : %d\n", b)

	/* calling a function to swap the values */
	swap(&a, &b)

	fmt.Printf("After swap, value of a : %d\n", a)
	fmt.Printf("After swap, value of b : %d\n", b)

}
