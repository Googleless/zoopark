package animal

import "fmt"

type Feline struct{}

func (f Feline) Move() {
	fmt.Println("Кошачьи ходят по земле.")
}

func (f Feline) Eat() {
	fmt.Println("Кошачьи питаются мясом, рыбой и птицами.")
}

func (f Feline) Fly() {
	fmt.Println("Кошачьи не умеют летать.")
}

func (f Feline) Swim() {
	fmt.Println("Кошачьи умеют плавать при необходимости.")
}

func (f Feline) Climb() {
	fmt.Println("Большинство кошачьих умеют очень хорошо карабкаться.")
}

func (f Feline) SoundFile() string {
	return "sound/Feline.mp3"
}

func (f Feline) DisplayInfo() {
	fmt.Println("Информация о кошачьих")
	f.Move()
	f.Eat()
	f.Fly()
	f.Swim()
	f.Climb()
}
