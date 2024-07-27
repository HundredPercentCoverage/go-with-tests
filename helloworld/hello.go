package hello

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"

	spanishHelloPrefix = "Hola "
	frenchHelloPrefix  = "Bonjour "
	englishHelloPrefix = "Hello "
)

// Return greeting prefix based on language
//
// Uses named return value so explicit 'return prefix' not required
func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func Hello(name, language string) string {
	if len(name) == 0 {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
