package main

import "fmt"

type Animal interface {
	move()
	eat()
	fly()
	swim()
	climb()
}

type Bear struct{}

type Feline struct{}

type Canine struct{}

type Bird struct{}

type Rodent struct{}

func move(animal Animal) {
	animal.move()
}

func eat(animal Animal) {
	animal.eat()
}
func fly(animal Animal) {
	animal.fly()
}

func swim(animal Animal) {
	animal.swim()
}

func climb(animal Animal) {
	animal.climb()
}

// Медведи
func (mb Bear) move() {
	fmt.Println("Медведи ходят по земле.")
}

func (eb Bear) eat() {
	fmt.Println("Медведи питаются мясом, рыбой и ягодами. Есть медведи исключительно на травоядной диете.")
}

func (fb Bear) fly() {
	fmt.Println("Медведи не умеют летать.")
}

func (sb Bear) swim() {
	fmt.Println("При необходимости, медведи умеют очень хорошо плавать.")
}

func (cb Bear) climb() {
	fmt.Println("Медведи умеют хорошо карабкаться по деревьям.")
}

// Кошачьи
func (mf Feline) move() {
	fmt.Println("Кошачьи ходят по земле.")
}

func (ef Feline) eat() {
	fmt.Println("Кошачьи питаются мясом, рыбой, и птицами, в зависимости от вида.")
}

func (ff Feline) fly() {
	fmt.Println("Кошачьи не умеют летать.")
}

func (sf Feline) swim() {
	fmt.Println("В большинстве своём, кошачьи предпочитают не залезать в воду, но они умеют плавать при необходимости.")
}

func (cf Feline) climb() {
	fmt.Println("Большинство кошачьих умеют очень хорошо охотиться. Конкретно леопарды питаются на деревьях чтобы скрыться от других хищников.")
}

// Собачьи
func (mc Canine) move() {
	fmt.Println("Собачьи ходят по земле.")
}

func (ec Canine) eat() {
	fmt.Println("Собачьи питаются мясом, мелкими животными, и травоядными.")
}

func (fc Canine) fly() {
	fmt.Println("Собачьи не умеют летать.")
}

func (sc Canine) swim() {
	fmt.Println("Основное количество собачьих умеют хорошо плавать.")
}

func (cc Canine) climb() {
	fmt.Println("Собачьи не умеют карабкаться по деревьям.")
}

// Птицы
func (mbi Bird) move() {
	fmt.Println("Птицы в большинстве своём летают.")
}

func (ebi Bird) eat() {
	fmt.Println("Птицы питаются грызунами и остатками чужой охоты.")
}

func (fbi Bird) fly() {
	fmt.Println("Основной способ передвижения для птиц — полёт.")
}

func (sbi Bird) swim() {
	fmt.Println("Большинство птиц не умеют плавать.")
}

func (cbi Bird) climb() {
	fmt.Println("Птицам необходимо карабкаться благодаря возможности летать. Однако те, которые летать не умеют, карабкаться так же не способны.")
}

// Грызуны
func (mr Rodent) move() {
	fmt.Println("Грызуны ходят по земле, а так же прячутся по норам, которые вырывают как свои жилища.")
}

func (er Rodent) eat() {
	fmt.Println("Грызуны питаются семенами, остатками человеческой еды, ягодами, и насекомыми.")
}

func (fr Rodent) fly() {
	fmt.Println("Грызуны не умеют летать. Такие грызуны как белка-летяга не летят, а планируют с помощью перепонок у своих конечностей.")
}

func (sr Rodent) swim() {
	fmt.Println("Грызуны умеют плавать до получаса прежде чем выдохнуться.")
}

func (cr Rodent) climb() {
	fmt.Println("Некоторым видам грызунов не обязательно карабкаться. Однако роду белок необходимо карабкаться, и делают они так с легкостью.")
}
func main() {
	grizzly := Bear{}
	lynx := Feline{}
	wolf := Canine{}
	parrot := Bird{}
	squirrel := Rodent{}

	move(grizzly)
	eat(grizzly)
	fly(grizzly)
	climb(grizzly)
	swim(grizzly)
	move(lynx)
	eat(lynx)
	fly(lynx)
	climb(lynx)
	swim(lynx)
	move(wolf)
	eat(wolf)
	fly(wolf)
	climb(wolf)
	swim(wolf)
	move(parrot)
	eat(parrot)
	fly(parrot)
	climb(parrot)
	swim(parrot)
	move(squirrel)
	eat(squirrel)
	fly(squirrel)
	climb(squirrel)
	swim(squirrel)
}
