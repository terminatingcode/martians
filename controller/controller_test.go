package controller

import (
	"fmt"
	"testing"

	r "github.com/terminatingcode/martians/robot"
)

func TestCreate(t *testing.T) {
	cases := []struct {
		x    int
		y    int
		want Controller
		err  error
	}{
		{0, 0, Controller{x: 0, y: 0}, nil},
		{10, 50, Controller{x: 10, y: 50}, nil},
		{51, 0, Controller{}, fmt.Errorf("horizontal limit %d greater than 50", 51)},
		{0, 51, Controller{}, fmt.Errorf("vertical limit %d greater than 50", 51)},
	}

	for _, c := range cases {
		controller, err := Create(c.x, c.y)
		if err != nil {
			if err.Error() != c.err.Error() {
				t.Errorf("Create(%d, %d) == %v, want %v", c.x, c.y, err.Error(), c.err.Error())
			}
		}
		if controller.x != c.want.x || controller.y != c.want.y {
			t.Errorf("Create(%d, %d) == %v, want %v", c.x, c.y, controller, c.want)
		}
	}
}

func TestConnectRobot(t *testing.T) {
	cases := []struct {
		x           int
		y           int
		orientation string
		want        int
		wantErr     bool
	}{
		{10, 50, "N", 1, false},
		{0, 0, "L", 0, true},
	}

	for _, c := range cases {
		controller, _ := Create(50, 50)
		_, err := controller.ConnectRobot(c.x, c.y, c.orientation)
		if err != nil && !c.wantErr {
			t.Errorf("ConnectRobot(%d, %d, %s) unexpected error %s", c.x, c.y, c.orientation, err.Error())
		}
		if len(controller.robots) != c.want {
			t.Errorf("ConnectRobot(%d, %d, %s) resulted in robots length %d", c.x, c.y, c.orientation, len(controller.robots))
		}
	}
}

func TestDirectRobot(t *testing.T) {
	robot, _ := r.Create(0, 0, "N")
	robot2, _ := r.Create(0, 1, "N")
	cases := []struct {
		robot       *r.Robot
		input       string
		want        string
		memorialLen int
		err         error
	}{
		{&robot, "R", "0 0 E", 0, nil},
		{&robot, "L", "0 0 N", 0, nil},
		{&robot, "F", "0 1 N", 0, nil},
		{&robot, "T", "0 1 N", 0, fmt.Errorf("invalid input T")},
		{&robot, "F", "0 1 N LOST", 1, nil},
		{&robot2, "F", "0 1 N", 1, nil},
	}

	controller, _ := Create(1, 1)
	for i, c := range cases {
		err := controller.DirectRobot(c.robot, c.input)
		if err != nil {
			if c.err != nil && err.Error() != c.err.Error() {
				t.Errorf("DirectRobot() recieved error %s wanted %s", err.Error(), c.err.Error())
			}
		}
		if c.robot.ToString() != c.want {
			t.Errorf("[%d] DirectRobot() resulted in robot %s wanted %s", i, c.robot.ToString(), c.want)
		}
		if len(controller.memorials) != c.memorialLen {
			t.Errorf("[%d] DirectRobot() resulted in memorials %d wanted %d", i, len(controller.memorials), c.memorialLen)
		}
	}
}
