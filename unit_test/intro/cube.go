package intro

import (
	"errors"
	"math"
)

type Cube struct {
	Side float64
}

func (c *Cube) Volume() (float64, error) {
	if c.Side < 0 {
		return 0.0, errors.New("tidak boleh negatif")
	}
	return math.Pow(c.Side, 3), nil
}

func (c *Cube) Area() (float64, error) {
	if c.Side < 0 {
		return 0.0, errors.New("tidak boleh negatif")
	}
	return math.Pow(c.Side, 2) * 6, nil
}
