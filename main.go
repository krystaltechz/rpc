package main

import (
	"fmt"
	"log"
)

func main() {
	for {
		var cmd string
		fmt.Println("Введите команду: ")
		_, err := fmt.Scanln(&cmd)

		if err != nil {
			log.Fatal(err)
		}

		str, err := f(cmd)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(str)
	}
}
