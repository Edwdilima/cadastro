package validators

import (
	"testing"
)

// TestTamanho verifica se o cpf tem 11 digitos
func TestTamanho(t *testing.T) {
	cpf := "933.764.280-85"
	esperado := true
	if Tamanho(cpf) != esperado {
		t.Errorf("CPF %v possui tamanho diferente! Esperado %v, recebido %v", cpf, esperado, Tamanho(cpf))
	}

	cpf2 := "933.764.280-8"
	esperado2 := false
	if Tamanho(cpf2) != esperado2 {
		t.Errorf("CPF %v possui tamanho diferente! Esperado %v, recebido %v", cpf2, esperado2, Tamanho(cpf2))
	}

	cpf3 := "2887836706"
	esperado3 := false
	if Tamanho(cpf3) != esperado3 {
		t.Errorf("CPF %v possui tamanho diferente! Esperado %v, recebido %v", cpf3, esperado3, Tamanho(cpf3))
	}

	cpf4 := "222.222.222-22"
	esperado4 := true
	if Tamanho(cpf4) != esperado4 {
		t.Errorf("CPF %v possui tamanho diferente! Esperado %v, recebido %v", cpf4, esperado4, Tamanho(cpf4))
	}
}

// TestFormated irá formatar o cpf
func TestFormated(t *testing.T) {
	cpf := "130.327.294-67"
	esperado := "13032729467"
	resultado := Formated(cpf)
	if resultado != esperado {
		t.Errorf("CPF inválido! Esperado %v, recebido %v", esperado, resultado)
	}

	cpf2 := "222.222.222-22"
	esperado2 := "22222222222"
	resultado2 := Formated(cpf2)
	if resultado2 != esperado2 {
		t.Errorf("CPF inválido! Esperado %v, recebido %v", esperado2, resultado2)
	}

	cpf3 := "490.785.580-08"
	esperado3 := "49078558008"
	resultado3 := Formated(cpf3)
	if resultado3 != esperado3 {
		t.Errorf("CPF inválido! Esperado %v, recebido %v", esperado3, resultado3)
	}

	cpf4 := "33333333333"
	esperado4 := "33333333333"
	resultado4 := Formated(cpf4)
	if resultado4 != esperado4 {
		t.Errorf("CPF inválido! Esperado %v, recebido %v", esperado4, resultado4)
	}
}

// TestVerificarFormatacao irá verificar o cpf convertido em um slice de int e em um int
func TestVerificarNumeros(t *testing.T) {

	cpf := "123.456.789-10"
	var sliceEsperado = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 0}

	resultado := Numbers(cpf)
	i := 0
	for i <= len(sliceEsperado)-1 {
		if resultado[i] != sliceEsperado[i] {
			t.Errorf("Esperado %v, recebido %v", sliceEsperado, resultado)
		}
		i++
	}

	cpf2 := "222.222.222-22"
	sliceEsperado2 := []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}

	resultado2 := Numbers(cpf2)
	for i <= len(sliceEsperado2)-1 {
		if resultado2[i] != sliceEsperado2[i] {
			t.Errorf("Esperado %v, recebido %v", sliceEsperado2, resultado2)
		}
		i++
	}

	cpf3 := "490.785.580-08"
	sliceEsperado3 := []int{4, 9, 0, 7, 8, 5, 8, 0, 8, 0, 8}

	resultado3 := Numbers(cpf3)
	for i <= len(sliceEsperado3)-1 {
		if resultado3[i] != sliceEsperado3[i] {
			t.Errorf("Esperado %v, recebido %v", sliceEsperado3, resultado3)
		}
		i++
	}

	cpf4 := "33333333333"
	sliceEsperado4 := []int{3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3}

	resultado4 := Numbers(cpf4)
	for i <= len(sliceEsperado4)-1 {
		if resultado4[i] != sliceEsperado4[i] {
			t.Errorf("Esperado %v, recebido %v", sliceEsperado4, resultado4)
		}
		i++
	}

}

// TestVerificaNumIguais irá checar se existe algum cpf com numeros iguais
func TestVerificaNumIguais(t *testing.T) {

	cpf := "222.222.222-22"
	esperado := true
	resultado := VerificarNumerosIguais(cpf)
	if resultado != esperado {
		t.Errorf("CPF possui numeros diferentes! Esperado %v, recebido %v", esperado, resultado)
	}

	cpf2 := "933.764.280-85"
	esperado2 := false
	resultado2 := VerificarNumerosIguais(cpf2)
	if resultado2 != esperado2 {
		t.Errorf("CPF possui numeros iguais! Esperado %v, recebido %v", esperado2, resultado2)
	}

	cpf3 := "490.785.580-08"
	esperado3 := false
	resultado3 := VerificarNumerosIguais(cpf3)
	if resultado3 != esperado3 {
		t.Errorf("CPF possui numeros iguais! Esperado %v, recebido %v", esperado3, resultado3)
	}

	cpf4 := "33333333333"
	esperado4 := true
	resultado4 := VerificarNumerosIguais(cpf4)
	if resultado4 != esperado4 {
		t.Errorf("CPF possui numeros iguais! Esperado %v, recebido %v", esperado4, resultado4)
	}

	cpf5 := "28878367060"
	esperado5 := false
	resultado5 := VerificarNumerosIguais(cpf5)
	if resultado5 != esperado5 {
		t.Errorf("CPF possui numeros iguais! Esperado %v, recebido %v", esperado5, resultado5)
	}

}

