package main

import "testing"

func TestAddSuccess(t *testing.T) {
	result := Add(20, 3)

	expect := 23

	if result != expect {
		t.Errorf("got %d, expected %d", result, expect)
	}

}
