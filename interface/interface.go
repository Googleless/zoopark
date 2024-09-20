package main

import (
	"fmt"
	"math"
)

// Интерфейс Shape, который задаёт поведение для всех фигур
type Shape interface {
	Area() float64      // Метод для вычисления площади
	Perimeter() float64 // Метод для вычисления периметра
}

// Структура Circle (Круг)
type Circle struct {
	Radius float64
}

// Реализация методов интерфейса для Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Структура Rectangle (Прямоугольник)
type Rectangle struct {
	Width, Height float64
}

// Реализация методов интерфейса для Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Функция, которая принимает интерфейс Shape и выводит информацию о фигуре
func printShapeInfo(s Shape) {
	fmt.Printf("Площадь: %.2f\n", s.Area())
	fmt.Printf("Периметр: %.2f\n", s.Perimeter())
}

func main() {
	// Создаём объекты разных фигур
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 4, Height: 3}

	// Используем интерфейс Shape для обработки разных фигур
	fmt.Println("Информация о круге:")
	printShapeInfo(circle)

	fmt.Println("\nИнформация о прямоугольнике:")
	printShapeInfo(rectangle)

}
