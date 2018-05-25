# Report about AWS billing

## Important note ##

In this guide I will use Cost and Usage Reports service, because it is
recomended by AWS, due to fact that standard reports are going to be
deprecated and unavailable in future

## Using Cost and Usage Reports (CUR) ##

Cost and Usage Report is a CSV file containing data about your usage

### Enable CUR ###

Go to

1. AWS Console
2. Billing
3. "Reports" in a side bar
4. Enable Cost and Usage Reports if needed
5. Create report
   * Choose report name (from aws: "Report name must be unique ...",
     possibly unique for account) (will be referred to as
     `reportname`)
   * Select preferable time unit
   * Set checkmark against "include resource IDs"
   * Go "Next"
6. Provide S3 bucket (apply sample policy if needed)
   * Choose preferable prefix path (ref. `prefix`)
   * Go "Next"
7. Review your report settings
8. Optional: Wait for 24h until the service will be enable for you

Usualy, the report path will be:

`bucket/prefix/reportname/timeinterval/randomdir/reportname-N.csv.gz`

* `bucket` is your s3 bucket unique name
* `perfix` is the prefix path chosen at step 6
* `reportname` is the report name chosen at the step 5
* `randomdir` is usually smth like
  `abcdefgh-igkl-mnop-qrst-uvwxyz123456` except that letters and
  digits are different
