package converter

import (
	"os"
)

func TxtToRTF(inputPath string, outputPath string) error {
	content, err := os.ReadFile(inputPath)
	if err != nil {
		return err
	}

	rtfContent := "{\\rtf1\\ansi\n"
	rtfContent += string(content)
	rtfContent += "\n}"

	return os.WriteFile(outputPath, []byte(rtfContent), 0644)
}
