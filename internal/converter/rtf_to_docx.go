package converter

import (
	"os"

	"github.com/unidoc/unioffice/document"
)

func RTFToDocx(inputPath string, outputPath string) error {
	content, err := os.ReadFile(inputPath)
	if err != nil {
		return err
	}

	doc := document.New()
	para := doc.AddParagraph()
	para.AddRun().AddText(string(content))

	return doc.SaveToFile(outputPath)
}
