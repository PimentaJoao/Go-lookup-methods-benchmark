package main

import "testing"

var workloads []int = []int{1}
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

func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		daySelectorMap(workloads[i%worksize])
	}
}

func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		daySelectorSlice(workloads[i%worksize])
	}
}
