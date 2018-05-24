package main

import (
	// "bufio"
	// "encoding/csv"
	"fmt"
	// "io"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/endpoints"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/s3manager"
	"log"
	"os"

	"github.com/qezz/xproject-reports/internal/report"
)

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

func main() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		log.Fatalln("unable to load SDK config:", err)
	}

	cfg.Region = endpoints.UsEast1RegionID

	s3path, err := report.FindReport(&cfg, "test-quicksight")

	fmt.Println("file s3 path:", s3path)
	return
	// svc := s3.New(cfg)
	downloader := s3manager.NewDownloader(cfg)

	filename := "test-output"
	f, err := os.Create(filename)
	if err != nil {
		fmt.Errorf("failed to create file %q, %v", filename, err)
		return
	}

	n, err := downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String("qezz-xproject-test-bucket-91739124"),
		Key:    aws.String(s3path),
	})

	if err != nil {
		fmt.Errorf("failed to download file, %v", err)
		return
	}

	fmt.Printf("file downloaded, %d bytes\n", n)
}
