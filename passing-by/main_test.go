package main

import "testing"

func BenchmarkPBR(b *testing.B) {
	for n := 0; n < b.N; n++ {
		passByReference(&obj)
	}
}

func BenchmarkPBV(b *testing.B) {
	for n := 0; n < b.N; n++ {
		passByValue(obj)
	}
}
