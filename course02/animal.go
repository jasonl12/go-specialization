package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"reflect"
	"strings"
	"syscall"
)

// "three fields, all of which are strings"
type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

// To call a method on the object (reflect.ValueOf, MethodByName, Call)
func invoke(any interface{}, value string) {
	reflect.ValueOf(any).MethodByName(value).Call([]reflect.Value{})
}

func main() {
	cow := Animal{
		food:       "grass",
		locomotion: "walk",
		noise:      "moo",
	}

	bird := Animal{
		food:       "worms",
		locomotion: "fly",
		noise:      "peep",
	}

	snake := Animal{
		food:       "mice",
		locomotion: "slither",
		noise:      "hsss",
	}

	m := map[string]string{
		"eat":   "Eat",
		"move":  "Move",
		"speak": "Speak",
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a request (or Ctrl+C to exit):")

	go func() {
		for {
			fmt.Print("> ")
			scanner.Scan()
			inputs := strings.Split(scanner.Text(), " ")

			if len(inputs) != 2 {
				fmt.Println("Every request must be a single line with 2 strings.")
				continue
			} else {
				name := strings.ToLower(inputs[0])
				request := strings.ToLower(inputs[1])
				v, found := m[request]
				if found {
					switch name {
					case "cow":
						invoke(&cow, v)
					case "bird":
						invoke(&bird, v)
					case "snake":
						invoke(&snake, v)
					default:
						fmt.Println("Invalid request data.")
						continue
					}
				} else {
					fmt.Println("Invalid request data.")
					continue
				}
			}
		}
	}()

	<-ch
	fmt.Println("\n\nProgram is going to exit...")
}
