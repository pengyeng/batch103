// This is the skeleton of the Batch Framework

package batch103

import (
	"log"
	"reflect"
)

type JobLauncher struct {
	// Different between lower and upper case
	ParameterData map[string]string
	chunkSize     int
}

// Chunk size property
func (j *JobLauncher) SetChunkSize(Size int) *JobLauncher {
	j.chunkSize = Size
	return j
}

func (j *JobLauncher) Run(r ReaderType, p []ProcessorType, w []WriterType) {
	log.Println("Job Launcher starts")

	r.SetParameters(j.ParameterData)
	// Reader Execution
	log.Println("==== Reader : ", reflect.TypeOf(r), " Start ====")
	var result = r.Read()
	log.Println("==== Reader : ", reflect.TypeOf(r), " End ====")

	var chunkProcessing bool
	var totalRecords = 0

	if (j.chunkSize != 0 && len(result) != 0){
		chunkProcessing = true
		totalRecords = len(result)
	}

	if chunkProcessing {

		log.Println("Chunk Processing...")
		var fromNum = 0
		var toNum = 1
		if len(result) < j.chunkSize {
			toNum = totalRecords
		} else {
			toNum = j.chunkSize
		}

		var chunkResult []BatchData
		for toNum <= totalRecords && fromNum <= totalRecords {
			log.Println("FromNum ", fromNum, " ToNum ", toNum, "Chunk Size", j.chunkSize)

			if toNum <= totalRecords {
				chunkResult = result[fromNum:toNum]
			} else {
				chunkResult = result[:totalRecords]
			}

			//Start List of Processors Execution
			for i := 0; i < len(p); i++ {
				//Pass Parameters into Processor
				p[i].SetParameters(j.ParameterData)
				p[i].Process(chunkResult)
			}
			//End List of Processors Execution

			//Start List of Writer Execution
			for i := 0; i < len(w); i++ {
				//Pass Parameters into Writer
				w[i].SetParameters(j.ParameterData)
				w[i].Write(chunkResult)
			}
			//End List of Writer Execution
			if (toNum + j.chunkSize) <= totalRecords+1 {
				log.Println("Not Finish")
				fromNum = fromNum + j.chunkSize
				toNum = toNum + j.chunkSize
			} else {
				log.Println("Last Chunk")
				fromNum = fromNum + j.chunkSize
				toNum = totalRecords
			}

		}
	} else {

		//Start List of Processors Execution
		for i := 0; i < len(p); i++ {
			p[i].SetParameters(j.ParameterData)
			log.Println("==== Processor : ", reflect.TypeOf(p[i]), " Start ====")
			result = p[i].Process(result)
			log.Println("==== Processor : ", reflect.TypeOf(p[i]), " End ====")

		}
		//End List of Processors Execution

		//Start List of Writer Execution
		for i := 0; i < len(w); i++ {

			w[i].SetParameters(j.ParameterData)
			log.Println("=== Writer : ", reflect.TypeOf(w[i]), " Start ===")
			w[i].Write(result)
			log.Println("=== Writer : ", reflect.TypeOf(w[i]), " End ===")
		}
		//End List of Writer Execution
	}
	generateReport(result)

	log.Println("Job Launcher End")
}

func generateReport(result []BatchData) {

	//Prepare Summary Report
	var processRejectedCount = 0
	var writeRejectedCount = 0

	for i := 0; i < len(result); i++ {
		if !(result[i].IsActive()) {
			if result[i].Stage() == "PROCESS" {
				processRejectedCount = processRejectedCount + 1
			}
			if result[i].Stage() == "WRITE" {
				writeRejectedCount = writeRejectedCount + 1
			}

		}

	}
	log.Println("===== Batch Execution Report =====")
	log.Println("STAGE : READ  Rejected : 0  Processed :", len(result))
	log.Println("STAGE : PROC  Rejected :", processRejectedCount, " Processed :", len(result)-processRejectedCount)
	log.Println("STAGE : WRITE Rejected :", writeRejectedCount, " Processed :", len(result)-processRejectedCount-writeRejectedCount)
	log.Println("==================================")
}
