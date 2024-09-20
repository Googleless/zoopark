package animal

import "fmt"

type Bird struct{}

func (b Bird) Move() {
	fmt.Println("Птицы в большинстве своём летают.")
}

func (b Bird) Eat() {
	fmt.Println("Птицы питаются грызунами и остатками чужой охоты.")
}

func (b Bird) Fly() {
	fmt.Println("Основной способ передвижения для птиц — полёт.")
}

func (b Bird) Swim() {
	fmt.Println("Большинство птиц не умеют плавать.")
}

func (b Bird) Climb() {
	fmt.Println("Птицам необходимо карабкаться благодаря возможности летать.")
}

func (b Bird) SoundFile() string {
	return "Bird.mp3"
}
