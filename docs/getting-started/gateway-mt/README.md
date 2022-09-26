---
description: Hosted S3 Compatible Multitenant Gateway
---

# Quickstart - AWS CLI and Hosted Gateway MT

Storj now offers a hosted multitenant gateway (Gateway MT)  that is backward compatible with S3. This means you’ll be able to integrate with the Storj network via HTTP, and you won’t have to run anything extra on your end.

{% hint style="info" %}
By using hosted Gateway MT you are opting in to **** [**server-side encryption**](../../concepts/encryption-key/design-decision-server-side-encryption.md).&#x20;
{% endhint %}

Using Gateway MT with AWS CLI is a 2-step process:

1. [Generate Credentials to the Gateway MT](./#generate-credentials-to-the-gateway-mt)
2. [Configure AWS CLI with your credentials](./#configure-aws-cli-with-your-credentials)

### Generate Credentials to the Gateway MT

**Navigate to the Access** page within your project and then click on **Create S3 Credentials**. A modal window will pop up where you should enter a name for this access grant.

![](<../../.gitbook/assets/image (24).png>)

![](<../../.gitbook/assets/image (17).png>)

**Assign the permissions** you want this access grant to have, then click on **Encrypt My Access**:

![](<../../.gitbook/assets/image (28).png>)

**Enter the Encryption Passphrase** you used for your other access grants. If this is your first access grant, we strongly encourage you to use a mnemonic phrase as your encryption passphrase (The GUI automatically generates one on the client-side for you.)

![](<../../.gitbook/assets/image (3) (1).png>)

{% hint style="warning" %}
**This passphrase is important!** Encryption keys derived from it are used to encrypt your data at rest, and your data will have to be re-uploaded if you want it to change!

Importantly, if you want two access grants to have access to the same data, **they must use the same passphrase**. You won't be able to access your data if the passphrase in your access grant is different than the passphrase you uploaded the data with.

Please note that **Storj does not know or store your encryption passphrase**, so if you lose it, you will not be able to recover your files.
{% endhint %}

Click either on the **Copy to clipboard** link or **Download .txt** and then confirm that you copied your Encryption Phrase to a safe place.

![](<../../.gitbook/assets/image (1).png>)

Click the **Create my Access** link to finish generating of S3 credentials.

![](<../../.gitbook/assets/image (41).png>)

Copy your **Access Key**, **Secret Key**, and **Endpoint** to a safe location or download them.

Now you are ready to configure AWS CLI.

### Configure AWS CLI with your credentials

{% hint style="info" %}
To continue make sure you have the AWS CLI installed on your machine.&#x20;
{% endhint %}

Verify your AWS CLI version by running `aws --version`in your terminal. AWS CLI current version is version 2. If you are using AWS CLI v1, you will need to install a plugin to be able to define the endpoint. See how [here](aws-cli-advanced-options.md#define-an-endpoint-with-aws-cli-v1).

2\. Configure your AWS CLI with the gateway MT credentials from the previous step by running `aws configure` in your terminal:

```
~ % aws configure 
AWS Access Key ID [****************e53q]: <<yourAccessKey>>
AWS Secret Access Key [****************bbxq]: <<yourSecretKey>>
Default region name [us-east-1]: 
Default output format [None]: 
~ % 
```

3\.  **Optional but strongly recommended**: Set the multipart threshold to 64 MB.&#x20;

You can now use AWS CLI. Some examples of use:

#### Make a bucket

```
~ % aws s3 --endpoint-url=https://gateway.storjshare.io mb s3://waterbear
```

#### Display buckets

```
aws s3 --endpoint-url=https://gateway.storjshare.io ls
```

#### Copy a file

```
aws s3 --endpoint-url=https://gateway.storjshare.io cp /tmp/test.zip s3://waterbear
```

#### List files in a bucket

```
aws s3 --endpoint-url=https://gateway.storjshare.io ls s3://waterbear
```

#### Copy a file from a bucket

```
aws s3 --endpoint-url=https://gateway.storjshare.io cp s3://waterbear/test.zip /tmp/Archive.zip
```

#### Delete a bucket

```
aws s3 --endpoint-url=https://gateway.storjshare.io rb s3://waterbear/
```

#### Delete a non-empty bucket

```
aws s3 --endpoint-url=https://gateway.storjshare.io rb --force s3://waterbear/
```
