package models

import (
	cur "github.com/aws/aws-sdk-go-v2/service/costandusagereportservice"
)

type CurInfo struct {
	Def          *cur.ReportDefinition
	ManifestPath string
	ReportPath   string
}
