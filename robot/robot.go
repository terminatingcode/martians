package robot

import "fmt"

var orientations = [4]string{"N", "E", "S", "W"}

type Robot struct {
	coordinates [2]int
	orientation int
	connected   bool
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
