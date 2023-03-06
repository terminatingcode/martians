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
