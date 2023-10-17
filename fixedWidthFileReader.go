package batch103

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type FixedWidthFileReader struct {
	FileReader
	Configuration string
}

type Fields struct {
	Fields      []Field `json:"Fields"`
	TotalLength int     `json:"TotalLength"`
}

type Field struct {
	Begin int `json:"Begin"`
	End   int `json:"End"`
}

func (r *FixedWidthFileReader) Read() ([]BatchData, error) {

	var result []BatchData
	var myFileUtils = &FileUtils{}
	var fileContent, err = myFileUtils.OpenFixedWidthFile(r.FileReader.GetFileName())
	if err != nil {
		return result, err
	}
	var jsonFile, jsonErr = os.Open(r.Configuration)
	if jsonErr != nil {
		return result, jsonErr
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var fields Fields
	json.Unmarshal([]byte(byteValue), &fields)

	if err != nil {
		return result, err
	}
	defer fileContent.Close()

	buf := make([]byte, fields.TotalLength)
	var batchData = &BatchData{}

	for {
		n, err := fileContent.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println("Printing n ", n)
		if n > 0 {
			var rowData []string
			var readLine = string(buf[:n])
			log.Println(readLine)
			for i := 0; i < len(fields.Fields); i++ {
				log.Println("retrieving ", readLine[fields.Fields[i].Begin:fields.Fields[i].End])
				rowData = append(rowData, readLine[fields.Fields[i].Begin:fields.Fields[i].End])
			}
			batchData = batchData.Create(rowData)
			result = append(result, *batchData)
		}
	}

	return result, nil
}
