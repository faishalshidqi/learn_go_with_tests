package helloWorld

import "fmt"

const (
	englishPrefix    = "Hello"
	spanishPrefix    = "Hola"
	frenchPrefix     = "Bonjour"
	indonesianPrefix = "Halo"
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "" {
		language = "english"
	}

	prefix := getPrefix(language)

	return fmt.Sprintf("%s, %s!", prefix, name)
}

func getPrefix(language string) (prefix string) {
	switch language {
	case "spanish":
		prefix = spanishPrefix
	case "french":
		prefix = frenchPrefix
	case "indonesian":
		prefix = indonesianPrefix
	default:
		prefix = englishPrefix
	}
	return
}
