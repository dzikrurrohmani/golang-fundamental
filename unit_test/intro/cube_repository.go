package intro

import (
	"errors"
	"math"
)

type CubeRepository interface {
	Volume() (float64, error)
	Area() (float64, error)
}

type cubeRepository struct {
	cube Cube
}

func (c *cubeRepository) Volume() (float64, error) {
	if c.cube.Side < 0 {
		return 0.0, errors.New("tidak boleh negatif")
	}
	return math.Pow(c.cube.Side, 3), nil
}

func (c *cubeRepository) Area() (float64, error) {
	if c.cube.Side < 0 {
		return 0.0, errors.New("tidak boleh negatif")
	}
	return math.Pow(c.cube.Side, 2) * 6, nil
}

func NewCubeRepository(cube Cube) CubeRepository {
	repo := new(cubeRepository)
	repo.cube = cube
	return repo
}