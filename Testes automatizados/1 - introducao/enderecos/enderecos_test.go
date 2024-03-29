package enderecos_test

import (
	/*
		Esse ponto (.) antes da string "introducao-testes/enderecos" é um import com alias (apelido).
		Isso significa que o pacote introducao-testes/enderecos será importado com o nome enderecos,
		ou seja, ao invés de usarmos introducao-testes/enderecos.TipoDeEndereco, usaremos enderecos.TipoDeEndereco.
	*/
	. "introducao-testes/enderecos"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {

	t.Parallel()

	cenarioDeTeste := []cenarioDeTeste{
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Praça das Rosas", "Tipo inválido"},
		{"Estrada da Ribeira", "Estrada"},
		{"Rua das Margaridas", "Rua"},
		{"AVENIDA REBOUÇAS", "Avenida"},
		{"", "Tipo inválido"},
	}

	for _, cenario := range cenarioDeTeste {
		tipoEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)

		if tipoEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s",
				cenario.retornoEsperado, tipoEnderecoRecebido)
		}
	}

}
