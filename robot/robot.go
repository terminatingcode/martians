package robot

import "fmt"

var orientations = [4]string{"N", "E", "S", "W"}

type Robot struct {
	coordinates [2]int
	orientation int
	connected   bool
}

func Create(x, y int, orientation string) (Robot, error) {
	o := orientate(orientation)
	if o == -1 {
		return Robot{}, fmt.Errorf("invalid input %s", orientation)
	}
	return Robot{
		coordinates: [2]int{x, y},
		orientation: o,
		connected:   true,
	}, nil
}

func orientate(orientation string) int {
	for i, o := range orientations {
		if o == orientation {
			return i
		}
	}
	return -1
}

func (r *Robot) Rotate(direction string) error {
	switch direction {
	case "L":
		r.orientation = (r.orientation + 4 - 1) % 4
	case "R":
		r.orientation = (r.orientation + 1) % 4
	default:
		return fmt.Errorf("invalid input %s", direction)
	}
	return nil
}

func (r *Robot) Forward() {
	if r.connected {
		switch r.orientation {
		case 0:
			r.coordinates[1] += 1
		case 1:
			r.coordinates[0] += 1
		case 2:
			r.coordinates[1] -= 1
		case 3:
			r.coordinates[0] -= 1
		}
	}
}

func (r Robot) IsConnected() bool {
	return r.connected
}

func (r *Robot) Disconnect() {
	r.connected = false
}

func (r Robot) ToString() string {
	connectedString := ""
	if !r.connected {
		connectedString = " LOST"
	}
	return fmt.Sprintf("%d %d %s%s", r.coordinates[0], r.coordinates[1], orientations[r.orientation], connectedString)
}
