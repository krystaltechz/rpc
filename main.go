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

		if err = f(cmd); err != nil {
			fmt.Println(err)
			break
		}

		fmt.Println("Вы выиграли | Ничья")
	}
}
