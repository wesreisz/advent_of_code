package main

import (
	"testing"
)

func Test_convertBinary2Decimal(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{
				input: "010101",
			},
			want: 21,
		},
		{
			args: args{
				input: "01010101010",
			},
			want: 682,
		},
		{
			args: args{
				input: "101010110101",
			},
			want: 2741,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertBinary2Decimal(tt.args.input); got != tt.want {
				t.Errorf("convertBinary2Decimal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loadData(t *testing.T) {
	input := loadfile("./sample.txt")
	data := loadData(input)
	expected_number_of_records := 12
	if len(data) != expected_number_of_records {
		t.Errorf("Did not load all the records. Expect %d got %d", expected_number_of_records, len(data))
	}
	expected_number_of_values := 5
	for i, val := range data {
		if len(val) != expected_number_of_values {
			t.Errorf("Did not load all the data in record %d. Expect %d got %d", i, expected_number_of_values, len(val))
		}
	}
}

func Test_getMostCommonBitOfRow(t *testing.T) {
	input := loadfile("./sample.txt")
	data := loadData(input)
	if getMostCommonBitOfRow(data, 0) != "1" {
		t.Errorf("First row should return a 1")
	}
	if getMostCommonBitOfRow(data, 1) != "0" {
		t.Errorf("Second row should return a 0")
	}
	if getMostCommonBitOfRow(data, 2) != "1" {
		t.Errorf("Third row should return a 1")
	}
	if getMostCommonBitOfRow(data, len(data[0])-1) != "0" {
		t.Errorf("Last row should return a 0")
	}
}

func Test_getGammaRate(t *testing.T) {
	input := loadfile("./sample.txt")
	data := loadData(input)
	result := getGammaRate(data)
	if result != 22 {
		t.Errorf("Expected GammaRate is 22 got %d", result)
	}
}

func Test_getLeastCommonBitOfRow(t *testing.T) {
	input := loadfile("./sample.txt")
	data := loadData(input)
	if getLeastCommonBitOfRow(data, 0) != "0" {
		t.Errorf("First row should return a 0")
	}
	if getLeastCommonBitOfRow(data, 1) != "1" {
		t.Errorf("Second row should return a 1")
	}
	if getLeastCommonBitOfRow(data, 2) != "0" {
		t.Errorf("Third row should return a 0")
	}
	if getLeastCommonBitOfRow(data, len(data[0])-1) != "1" {
		t.Errorf("Last row should return a 1")
	}
}

func Test_getEpsilonRate(t *testing.T) {
	input := loadfile("./sample.txt")
	data := loadData(input)
	result := getEpsilonRate(data)
	if result != 9 {
		t.Errorf("Expected EpsilonRate is 9 got %d", result)
	}
}
