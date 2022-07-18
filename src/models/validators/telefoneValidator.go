package validators

import(
	"strings"
	"strconv"
)



func NumeroValido(telefone string) (int, bool){
	var telefoneGerado int
	telefoneGerado, _ = strconv.Atoi(Format(telefone)) 

	if !TamanhoNumero(telefone){
		return 0, false
	}else{
		return telefoneGerado, true
	}
}

func TamanhoNumero(telefone string) bool{
	if len(Format(telefone)) != 11 {
		return false
	} else {
		return true
	}
}

func Format(telefone string) string{

	telefone = strings.Replace(telefone, "(", "", -1)
	telefone = strings.Replace(telefone, ")", "", -1)
	telefone = strings.Replace(telefone, " ", "", -1)
	telefone = strings.Replace(telefone, "-", "", -1)

	return telefone
}