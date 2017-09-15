package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for i := 0; i <= 100000000; i++ {
		//fmt.Printf("The number is: %d \n", i)
	}
	fmt.Println(time.Since(start))
}
