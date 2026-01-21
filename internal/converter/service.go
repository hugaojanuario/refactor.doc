package converter

func ConvertTxtToDocx(inputTxt string, outputDocx string) error {
	rtfPath := "output/tempo.rtf"

	err := TxtToRTF(inputTxt, rtfPath)
	if err != nil {
		return err
	}

	err = RTFToDocx(rtfPath, outputDocx)
	if err != nil {
		return err
	}

	return nil
}
