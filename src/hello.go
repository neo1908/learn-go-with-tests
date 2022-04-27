package src

import "fmt"

const helloPrefix = "Hello, "
const helloSuffix = "!"
const helloDefault = "World"

func Hello(name string) string {
	if name == "" {
		name = helloDefault
	}

	return helloPrefix + name + helloSuffix
}

func main() {
	fmt.Println(Hello(""))
}
