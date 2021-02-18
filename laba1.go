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
func SetArr(n int) []Realiz {
	var array []Realiz

	for i := 0; i < n; i++ {
		r := Realiz{}
		r.Set()
		array = append(array, r)
	}
	return array
}
func main() {
	var arr []Realiz
	arr = SetArr(10)
	fmt.Print(arr)
}
