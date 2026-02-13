package main

import (
	"gitbhub.com/basthianotto/godesde0/goroutines"
)

func main() {
	/*
		numeroInt, numeroString := ejercicios.DevuelveValores("99")
		fmt.Println(strconv.Itoa(numeroInt) + " " + numeroString)

		numeroInt, numeroString = ejercicios.DevuelveValores("101")
		fmt.Println(strconv.Itoa(numeroInt) + " " + numeroString)

		ejercicios.TablaMultiplicacion()*/
	//files.SumaTabla()
	//files.LeoArchivo()
	//funciones.Calculos()
	//funciones.LlamarClousure()
	//arreglos_slices.MuestroSlice()
	//mapas.MostrarMapas()
	//users.AltaUsuario()
	/*Pedro := new(modelos.Hombre)
	  Julia := new(modelos.Mujer)
	  ejer_interfaces.HumanosRespirando(Pedro)
	  ejer_interfaces.HumanosRespirando(Julia)
	*/
	//defer_panic.VemosDefer()
	//defer_panic.EjemploPanic()
	canal1 := make(chan bool)
	go goroutines.MiNombreLentooo("Basthian Duarte", canal1)
	defer func() {
		<-canal1
	}()

	/*var x string
	fmt.Scanln(&x)*/
}
