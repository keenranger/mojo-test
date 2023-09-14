package main

import (
	"math/rand"
	"time"
)

const (
	Length = 1e9
)

func main() {
	var total float32
	randoms := make([]float32, Length)
	for i := 0; i < Length; i++ {
		randoms[i] = rand.Float32()
	}
	start := time.Now()
	for i := 0; i < Length; i++ {
		if randoms[i] > 0.5 {
			total += randoms[i]
		}
	}
	elapsed := time.Since(start)
	println(total)
	println(elapsed.Milliseconds())
}
