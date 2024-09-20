package animal

import "fmt"

type Canine struct{}

func (c Canine) Move() {
	fmt.Println("Собачьи ходят по земле.")
}

func (c Canine) Eat() {
	fmt.Println("Собачьи питаются мясом, мелкими животными, и травоядными.")
}

func (c Canine) Fly() {
	fmt.Println("Собачьи не умеют летать.")
}

func (c Canine) Swim() {
	fmt.Println("Основное количество собачьих умеют хорошо плавать.")
}

func (c Canine) Climb() {
	fmt.Println("Собачьи не умеют карабкаться по деревьям.")
}

func (c Canine) SoundFile() string {
	return "Canine.mp3"
}
