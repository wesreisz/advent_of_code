package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var debug bool = false

func main() {
	sub := NewSubmarine("./input.txt")
	sub.execute()
	fmt.Println("\n------------------------------------------------------------------------------------------------")
	fmt.Println("depth: ", sub.depth)
	fmt.Println("horizontal position: ", sub.horizontal_position)
	fmt.Println("What do you get if you multiply your final horizontal position by your final depth?", sub.depth*sub.horizontal_position)
	fmt.Println("-----------------------------------------------------------------------------------------------")
}

type Submarine struct {
	instructions        []string
	aim                 int
	depth               int
	horizontal_position int
}

func NewSubmarine(input string) *Submarine {
	sub := new(Submarine)
	sub.aim = 0
	sub.depth = 0
	sub.horizontal_position = 0

	sub.instructions = loadfile(input)

	return sub
}

func (sub *Submarine) execute() {
	for _, val := range sub.instructions {
		sub.move(val)

		if debug {
			fmt.Println(" ", val)
			fmt.Println(" - aim", sub.aim)
			fmt.Println(" - depth", sub.depth)
			fmt.Println(" - horizontal position", sub.horizontal_position)
		}
	}
}

func (sub *Submarine) move(cmd string) string {
	s := strings.Split(cmd, " ")
	command_verb := s[0]
	command_movement, _ := strconv.Atoi(s[1])
	if "forward" == command_verb {
		sub.move_forward(command_movement)
	} else {
		if "down" == command_verb {
			sub.move_down(command_movement)
		} else {
			sub.move_up(command_movement)
		}
	}

	return cmd
}

func (sub *Submarine) move_forward(command_movement int) {
	sub.horizontal_position = sub.horizontal_position + command_movement
	sub.depth += sub.aim * command_movement
}

func (sub *Submarine) move_down(command_movement int) {
	sub.aim = sub.aim + command_movement
}

func (sub *Submarine) move_up(command_movement int) {
	sub.aim = sub.aim - command_movement
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
