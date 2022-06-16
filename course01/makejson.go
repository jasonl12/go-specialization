package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	m := make(map[string]string)
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\nEnter a name: ")
	name, _ := reader.ReadString('\n')

	fmt.Print("Enter an address: ")
	address, _ := reader.ReadString('\n')

	m["name"] = strings.TrimSuffix(name, "\n")
	m["address"] = strings.TrimSuffix(address, "\n")

	b, err := json.Marshal(m)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	os.Stdout.Write(b)
	os.Stdout.Write([]byte("\n"))
}
