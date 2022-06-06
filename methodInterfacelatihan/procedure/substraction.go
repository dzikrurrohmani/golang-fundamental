package procedure

type Substraction struct{
	Num1, Num2 float64
}

func (s Substraction) DoCalc() float64 {
	return s.Num1-s.Num2
}