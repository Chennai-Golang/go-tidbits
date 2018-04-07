//go:generate go run sample/main.go
package main

import "fmt"

func init() {
	fmt.Print("Hello ")
}

func main() {
	fmt.Println("World!")
}
