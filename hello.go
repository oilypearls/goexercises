package main

import "fmt"

// separate your main code from outside (ie. sideeffects)

const englishHelloPrefix = "Hello, "
const world = "World"

func HelloWorld(name string) string {
	/*if name == "" {
		return englishHelloPrefix + world
	} else {
		return englishHelloPrefix + name
	}*/ // old inelegant code
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name

}

func main() {
	fmt.Println(HelloWorld("madre"))
}
