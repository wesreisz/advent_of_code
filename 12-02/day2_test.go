package main

import (
	"testing"
)

func TestResult(t *testing.T) {
	sub := NewSubmarine("./sample.txt")
	sub.execute()

	if sub.horizontal_position != 15 {
		t.Errorf("horizontal_position: got %d, wanted %d", sub.horizontal_position, 15)
	}

	if sub.depth != 60 {
		t.Errorf("depth: got %d, wanted %d", sub.depth, 60)
	}

	got := sub.horizontal_position * sub.depth
	want := 900

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestDown(t *testing.T) {
	sub := NewSubmarine("./sample.txt")
	sub.move_down(5)

	got := sub.aim
	want := 5

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestForwardStep1(t *testing.T) {
	sub := NewSubmarine("./sample.txt")
	sub.move_forward(5)

	if sub.horizontal_position != 5 {
		t.Errorf("got %d, wanted %d", sub.horizontal_position, 5)
	}
}

func TestDownStep2(t *testing.T) {
	sub := NewSubmarine("./sample.txt")
	sub.move_forward(5)
	sub.move_down(5)

	if sub.aim != 5 {
		t.Errorf("got %d, wanted %d", sub.horizontal_position, 5)
	}
}

func TestForwardStep3(t *testing.T) {
	sub := NewSubmarine("./sample.txt")
	sub.move_forward(5)
	sub.move_down(5)
	sub.move_forward(8)

	if sub.horizontal_position != 13 {
		t.Errorf("got %d, wanted %d", sub.horizontal_position, 13)
	}
	if sub.depth != 40 {
		t.Errorf("got %d, wanted %d", sub.depth, 40)
	}
}

func TestForwardStep3b(t *testing.T) {
	sub := NewSubmarine("./sample.txt")
	sub.horizontal_position = 5
	sub.aim = 5
	sub.depth = 0
	sub.move_forward(8)

	if sub.depth != 40 {
		t.Errorf("got %d, wanted %d", sub.depth, 40)
	}
}

func TestUpStep4(t *testing.T) {
	sub := NewSubmarine("./sample.txt")
	sub.horizontal_position = 5
	sub.aim = 5
	sub.depth = 0
	sub.move_up(3)

	if sub.aim != 2 {
		t.Errorf("got %d, wanted %d", sub.aim, 2)
	}

}

func TestDownStep5(t *testing.T) {
	sub := NewSubmarine("./sample.txt")
	sub.horizontal_position = 5
	sub.aim = 2
	sub.depth = 0
	sub.move_down(8)

	if sub.aim != 10 {
		t.Errorf("got %d, wanted %d", sub.aim, 10)
	}
}

func TestForwardStep6(t *testing.T) {
	sub := NewSubmarine("./sample.txt")
	sub.horizontal_position = 13
	sub.aim = 10
	sub.depth = 40
	sub.move_forward(2)

	if sub.horizontal_position != 15 {
		t.Errorf("got %d, wanted %d", sub.horizontal_position, 15)
	}

	if sub.depth != 60 {
		t.Errorf("got %d, wanted %d", sub.depth, 60)
	}
}
