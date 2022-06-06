package data

type Rectangle struct {
	Length, Width float32
}

func(r *Rectangle) Area() float32 {
	return r.Length*r.Width
}


func(r *Rectangle) Perimeter() float32 {
	return 2 * (r.Length+r.Width)
}