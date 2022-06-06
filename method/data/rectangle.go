package data

type Rectangle struct {
	Length, Width float32
}

func(r *Rectangle) Area() float64 {
	return float64(r.Length)*float64(r.Width)
}


func(r *Rectangle) Perimeter() float32 {
	return 2 * (r.Length+r.Width)
}