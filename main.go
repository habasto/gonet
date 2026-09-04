package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("gonet - Network Toolkit")

	if len(os.Args) < 2 {
		fmt.Println("Uso: gonet <comando> [argumentos]")
		return
	}

	comando := os.Args[1]

	switch comando {

	case "ping": 
		fmt.Println("Ejecutando ping")

		if len(os.Args) < 3 {
			fmt.Println("Falta destino")
			return
		}

                	destino := os.Args[2]
                	fmt.Println(comando, "a:", destino)

	case "dns":
                fmt.Println("Ejecutando consulta DNS")

	case "tcp":
                fmt.Println("Ejecutando prueba TCP")

	default:
		fmt.Println("Comando desconocido:", comando)
	}

}
