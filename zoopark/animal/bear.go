package animal

import "fmt"

type Bear struct{}

func (b Bear) Move() {
	fmt.Println("Медведи ходят по земле.")
}

func (b Bear) Eat() {
	fmt.Println("Медведи питаются мясом, рыбой и ягодами.")
}

func (b Bear) Fly() {
	fmt.Println("Медведи не умеют летать.")
}

func (b Bear) Swim() {
	fmt.Println("При необходимости, медведи умеют очень хорошо плавать.")
}

func (b Bear) Climb() {
	fmt.Println("Медведи умеют хорошо карабкаться по деревьям.")
}

func (b Bear) SoundFile() string {
	return "Bear.mp3"
}
