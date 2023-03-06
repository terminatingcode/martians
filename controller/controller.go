package controller

import (
	"fmt"
	"strings"

	r "github.com/terminatingcode/martians/robot"
)

type Controller struct {
	robots    []*r.Robot
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
		robots:    []*r.Robot{},
		x:         x,
		y:         y,
		memorials: map[int]int{},
	}, nil
}

func (c *Controller) ConnectRobot(x, y int, orientation string) (*r.Robot, error) {
	robot, err := r.Create(x, y, orientation)
	if err != nil {
		return &robot, err
	}
	c.robots = append(c.robots, &robot)
	return &robot, nil
}

func (c *Controller) DirectRobot(robot *r.Robot, input string) error {
	switch input {
	case "L":
		fallthrough
	case "R":
		robot.Rotate(input)
		return nil
	case "F":
		if robot.IsConnected() {
			coordinates := robot.Location()
			robot.Forward(c.x, c.y, c.isMemorial(coordinates[0], coordinates[1]))
			if !robot.IsConnected() {
				c.memorials[coordinates[0]] = coordinates[1]
				return nil
			}
		}
		return nil
	default:
		return fmt.Errorf("invalid input %s", input)
	}
}

func (c Controller) isMemorial(x, y int) bool {
	if cy, found := c.memorials[x]; found {
		if cy == y {
			return true
		}
	}
	return false
}

func (c Controller) ToString() string {
	var sb strings.Builder
	for _, r := range c.robots {
		sb.WriteString(fmt.Sprintf("%s\n", r.ToString()))
	}
	return sb.String()
}
