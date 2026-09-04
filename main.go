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

	fmt.Println("Comando:", comando)

	if len(os.Args) >=3 {
		destino := os.Args[2]
		fmt.Println("Destino:", destino)
	}
}
