package robot

import (
	"fmt"
	"testing"
)

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
		robot Robot
		want  [2]int
	}{
		{Robot{orientation: 0, coordinates: [2]int{0, 0}, connected: true}, [2]int{0, 1}},
		{Robot{orientation: 1, coordinates: [2]int{0, 0}, connected: true}, [2]int{1, 0}},
		{Robot{orientation: 2, coordinates: [2]int{0, 0}, connected: true}, [2]int{0, -1}},
		{Robot{orientation: 3, coordinates: [2]int{0, 0}, connected: true}, [2]int{-1, 0}},
	}
	for _, c := range cases {
		c.robot.Forward()
		if c.robot.coordinates != c.want {
			t.Errorf("Robot[%v].Forward() == %d, want %d", c.robot, c.robot.coordinates, c.want)
		}
	}
}

func TestConnected(t *testing.T) {
	cases := []struct {
		robot Robot
		want  bool
	}{
		{Robot{connected: true}, false},
	}
	for _, c := range cases {
		c.robot.Disconnect()
		if c.robot.connected != c.want {
			t.Errorf("Robot[%v].Disconnect() == %t, want %t", c.robot, c.robot.connected, c.want)
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
