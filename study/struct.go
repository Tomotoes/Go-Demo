package main

import (
	"fmt"
	"math"
)

type Vertext struct {
	X, Y float64
}

func (v Vertext) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertext, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertext) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Abs(v Vertext) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertext{3, 4}
	v.Scale(10) // GO 编译器会把 v.Scale 解释成 (&v).Scale
	fmt.Println(v)
	Scale(&v, 10)
	fmt.Println(v)
	fmt.Println(Abs(v))

	p := &Vertext{3, 4}
	p.Scale(10)
	fmt.Println(p)
	Scale(p, 10)
	fmt.Println(p)
	fmt.Println(Abs(*p))

//	结论: 在结构体成员方法中 , 会隐式转换类型
//       在普通方法中, 必须类型一致 , 一切以实参为主

}
