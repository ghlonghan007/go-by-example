package main

import "fmt"

func main() {
	a, b := 1, 2
	fmt.Printf("a: %v\nb: %v\n", a, b)
	const c string = "helloworld"
	// c.p
	fmt.Printf("c: %v\n", c)
	const d  = 101010
	fmt.Printf("d: %v\n", d)
}
