package procedure

type Addition struct{
	Num1, Num2 float64
}

func (a Addition) DoCalc() float64 {
	return a.Num1+a.Num2
}