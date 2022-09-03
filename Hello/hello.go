package main


import "fmt"

const helloPrefix = "Hello, "
const helloSuffix = "!"

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return helloPrefix + name + helloSuffix
}

func main() {
	fmt.Println(Hello("world!"))
}