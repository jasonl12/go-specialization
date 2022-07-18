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

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Data struct {
	food       string
	locomotion string
	noise      string
}

func (d Data) Eat() {
	fmt.Println(d.food)
}

func (d Data) Move() {
	fmt.Println(d.locomotion)
}

func (d Data) Speak() {
	fmt.Println(d.noise)
}

// "type of the new animal, either cow, bird, or snake"
func NewData(info string) *Data {
	if info == "cow" {
		return &Data{
			food:       "grass",
			locomotion: "walk",
			noise:      "moo",
		}
	} else if info == "bird" {
		return &Data{
			food:       "worms",
			locomotion: "fly",
			noise:      "peep",
		}
	} else if info == "snake" {
		return &Data{
			food:       "mice",
			locomotion: "slither",
			noise:      "hsss",
		}
	} else {
		return nil
	}
}

func invoke(any interface{}, value string) {
	reflect.ValueOf(any).MethodByName(value).Call([]reflect.Value{})
}

func main() {
	// "command must be either newanimal or query"
	commands := map[string]bool{
		"newanimal": true,
		"query":     true,
	}

	// "requestd about the animal, either eat, move, or speak"
	attributes := map[string]string{
		"eat":   "Eat",
		"move":  "Move",
		"speak": "Speak",
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	records := map[string]*Data{}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter command (or Ctrl+C to exit):")

	go func() {
		for {
			fmt.Print("> ")
			scanner.Scan()
			inputs := strings.Split(scanner.Text(), " ")

			if len(inputs) != 3 {
				fmt.Println("Command must be a single line containing 3 strings.")
				continue
			} else {
				comm := strings.ToLower(inputs[0])
				name := strings.ToLower(inputs[1])
				info := strings.ToLower(inputs[2])

				_, found := commands[comm]
				if found {
					switch comm {
					case "newanimal":
						d := NewData(info)
						if d == nil {
							fmt.Println("Types of animals are restricted to either cow, bird, or snake.")
							continue
						}
						records[name] = d
					case "query":
						var a Animal
						v1, f1 := records[name]
						v2, f2 := attributes[info]
						if f1 && f2 {
							a = v1
							invoke(a, v2)
						} else {
							fmt.Printf("\"%s %s\" not found on the query\n", name, info)
							continue
						}
					}
				} else {
					fmt.Println("Command must be either \"newanimal\" or \"query\".")
					continue
				}
			}
		}
	}()

	<-ch
	fmt.Println("\n\nProgram is going to exit...")
}
