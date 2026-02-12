package files

import (
	"fmt"
	"os"

	"gitbhub.com/basthianotto/godesde0/ejercicios"
)

func GrabaTabla() {
	texto := ejercicios.TablaMultiplicacion()
	archivo, err := os.Create("./files/txt/tabla.txt")
	if err != nil {
		fmt.Println("Error al crear el archivo " + err.Error())
		return
	}

	fmt.Fprint(archivo, texto)
	archivo.Close()
}
