package rbo

import "testing"

func TestRevBits(t *testing.T) {
	const (
		k_in uint8  = 64
		x_in uint64 = 1<<64 - 1
	)

	if x := RevBits(k_in, x_in); x != x_in {
		t.Errorf("RevBits(%v,%v) = %v, want %v", k_in, x_in, x, x_in)
	}
}
