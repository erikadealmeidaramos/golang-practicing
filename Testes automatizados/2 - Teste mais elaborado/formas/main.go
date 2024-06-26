package formas

import (
	"math"
)

type Retangulo struct {
	Altura  float64
	Largura float64
}

type Circulo struct {
	Raio float64
}

type Forma interface {
	area() float64
}

func (r Retangulo) area() float64 {
	return r.Altura * r.Largura
}

func (c Circulo) area() float64 {
	return math.Pow(c.Raio, 2) * math.Pi
}
