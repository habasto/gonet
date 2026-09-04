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

	case "dns":
                fmt.Println("Ejecutando consulta DNS")

	case "tcp":
                fmt.Println("Ejecutando prueba TCP")

	default:
		fmt.Println("Comndo desconocido:", comando)
	}


	fmt.Println("Comando:", comando)

	if len(os.Args) >=3 {
		destino := os.Args[2]
		fmt.Println("Destino:", destino)
	}
}
