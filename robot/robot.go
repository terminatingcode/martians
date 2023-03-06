package robot

import "fmt"

var orientations = [4]string{"N", "E", "S", "W"}

type Robot struct {
	coordinates []int
	orientation int
	live        bool
}

func (r *Robot) Rotate(direction string) error {
	orientation := r.orientation
	switch direction {
	case "L":
		r.orientation = (orientation + 4 - 1) % 4
	case "R":
		r.orientation = (orientation + 1) % 4
	default:
		return fmt.Errorf("invalid input %s", direction)
	}
	return nil
}
