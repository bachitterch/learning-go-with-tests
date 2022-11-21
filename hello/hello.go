package hello

import (
	"fmt"
)

const (
	french  = "French"
	spanish = "Spanish"
)

const (
	englishHelloPrefix = "Hello, "
	frenchHelloPreifx  = "Bonjour, "
	spanishHelloPrefix = "Hola, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPreifx
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("world", spanish))
}
