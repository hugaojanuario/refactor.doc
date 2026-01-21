package converter

import (
	"os"

	"github.com/nguyenthenguyen/docx"
)

func TxtToDocx(inputPath string, outputPath string) error {
	text, err := os.ReadFile(inputPath)
	if err != nil {
		return err
	}

	// abre o template
	template, err := docx.ReadDocxFile("template/template.docx")
	if err != nil {
		return err
	}
	defer template.Close()

	doc := template.Editable()

	// substitui o placeholder
	doc.Replace("{{CONTENT}}", string(text), -1)

	return doc.WriteToFile(outputPath)
}
