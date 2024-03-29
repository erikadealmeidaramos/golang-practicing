package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retângulo", func(t *testing.T) {
		r := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := r.area()
		if areaRecebida != areaEsperada {

			/*
				t.Fatal() é equivalente a chamar t.Error() seguido por os.Exit(1).a
				Quando chamado, ele registra o teste como falhado e interrompe a execução do teste.
				Difere de t.Error() pois, com este, o resto dos testes continuam a ser executados.
			*/
			t.Fatalf("A área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
	})

	t.Run("Círculo", func(t *testing.T) {
		c := Circulo{10}
		areaEsperada := float64(math.Pi * 100)
		areaRecebida := c.area()
		if areaRecebida != areaEsperada {
			t.Fatalf("A área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
	})
}
