package main

import "testing"

// from fib_test.go
func BenchmarkFib10(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		Fib(10)
	}
}

func TestFib10(t *testing.T) {
	expected := []int{100, 200, 300, 400, 500}
	inputs := []int{1, 2, 3, 4, 5}
	for k, exp := range expected {
		result := Fib(inputs[k])
		if result != exp {
			t.Errorf("Something is wrong expected %d received %d", exp, result)
		}
	}
}

// Calling bench go test -bench=.

/*
func benchmarkFib(i int, b *testing.B) {
        for n := 0; n < b.N; n++ {
                Fib(i)
        }
}

func BenchmarkFib1(b *testing.B)  { benchmarkFib(1, b) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(2, b) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(3, b) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(10, b) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(20, b) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(40, b) }
*/
