package hello

const testVersion = 2

func HelloWorld(word string) string {
	if word == "" {
		return "Hello, World!"
	}
	return "Hello, " + word + "!"
}
