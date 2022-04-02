package tester

import (
	"testing"
)

func Benchmark100PrimesWith0MSSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		primesAndSleep(100, 0)
	}
}

func Benchmark100PrimesWith5MSSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		primesAndSleep(100, 5)
	}
}
func Benchmark100PrimesWith10MSSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		primesAndSleep(100, 10)
	}
}

func Benchmark100GoPrimesWith0MSSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goPrimesAndSleep(100, 0)
	}
}

func Benchmark100GoPrimesWith5MSSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goPrimesAndSleep(100, 5)
	}
}

func Benchmark100GoPrimesWith10MSSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goPrimesAndSleep(100, 10)
	}
}