* `timeinterval` for may, 2018 it is `20180501-20180601`
* `N` is usually "1" (I don't know can it even be not 1)

To get the most recent report you should parse the
`bucket/prefix/reportname/timeinterval/reportname-Manifest.json` file
and get the `assemblyId` field

<details><summary>JSON Manifest example</summary>

```json
{
  "assemblyId":"afe8b092-dc87-405a-a0b7-fc743e14adff",
  "account":"141375116029",
  "columns":[{
    "category":"identity",
    "name":"LineItemId"
  },{
    "category":"identity",
    "name":"TimeInterval"
  },{
    "category":"bill",
    "name":"InvoiceId"
  },{
    "category":"bill",
    "name":"BillingEntity"
  },{
    "category":"bill",
    "name":"BillType"
  },{
    "category":"bill",
    "name":"PayerAccountId"
  },{
    "category":"bill",
    "name":"BillingPeriodStartDate"
  },{
    "category":"bill",
    "name":"BillingPeriodEndDate"
  },{
    "category":"lineItem",
    "name":"UsageAccountId"
  },{
    "category":"lineItem",
    "name":"LineItemType"
  },{
    "category":"lineItem",
    "name":"UsageStartDate"
  },{
    "category":"lineItem",
    "name":"UsageEndDate"
  },{
    "category":"lineItem",
    "name":"ProductCode"
  },{
    "category":"lineItem",
    "name":"UsageType"
  },{
    "category":"lineItem",
    "name":"Operation"
  },{
    "category":"lineItem",
    "name":"AvailabilityZone"
  },{
    "category":"lineItem",
    "name":"ResourceId"
  },{
    "category":"lineItem",
    "name":"UsageAmount"
  },{
    "category":"lineItem",
    "name":"NormalizationFactor"
  },{
    "category":"lineItem",
    "name":"NormalizedUsageAmount"
  },{
    "category":"lineItem",
    "name":"CurrencyCode"
  },{
    "category":"lineItem",
    "name":"UnblendedRate"
  },{
    "category":"lineItem",
    "name":"UnblendedCost"
  },{
    "category":"lineItem",
    "name":"BlendedRate"
  },{
    "category":"lineItem",
    "name":"BlendedCost"
  },{
    "category":"lineItem",
    "name":"LineItemDescription"
  },{
    "category":"lineItem",
    "name":"TaxType"
  },{
    "category":"lineItem",
    "name":"LegalEntity"
  },{
    "category":"product",
    "name":"ProductName"
  },{
    "category":"product",
    "name":"alarmType"
  },{
    "category":"product",
    "name":"availability"
  },{
    "category":"product",
    "name":"capacitystatus"
  },{
    "category":"product",
    "name":"clockSpeed"
  },{
    "category":"product",
    "name":"currentGeneration"
  },{
    "category":"product",
    "name":"databaseEngine"
  },{
    "category":"product",
    "name":"deploymentOption"
  },{
    "category":"product",
    "name":"durability"
  },{
    "category":"product",
    "name":"ecu"
  },{
    "category":"product",
    "name":"engineCode"
  },{
    "category":"product",
    "name":"fromLocation"
  },{
    "category":"product",
    "name":"fromLocationType"
  },{
    "category":"product",
    "name":"group"
  },{
    "category":"product",
    "name":"groupDescription"
  },{
    "category":"product",
    "name":"instanceFamily"
  },{
    "category":"product",
    "name":"instanceType"
  },{
    "category":"product",
    "name":"instanceTypeFamily"
  },{
    "category":"product",
    "name":"licenseModel"
  },{
    "category":"product",
    "name":"location"
  },{
    "category":"product",
    "name":"locationType"
  },{
    "category":"product",
    "name":"maxIopsBurstPerformance"
  },{
    "category":"product",
    "name":"maxIopsvolume"
  },{
    "category":"product",
    "name":"maxThroughputvolume"
  },{
    "category":"product",
    "name":"maxVolumeSize"
  },{
    "category":"product",
    "name":"memory"
  },{
    "category":"product",
    "name":"minVolumeSize"
  },{
    "category":"product",
    "name":"networkPerformance"
  },{
    "category":"product",
    "name":"normalizationSizeFactor"
  },{
    "category":"product",
    "name":"operatingSystem"
  },{
    "category":"product",
    "name":"operation"
  },{
    "category":"product",
    "name":"physicalProcessor"
  },{
    "category":"product",
    "name":"preInstalledSw"
  },{
    "category":"product",
    "name":"processorArchitecture"
  },{
    "category":"product",
    "name":"processorFeatures"
  },{
    "category":"product",
    "name":"productFamily"
  },{
    "category":"product",
    "name":"region"
  },{
    "category":"product",
    "name":"requestType"
  },{
    "category":"product",
    "name":"servicecode"
  },{
    "category":"product",
    "name":"servicename"
  },{
    "category":"product",
    "name":"sku"
  },{
    "category":"product",
    "name":"storage"
  },{
    "category":"product",
    "name":"storageClass"
  },{
    "category":"product",
    "name":"storageMedia"
  },{
    "category":"product",
    "name":"tenancy"
  },{
    "category":"product",
    "name":"toLocation"
  },{
    "category":"product",
    "name":"toLocationType"
  },{
    "category":"product",
    "name":"transferType"
  },{
    "category":"product",
    "name":"usagetype"
  },{
    "category":"product",
    "name":"vcpu"
  },{
    "category":"product",
    "name":"volumeType"
  },{
    "category":"pricing",
    "name":"publicOnDemandCost"
  },{
    "category":"pricing",
    "name":"publicOnDemandRate"
  },{
    "category":"pricing",
    "name":"term"
  },{
    "category":"pricing",
    "name":"unit"
  },{
    "category":"reservation",
    "name":"AmortizedUpfrontCostForUsage"
  },{
    "category":"reservation",
    "name":"AmortizedUpfrontFeeForBillingPeriod"
  },{
    "category":"reservation",
    "name":"EffectiveCost"
  },{
    "category":"reservation",
    "name":"EndTime"
  },{
    "category":"reservation",
    "name":"ModificationStatus"
  },{
    "category":"reservation",
    "name":"NormalizedUnitsPerReservation"
  },{
    "category":"reservation",
    "name":"RecurringFeeForUsage"
  },{
    "category":"reservation",
    "name":"StartTime"
  },{
    "category":"reservation",
    "name":"TotalReservedNormalizedUnits"
  },{
    "category":"reservation",
    "name":"TotalReservedUnits"
  },{
    "category":"reservation",
    "name":"UnitsPerReservation"
  },{
    "category":"reservation",
    "name":"UnusedAmortizedUpfrontFeeForBillingPeriod"
  },{
    "category":"reservation",
    "name":"UnusedNormalizedUnitQuantity"
  },{
    "category":"reservation",
    "name":"UnusedQuantity"
  },{
    "category":"reservation",
    "name":"UnusedRecurringFee"
  },{
    "category":"reservation",
    "name":"UpfrontValue"
  }],
  "charset":"UTF-8",
  "compression":"GZIP",
  "contentType":"text/csv",
  "reportId":"1fbc20f088082669ee179c61999326b3b329325738afe67a98931c03b07cc613",
  "reportName":"test-quicksight",
  "billingPeriod":{
    "start":"20180501T000000.000Z",
    "end":"20180601T000000.000Z"
  },
  "bucket":"qezz-xproject-test-bucket-91739124",
  "reportKeys":["quicksight/test-quicksight/20180501-20180601/afe8b092-dc87-405a-a0b7-fc743e14adff/test-quicksight-1.csv.gz"],
  "additionalArtifactKeys":[{
    "artifactType":"RedshiftCommands",
    "name":"quicksight/test-quicksight/20180501-20180601/afe8b092-dc87-405a-a0b7-fc743e14adff/test-quicksight-RedshiftCommands.sql"
  },{
    "artifactType":"RedshiftManifest",
    "name":"quicksight/test-quicksight/20180501-20180601/afe8b092-dc87-405a-a0b7-fc743e14adff/test-quicksight-RedshiftManifest.json"
  }]
}
```

</details>

### How to get the most recent report file from S3 ###

1. Setup the SDK ([guide](https://github.com/pavlov-tony/xproject/issues/3#issuecomment-388046256)).
   Environment variables are preferable.
2. Get the Manifest
3. Get the report

See example: [aws-get-s3-example.go]

Usage `go run aws-get-s3-example.go report-name` 



## Costs for the whole month




