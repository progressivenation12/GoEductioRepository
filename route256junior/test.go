package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Устанавливаем seed случайным значением, чтобы генерировать различные числа каждый раз при запуске программы.
	rand.Seed(time.Now().UnixNano())

	min := 10
	max := 20

	// Генерируем случайное число в диапазоне от min до max.
	randomNumber := rand.Intn(max-min+1) + min

	fmt.Println(randomNumber)
}
