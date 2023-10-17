package batch103

import (
	"io"
	"log"
)

type DelimitedFileReader struct {
	FileReader
}

func (r *DelimitedFileReader) Read() ([]BatchData, error) {

	var result []BatchData

	// Checking Input File and Configuration File
	if r.FileReader.GetFileName() == "" {
		var inputError = &InputError{}
		return result, inputError
	}

	csvFileReader, err := r.OpenCSVFile()
	if err != nil {
		return result, err
	}

	for {

		record, err := csvFileReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		//Inserting File Record into Result set
		var rowData []string
		for i := 0; i < len(record); i++ {
			rowData = append(rowData, record[i])
		}
		var batchData = &BatchData{}
		batchData = batchData.Create(rowData)
		result = append(result, *batchData)
	}

	return result, nil
}
