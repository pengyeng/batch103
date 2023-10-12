package batch103

import (
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type AWSUtils struct {
	region string
}

func (a *AWSUtils) SetRegion(Region string) {
	a.region = Region
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

func (a *AWSUtils) DownloadFileFromS3Bucket(bucket string, fileName string) error {
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: "default",
		Config: aws.Config{
			Region: aws.String(a.region),
		},
	})
	if err != nil {
		return err
	}
	bucketName := bucket
	downloader := s3manager.NewDownloader(sess)
	err = downloadFileFromS3(downloader, bucketName, fileName)
	return err

}

func (a *AWSUtils) UploadFileToS3Bucket(bucket string, fileName string) error {
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: "default",
		Config: aws.Config{
			Region: aws.String(a.region),
		},
	})
	if err != nil {
		return err
	}

	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	fileSize := fileInfo.Size()

	s3Svc := s3.New(sess)

	input := &s3.PutObjectInput{
		Body:          file,
		Bucket:        aws.String(bucket),
		Key:           aws.String(filepath.Base(fileName)),
		ContentLength: aws.Int64(fileSize),
	}

	_, err = s3Svc.PutObject(input)
	if err != nil {
		return err
	}

	return nil
}
