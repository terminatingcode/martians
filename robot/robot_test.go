package robot

import (
	"fmt"
	"testing"
)

func TestCreate(t *testing.T) {
	cases := []struct {
		x           int
		y           int
		orientation string
		want        Robot
		err         error
	}{
		{0, 0, "N", Robot{orientation: 0, coordinates: [2]int{0, 0}, connected: true}, nil},
		{10, 0, "E", Robot{orientation: 1, coordinates: [2]int{10, 0}, connected: true}, nil},
		{0, 10, "S", Robot{orientation: 2, coordinates: [2]int{0, 10}, connected: true}, nil},
		{10, 10, "W", Robot{orientation: 3, coordinates: [2]int{10, 10}, connected: true}, nil},
		{0, 0, "L", Robot{}, fmt.Errorf("invalid input %s", "L")},
	}

	for _, c := range cases {
		robot, err := Create(c.x, c.y, c.orientation)
		if err != nil {
			if err.Error() != c.err.Error() {
				t.Errorf("Create(%d, %d, %s) == %v, want %v", c.x, c.y, c.orientation, err.Error(), c.err.Error())
			}
		}
		if robot != c.want {
			t.Errorf("Create(%d, %d, %s) == %v, want %v", c.x, c.y, c.orientation, robot, c.want)
		}
	}
}

func TestRotate(t *testing.T) {
	cases := []struct {
		in   string
		want int
		err  error
	}{
		{"L", 3, nil},
		{"R", 0, nil},
		{"R", 1, nil},
		{"L", 0, nil},
		{"N", 0, fmt.Errorf("invalid input %s", "N")},
	}

	robot := Robot{}
	for _, c := range cases {
		err := robot.Rotate(c.in)
		if err != nil {
			if err.Error() != c.err.Error() {
				t.Errorf("Rotate(%s) == %s, want %s", c.in, err.Error(), c.err.Error())
			}
		}
		if robot.orientation != c.want {
			t.Errorf("Rotate(%s) == %d, want %d", c.in, robot.orientation, c.want)
		}
	}
}

func TestForward(t *testing.T) {
	cases := []struct {
		robot      Robot
		hasWarning bool
		want       Robot
	}{
		{Robot{orientation: 0, coordinates: [2]int{0, 0}, connected: true}, false, Robot{orientation: 0, coordinates: [2]int{0, 1}, connected: true}},
		{Robot{orientation: 1, coordinates: [2]int{0, 0}, connected: true}, false, Robot{orientation: 1, coordinates: [2]int{1, 0}, connected: true}},
		{Robot{orientation: 2, coordinates: [2]int{1, 1}, connected: true}, false, Robot{orientation: 2, coordinates: [2]int{1, 0}, connected: true}},
		{Robot{orientation: 3, coordinates: [2]int{1, 1}, connected: true}, false, Robot{orientation: 3, coordinates: [2]int{0, 1}, connected: true}},
		{Robot{orientation: 1, coordinates: [2]int{5, 5}, connected: true}, false, Robot{orientation: 1, coordinates: [2]int{5, 5}, connected: false}},
		{Robot{orientation: 1, coordinates: [2]int{5, 5}, connected: true}, false, Robot{orientation: 1, coordinates: [2]int{5, 5}, connected: false}},
		{Robot{orientation: 2, coordinates: [2]int{0, 0}, connected: true}, false, Robot{orientation: 2, coordinates: [2]int{0, 0}, connected: false}},
		{Robot{orientation: 3, coordinates: [2]int{0, 0}, connected: true}, false, Robot{orientation: 3, coordinates: [2]int{0, 0}, connected: false}},
		{Robot{orientation: 3, coordinates: [2]int{0, 0}, connected: true}, true, Robot{orientation: 3, coordinates: [2]int{0, 0}, connected: true}},
	}
	for _, c := range cases {
		c.robot.Forward(5, 5, c.hasWarning)
		if c.robot != c.want {
			t.Errorf("Robot[%v].Forward() want %v", c.robot, c.want)
		}
	}
}

func TestToString(t *testing.T) {
	cases := []struct {
		robot Robot
		want  string
	}{
		{Robot{connected: true}, "0 0 N"},
		{Robot{coordinates: [2]int{5, 10}, orientation: 1}, "5 10 E LOST"},
		{Robot{coordinates: [2]int{-1, 0}, orientation: 2}, "-1 0 S LOST"},
		{Robot{coordinates: [2]int{0, 0}, orientation: 3, connected: true}, "0 0 W"},
	}
	for _, c := range cases {
		got := c.robot.ToString()
		if got != c.want {
			t.Errorf("Robot[%v].ToString() == %s, want %s", c.robot, got, c.want)
		}
	}
}
