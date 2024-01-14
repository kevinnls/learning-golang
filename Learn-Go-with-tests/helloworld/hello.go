package main

import (
	"fmt"
	internal "helloworld/internal"
)

func Hello() string {
	return "Hello World!"
}

func main() {
	fmt.Println(Hello())
	fmt.Println(internal.ModularHello())
}
