package coordinate

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Coordinate struct {
	x, y, z float64
}

func Parse(s string) (Coordinate, error) {
	// Remove parentheses
	s = strings.TrimSpace(s)
	if len(s) < 2 || s[0] != '(' || s[len(s)-1] != ')' {
		return Coordinate{}, errors.New("invalid format: input must be in the form '(x, y, z)'")
	}
	s = s[1 : len(s)-1] // Remove '(' and ')'

	// Split into components
	parts := strings.Split(s, ",")
	if len(parts) != 3 {
		return Coordinate{}, errors.New("invalid format: expected exactly 3 components")
	}

	// Parse x, y, z
	x, err := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
	if err != nil {
		return Coordinate{}, fmt.Errorf("invalid x value: %v", err)
	}

	y, err := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
	if err != nil {
		return Coordinate{}, fmt.Errorf("invalid y value: %v", err)
	}

	z, err := strconv.ParseFloat(strings.TrimSpace(parts[2]), 64)
	if err != nil {
		return Coordinate{}, fmt.Errorf("invalid z value: %v", err)
	}

	return Coordinate{x: x, y: y, z: z}, nil
}
