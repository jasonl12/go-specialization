package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const stringSize = 20

type Name struct {
	fname string
	lname string
}

func limitString(fname, lname string) (string, string) {
	first := fname
	last := lname

	if len(fname) > stringSize {
		first = fname[:stringSize]
	}

	if len(lname) > stringSize {
		last = lname[:stringSize]
	}

	return first, last
}

func getFile() (string, error) {
	path, _ := os.Getwd()
	fmt.Printf("\nWorking directory: %s\n", path)
	fmt.Print("Enter the name of text file: ")

	reader := bufio.NewReader(os.Stdin)
	fileName, _ := reader.ReadString('\n')
	fileName = strings.TrimSuffix(fileName, "\n")

	if _, err := os.Stat(fileName); err != nil {
		return fileName, err
	}

	return fileName, nil
}

func readFile(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	var names []Name
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		values := strings.Split(scanner.Text(), " ")
		fn, ln := limitString(values[0], values[1])
		n := Name{fname: fn, lname: ln}
		names = append(names, n)
	}

	for _, item := range names {
		fmt.Printf("%+v\n", item)
	}

	return nil
}

/*
file.txt

Golang Google
Robert Griesemer
Rob Pike
Ken Thompson
Russ Cox
aaaabbbbccccxxxxyyyyzzzz foobar
 foobar
Gophers abcdefghijklmnopqrstuvwxyz
*/
func main() {
	file, err := getFile()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%q file not exists\n", file)
		os.Exit(1)
	}

	if err := readFile(file); err != nil {
		log.Fatal(err)
	}
}
