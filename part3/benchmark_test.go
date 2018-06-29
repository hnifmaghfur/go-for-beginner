package stringutils

import "testing"

func BenchmarkSwapCase(b *testing.B) {
	for i:= 0; i< b.N; i++ {
		SwapCase("Hello, World")
	}
}

func BenchmarkReverse(b *testing.B) {
	for i:= 0; i< b.N; i++ {
		Reverse("Hello, World")
	}
}
