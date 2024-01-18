package main

import "fmt"

type Greeter interface {
	Greet() string
}

type englishGreet struct {
}

func (e *englishGreet) Greet() string {
	return "hello\n"
}

func main() {
	g := englishGreet{}
	msg := g.Greet()
	fmt.Printf(msg)
}
