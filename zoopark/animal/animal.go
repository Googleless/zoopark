package animal

type Animal interface {
	Move()
	Eat()
	Fly()
	Swim()
	Climb()
	SoundFile() string
	DisplayInfo()
}
