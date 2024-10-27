package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
	"hello-go/zoopark/animal"
	"hello-go/zoopark/database"
	"hello-go/zoopark/sound"
)

func main() {
	connStr := "user=postgres password=123456789 dbname=postgres sslmode=disable"
	database, err := db.InitializeDatabase(connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer func(database *db.DB) {
		err := database.Close()
		if err != nil {

		}
	}(database)

	grizzly := animal.Bear{}
	lynx := animal.Feline{}
	wolf := animal.Canine{}
	parrot := animal.Bird{}
	squirrel := animal.Rodent{}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("О каких животных Вы хотите узнать?\n")
	fmt.Println("Медведи.")
	fmt.Println("Кошачьи.")
	fmt.Println("Собачьи.")
	fmt.Println("Птицы.")
	fmt.Println("Грызуны.")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	fmt.Println("Вы ввели: ", input)

	switch input {
	case "Медведи":
		grizzly.Move()
		grizzly.Eat()
		grizzly.Fly()
		grizzly.Climb()
		grizzly.Swim()
		err := sound.PlaySound(grizzly)
		if err != nil {
			return
		}

	case "Кошачьи":
		lynx.Move()
		lynx.Eat()
		lynx.Fly()
		lynx.Climb()
		lynx.Swim()
		err := sound.PlaySound(lynx)
		if err != nil {
			return
		}

	case "Собачьи":
		wolf.Move()
		wolf.Eat()
		wolf.Fly()
		wolf.Climb()
		wolf.Swim()
		err := sound.PlaySound(wolf)
		if err != nil {
			return
		}

	case "Птицы":
		parrot.Move()
		parrot.Eat()
		parrot.Fly()
		parrot.Climb()
		parrot.Swim()
		err := sound.PlaySound(parrot)
		if err != nil {
			return
		}

	case "Грызуны":
		squirrel.Move()
		squirrel.Eat()
		squirrel.Fly()
		squirrel.Climb()
		squirrel.Swim()
		err := sound.PlaySound(squirrel)
		if err != nil {
			return
		}

	default:
		fmt.Println("Неизвестный ввод.")
	}
}
