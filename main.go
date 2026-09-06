package main

import (
	"errors"
	"fmt"
	"log/slog"
)

func main() {
	fmt.Println("Привет, мир!")
	println("Hello, PTU")
	print("Hello")
	slog.Error("Hello bot", "error", errors.New("hello"))
}
