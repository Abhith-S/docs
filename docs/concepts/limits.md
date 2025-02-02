---
description: Account Limits, and how to increase them
---

# Usage Limits

Usage Limits allow us to ensure a consistent level of service for all customers. We have limits for usage established per Project on all Storj DCS Satellites. All limits are set to default values as follows:

## **PRO Account (Paid Tier)**

### Credit Card Payment method

Adding a credit card as payment method will result in your per project limits being automatically raised to Pro Account limits:

* 3 projects
* 100 buckets per project
* 25 TB storage per project
* 100 TB egress bandwidth per project
* 100 request per second rate limit

### STORJ token Payment method

However, if your only payment method on file is STORJ token, you can submit a [request to increase your projects limits](../billing-payment-and-accounts-1/pricing/usage-limit-increases.md). You should [deposit](../getting-started/satellite-developer-account/my-account/billing.md#adding-storj-tokens) a sufficient amount of tokens to your account to cover potential usage up to the requested limits for at least for the next 3 months. You will be charged only for actual usage above a [Free Tier](limits.md#free-tier).

{% hint style="info" %}
Note that you could avoid having to wait for a manual project limit increase to be applied by adding a credit card as backup payment method. In this case, your usage would still get paid from your STORJ balance if it is sufficient to cover your outstanding charges. Only usage that exceeds what can be covered by your STORJ balance at time of invoicing would get charged to the credit card.
{% endhint %}

## **Free Tier**

* 1 Project
* 100 Buckets per Project
* 150 GB Storage per Project
* 150 GB Egress Bandwidth per Project
* 100 request per second rate limit
* 10,000 Segments per Month

If you would like to increase your limits to higher values and your only payment method is STORJ token, you may[ contact our support team through the Storj DCS support portal](https://supportdcs.storj.io/hc/en-us/requests/new?ticket\_form\_id=360000683212).​

## Rationales behind limits

Storage and bandwidth limits are imposed by most cloud infrastructure providers as a normal part of capacity planning and to ensure achievement of SLAs. In distributed and decentralized storage systems they are equally important, if not more so. Just like any provider, the aggregate amount of available storage and bandwidth must be shared across all users. With a distributed and decentralized storage system like Storj DCS, the storage and bandwidth are provided by a network of third parties running storage node software. One of the key aspects to success is the balance of supply and demand. If there are too many users over-utilizing available resources, the user experience will be poor.

If there are too many storage nodes, there won’t be enough use to provide a meaningful ROI for Storage Node Operators. This can lead to storage node churn, increasing load on the network, and potentially impacting durability. Usage limits are one of the tools that maintain the balance.

We set rate limits between the uplink and the satellite to ensure a good quality of service for all uplink users. Without the rate limit it would be possible for users to inadvertently consume most of the database resources available on the satellite and cause issues for other users.  We selected a limit that would be mostly unnoticed by end users (as the typical use case shouldn’t hit the limit). The current default limits are based on requests per second for all meta info calls: list, get, delete, put.

We set the default limits for the number of buckets per project to ensure performance.&#x20;

In addition, we limit the number of Projects per Developer Account to minimize complexity.&#x20;

We have also set the default limit for the number of segments to a level that is healthy for the network. Note that increasing the Segment Project Limit may incur additional fees. Read more about [Usage Limit Increases](../billing-payment-and-accounts-1/pricing/usage-limit-increases.md).&#x20;

Customers can request a limit increase when needed by filling out the [limit increase request form](https://supportdcs.storj.io/hc/en-us/requests/new?ticket\_form\_id=360000683212) on our Storj DCS support portal if their only payment method on file is STORJ token.

An automatic limit increase to Pro Account can be accomplished by adding a credit card as payment method. Please only make such requests if your use case really requires more than the current default limits. Requests will be evaluated taking into account the intended use case and availability on the network.
