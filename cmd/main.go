package main

import "fmt"

func main() {
	funcion := func(y float64, t float64) float64 {
		return 5*y*t - 1
	}
	funcionSuperior := func(yDerivada float64, y float64, t float64) float64 {
		return yDerivada*t + y
	}
	y := 2.0
	h := .2
	t := 0.0
	fmt.Println("RungeKutta2")
	fmt.Println(RungeKutta2(funcion, y, h, t))
	fmt.Println("RungeKutta3")
	fmt.Println(RungeKutta3(funcion, y, h, t))
	fmt.Println("RungeKutta1_3Simpson")
	fmt.Println(RungeKutta1_3Simpson(funcion, y, h, t))
	fmt.Println("RungeKutta1_8Simpson")
	fmt.Println(RungeKutta1_8Simpson(funcion, y, h, t))
	fmt.Println("RungeKuttaSuperior")
	fmt.Println(RungeKuttaSuperior(funcionSuperior, 2, 1, .5, 0))
	fmt.Println("Euler_Modificado")
	fmt.Println(Euler_Modificado(func(y float64, t float64) float64 { return (5*y*t - 1) / 3 }, 1.2, 2, .2, 0))
}

func RungeKutta2(f func(float64, float64) float64, y float64, h float64, t float64) (float64, float64) {
	k1 := h * f(y, t)
	k2 := h * f(y+k1, t+h)
	fmt.Printf("k1: %v\nk2: %v\n", k1, k2)
	yn := y + (k1+k2)/2
	tn := t + h
	return yn, tn
}

func RungeKutta3(f func(float64, float64) float64, y float64, h float64, t float64) (float64, float64) {
	k1 := h * f(y, t)
	k2 := h * f(y+k1/2, t+h/2)
	k3 := h * f(y-k1+2*k2, t+h)
	fmt.Printf("k1: %v\nk2: %v\nk3: %v\n", k1, k2, k3)
	yn := y + (k1+4*k2+k3)/6
	tn := t + h
	return yn, tn
}

func RungeKutta1_3Simpson(f func(float64, float64) float64, y float64, h float64, t float64) (float64, float64) {
	k1 := h * f(y, t)
	k2 := h * f(y+k1/2, t+h/2)
	k3 := h * f(y+k2/2, t+h/2)
	k4 := h * f(y+k3, t+h)
	fmt.Printf("k1: %v\nk2: %v\nk3: %v\nk4: %v\n", k1, k2, k3, k4)
	yn := y + (k1 + 2*k2 + 2*k3 + k4)
	tn := t + h
	return yn, tn
}

func RungeKutta1_8Simpson(f func(float64, float64) float64, y float64, h float64, t float64) (float64, float64) {
	k1 := h * f(y, t)
	k2 := h * f(y+k1/3, t+h/3)
	k3 := h * f(y+k1/3+k2/3, t+h*(2/3))
	k4 := h * f(y+k1-k2+k3, t+h)
	yn := y + (k1+3*k2+3*k3+k4)/8
	tn := t + h
	return yn, tn
}

func RungeKuttaSuperior(f func(float64, float64, float64) float64, yDerivada float64, y float64, h float64, t float64) (float64, float64, float64) {
	V, U, q, a, b := yDerivada, y, t, 1.0, 1.0
	k1 := h * V
	m1 := h * f(a*V, b*U, q)
	k2 := h * (V + m1)
	m2 := h * f(a*(V+m1), b*(U+k1), q+h)
	fmt.Printf("k1: %v\nm1: %v\nk2: %v\nm2: %v\n", k1, m1, k2, m2)
	yn := U + (k1+k2)/2
	yDerivadaN := V + (m1+m2)/2
	tn := t + h
	return yDerivadaN, yn, tn
}

func Euler_Modificado(f func(float64, float64) float64, y0 float64, y1 float64, h float64, t float64) (float64, float64) {
	tn := t + h
	yDerivada := y0 + h/2*(f(y0, t)+f(y1, tn))
	return yDerivada, tn
}
