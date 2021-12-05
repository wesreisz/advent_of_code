package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var debug = false

func main() {
	fmt.Println("------------------------------------------------------------------------------------------------")
	fmt.Println("Day 3")
	fmt.Println("------------------------------------------------------------------------------------------------")
	input := loadfile("./input.txt")
	data := loadData(input)

	gamma_rate := getGammaRate(data)
	epsilon_rate := getEpsilonRate(data)
	fmt.Println("gamma_rate:", gamma_rate)
	fmt.Println("epsilon_rate:", epsilon_rate)
	fmt.Println("power consumption:", gamma_rate*epsilon_rate)
}
func convertBinary2Decimal(input string) int64 {
	output, err := strconv.ParseInt(input, 2, 64)
	if err != nil {
		log.Fatal("Unable to parse string", err)
	}
	return output
}
func loadfile(file_name string) []string {
	lines := make([]string, 0)

	file, _ := os.Open(file_name)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		i += 1
	}
	return lines
}
func loadData(input []string) [][]string {
	records := make([][]string, 0)

	for _, val := range input {
		records = append(records, strings.Split(val, ""))
	}
	return records
}

func getMostCommonBitOfRow(input_data [][]string, row_num int) string {
	bitmap := make(map[string]int)
	for i := 1; i < len(input_data); i++ {
		val := input_data[i][row_num]
		bitmap[val] += 1
	}

	max_key := ""
	for key, val := range bitmap {
		if val > bitmap[max_key] {
			max_key = key
		}
	}
	if debug {
		fmt.Println("max record: ", max_key, "man count: ", bitmap[max_key])
	}
	return max_key
}

func getLeastCommonBitOfRow(input_data [][]string, row_num int) string {
	bitmap := make(map[string]int)
	for i := 1; i < len(input_data); i++ {
		val := input_data[i][row_num]
		bitmap[val] += 1
	}

	min_key := ""
	for key, val := range bitmap {
		if min_key == "" {
			min_key = key
			continue
		}
		if val < bitmap[min_key] {
			min_key = key
		}
	}
	if debug {
		fmt.Println("Least record: ", min_key, "man count: ", bitmap[min_key])
	}
	return min_key
}

func getGammaRate(input_data [][]string) int64 {
	size := len(input_data[0])
	result := make([]string, size)
	for i := 0; i < size; i++ {
		result[i] = getMostCommonBitOfRow(input_data, i)
	}
	return convertBinary2Decimal(strings.Join(result, ""))
}

func getEpsilonRate(input_data [][]string) int64 {
	size := len(input_data[0])
	result := make([]string, size)
	for i := 0; i < size; i++ {
		result[i] = getLeastCommonBitOfRow(input_data, i)
	}
	return convertBinary2Decimal(strings.Join(result, ""))
}
