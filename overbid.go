package main

import (
	"fmt"
	"time"
)

func overbid(x float32) float32 {

	if x <= 25 {
		return x * 1.5
	} else if x <= 50 {
		return x * 1.3
	} else if x <= 75 {
		return x * 1.15
	} else if x <= 100 {
		return x * 1.1
	} else if x <= 150 {
		return x * 1.05
	}

	return 1.0
}

func main() {
	now := time.Now()

	fmt.Println("\n\nWelcome to the overbider!\n")
	fmt.Println("process elapse\n", time.Since(now), "\n")

	bidArray := [5]float32{10.0, 25.0, 75.0, 100.0, 125.0}

	for _, i := range bidArray {
		i = overbid(i)
		fmt.Println(i)
	}
}
