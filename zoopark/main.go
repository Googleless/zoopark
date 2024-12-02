package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/gen2brain/beeep"
	"hello-go/zoopark/animal"
	"hello-go/zoopark/database"
	"hello-go/zoopark/sound"
)

func main() {
	// Инициализация базы данных
	connStr := "user=postgres password=123456789 dbname=postgres sslmode=disable"
	db, err := database.InitializeDatabase(connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Животные
	animals := map[string]animal.Animal{
		"Медведи": animal.Bear{},
		"Кошачьи": animal.Feline{},
		"Собачьи": animal.Canine{},
		"Птицы":   animal.Bird{},
		"Грызуны": animal.Rodent{},
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("О каких животных Вы хотите узнать?\n")
	for animal := range animals {
		fmt.Println(animal)
	}

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	fmt.Println("Вы ввели:", input)

	selectedAnimal, exists := animals[input]
	if !exists {
		fmt.Println("Неизвестный ввод.")
		return
	}

	// Запуск горутины для обработки животного
	var wg sync.WaitGroup
	wg.Add(1)
	go func(a animal.Animal) {
		defer wg.Done()
		a.DisplayInfo()
		err := sound.PlaySound(a)
		if err != nil {
			fmt.Printf("Ошибка воспроизведения звука для %T: %v\n", a, err)
		}
		notifyAnimalProcessed(a)
	}(selectedAnimal)

	// Ожидание завершения всех горутин
	wg.Wait()
	fmt.Println("Обработка завершена.")
}

// Уведомление о завершении обработки животного
func notifyAnimalProcessed(a animal.Animal) {
	go func() {
		err := beeep.Notify(fmt.Sprintf("Обработан: %T", a), "Информация и звук успешно обработаны.", "")
		if err != nil {
			fmt.Printf("Ошибка отправки уведомления: %v\n", err)
		}
	}()
}
