package main

import "testing"

func TestKilometersToMiles(t *testing.T) {
	got := kilometersToMiles(10)
	want := 6.21371

	if got != want {
		t.Fatalf("expected %v, got %v", want, got)
	}
}
