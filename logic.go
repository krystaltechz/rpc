package main

import (
	"errors"
	"math/rand"
)

// TODO: написать func(cmd string) {}
// которая принимает на вход мой ход (камень, ножницы, бумага)
// внутри рандомно выбирает ход компьютера
// из 3 вариантов (камень, ножницы, бумага)
// возвращает строку ответ: "Проиграл" или "Выиграл"

func f(cmd string) error {
	if cmd != "ножницы" && cmd != "камень" && cmd != "бумага" {
		return errors.New("Неверная команда")
	}
	arr := []string{"камень", "ножницы", "бумага"}
	computer := arr[rand.Intn(3)]

	if cmd == computer {
		return nil
	}

	if cmd == "камень" && computer == "ножницы" ||
		cmd == "бумага" && computer == "камень" ||
		cmd == "ножницы" && computer == "бумага" {
		return nil
	} else {
		return errors.New("Вы проиграли")
	}
}
