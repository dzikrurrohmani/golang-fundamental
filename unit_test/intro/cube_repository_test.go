package intro

import "testing"

func TestCubeRepositoryVolume(t *testing.T) {
	t.Run("When the input is correct", func(t *testing.T) {
		cubeRepo := NewCubeRepository(Cube{1})
		// cubeService := NewCubeService(cubeRepo)

		expected := float64(1)
		actual, err := cubeRepo.Volume()
		// actual, err := cubeService.GetCubeVolume()

		if expected != actual {
			t.Errorf("Seharusnya %.2f", actual)
		}

		if err != nil {
			t.Error("Volume result error")
		}
	})

	t.Run("When the input is incorrect", func(t *testing.T) {
		cubeRepo := NewCubeRepository(Cube{-1})
		// cubeService := NewCubeService(cubeRepo)

		actual, err := cubeRepo.Volume()
		// actual, err := cubeService.GetCubeVolume()

		if err == nil {
			t.Error("Seharusnya error")
		}

		if actual != 0 {
			t.Error("Seharusnya keluaran 0")
		}
	})
}

func TestCubeRepositoryArea(t *testing.T) {
	t.Run("When the input is correct", func(t *testing.T) {
		cubeRepo := NewCubeRepository(Cube{1})
		// cubeService := NewCubeService(cubeRepo)

		expected := float64(6)
		actual, err := cubeRepo.Area()
		// actual, err := cubeService.GetCubeArea()

		if expected != actual {
			t.Errorf("Seharusnya %.2f", actual)
		}

		if err != nil {
			t.Error("Volume result error")
		}
	})

	t.Run("When the input is incorrect", func(t *testing.T) {
		cubeRepo := NewCubeRepository(Cube{-1})
		// cubeService := NewCubeService(cubeRepo)

		actual, err := cubeRepo.Area()
		// actual, err := cubeService.GetCubeArea()

		if err == nil {
			t.Error("Seharusnya error")
		}

		if actual != 0 {
			t.Error("Seharusnya keluaran 0")
		}
	})

}
