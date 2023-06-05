package main
import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello world")
	var name string
	name = os.Getenv("USER")
	fmt.Println("hello", name)
}
