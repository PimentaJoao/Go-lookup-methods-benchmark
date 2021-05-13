package main

import "testing"

// "math/rand"
// "time"

// var workloads []int

// const min int = 1
// const max int = 7

// func setup() {
// 	rand.Seed(time.Now().UnixNano())
// 	for i := 0; i < 100; i++ {
// 		workloads = append(workloads, rand.Intn(max-min+1)+max)
// 	}
// }

var workloads []int = []int{4, 2, 6, 4, 5, 6, 7, 2, 3, 4, 1, 5, 4, 3, 2, 6, 2, 5, 7, 1, 2, 3, 5, 1, 2}
var worksize int = len(workloads)

func BenchmarkIf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		daySelectorIf(workloads[i%worksize])
	}
}

func BenchmarkSwitch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		daySelectorSwitch(workloads[i%worksize])
	}
}
