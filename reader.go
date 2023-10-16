package main

import (
	"encoding/csv"
	"log"
	"os"
)

type ReaderType interface {
	Read() ([]BatchData, error)
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

// Method under Base Reader
func (b *BaseReader) SetParameters(values map[string]string) {
	b.parameters = values
}

/* Parameters Getter */
func (b BaseReader) Parameters() map[string]string {
	return b.parameters
}

// Method under File Reader
func (f *FileReader) SetFileName(value string) {
	f.filename = value
}

func (f *FileReader) GetFileName() string {
	return f.filename
}

func (f *FileReader) OpenCSVFile() (csv.Reader, error) {

	csvFile, err := os.Open(f.filename)
	csvReader := csv.NewReader(csvFile)

	if err != nil {
		log.Fatalln("Couldn't open csv file", err)
	}

	return *csvReader, err
}
