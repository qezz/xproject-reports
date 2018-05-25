package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"
)

func main() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		log.Fatalln("unable to load SDK config:", err)
	}

	sc := servicecatalog.New(cfg)
	input := servicecatalog.ScanProvisionedProductsInput{}
	req := sc.ScanProvisionedProductsRequest(&input)
	resp, err := req.Send()
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(resp)
}
