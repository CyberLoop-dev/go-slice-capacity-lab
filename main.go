package main

import "fmt"

func describe(label string, values []int) {
	var address any = "<empty>"
	if len(values) > 0 {
		address = fmt.Sprintf("%p", &values[0])
	}
	fmt.Printf("%-22s values=%v len=%d cap=%d backing=%v\n",
		label, values, len(values), cap(values), address)
}

func runDemo() {
	base := make([]int, 2, 4)
	base[0], base[1] = 10, 20
	alias := base[:]
	describe("base", base)
	describe("alias", alias)

	withinCapacity := append(base, 30)
	describe("append within cap", withinCapacity)
	alias[0] = 99
	describe("base after alias edit", base)
	describe("appended view", withinCapacity)

	full := make([]int, 2, 2)
	full[0], full[1] = 1, 2
	reallocated := append(full, 3)
	describe("full slice", full)
	describe("after reallocation", reallocated)
}

func main() { runDemo() }
