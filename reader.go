package batch103

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/service/s3"		
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

func downloadFileFromS3(downloader *s3manager.Downloader, bucketName string, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = downloader.Download(
		file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(fileName),
		},
	)

	return err
}

func (f *FileReader) DownloadFileFromS3Bucket(region string, bucket string) error {
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: "default",
		Config: aws.Config{
			Region: aws.String(region),
		},
	})
	if err != nil {
		return err
	}
	fileName := f.filename
	bucketName := bucket
	downloader := s3manager.NewDownloader(sess)
	err = downloadFileFromS3(downloader, bucketName, fileName)
	return err
}

func (f *FileReader) OpenCSVFile() (csv.Reader, error) {

	csvFile, err := os.Open(f.filename)
	csvReader := csv.NewReader(csvFile)

	if err != nil {
		log.Fatalln("Couldn't open csv file", err)
	}

	return *csvReader, err
}
