# Reference

## Reports ##

In this section the reports from settings page are described

<details><summary>Screenshots</summary>

Billings -> Preferences -> `[ ]` Receive Billing Reports

![SettingsPage](qezz.github.io/shared/aws-settings-reports.png)

In some time you will see these reports at the provided s3 bucket

![BucketRootPage](qezz.github.io/shared/aws-settings-reports-2.png)

</details>

### Detailed Billing Report with Resources and Tags ###

> **Important**

> The Detailed Billing Report with Resources and Tags will be
> unavailable at a later date. We strongly recommend that you use the
> Cost and Usage Report instead.

## Rates ##

### Unblended rates (from aws.amazon.com/premiumsupport/knowledge-center/blended-rates-intro/) ###

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

### Blended rates (from docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/con-bill-blended-rates.html#Blended_CB) ###

Blended rates are the averaged rates of the Reserved Instances and
On-Demand Instances that are used by member accounts in an
organization in AWS Organizations. AWS calculates blended costs by
multiplying the blended rate for each service with an account’s usage
of that service. 

[How AWS calculates blended rates for the following services](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/con-bill-blended-rates.html#Blended_CB)

