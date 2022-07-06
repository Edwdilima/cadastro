package validators

import (
	"testing"
)

// TestNumeroValido verifica se o telefone é válido e retorna o valor inteiro do número
func TestNumeroValido(t *testing.T){

	telefone := "(11) 9 4002-8922"
	esperado := 11940028922
	validacao := true
	resultado, isValid := NumeroValido(telefone)

	if isValid != validacao {
		t.Errorf("Número %v inválido!", telefone)
	}
	if(resultado != esperado){
		t.Errorf("Esperado: %v, Obtido: %v!",esperado, resultado)
	}

	telefone2 := "119400289"
	esperado2 := 0
	validacao2 := false
	resultado2, isValid2 := NumeroValido(telefone2)

	if isValid2 != validacao2 {
		t.Errorf("Número %v inválido!", telefone2)
	}
	if(resultado2 != esperado2){
		t.Errorf("Esperado: %v, Obtido: %v!",esperado2, resultado2)
	}
}

// TestTamanhoaNumero verifica se o telefone tem 11 digitos  ex: 83986263622
func TestTamanhoNUmero(t *testing.T){

	telefone := "(11) 9 4002-8922"
	esperado := true
	resultado := TamanhoNumero(telefone)
	if resultado != esperado {
		t.Errorf("Telefone %v possui tamanho diferente! Esperado %v, recebido %v", telefone, esperado, TamanhoNumero(telefone))
	}

	telefone2 := "11940028922"
	esperado2 := true
	resultado2 := TamanhoNumero(telefone2)
	if resultado2 != esperado2 {
		t.Errorf("Telefone %v possui tamanho diferente! Esperado %v, recebido %v", telefone2, esperado2, TamanhoNumero(telefone2))
	}

	telefone3 := "839862636"
	esperado3 := false
	resultado3 := TamanhoNumero(telefone3)
	if resultado3 != esperado3 {
		t.Errorf("Telefone %v possui tamanho diferente! Esperado %v, recebido %v", telefone3, esperado3, TamanhoNumero(telefone3))
	}

	telefone4 := "(83)986-2636"
	esperado4 := false
	resultado4 := TamanhoNumero(telefone4)
	if resultado4 != esperado4 {
		t.Errorf("Telefone %v possui tamanho diferente! Esperado %v, recebido %v", telefone4, esperado4, TamanhoNumero(telefone4))
	}

}


// TestFormat remove os caracteres especiais
func TestFormat(t *testing.T){

	telefone := "(11) 9 4002-8922"
	esperado := "11940028922"
	resultado := Format(telefone)
	if resultado != esperado {
		t.Errorf("Não foi possivél formatar o telefone %v! Esperado %v, recebido %v", telefone, esperado, Format(telefone))
	}

	telefone2 := "11940028922"
	esperado2 := "11940028922"
	resultado2 := Format(telefone2)
	if resultado2 != esperado2 {
		t.Errorf("Não foi possivél formatar o telefone %v! Esperado %v, recebido %v", telefone, esperado, Format(telefone))
	}

	telefone3 := "(83)986263622"
	esperado3 := "83986263622"
	resultado3 := Format(telefone3)
	if resultado3 != esperado3 {
		t.Errorf("Não foi possivél formatar o telefone %v! Esperado %v, recebido %v", telefone, esperado, Format(telefone))
	}

	telefone4 := "8398626-3622"
	esperado4 := "83986263622"
	resultado4 := Format(telefone4)
	if resultado4 != esperado4 {
		t.Errorf("Não foi possivél formatar o telefone %v! Esperado %v, recebido %v", telefone, esperado, Format(telefone))
	}

}

