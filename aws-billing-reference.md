# Reference

## Cost and Usage Reports (CUR) (recomended) ##

See [aws-billing.md](aws-billing.md)

## AWS Service Catalog (Scan Provision)

Well... It looks like the feature for organizations only

See example: [aws-scan-provision-example.go](aws-scan-provision-example.go)

From the landing page: (https://aws.amazon.com/servicecatalog/)

AWS Service Catalog allows **organizations** to create and manage
catalogs of IT services that are approved for use on AWS. These IT
services can include everything from virtual machine images, servers,
software, and databases to complete multi-tier application
architectures. AWS Service Catalog allows you to centrally manage
commonly deployed IT services, and helps you achieve consistent
governance and meet your compliance requirements, while enabling users
to quickly deploy only the approved IT services they need.

* Landing page: https://aws.amazon.com/servicecatalog/
* Landing page / Details: https://aws.amazon.com/servicecatalog/details/
* Doc / Service Catalog: https://docs.aws.amazon.com/servicecatalog/latest/adminguide/introduction.html


## Reports (are not recomended to use!) ##

In this section the reports from settings page are described

<details><summary>Screenshots</summary>

Billings -> Preferences -> `[ ]` Receive Billing Reports

![SettingsPage](https://qezz.github.io/shared/aws-settings-reports.png)

In some time you will see these reports at the provided s3 bucket

![BucketRootPage](https://qezz.github.io/shared/aws-settings-reports-2.png)

</details>

### Detailed Billing Report with Resources and Tags ###

> **Important**

> The Detailed Billing Report with Resources and Tags will be
> unavailable at a later date. We strongly recommend that you use the
> Cost and Usage Report instead.

## Rates ##

### Unblended rates (from [](aws.amazon.com/premiumsupport/knowledge-center/blended-rates-intro/) ) ###

For billing purposes, AWS treats all accounts in an organization as if
they're one account. The pricing tiers and capacity reservations of
the accounts in the organization are combined into one consolidated
bill, which can lower the effective price per hour for some services.

This effective price per hour is displayed as the "blended" rate in
your Billing and Cost Management console. The blended rate shown in
the console is for informational purposes only. For more information,
see Blended Rates.

If you're the master payer account for an organization, and you want
to restrict the pricing benefit of Reserved Instances to the accounts
they were purchased on, disable Reserved Instance sharing for your
organization.

### Blended rates (from [](docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/con-bill-blended-rates.html#Blended_CB)) ###

Blended rates are the averaged rates of the Reserved Instances and
On-Demand Instances that are used by member accounts in an
organization in AWS Organizations. AWS calculates blended costs by
multiplying the blended rate for each service with an account’s usage
of that service. 

[How AWS calculates blended rates for the following services](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/con-bill-blended-rates.html#Blended_CB)

