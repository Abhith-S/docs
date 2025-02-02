# Simplify File Management with S3 Browser and Storj DCS

S3 Browser is a Windows-based client that provides simple and reliable file management for AWS S3 storage and [AWS S3 compatible storage such as Storj DCS](https://www.storj.io/blog/what-is-s3-compatibility). Via the intuitive web file management interface, users can store and retrieve files from their Storj DCS bucket anytime and anywhere.

While S3 Browser is free for personal use, users who wish to utilize the S3 Browser in commercial, business, government, or military institutions, or for any other profit activity, must purchase a pro license.

Keep in mind that to get the best performance in using S3 Browser to manage your Storj buckets, you’ll want to apply some additional configurations to your S3 Browser instance.

S3 Browser can be configured to download large files via multiple parallel threads. By default, S3 Browsers will download everything using 5MB chunks, whereas, you have the configuration option to increase that download size to 64MB, the segment size for Storj. We suggest configuring your S3 Bucket instance to _**Enable Multipart downloads with part size (in megabytes)**_ of 64. You can find more on configuring this option [here](https://s3browser.com/multipart-downloads.aspx).

{% hint style="info" %}
One license allows you to install one instance of S3 Browser on a single computer. Your license can be transferred if you change your PC. The license is a lifetime license and includes one year of free upgrades and support. Users are also limited to two accounts added within the free version of S3 Browser.
{% endhint %}

### Downloading S3 Browser

![](https://lh5.googleusercontent.com/gpQ2ngjwlFco3En36k6AgjFCRp8OiDgYTeRnNGbUuukhBmqYCxVEf4ahWxUxRWjLPZfkNtv3Y21Rnf21copq3HtTHMXNixWYPtz62jzBfJMv7z2cWCfrt3NAiVsh9\_JZPmX0C9Iv)

As noted, S3 Browser is only available for Windows, supporting **Windows XP/Vista/7/8/10/11 and Windows Server 2003/2008/2012/2016/2019/2022.**

Users can download the S3 Browser client by navigating to the S3 Browser homepage at [https://s3browser.com/](https://s3browser.com/) and selecting the _Download_ icon, or at [https://s3browser.com/download.aspx](https://s3browser.com/download.aspx).

Some stats for the S3 Browser Download:&#x20;

**S3 Browser Version** 10.3.1&#x20;

**Size of file**: 5.37 MB (5 631 160 bytes)&#x20;

**SHA256**: 0b813e6f4d5cc9d2898fd9045f577d0f5e750dd960408abf3894b447033143e2&#x20;

**Operating System**: Windows XP/Vista/7/8/10/11 and Windows Server 2003/2008/2012/2016/2019/2022

{% hint style="info" %}
There is no option to download S3 Browser via CLI
{% endhint %}

### Generate Credentials to the Gateway MT

Users interested in accessing their Storj DCS bucket(s) via S3 Browser can do so via the hosted AWS multi-tenant gateway known as Gateway MT. This backward-compatible hosted gateway is one of the most versatile ways to get up and running with Storj DCS when using platforms such as S3 Browser or other file manager platforms that support Storj DCS.

**Gateway MT offers the following:**

* Encryption, erasure coding, and upload to nodes occur server side
* Supports parallelism for upload and multi transfer for download
* A 1GB upload will result in 1GB of data being uploaded to storage nodes across the network based on the S3 standard

**Navigate to the Access** page within your project and then click on **Create S3 Credentials**. A modal window will pop up where you should enter a name for this access grant.

![](<../.gitbook/assets/image (24).png>)

![](<../.gitbook/assets/image (17).png>)

**Assign the permissions** you want this access grant to have, then click on **Encrypt My Access**:

![](<../.gitbook/assets/image (28).png>)

**Enter the Encryption Passphrase** you used for creating your other access grants. If this is your first access grant, we strongly encourage you to use a mnemonic phrase as your encryption passphrase (The GUI automatically generates one on the client-side for you.)

![](<../.gitbook/assets/image (3) (1).png>)

Click either on the **Copy to clipboard** link or **Download .txt** and then confirm that you copied your Encryption Phrase to a safe place.

![](<../.gitbook/assets/image (1).png>)

Click the **Create my Access** link to finish generating of S3 credentials.

![](<../.gitbook/assets/image (41).png>)

Copy your **Access Key**, **Secret Key**, and **Endpoint** to a safe location or download them. We’ll be using this shortly!

{% hint style="info" %}
Storj does not know or store your encryption passphrase. However, if you are still reluctant to enter your passphrase into our web application, that’s completely understandable, and you should go back to the main **Access** page and select **Create Keys for CLI** and follow [these instructions](../getting-started/quickstart-uplink-cli/generate-access-grants-and-tokens/generate-a-token.md).
{% endhint %}

### Configuring Storj + S3 Browser

Now that your S3 Browser client is downloaded and installed and you’ve generated and saved your S3 Gateway Credentials, it’s time to configure S3 Browser to interface with your Storj DCS bucket.

Select the **Accounts** menu item at the top left of the S3 Browser client. Select **Add New Account**. Add any name to your account in the **Display Name** selection. In the dropdown menu titled **Account type** select **S3 Compatible Storage**.

![](<../.gitbook/assets/S3 Browser #1.png>)

In the **REST Endpoint** section enter the **S3 Gateway Credentials End Point** without `https://` and trailing `/`.

In the **Access Key ID** section enter the **S3 Gateway Credentials Access Key**.

In the **Secret Access Key** section enter the **S3 Gateway Credentials Secret Key**.

Optionally, you’ll be able to protect your Access Keys with a master password by selecting the **Encrypt Access Keys with a Password** checkbox.

Check the box **Use secure transfer (SSL/TSL)** to secure all transfers via (SSL/TLS).

![](<../.gitbook/assets/image (23).png>)

Finally, hit **Connect**.

If you’ve added in your S3 Gateway Credentials properly, you’ll see the following:

![](<../.gitbook/assets/S3 Browser #2.png>)

### Uploading Files to Storj DCS Through S3 Browser

Within the S3 Browser, you’ll be able to upload files directly to your Storj DCS bucket once you’ve effectively tied in StorJ DCS to S3 Browser.

Start by selecting which Storj bucket you wish to upload data into by selecting the bucket at the top left. Once you’ve selected your bucket, select the **Upload** icon.

Here, you’ll be prompted to select whether you’d like to **upload file(s)** or **upload folder(s)**.

Following a selection of **upload file(s)** or **upload folder(s)** you’ll be prompted with a file navigator to select the file or folder you wish to upload.

### Downloading Files From Storj DCS Through S3 Browser

Within the S3 Browser, you’ll be able to download files directly from your Storj DCS bucket once you’ve effectively tied in StorJ DCS to S3 Browser.

Start by selecting which Storj bucket you wish to download data from by selecting the bucket at the top left.

Once you’ve selected your bucket, select the **Download** icon.

Here, you’ll be prompted to select whether you’d like to **download file(s)** or **download folder(s)**

Following a selection of **download file(s)** or **download folder(s)** you’ll be prompted with a file navigator to select the file or folder you wish to download.
