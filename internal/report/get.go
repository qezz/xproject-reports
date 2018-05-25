package report

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/qezz/xproject-reports/internal/models"

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

func FindReport(cfg *aws.Config, reportname string) (models.CurInfo, error) {
	cfg.Region = endpoints.UsEast1RegionID

	curclient := cur.New(*cfg)

	params := cur.DescribeReportDefinitionsInput{}
	req := curclient.DescribeReportDefinitionsRequest(&params)
	resp, err := req.Send()
	if err != nil {
		return models.CurInfo{}, NewFindReportError(&err, reportname)
	}

	log.Printf("All reports:\n%v\n", resp)

	for _, repdef := range resp.ReportDefinitions {
		if reportname == *repdef.ReportName {
			mp, rp := getMostRecentReportPath(cfg, repdef)

			return models.CurInfo{
				Def:          &repdef,
				ManifestPath: mp,
				ReportPath:   rp,
			}, nil
		}
	}

	e := errors.New("Report not found")
	return models.CurInfo{}, NewFindReportError(&e, reportname) // FIXME: fix this error with good one
}

type ManifestPath = string
type ReportPath = string

// FIXME: Fix return types, add possible error
func getMostRecentReportPath(cfg *aws.Config, rd cur.ReportDefinition) (ManifestPath, ReportPath) {
	timeInterval := "20180501-20180601" // FIXME: do not hardcode this, calculate in runtime

	manifestPath := *rd.S3Prefix + "/" +
		*rd.ReportName + "/" + timeInterval + "/" + *rd.ReportName + "-Manifest.json"

	cfg.Region = string(rd.S3Region)

	// get manifest
	downloader := s3manager.NewDownloader(*cfg)

	filename := "tmp-manifest" // FIXME: use tempfile, or do not save it in fs at all

	f, err := os.Create(filename)

	if err != nil {
		log.Printf("failed to create file %q, %v", filename, err)
		// FIXME: return errors
		return "", ""
	}

	_, err = downloader.Download(f, &s3.GetObjectInput{
		Bucket: rd.S3Bucket,
		Key:    aws.String(manifestPath),
	})

	if err != nil {
		log.Printf("failed to download file, %v", err)
		// FIXME: return errors
		return "", ""
	}

	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("failed to read file:", err)
	}

	// FIXME: rewrite this
	var ff interface{}
	err = json.Unmarshal(dat, &ff)

	zz := ff.(map[string]interface{})
	yy := zz["reportKeys"].([]interface{}) // NOTE: I don't know why does this return an array, since report name is unique
	ss := yy[0].(string)

	return manifestPath, ss
}
