package report

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/endpoints"
	cur "github.com/aws/aws-sdk-go-v2/service/costandusagereportservice"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/s3manager"
)

type FindReportError struct {
	Inner      *error
	ReportName string
}

func (e FindReportError) Error() string {
	return fmt.Sprintf("error finding report %v: %v", e.ReportName, e.Inner)
}

func NewFindReportError(err *error, reportname string) FindReportError {
	return FindReportError{
		Inner:      err,
		ReportName: reportname,
	}
}

func FindReport(cfg *aws.Config, reportname string) (string, error) {
	cfg.Region = endpoints.UsEast1RegionID

	curclient := cur.New(*cfg)

	params := cur.DescribeReportDefinitionsInput{}
	req := curclient.DescribeReportDefinitionsRequest(&params)
	resp, err := req.Send()
	if err != nil {
		return "", NewFindReportError(&err, reportname)
	}

	fmt.Printf("%v\n", resp)
	// reps := utils.GetProperty(resp, "ReportDefinitions").([]interface{})

	for _, repdef := range resp.ReportDefinitions {
		fmt.Println(repdef)
		if reportname == *repdef.ReportName {
			return getMostRecentReportPath(cfg, repdef), nil
		}
	}

	e := errors.New("Report not found")
	return "", NewFindReportError(&e, reportname) // FIXME: fix this error with good one
}

// FIXME: return errors
func getMostRecentReportPath(cfg *aws.Config, rd cur.ReportDefinition) string {
	timeInterval := "20180501-20180601" // FIXME: do not hardcode this, calculate in runtime

	manifestPath := // *rd.S3Bucket + "/" +
		*rd.S3Prefix + "/" +
			*rd.ReportName + "/" + timeInterval + "/" + *rd.ReportName + "-Manifest.json"

	cfg.Region = string(rd.S3Region)

	// get manifest
	downloader := s3manager.NewDownloader(*cfg)

	filename := "tmp-manifest" // FIXME: use tempfile, or do not save it in fs at all

	f, err := os.Create(filename)
	// FIXME: return errors
	if err != nil {
		log.Printf("failed to create file %q, %v", filename, err)
		return ""
	}

	_, err = downloader.Download(f, &s3.GetObjectInput{
		Bucket: rd.S3Bucket,
		Key:    aws.String(manifestPath),
	})
	fmt.Println("s3 manifest:", manifestPath)
	// FIXME: return errors
	if err != nil {
		log.Printf("failed to download file, %v", err)
		return ""
	}

	// read field
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("failed to read file:", err)
	}

	// FIXME: rewrite this
	var ff interface{}
	err = json.Unmarshal(dat, &ff)

	zz := ff.(map[string]interface{})
	yy := zz["reportKeys"].([]interface{})
	ss := yy[0].(string)

	return ss
}
