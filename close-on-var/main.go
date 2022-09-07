package main

import "fmt"

type Foo struct {
	x int
}

var foo = Foo{1}

func bar() {
	fmt.Println(foo)
}

func main() {
	bar()
	foo = Foo{42}
	bar()
}
