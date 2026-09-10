package main

import (
	"crypto/rand"
	"errors"
	"math/big"
)

// TODO: написать func(cmd string) {}
// которая принимает на вход мой ход (камень, ножницы, бумага)
// внутри рандомно выбирает ход компьютера
// из 3 вариантов (камень, ножницы, бумага)
// возвращает строку ответ: "Проиграл" или "Выиграл"

func f(cmd string) (string, error) {
	if cmd != "ножницы" && cmd != "камень" && cmd != "бумага" {
		return "", errors.New("Неверная команда")
	}
	arr := []string{"камень", "ножницы", "бумага"}

	generated, err := generateInt()
	if err != nil {
		return "", err
	}

	computer := arr[generated]

	if cmd == computer {
		return "Ничья", nil
	}

	if cmd == "камень" && computer == "ножницы" ||
		cmd == "бумага" && computer == "камень" ||
		cmd == "ножницы" && computer == "бумага" {
		return "Вы выиграли", nil
	} else {
		return "Вы проиграли", nil
	}
}

// Обновил поведение генерации рандомных чисел
func generateInt() (int64, error) {
	max := big.NewInt(3)

	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return 0, errors.New("Ошибка при генерации числа")
	}

	return n.Int64(), nil
}
