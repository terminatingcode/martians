package controller

import (
	"fmt"

	r "github.com/terminatingcode/martians/robot"
)

type Controller struct {
	robots    []r.Robot
	x         int
	y         int
	memorials map[int]int
}

func Create(x, y int) (Controller, error) {
	if x > 50 {
		return Controller{}, fmt.Errorf("horizontal limit %d greater than 50", x)
	}
	if y > 50 {
		return Controller{}, fmt.Errorf("vertical limit %d greater than 50", y)
	}

	return Controller{
		robots:    []r.Robot{},
		x:         x,
		y:         y,
		memorials: make(map[int]int),
	}, nil
}
