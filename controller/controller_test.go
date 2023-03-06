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
