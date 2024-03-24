package main

import (
	"fmt"
	"math"
)

type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("A área da forma é %0.2f\n", f.area())
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func (c circulo) area() float64 {
	return math.Pow(c.raio, 2) * math.Pi
}

func main() {
	r := retangulo{10, 15}
	c := circulo{5}

	escreverArea(r)
	escreverArea(c)
}
