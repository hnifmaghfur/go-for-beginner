package stringutils

import (
	"testing"
	"time"
)

func TestSwapcaseInParallel(t *testing.T){
	t.Parallel()

	//delaying 1 second for the sake of demonstration
	time.Sleep(1*time.Second)
	input, expected := "Hello, World", "hELLO, wORLD"
	result := SwapCase(input)

	if(result != expected){
		t.Errorf("SwapCase(%q) == %q, expected %q",input,result,expected)
	}
}

func TestReverseInParallel(t *testing.T){
	t.Parallel()

	//delaying 2 second for the sake of demonstration
	time.Sleep(2*time.Second)
	input, expected := "Hello, World", "dlroW ,olleH"
	result := Reverse(input)

	if(result != expected){
		t.Errorf("Reverse(%q) == %q, expected %q",input,result,expected)
	}

}