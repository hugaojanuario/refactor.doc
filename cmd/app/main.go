package main

import (
	"fmt"
	"log"

	"github.com/hugaojanuario/refactor.doc/internal/converter"
)

func main() {
	fmt.Println("Iniciando conversao de arquivos...")

	err := converter.ConvertTxtToDocx(
		"input/exemplo.txt",
		"output/exemplo.docx",
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conversao finalizada com sucesso!")
}
