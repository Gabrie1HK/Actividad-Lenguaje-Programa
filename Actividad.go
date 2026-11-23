// Elaborado por Diego Pedra y Gabriel Espinoza
package main

import (
	"fmt"
	"math"
)

type Figura interface {
	Area() float64
	Perimetro() float64
	Nombre() string
}

type Rectangulo struct {
	largo float64
	ancho float64
}

func (r Rectangulo) Area() float64 {
	return r.largo * r.ancho
}

func (r Rectangulo) Perimetro() float64 {
	return 2 * (r.largo + r.ancho)
}

func (r Rectangulo) Nombre() string {
	return "Rectángulo"
}

type Circulo struct {
	radio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * c.radio * c.radio
}

func (c Circulo) Perimetro() float64 {
	return 2 * math.Pi * c.radio
}

func (c Circulo) Nombre() string {
	return "Círculo"
}

type Triangulo struct {
	a float64
	b float64
	c float64
}

func (t Triangulo) Perimetro() float64 {
	return t.a + t.b + t.c
}

func (t Triangulo) Area() float64 {
	s := t.Perimetro() / 2
	return math.Sqrt(s * (s - t.a) * (s - t.b) * (s - t.c))
}

func (t Triangulo) Nombre() string {
	return "Triángulo"
}

func ImprimirFigura(f Figura) {
	fmt.Printf("Figura: %s | Área: %.2f | Perímetro: %.2f\n", f.Nombre(), f.Area(), f.Perimetro())
}

func main() {
	rect := Rectangulo{largo: 3, ancho: 4}
	circ := Circulo{radio: 5}
	tri := Triangulo{a: 3, b: 4, c: 5}

	ImprimirFigura(rect)
	ImprimirFigura(circ)
	ImprimirFigura(tri)
}