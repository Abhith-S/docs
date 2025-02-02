---
description: >-
  How to Integrate Filezilla Pro with Storj DCS to Easily Find, Transfer and
  Download All Of Your Files
---

# Filezilla Pro Integration Guide

## How to Integrate Filezilla Pro with Storj DCS to Easily Find, Transfer and Download All Of Your Files

At a quick glance, FileZilla provides a fast and reliable cross-platform (Windows, Linux and Mac OS X) FTP, FTPS and SFTP client that supports connections, file transfers and file structure browsing for many of today’s cloud data storage services, including Storj. This integration is beneficial for developers as it allows them to use the FileZilla Pro client to transfer files point-to-point using the decentralized cloud.

However, there are some caveats around using FileZilla that Storj DCS users should take into consideration, namely what version of FileZilla supports integration with Storj.

## What Is the Difference Between FileZilla and FileZilla Pro and How Does This Influence Integration with Storj?

FileZilla is available in a free version, known as FileZilla Standard. It not only supports Storj DCS and FTP, but also FTP over TLS (FTPS) and SFTP. It is open-source software distributed free of charge under the terms of the GNU General Public License.

****[**FileZillaPro**](https://filezillapro.com/) is a paid upgrade which delivers all of the base functionality of [**FileZilla Standard**](https://filezilla-project.org/), while adding additional support for many of today’s popular cloud data storage services like Amazon S3, OneDrive, Dropbox, WebDAV, Microsoft Azure, OneDrive for Business and SharePoint, Google Cloud, Backblaze and, of course, Storj DCS.

You can use FileZilla Standard with the Storj DCS native connector—as long as you don’t use a package manager to [download it](set-up-filezilla-for-decentralized-file-transfer.md#downloading-filezilla). To learn how to use the Storj integration with FileZilla Standard, check out this [how-to doc](set-up-filezilla-for-decentralized-file-transfer.md). To leverage the FileZillaPro functionality, you can integrate it with Storj using a native connector or our backwards S3-compatible Gateway MT. Below we will focus on the integration between Storj and FileZilla Pro.

## Storj + FileZilla Pro Integration&#x20;

FileZilla Pro gives users the option to send files to a Storj DCS account in two ways, either via native uplink or via [Gateway MT](https://docs.storj.io/dcs/getting-started/gateway-mt/). Let’s take a look at some of the specs of both Native Uplink and Gateway MT, providing a clearer understanding of which integration method will work better for unique use cases.

#### Native Uplink specs regarding integration with Storj DCS:

* Native Integration (Fastest for downloading large files)
* Encrypt, erasure code, and transfer from the storage nodes directly from your computer. This is ideal for downloading large files fast.
* Supports parallelism for downloads
* Has a 2.68x upload multiplier for uploads and does not support segment parallelism

#### GatewayMT specs regarding integration with Storj DCS:

* Gateway MT (Fastest for uploading large files)
* Encryption, erasure coding, and upload to storage nodes occur server side
* Supports parallelism for upload and multi-transfer for download
* A 1GB upload will result in 1GB of data being uploaded to storage nodes across the network, based on S3 standard

There are benefits to each method of integration. To provide users with the best value as they look to make the most out of their Storj and FileZillaPro integration, we’ve put together dedicated sections on integrating Storj DCS with FileZillaPro via native uplink as well as through Gateway MT.

### Storj + FileZilla Pro Via Native Uplink

Navigate to the [**Access**](../getting-started/satellite-developer-account/access-grants.md) **** page within your project and then click on **Create S3 Credentials**. A modal window will pop up where you should enter a name for this access grant.

![](<../.gitbook/assets/image (24).png>)

![](<../.gitbook/assets/image (2).png>)

{% hint style="info" %}
If you click **Encrypt My Access**, our client-side javascript will finalize your access grant with your encryption passphrase. Your data will remain end-to-end encrypted until you explicitly register your access grant with [Gateway MT](../getting-started/gateway-mt/) for S3 compatibility. Only then will your access grant be shared with our servers. Storj does not know or store your encryption passphrase.

However, if you are still reluctant to enter your passphrase into our web application, that's completely understandable, and you should cancel creation of Access Grant in Web UI, select **Create Keys for CLI** and follow these [instructions](../getting-started/quickstart-uplink-cli/generate-access-grants-and-tokens/generate-a-token.md).

**The instructions below assume you selected **_**Encrypt My Access.**_
{% endhint %}

**Assign the permissions** you want this access grant to have, then click on **Encrypt My Access**:

![](<../.gitbook/assets/image (28).png>)

Select a **Passphrase** type: Either **Enter** your own _**Encryption Passphrase**_ or **Generate** a 12-Word _**Mnemonic Passphrase**_. Make sure you **save your encryption passphrase** as you'll not be able to reset this after it's created.

**Enter the Encryption Passphrase** you used for your other access grants. If this is your first access grant, we strongly encourage you to use a mnemonic phrase as your encryption passphrase (The GUI automatically generates one on the client-side for you.)

![](<../.gitbook/assets/image (3) (1).png>)

{% hint style="warning" %}
**This passphrase is important!** Encryption keys derived from it are used to encrypt your data at rest, and your data will have to be re-uploaded if you want it to change!

Importantly, if you want two access grants to have access to the same data, **they must use the same passphrase**. You won't be able to access your data if the passphrase in your access grant is different than the passphrase you uploaded the data with.

Please note that **Storj does not know or store your encryption passphrase**, so if you lose it, you will not be able to recover your files.
{% endhint %}

Click either on the **Copy to clipboard** link or **Download .txt** and then confirm that you copied your Encryption Phrase to a safe place.

![](<../.gitbook/assets/image (1).png>)

Click the **Create my Access** link to finish generating of Access Grant.

![](<../.gitbook/assets/image (16).png>)

{% hint style="danger" %}
Please note that Storj does not know or store your encryption passphrase, so if you lose it, you will not be able to recover your files. Please store it in a safe place.
{% endhint %}

Now that the Access Grant has been generated, this will allow for integration with FileZilla Pro via native uplink. Let's take a look.

Once the FileZilla Pro client is open, select the **Open the Site Manager** icon at the top left of the FileZilla Pro client. Once open, start by selecting the **New Site** button and _**Storj - Decentralized Cloud Storage**_ as the protocol.

![](<../.gitbook/assets/image (6).png>)

Now, add the appropriate **Satellite** url (without adding `htpps://`) and simply copy your Access Grant that was previously generated within your Storj DCS account to the **Access Grant** field:

![](<../.gitbook/assets/image (137) (1) (2).png>)

Hit **Connect**, and access to your Storj DCS account should be established.

### Storj + FileZilla Pro via Gateway MT

In this section, we’ll go through the Storj FileZilla Pro integration leveraging Gateway MT.

**Navigate to the Access** page within your project and then click on **Create S3 Credentials**. A modal window will pop up where you should enter a name for this access grant.

![](<../.gitbook/assets/image (24).png>)

![](<../.gitbook/assets/image (17).png>)

**Assign the permissions** you want this access grant to have, then click on **Encrypt My Access**:

![](<../.gitbook/assets/image (28).png>)

**Enter the Encryption Passphrase** you used for your other access grants. If this is your first access grant, we strongly encourage you to use a mnemonic phrase as your encryption passphrase (The GUI automatically generates one on the client-side for you.)

![](<../.gitbook/assets/image (3) (1).png>)

{% hint style="warning" %}
**This passphrase is important!** Encryption keys derived from it are used to encrypt your data at rest, and your data will have to be re-uploaded if you want it to change!

Importantly, if you want two access grants to have access to the same data, **they must use the same passphrase**. You won't be able to access your data if the passphrase in your access grant is different than the passphrase you uploaded the data with.

Please note that **Storj does not know or store your encryption passphrase**, so if you lose it, you will not be able to recover your files.
{% endhint %}

Click either on the **Copy to clipboard** link or **Download .txt** and then confirm that you copied your Encryption Phrase to a safe place.

![](<../.gitbook/assets/image (1).png>)

Click the **Create my Access** link to finish generating of S3 credentials.

![](<../.gitbook/assets/image (41).png>)

Copy your **Access Key**, **Secret Key**, and **Endpoint** to a safe location or download them.

### Setting up regions

1. In the FileZilla menu bar, click on _**Edit**_ > _**Settings**_ ...
2. Select _**Transfers**_ > _**S3: Providers**_ from the menu on the left.\


![](https://lh3.googleusercontent.com/zvXGMbx5aGkx4JfZ0FqbITMSSXzDglshM37zoyX1pvzxZP3uq4EhHtvV3R2cXF3\_KxCeSGIm3IVjka7NCnGXLZ5fA3w1Iy6lduvJ0aIrux6OrfjaYBh2\_Rqpc4d1sJ6dfLffl09S)

3\. Click on the **Add** button under the **Providers** list.

4\. Enter _**Storj**_ as the name of the hosting provider.

5\. Press ENTER.

6\. Highlight the new hosting provider (_**Storj**_).

7\. Click on the **Add** button under the **Regions** list.

8\. Enter a name (for example, "US1") for the region and then press ENTER.

9\. Optionally, you can enter a description here.

10\. Click on the **Endpoints** column of the new region row and enter the _**Endpoint address**_ for Storj DCS, generated earlier during the creation of the Storj Gateway MT credentials (without adding `https://`).

![](<../.gitbook/assets/image (140) (2).png>)

11\. Click on **OK**.

{% hint style="success" %}
Instead of using regional endpoints, you can also use the common `gateway.storjshare.io` to route access automatically to the closest location.
{% endhint %}

### Adding a new site to FileZillaPro

1. In the menu bar, click on **File** > **Site Manager** …
2. Click on **New Site**.
3. Enter a name for the Site - for example, _**US1 GatewayMT**_.
4. In the protocol section, select **S3 - Amazon Simple Storage Service** from the **Protocol** drop-down list.

![](<../.gitbook/assets/image (124) (2).png>)

5\. Provide the **Access key ID** and **Secret Access Key** in the parameters for the new Site.

6\. Click **Connect** to connect to Storj via Gateway MT.

{% hint style="info" %}
If this is the first time you connect, you may see a message like this:

![](<../.gitbook/assets/image (26) (1).png>)

You will need to confirm it by clicking the **OK** button.
{% endhint %}

You should now see your buckets:

![](<../.gitbook/assets/image (143).png>)
