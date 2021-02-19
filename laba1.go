package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Realiz struct {
	R float64
	X float64
}

func (s *Realiz) Set() {
	k := math.Sqrt(8 / math.Pi)
	s.R = rand.Float64()
	s.X = 1 / k * math.Log((1+s.R)/(1-s.R))
}

//Заполнение массива
func SetArr(n int) []Realiz {
	var array []Realiz

	for i := 0; i < n; i++ {
		r := Realiz{}
		r.Set()
		array = append(array, r)
	}
	return array
}

// Получение оценки матиматического ожидания
func GetOcMatOj(arr []Realiz) float64 {
	s := 0.0
	for i := 0; i < len(arr); i++ {
		s += arr[i].R
	}
	s = 1.0 / float64(len(arr)) * s
	return s
}

// Получение оценки дисперсии
func GetOcDisp(arr []Realiz, mj float64) float64 {
	s := 0.0
	for i := 0; i < len(arr); i++ {
		s += math.Pow((arr[i].R - mj), 2.0)
	}
	s = 1.0 / float64(len(arr)) * s
	return s
}

func main() {
	var arr []Realiz
	arr = SetArr(10)
	mtj := GetOcMatOj(arr)
	disp := GetOcDisp(arr, mtj)
	otkl := math.Sqrt(disp)
	fmt.Print(arr)
	fmt.Printf("\nMatOj = %g", mtj)
	fmt.Printf("\nDisp = %g", disp)
	fmt.Printf("\nSrKvadOtkl = %g", otkl)
}
