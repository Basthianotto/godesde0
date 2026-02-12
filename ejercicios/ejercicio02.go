package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TablaMultiplicacion() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese número:")
	var numeroTXT string
	var numeroINT int
	var err error
	for {

		if scanner.Scan() {
			numeroTXT = scanner.Text()
			numeroINT, err = strconv.Atoi(numeroTXT)
			if err != nil {
				fmt.Println("El dato ingresado es incorrecto :" + err.Error())
				fmt.Println("Ingrese nuevamente:")
				continue
			} else {
				var texto string
				for i := 1; i <= 10; i++ {
					multiplicado := numeroINT * i
					texto += fmt.Sprintln(numeroTXT + "*" + strconv.Itoa(i) + "=" + strconv.Itoa(multiplicado))
					//fmt.Println(numeroTXT + "*" + strconv.Itoa(i) + "=" + strconv.Itoa(multiplicado))
					//fmt.Println("Ingrese nuevamente:")

				}
				fmt.Println(texto)
				return texto
			}
		}

	}

}
