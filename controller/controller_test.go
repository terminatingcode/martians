package controller

import (
	"fmt"
	"testing"
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
		controller, err := Create(50, 50)
		controller.ConnectRobot(c.x, c.y, c.orientation)
		if err != nil && !c.wantErr {
			t.Errorf("ConnectRobot(%d, %d, %s) unexpected error %s", c.x, c.y, c.orientation, err.Error())
		}
		if len(controller.robots) != c.want {
			t.Errorf("ConnectRobot(%d, %d, %s) resulted in robots length %d", c.x, c.y, c.orientation, len(controller.robots))
		}
	}
}
