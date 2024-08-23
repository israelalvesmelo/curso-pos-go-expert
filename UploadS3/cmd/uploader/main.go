package main

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	awsRegion = "us-east-1"
)

var (
	s3Client *s3.S3
	s3Bucket string
	wg       sync.WaitGroup
)

func init() {
	session, err := session.NewSession(
		&aws.Config{
			Region: aws.String(awsRegion),
			Credentials: credentials.NewStaticCredentials(
				"AKIA3Z3Z3Z3Z3Z3Z3Z3Z3",
				"Z3Z3Z3Z3Z3Z3Z3Z3Z3Z3Z3Z3Z3Z3Z3Z3Z3",
				"",
			),
		},
	)
	if err != nil {
		panic(err)
	}
	s3Client = s3.New(session)
	s3Bucket = "goexpert-bucket-example"
}

func main() {
	dir, err := os.Open("./tmp")
	if err != nil {
		panic(err)
	}
	defer dir.Close()
	uploadControl := make(chan struct{}, 100)
	errorFileUpload := make(chan string, 10)

	go func() {
		for {
			select {
			case file := <-errorFileUpload:
				fmt.Printf("Error uploading file %s. Retrying...\n", file)
				wg.Add(1)
				uploadControl <- struct{}{}
				go uploadFile(file, uploadControl, errorFileUpload)
			}
		}
	}()

	for {
		files, err := dir.Readdir(1) // lê o diretório
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("Error reading directory: %s\n", err)
			continue
		}
		wg.Add(1)
		uploadControl <- struct{}{}
		go uploadFile(files[0].Name(), uploadControl, errorFileUpload)
	}
	wg.Wait()
}

func uploadFile(fileName string, uploadControl <-chan struct{}, errorFileUpload chan<- string) {
	defer wg.Done()
	completeFileName := fmt.Sprintf("./tmp/%s", fileName)
	fmt.Printf("Uploading file %s to bucket\n", completeFileName)
	f, err := os.Open(completeFileName)
	if err != nil {
		fmt.Printf("Error opening file %s\n", completeFileName)
		<-uploadControl
		errorFileUpload <- fileName
		return
	}
	defer f.Close()
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(fileName),
		Body:   f,
	})

	if err != nil {
		fmt.Printf("Error uploading file %s\n", completeFileName)
		<-uploadControl
		errorFileUpload <- fileName

	}

	fmt.Printf("File %s uploaded successfully\n", completeFileName)
	<-uploadControl

}
