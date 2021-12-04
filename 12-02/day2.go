package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var debug bool = true

func main() {
	sub := NewSubmarine("./input.txt")

	fmt.Println("depth: ", sub.depth)
	fmt.Println("horizontal position: ", sub.horizontal_position)
	fmt.Println("What do you get if you multiply your final horizontal position by your final depth?", sub.depth*sub.horizontal_position)
}

type Submarine struct {
	instructions        []string
	depth               int
	horizontal_position int
}

func NewSubmarine(input string) *Submarine {
	sub := new(Submarine)
	sub.instructions = loadfile(input)
	for _, val := range sub.instructions {
		sub.move(val)
	}
	return sub
}

func (sub *Submarine) move(cmd string) string {
	s := strings.Split(cmd, " ")
	if debug {
		fmt.Println("val: ", s)
	}
	command_verb := s[0]
	command_movement, _ := strconv.Atoi(s[1])
	if "forward" == command_verb {
		sub.horizontal_position += command_movement
	} else {
		if "down" == command_verb {
			sub.depth = sub.depth + command_movement
		} else {
			sub.depth = sub.depth - command_movement
		}
	}

	return cmd
}

func loadfile(file_name string) []string {
	lines := make([]string, 0)

	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		input_num := scanner.Text()
		lines = append(lines, input_num)

		i += 1
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}
