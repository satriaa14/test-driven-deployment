package models

import (
	"errors"
)

type Nilai struct {
	A    float64 `json:"a"`
	B    float64 `json:"b"`
	Tipe string  `json:"tipe"`
}

func (n *Nilai) Perhitungan() (float64, error) {

	switch n.Tipe {
	case "addition":
		return n.A + n.B, nil
	case "substraction":
		return n.A - n.B, nil
	case "multiplier":
		return n.A * n.B, nil
	case "division":
		return n.A / n.B, nil
	default:
		return 0, errors.New("Tipe tidak dikenal")
	}
}
