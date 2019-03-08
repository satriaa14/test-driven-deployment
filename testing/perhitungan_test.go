package test

import (
	"fmt"
	"testing"

	"github.com/satriaa14/test-driven-deployment/models"
)

var dataPerhitungan = []struct {
	inputPerhitungan models.Nilai
	hasilMaunya      float64
}{
	{inputPerhitungan: models.Nilai{A: 1, B: 1, Tipe: "addition"},
		hasilMaunya: 2},
	{inputPerhitungan: models.Nilai{A: 1, B: 1, Tipe: "substraction"},
		hasilMaunya: 0},
	{inputPerhitungan: models.Nilai{A: 1, B: 1, Tipe: "multiplier"},
		hasilMaunya: 1},
	{inputPerhitungan: models.Nilai{A: 1, B: 1, Tipe: "division"},
		hasilMaunya: 1},
	{inputPerhitungan: models.Nilai{A: 1, B: 1, Tipe: "etc"},
		hasilMaunya: 0},

	{inputPerhitungan: models.Nilai{A: 5, B: 2, Tipe: "addition"},
		hasilMaunya: 7},
	{inputPerhitungan: models.Nilai{A: 5, B: 2, Tipe: "substraction"},
		hasilMaunya: 3},
	{inputPerhitungan: models.Nilai{A: 5, B: 2, Tipe: "multiplier"},
		hasilMaunya: 10},
	{inputPerhitungan: models.Nilai{A: 5, B: 2, Tipe: "division"},
		hasilMaunya: 2.5},
	{inputPerhitungan: models.Nilai{A: 5, B: 2, Tipe: "etc"},
		hasilMaunya: 0},
}

func TestPerhitungan(t *testing.T) {
	t.Run("Test untuk Fungsi Perhitungan", func(t *testing.T) {

		for _, input := range dataPerhitungan {
			hasilDapatnya, theError := input.inputPerhitungan.Perhitungan()
			if hasilDapatnya != input.hasilMaunya {
				t.Fatalf("%v %v %v = Got : %v, Want : %v, Error : %v\n",
					input.inputPerhitungan.A,
					input.inputPerhitungan.Tipe,
					input.inputPerhitungan.B,
					hasilDapatnya,
					input.hasilMaunya, theError)
			}
			fmt.Printf("%v %v %v = Got : %v, Want : %v, Error : %v\n",
				input.inputPerhitungan.A,
				input.inputPerhitungan.Tipe,
				input.inputPerhitungan.B,
				hasilDapatnya,
				input.hasilMaunya, theError)
		}

	})
}
