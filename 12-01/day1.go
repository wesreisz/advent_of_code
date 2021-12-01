package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	depth_increases := 0
	depth_splice := loadfile("./input-day1.txt")
	previous_depth := int64(0)

	//fmt.Println(previous_depth)

	for current_pos, current_depth := range depth_splice {
		if current_pos <= 0 {
			continue
		}
		previous_depth = depth_splice[current_pos-1]
		if current_depth > previous_depth {
			depth_increases += 1
		}
	}

	fmt.Println("Depth Increases: ", depth_increases)
}

func loadfile(file_name string) []int64 {
	depths := make([]int64, 0)

	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		input_num, err2 := strconv.ParseInt(scanner.Text(), 10, 64)
		if err2 != nil {
			log.Fatal(err2)
		}
		depths = append(depths, input_num)

		i += 1
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return depths
}
