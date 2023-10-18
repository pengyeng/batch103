# batch103
Simple SpringBatch like Batch Framework for Go

batch103 consists of the following components:
- JobLauncher
- Reader
- Processor
- Writer

# Job Launcher #
Job Launcher is main skeleton of the batch framework. Job Launcher offers one method Run(r ReaderType, p []ProcessorType, w []WriterType); which take in Reader, one/multiple Processor(s) and one/multiple Writer(s).

# Reader # 
Reader is the component that responsible to receive data from any sources (e.g. File, Database) that you can implement. You can implement your custom reader by conforming to the Reader interface.

Below is a sample implementation of Reader (sampleReader.go):
```
type SampleReader struct {}

func (r *DelimitedFileReader) Read() ([]batch103.BatchData, error) {
    var result []batch103.BatchData
    input := []string{"Field1", "Field2", "Field3", "Field4"}
    var batchData = &batch103.BatchData{}
	batchData = batchData.Create(input)
	result = append(result, *batchData)
    
    return result, nil
}    
```

Alternatively, you can also use standard file reader which has already been implemented. Below are the standard readers avaiable:
- FixedWidthFileReader 
- DelimitedFileReader

Below is a sample implementation of using standard file reader in Job Launcher:
```
func main() {
    var myJobLauncher batch103.JobLauncher
    var myReader = &batch103.DelimitedFileReader{}
    myReader.FileReader.SetFileName("output.csv")
    myJobLauncher.Run(myReader, myProcessorList, myWriterList)
}    
```

# Processor # 
Processor is the component that responsible to process data received from Reader. Slice of BatchData return from Reader will be passed as a parameter to Processor. You may pass an empty slice to Job Launcher if you don't need a Processor. If Processor is required, You need to implement your custom processor by conforming to the Processor interface. 

Below is a sample implementation of Reader (sampleProcessor.go):
```
type SampleProcessor struct {
	batch103.BaseProcessor
}

func (r *SampleProcessor) Process(data []batch103.BatchData) ([]batch103.BatchData, error) {
	var processRecords []batch103.BatchData
	for i := 0; i < len(data); i++ {
		if data[i].IsActive() {
			var record = data[i].GenericData
			if record[1] == "TREASURE" {
				log.Println("TREASURE FOUND : ", record[0])
			} else {
				data[i].Reject(batch103.StgProcess)
			}
			processRecords = append(processRecords, data[i])
		}
	}
	return processRecords, nil
}
```
# Writer # 
Writer is the component that responsible to write the processed data received from Processor to database or files. Slice of BatchData return from Reader will be passed as a parameter to Writer. You may pass an empty slice to Job Launcher if you don't need a Writer. If Writer is required, You need to implement your custom writer by conforming to the Writer interface. 

Below is a sample implementation of Writer (sampleWriter.go):
```
type SampleWriter struct {
	batch103.BaseWriter
}

func (w *SampleWriter) Write(data []batch103.BatchData) error {

	var concatenatedContent string
	for i := 0; i < len(data); i++ {
		if data[i].IsActive() {
			var record = data[i].GenericData
			concatenatedContent = concatenatedContent + record[0] + "," + record[1] + "," + record[2] + "\n"
		}
	}
	fileContent := []byte(concatenatedContent)
	err := os.WriteFile("output.csv", fileContent, 0644)
	if err != nil {
		return err
	}
	return nil
}
```