package converter

func ConvertTxtToDocx(inputTxt string, outputDocx string) error {
	// Gera RTF (opcional)
	err := TxtToRTF(inputTxt, "output/temp.rtf")
	if err != nil {
		return err
	}

	// Gera DOCX (principal)
	return TxtToDocx(inputTxt, outputDocx)
}
