package stringutils

import (
	"testing"
)

func TestSwapCase(t *testing.T) {
	input,expected := "Hello, World", "hELLO, wORLD"
	result := SwapCase(input)

	if result != expected{
		t.Errorf("SwapCase(%q) == %q, expected %q",input,result,expected)
	}
}

func TestReverse(t *testing.T) {
	input,expected := "Hello, World", "dlroW ,olleH"
	result := Reverse(input)

	if result != expected{
		t.Errorf("SwapCase(%q) == %q, expected %q",input,result,expected)
	}
}
