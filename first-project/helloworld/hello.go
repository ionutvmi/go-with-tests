package helloworld

const PREFIX = "Hello, "
const PREFIX_ES = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) string {
	if language == "Spanish" {
		return PREFIX_ES
	}

	return PREFIX
}
