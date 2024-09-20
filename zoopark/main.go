package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"hello-go/zoopark/animal"
	"hello-go/zoopark/sound"
)

func main() {
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
		sound.PlaySound(grizzly)

	case "Кошачьи":
		lynx.Move()
		lynx.Eat()
		lynx.Fly()
		lynx.Climb()
		lynx.Swim()
		sound.PlaySound(lynx)

	case "Собачьи":
		wolf.Move()
		wolf.Eat()
		wolf.Fly()
		wolf.Climb()
		wolf.Swim()
		sound.PlaySound(wolf)

	case "Птицы":
		parrot.Move()
		parrot.Eat()
		parrot.Fly()
		parrot.Climb()
		parrot.Swim()
		sound.PlaySound(parrot)

	case "Грызуны":
		squirrel.Move()
		squirrel.Eat()
		squirrel.Fly()
		squirrel.Climb()
		squirrel.Swim()
		sound.PlaySound(squirrel)

	default:
		fmt.Println("Неизвестный ввод.")
	}
}
