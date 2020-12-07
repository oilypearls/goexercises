package main

import "fmt"

// separate your main code from outside (ie. sideeffects)

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const germanHelloPrefix = "Hallo, "
const englishWorld = "World"
const frenchWorld = "Monde"
const spanishWorld = "Mundo"
const germanWorld = "Welt"

func SetLanguageGreeting(lang string) string {
	greeting := ""
	switch lang {
	case "English":
		greeting = englishHelloPrefix
	case "Spanish":
		greeting = spanishHelloPrefix
	case "French":
		greeting = frenchHelloPrefix
	case "German":
		greeting = germanHelloPrefix
	default:
		greeting = englishHelloPrefix
	}
	return greeting
}

func SetNameHelper(lang string) string {
	name := ""
	switch lang {
	case "English":
		name = englishWorld
	case "Spanish":
		name = spanishWorld
	case "French":
		name = frenchWorld
	case "German":
		name = germanWorld
	default:
		name = englishWorld
	}
	return name
}

func SetName(lang string, name string) string {
	if name == "" {
		name = SetNameHelper(lang)
	}
	return name
}

func HelloWorld(name string, lang string) string {
	return SetLanguageGreeting(lang) + SetName(lang, name)

}

func main() {
	fmt.Println(HelloWorld("madre", "English"))
}
