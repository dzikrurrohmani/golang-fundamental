package procedure

import "fmt"

type Calculator struct {
	Tampil CalcOps
}

func (c Calculator) PrintResult() {
	fmt.Println(c.Tampil.DoCalc())
}