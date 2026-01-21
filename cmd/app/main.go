package main

import (
	"fmt"
	"github.com/hugaojanuario/refactor.doc/internal/converter"
	"log"
)

func main() {
	fmt.Println("Iniciando conversao de arquivos...")

	err := converter.ConvertTxtToDocx(
		"input/exemplo.txt",
		"output/exemplo.docx")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conversão finalizada com sucesso!")
}
