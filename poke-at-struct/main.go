package main

import "fmt"

type Foo struct {
	x int
	y int
}

func main() {
	foo1 := Foo{
		x: 1,
		y: 2,
	}
	p := &foo1.y

	foo2 := Foo{
		x: 100,
		y: 200,
	}

	fmt.Println("foo1", foo1)
	fmt.Println("foo2", foo2)

	foo1 = foo2
	fmt.Println("foo1", foo1)

	*p = 42
	fmt.Println("foo1", foo1)
}
