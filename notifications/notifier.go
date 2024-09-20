package main

import (
	"fmt"
	"log"

	"github.com/gen2brain/beeep"
)

func main() {
	// Уведомление
	title := "Привет!"
	message := "Это пример использования библиотеки beeep на Go."

	err := beeep.Notify(title, message, "")
	if err != nil {
		log.Fatal(err)
	}

	// Воспроизведение звука
	err = beeep.Alert("Внимание!", "Это звуковое уведомление.", "")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Уведомления показаны!")
}
