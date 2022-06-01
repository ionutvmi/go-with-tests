package helloworld

const PREFIX = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}

	return PREFIX + name
}
