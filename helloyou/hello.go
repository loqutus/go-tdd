package main

func getPrefix(language string) (prefix string) {
	switch language {
	case "Spanish":
		prefix = "Hola, "
	case "French":
		prefix = "Bonjour, "
	default:
		prefix = "Hello, "
	}
	return
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return getPrefix(language) + name
}
