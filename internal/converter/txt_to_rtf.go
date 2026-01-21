package converter

import "os"

func TxtToRTF(inputPath string, outputPath string) error {
	text, err := os.ReadFile(inputPath)
	if err != nil {
		return err
	}

	rtf := "{\\rtf1\\ansi\n" + string(text) + "\n}"

	return os.WriteFile(outputPath, []byte(rtf), 0644)
}
