package ejercicios

import (
	"fmt"
	"strconv"
)

func DevuelveValores(valor string) (int, string) {
	var err error
	numero, err := strconv.Atoi(valor)
	var texto string
	if numero > 100 {
		texto = "Es mayor a 100"
	} else {
		texto = "Es menor a 100"
	}

	if err != nil {
		fmt.Println(err)
	}

	return numero, texto
}
