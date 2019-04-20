package main

import "testing"

func TestMyIP(t *testing.T) {
	want := "110.54.240.61"
	got := IP()

	if want != got {
		t.Errorf("want: %v, but got: %v instead.", want, got)
	}
}
