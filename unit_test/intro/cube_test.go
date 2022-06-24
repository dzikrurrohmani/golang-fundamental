// package intro

// import (
// 	"testing"
// )

// func TestVolumePositif(t *testing.T) {
// 	temp := Cube{}
// 	temp.Side = 1
// 	expected := 1.0
// 	actual, _ := temp.Volume()
// 	if expected != actual {
// 		t.Error("tidak boleh negatif")
// 	}
// }

// func TestVolumeNegatif(t *testing.T) {
// 	temp := Cube{}
// 	temp.Side = -1
// 	expected := 0.0
// 	actual, _ := temp.Volume()
// 	if expected != actual {
// 		t.Error("tidak boleh negatif")
// 	}
// }

// func TestAreaPositif(t *testing.T) {
// 	temp := Cube{}
// 	temp.Side = 1
// 	expected := 6.0
// 	actual, _ := temp.Area()
// 	if expected != actual {
// 		t.Error("tidak boleh negatif")
// 	}
// }

// func TestAreaNegatif(t *testing.T) {
// 	temp := Cube{}
// 	temp.Side = -1
// 	expected := 0.0
// 	actual, _ := temp.Area()
// 	if expected != actual {
// 		t.Error("tidak boleh negatif")
// 	}
// }

package intro

import "testing"

func TestVolumePositif(t *testing.T) {
	t.Run("When the input is correct", func(t *testing.T) {
		kubus := Cube{
			Side: 5,
		}

		expected := float64(125)
		actual, err := kubus.Volume()

		if expected != actual {
			t.Errorf("Seharusnya %.2f", actual)
		}

		if err != nil {
			t.Error("Volume result error")
		}
	})

	t.Run("When the input is incorrect", func(t *testing.T) {
		kubus := Cube{
			Side: -5,
		}

		// expected := 0.0
		actual, err := kubus.Volume()

		if err == nil {
			t.Error("tidak boleh negatif")
		}

		if actual < 0 {
			t.Error("tidak boleh negatif")
		}
	})
}

func TestAreaPositif(t *testing.T) {
	t.Run("When the input is correct", func(t *testing.T) {
		kubus := Cube{
			Side: 5,
		}

		expected := float64(150)
		actual, err := kubus.Area()

		if expected != actual {
			t.Errorf("Seharusnya %.2f", actual)
		}

		if err != nil {
			t.Error("Volume result error")
		}
	})

	t.Run("When the input is incorrect", func(t *testing.T) {
		kubus := Cube{
			Side: -5,
		}

		// expected := 0.0
		actual, err := kubus.Area()

		if err == nil {
			t.Error("tidak boleh negatif")
		}

		if actual < 0 {
			t.Error("tidak boleh negatif")
		}
	})

}

// func BenchmarkVolumeTest(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		// Volume()
// 	}
// }
