package main

import "fmt"

const Pi = 3.14

func main() {
	const world = "世界"
	fmt.Println("hello", "world")
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("GO rules?", Truth)
}
