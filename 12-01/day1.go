package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	debug := false
	depth_increases := 0
	depth_splice := loadfile("./input-day1.txt")
	depth_window_splice := make([]int64, 0)

	//load depth window splices
	for current_pos, _ := range depth_splice {
		current_depth_window := int64(0)
		if current_pos+2 < len(depth_splice) {
			if debug {
				fmt.Println(" - step 1", depth_splice[current_pos])
				fmt.Println(" - step 2", depth_splice[current_pos+1])
				fmt.Println(" - step 2", depth_splice[current_pos+2])
			}
			current_depth_window = int64(depth_splice[current_pos] + depth_splice[current_pos+1] + depth_splice[current_pos+2])
			depth_window_splice = append(depth_window_splice, current_depth_window)
			if debug {
				fmt.Println("Current Depth: ", current_depth_window)
			}
		}
	}

	previous_depth_window := int64(0)
	for pos, current := range depth_window_splice {
		if pos <= 0 {
			continue
		}
		if current > previous_depth_window {
			depth_increases += 1
		}
		previous_depth_window = current
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
