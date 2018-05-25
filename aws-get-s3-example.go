package main

import (
	// "bufio"
	// "encoding/csv"
	"fmt"
	// "io"
	"github.com/aws/aws-sdk-go-v2/aws"
	// "github.com/aws/aws-sdk-go-v2/aws/endpoints"
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
	if len(os.Args) < 2 {
		log.Println("Usage: go run aws-get-s3-example.go report-name")
		return
	}
	reportname := os.Args[1]

	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		log.Fatalln("unable to load SDK config:", err)
	}

	// Get report s3 path
	// Report name is unique so we can use it as the only parameter
	report, err := report.FindReport(&cfg, reportname)

	fmt.Println("S3 report path:", report.ReportPath)

	downloader := s3manager.NewDownloader(cfg)

	filename := "test-output"
	f, err := os.Create(filename)
	if err != nil {
		log.Printf("failed to create file %q, %v", filename, err)
		return
	}

	n, err := downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(*report.Def.S3Bucket),
		Key:    aws.String(report.ReportPath),
	})

	if err != nil {
		log.Printf("failed to download file, %v", err)
		return
	}

	fmt.Printf("file downloaded, %d bytes\n", n)
}
