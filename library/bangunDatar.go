package library

import "math"

type Segitiga struct {
	Alas   int `json:"alas"`
	Tinggi int `json:"tinggi"`
}

type PersegiPanjang struct {
	Panjang int `json:"panjang"`
	Lebar   int `json:"lebar"`
}

type Persegi struct {
	Sisi int `json:"sisi"`
}

type JajarGenjang struct {
	Sisi   float64 `json:"sisi"`
	Alas   float64 `json:"alas"`
	Tinggi float64 `json:"tinggi"`
}

type Lingkaran struct {
	JariJari float64 `json:"jari"`
}

type HitungBangunDatar interface {
	Luas() int
	Keliling() int
}

func (s Segitiga) Luas() int {
	return (s.Alas * s.Tinggi) / 2
}

func (s Segitiga) Keliling() int {
	return s.Alas + s.Tinggi
}

func (p PersegiPanjang) Luas() int {
	return p.Panjang * p.Lebar
}

func (p PersegiPanjang) Keliling() int {
	return 2 * (p.Panjang + p.Lebar)
}

func (p Persegi) Luas() int {
	return int(math.Pow(float64(p.Sisi), 2))
}

func (p Persegi) Keliling() int {
	return 4 * p.Sisi
}

func (j JajarGenjang) Luas() int {
	return int(0.5 * j.Alas * j.Tinggi)
}

func (j JajarGenjang) Keliling() int {
	return int(2 * j.Sisi * j.Alas)
}

func (l Lingkaran) Luas() int {
	return int(math.Phi * math.Pow(l.JariJari, 2))
}

func (l Lingkaran) Keliling() int {
	return int(math.Phi * l.JariJari)
}
