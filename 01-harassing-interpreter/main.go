package main

import "time"

func main(){
	total := 0
	start := time.Now()
	for i := 0; i < 1000000000; i++ {
		total += i
	}
	elapsed := time.Since(start)
	println(total)
	println(elapsed.Milliseconds())
}