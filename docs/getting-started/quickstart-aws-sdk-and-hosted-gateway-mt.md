---
description: Hosted S3 Compatible Multitenant Gateway - AWS SDK
---

# Quickstart - AWS SDK and Hosted Gateway MT

Storj now offers a hosted multitenant gateway (Gateway MT)  that is backward compatible with S3. This means you’ll be able to integrate with the Storj network via HTTP, and you won’t have to run anything extra on your end.

{% hint style="info" %}
By using hosted Gateway MT you are opting-in to **** [**server-side encryption**](../concepts/encryption-key/design-decision-server-side-encryption.md).&#x20;
{% endhint %}

Using Gateway MT with AWS SDK is a 2-step process:

1. [Generate Credentials to the Gateway MT](gateway-mt/#generate-credentials-to-the-gateway-mt)
2. [Configure AWS SDK with your credentials](quickstart-aws-sdk-and-hosted-gateway-mt.md#1.-install-or-include-the-amazon-s3-sdk)

## Generate Credentials to the Gateway MT

**Navigate to the Access** page within your project and then click on **Create S3 Credentials**. A modal window will pop up where you should enter a name for this access grant.

![](<../.gitbook/assets/image (24).png>)

![](<../.gitbook/assets/image (17).png>)

**Assign the permissions** you want this access grant to have, then click on **Encrypt My Access**:

![](<../.gitbook/assets/image (28).png>)

**Enter the Encryption Passphrase** you used for your other access grants. If this is your first access grant, we strongly encourage you to use a mnemonic phrase as your encryption passphrase (The GUI automatically generates one on the client-side for you.)

![](<../.gitbook/assets/image (3).png>)

{% hint style="warning" %}
**This passphrase is important!** Encryption keys derived from it are used to encrypt your data at rest, and your data will have to be re-uploaded if you want it to change!

Importantly, if you want two access grants to have access to the same data, **they must use the same passphrase**. You won't be able to access your data if the passphrase in your access grant is different than the passphrase you uploaded the data with.

Please note that **Storj does not know or store your encryption passphrase**, so if you lose it, you will not be able to recover your files.
{% endhint %}

Click either on the **Copy to clipboard** link or **Download .txt** and then confirm that you copied your Encryption Phrase to a safe place.

![](<../.gitbook/assets/image (1).png>)

Click the **Create my Access** link to finish generating of S3 credentials.

![](<../.gitbook/assets/image (12).png>)

Copy your **Access Key**, **Secret Key**, and **Endpoint** to a safe location or download them.

Now you are ready to configure AWS SDK

## Gateway MT with Amazon S3 SDK (Node.js)

### 1. Install or include the Amazon S3 SDK

e.g. with npm

```javascript
npm install --save aws-sdk
```

### 2. Import the S3 client

```javascript
import S3 from "aws-sdk/clients/s3";
```

### 3. Create client object with MT credentials

```javascript
const accessKeyId = "access key here";
const secretAccessKey = "secret access key here";
const endpoint = "https://gateway.storjshare.io";

const s3 = new S3({
  accessKeyId,
  secretAccessKey,
  endpoint,
  s3ForcePathStyle: true,
  signatureVersion: "v4",
  connectTimeout: 0,
  httpOptions: { timeout: 0 }
});
```

### 4. List objects and log to console

```javascript
(async () => {

  const { Buckets } = await s3.listBuckets({}).promise();
  
  console.log(Buckets);

})();
```

### 5. Upload an object

```javascript
(async () => {

  // `file` can be a readable stream in node or a `Blob` in the browser

  const params = {
    Bucket: "my-bucket",
    Key: "my-object",
    Body: file
  };

  await s3.upload(params, {
    partSize: 64 * 1024 * 1024
  }).promise();
  
})();
```

### 6. Get URL that points to an object

The `getSignedUrl` function creates a cryptographically signed url. No contact with the gateway is needed here; this happens instantaneously.

```javascript
const params = {
  Bucket: "my-bucket",
  Key: "my-object"
}

const url = s3.getSignedUrl("getObject", params);

// e.g. create an <img> where src points to url
```
