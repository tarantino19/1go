package main

import "testing"

func TestCountWords(t *testing.T) {
	input := "one two three four five"
	wants := 5

	result := CountWords([]byte(input))

	//assertion 1
	if result != wants {
		t.Logf("expected: %v, got: %v", wants, result)
		t.Fail()
	}

	//assertion 2
	input = ""
	wants = 0

	result = CountWords([]byte(input))

	if result != wants {
		t.Logf("expected: %v, got: %v", wants, result)
		t.Fail()
	}

	//assertion 3
	input = " "
	wants = 0

	result = CountWords([]byte(input))

	if result != wants {
		t.Logf("expected: %v, got: %v", wants, result)
		t.Fail()
	}

}
