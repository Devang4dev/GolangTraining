package calc

import (
	"fmt"
	"math"
)

func Add(num1, num2 float32) {

	fmt.Println("addition =", num1+num2)

}

func Substraction(num1, num2 int) {

	fmt.Println("substraction =", num1-num2)
}

func Multiplication(num1, num2 int) {

	fmt.Println("multiplication =", num1*num2)
}

func Division(num1, num2 int) {

	fmt.Println("division =", num1/num2, "remin", num1%num2)

}

var x float64

func Squareroot(x float64) {

	fmt.Println("squareroot =", math.Sqrt(x))
}
