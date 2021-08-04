package batch103

import (
	"encoding/csv"
	"log"
	"os"
)

type ReaderType interface {
	Read() []BatchData
	SetParameters(values map[string]string)
}

type BaseReader struct {
	parameters map[string]string
}

// File Reader inherits from Base Reader
type FileReader struct {
	BaseReader
	filename string
}

//Method under Base Reader
func (b *BaseReader) SetParameters(values map[string]string) {
	b.parameters = values
}

/* Parameters Getter */
func (b BaseReader) Parameters() map[string]string {
	return b.parameters
}

//Method under File Reader
func (f *FileReader) SetFileName(value string) {
	f.filename = value
}

func (f *FileReader) OpenCSVFile() csv.Reader {

	csvFile, err := os.Open(f.filename)

	if err != nil {
		log.Fatalln("Couldn't open csv file", err)
	}

	csvReader := csv.NewReader(csvFile)

	return *csvReader
}
