package main

import (
	"fmt"
	"strconv"

	"gitbhub.com/basthianotto/godesde0/ejercicios"
)

func main() {

	numeroInt, numeroString := ejercicios.DevuelveValores("99")
	fmt.Println(strconv.Itoa(numeroInt) + " " + numeroString)

	numeroInt, numeroString = ejercicios.DevuelveValores("101")
	fmt.Println(strconv.Itoa(numeroInt) + " " + numeroString)
}
