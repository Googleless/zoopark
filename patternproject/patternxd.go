package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Command interface {
	Execute()
}

type LightReceiver struct{}

func (l *LightReceiver) TurnOn() {
	fmt.Println("Свет включен")
}

func (l *LightReceiver) TurnOff() {
	fmt.Println("Свет выключен")
}

type LightOnCommand struct {
	light *LightReceiver
}

func (c *LightOnCommand) Execute() {
	c.light.TurnOn()
}

type LightOffCommand struct {
	light *LightReceiver
}

func (c *LightOffCommand) Execute() {
	c.light.TurnOff()
}

type RemoteControl struct {
	command Command
}

func (r *RemoteControl) SetCommand(command Command) {
	r.command = command
}

func (r *RemoteControl) PressButton() {
	r.command.Execute()
}

func main() {
	light := &LightReceiver{}
	lightOn := &LightOnCommand{light: light}
	lightOff := &LightOffCommand{light: light}
	remote := &RemoteControl{}

	lightswitch := false // состояние света
	reader := bufio.NewReader(os.Stdin)

	for {
		if lightswitch {
			fmt.Println("Свет включен. Напишите 'Выключить' для выключения или 'Выход' для завершения.")
		} else {
			fmt.Println("Свет выключен. Напишите 'Включить' для включения или 'Выход' для завершения.")
		}

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "Включить":
			if !lightswitch {
				remote.SetCommand(lightOn)
				remote.PressButton()
				lightswitch = true
			} else {
				fmt.Println("Свет уже включен.")
			}
		case "Выключить":
			if lightswitch {
				remote.SetCommand(lightOff)
				remote.PressButton()
				lightswitch = false
			} else {
				fmt.Println("Свет уже выключен.")
			}
		case "Выход":
			fmt.Println("Программа завершена.")
			return
		default:
			fmt.Println("Неизвестная команда. Используйте 'Включить', 'Выключить' или 'Выход'.")
		}
	}
}
