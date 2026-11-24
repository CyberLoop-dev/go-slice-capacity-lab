package main

import "testing"

func TestSlicesShareBackingArrayUntilGrowth(t *testing.T) {
	base := make([]int, 2, 4)
	base[0], base[1] = 10, 20
	alias := base[:]
	extended := append(base, 30)
	alias[0] = 99
	if base[0] != 99 || extended[0] != 99 {
		t.Fatal("append within capacity should keep the same backing array")
	}

	full := make([]int, 2, 2)
	full[0], full[1] = 1, 2
	grown := append(full, 3)
	grown[0] = 7
	if full[0] != 1 {
		t.Fatal("append beyond capacity should use a new backing array")
	}
}
