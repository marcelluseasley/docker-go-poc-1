package foo

import "testing"




func TestAdd(t *testing.T) {
	// Test case 1: Testing with integers
	resultInt := Add(2, 3)
	expected := 5
	if resultInt != expected {
		t.Errorf("Add(2, 3) = %d; expected %d", resultInt, expected)
	}

	// Test case 2: Testing with floats
	resultFloat64 := Add(2.5, 3.7)
	expectedFloat := 6.2
	if resultFloat64 != expectedFloat {
		t.Errorf("Add(2.5, 3.7) = %f; expected %f", resultFloat64, expectedFloat)
	}
}
