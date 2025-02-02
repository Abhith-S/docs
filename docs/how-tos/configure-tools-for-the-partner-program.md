---
description: >-
  How to configure the User agent for partner value attribution from
  commandline, from setup, in code and from configuration yaml.
---

# Configure Tools for the Partner Program

## Partner Program

The Storj Partner Ecosystem enables developers to build Storj DCS Connectors, which their customers can use to store data on Storj DSC.&#x20;

The data itself is client-side encrypted, however we are able to measure the aggregate volume of storage and bandwidth usage.  When a user of a Storj DCS Connector stores data in a bucket, we are able to give the partner attribution for the stored data and the used bandwidth for the Connector Integration, and provide programmatic revenue share.

You can learn more about our partner program [here](https://storj.io/partners/).

### Value Attribution

Value attribution is done on a per bucket basis. To recognize which partner the shared revenue should go to, we use a user agent that identifies it. A bucket can only have one user agent value and it can be set only once, and only on an empty bucket. This has the following consequences:

* Uploading an object to a bucket can only be done **after setting the user agent.** Otherwise, the bucket won't be empty and you will be unable to set the user agent.
* If you upload an object to a bucket with a defined user agent, the shared revenue will go to the corresponding partner. If it is not your user agent, it won't be in your shared revenue.

## Setting the User Agent

{% hint style="warning" %}
Before continuing, beware that partner value attribution is only possible if you are registered as such by Storj. You can access an up to date list of recognized user agents [here](https://github.com/storj/storj/blob/master/satellite/rewards/partners.go#L28).
{% endhint %}

### Uplink CLI

UserAgent can only be configured during setup:

```
uplink setup --client.user-agent "MyCompany"
```

or by adding or updating the following lines of the uplink configuration yaml:

```yaml
# User-Agent used for connecting to the satellite
client.user-agent: "MyCompany"
```

### **S3 Gateway**

UserAgent can only be configured during setup:

```
gateway setup --client.user-agent "MyCompany"
```

or by adding or updating the following lines of the uplink configuration yaml:

```yaml
# User-Agent used for connecting to the satellite
client.user-agent: "MyCompany"
```

### **In code**

UserAgent can be configured from Go code:

```go
uplink.Config{UserAgent: "MyCompany"}.OpenProject(...)
```