// TestVerificaDigitosVerificadores irá checar se os digitos verificadores são válidos
func TestVerificaDigitosVerificadores(t *testing.T) {
	cpf := "130.327.294-67"
	esperado := true
	resultado := VerificacaoPorDigito(Numbers(cpf))

	if resultado != esperado {
		t.Errorf("CPF %v possui digitos verificadores inválidos! Esperado %v, recebido %v", cpf, esperado, resultado)
	}

	cpf2 := "222.222.222-22"
	esperado2 := true
	resultado2 := VerificacaoPorDigito(Numbers(cpf2))

	if resultado2 != esperado2 {
		t.Errorf("CPF %v possui digitos verificadores inválidos! Esperado %v, recebido %v", cpf2, esperado2, resultado2)
	}

	cpf3 := "933.764.280-87"
	esperado3 := false
	resultado3 := VerificacaoPorDigito(Numbers(cpf3))

	if resultado3 != esperado3 {
		t.Errorf("CPF %v possui digitos verificadores inválidos! Esperado %v, recebido %v", cpf3, esperado3, resultado3)
	}

	cpf4 := "490.785.580-08"
	esperado4 := true
	resultado4 := VerificacaoPorDigito(Numbers(cpf4))

	if resultado4 != esperado4 {
		t.Errorf("CPF %v possui digitos verificadores inválidos! Esperado %v, recebido %v", cpf4, esperado4, resultado4)
	}

	cpf5 := "28878367060"
	esperado5 := true
	resultado5 := VerificacaoPorDigito(Numbers(cpf5))

	if resultado5 != esperado5 {
		t.Errorf("CPF %v possui digitos verificadores inválidos! Esperado %v, recebido %v", cpf5, esperado5, resultado5)
	}

	cpf6 := "33333333333"
	esperado6 := true
	resultado6 := VerificacaoPorDigito(Numbers(cpf6))

	if resultado6 != esperado6 {
		t.Errorf("CPF %v possui digitos verificadores inválidos! Esperado %v, recebido %v", cpf6, esperado6, resultado6)
	}

}

// TestVerificarCPF irá verificar se um cpf é válido
func TestVerificarCPF(t *testing.T) {

	cpf2 := "222.222.222-22"
	esperado2 := false
	resultado2, cpfInInt2 := VerificarCPF(cpf2)
	if resultado2 != esperado2 {
		t.Errorf("CPF %v inválido! Esperado %v, recebido %v", cpf2, esperado2, resultado2)
	}
	if cpfInInt2 != 0 {
		t.Errorf("CPF %v convertido incorretamente! Esperado %v, recebido %v", cpf2, 0, cpfInInt2)
	}

	cpf3 := "933.764.280-87"
	esperado3 := false
	resultado3, cpfInInt3 := VerificarCPF(cpf3)
	if resultado3 != esperado3 {
		t.Errorf("CPF %v inválido! Esperado %v, recebido %v", cpf3, esperado3, resultado3)
	}
	if cpfInInt3 != 0 {
		t.Errorf("CPF %v convertido incorretamente! Esperado %v, recebido %v", cpf3, 0, cpfInInt3)
	}

	cpf4 := "490.785.580-08"
	esperado4 := true
	resultado4, cpfInInt4 := VerificarCPF(cpf4)
	if resultado4 != esperado4 {
		t.Errorf("CPF %v inválido! Esperado %v, recebido %v", cpf4, esperado4, resultado4)
	}
	if cpfInInt4 != 49078558008 {
		t.Errorf("CPF %v convertido incorretamente! Esperado %v, recebido %v", cpf4, 49078558008, cpfInInt4)
	}

	cpf5 := "28878367060"
	esperado5 := true
	resultado5, cpfInInt5 := VerificarCPF(cpf5)
	if resultado5 != esperado5 {
		t.Errorf("CPF %v inválido! Esperado %v, recebido %v", cpf5, esperado5, resultado5)
	}
	if cpfInInt5 != 28878367060 {
		t.Errorf("CPF %v convertido incorretamente! Esperado %v, recebido %v", cpf5, 28878367060, cpfInInt5)
	}

	cpf6 := "33333333333"
	esperado6 := false
	resultado6, cpfInInt6 := VerificarCPF(cpf6)
	if resultado6 != esperado6 {
		t.Errorf("CPF %v inválido! Esperado %v, recebido %v", cpf6, esperado6, resultado6)
	}
	if cpfInInt6 != 0 {
		t.Errorf("CPF %v convertido incorretamente! Esperado %v, recebido %v", cpf6, 0, cpfInInt6)
	}

	cpf7 := "2887836706"
	esperado7 := false
	resultado7, cpfInInt7 := VerificarCPF(cpf7)
	if resultado7 != esperado7 {
		t.Errorf("CPF %v inválido! Esperado %v, recebido %v", cpf7, esperado7, resultado7)
	}
	if cpfInInt7 != 0 {
		t.Errorf("CPF %v convertido incorretamente! Esperado %v, recebido %v", cpf7, 0, cpfInInt7)
	}

}
