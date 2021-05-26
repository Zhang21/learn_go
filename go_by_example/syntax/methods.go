package main

import "fmt"

// Go supports methods defines on struct types.
// Methods can be defined for either pointer or value receiver types.
// Go automatically handles conversion between values and pointers for method call. You may want to use a pointer receiver type to avoid copying on method calls
// or to allow the method to mutate the receiving struct.

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())
}
