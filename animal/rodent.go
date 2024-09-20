package animal

import "fmt"

type Rodent struct{}

func (r Rodent) Move() {
	fmt.Println("Грызуны ходят по земле, а также прячутся по норам, которые вырывают как свои жилища.")
}

func (r Rodent) Eat() {
	fmt.Println("Грызуны питаются семенами, остатками человеческой еды, ягодами и насекомыми.")
}

func (r Rodent) Fly() {
	fmt.Println("Грызуны не умеют летать. Такие грызуны как белка-летяга не летят, а планируют с помощью перепонок у своих конечностей.")
}

func (r Rodent) Swim() {
	fmt.Println("Грызуны умеют плавать до получаса прежде чем выдохнуться.")
}

func (r Rodent) Climb() {
	fmt.Println("Некоторым видам грызунов не обязательно карабкаться. Однако роду белок необходимо карабкаться, и делают они так с легкостью.")
}

func (r Rodent) SoundFile() string {
	return "Rodent.mp3"
}
