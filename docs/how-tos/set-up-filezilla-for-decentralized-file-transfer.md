---
description: >-
  This guide walks users through the process around setting up FileZilla to
  transfer files over Storj DCS
---

# FileZilla Native Integration

{% hint style="info" %}
The native integration uses [**end-to-end encryption**](../concepts/encryption-key/design-decision-end-to-end-encryption.md) for your object data, including metadata and path data.
{% endhint %}

{% hint style="warning" %}
This is the only integration available for the **free version of Filezilla**. If you wish to use the Hosted Gateway MT you will need the [paid version of Filezilla](filezilla-pro-integration-guide.md).&#x20;
{% endhint %}

## **Background**

The _**FileZilla**_ Client is a fast and reliable cross-platform (Windows, Linux and Mac OS X) FTP, FTPS and SFTP client with lots of useful features and an intuitive graphical user interface.

It includes a site manager to store all your connection details and logins, as well as an Explorer-style interface that shows the local and remote folders and can be customized independently.

With the launch of the native Storj DCS Integration into the FileZilla client, developers can use the client configured to transfer files, point-to-point using the decentralized cloud.

## Getting Started

![Getting Started Guide to Configure Storj DCS with Filezilla](<../.gitbook/assets/image (77).png>)

### Create an Access Grant

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

### Downloading FileZilla

To download the latest release of FileZilla, navigate to [https://filezilla-project.org/download.php?show\_all=1](https://filezilla-project.org/download.php?show\_all=1) and select the version appropriate for your operating system, then install FileZilla.

### Creating a new Site

Open the Site Manager by clicking on the leftmost icon.

![](<../.gitbook/assets/image (101).png>)

Select the 'New Site' option

![](<../.gitbook/assets/image (116).png>)

### Configure the Satellite and Access Grant

Next, select Protocol:  "Storj - Decentralized Cloud Storage" from the Protocol dropdown in the "General" tab.&#x20;

Now enter the **Satellite** and **Access Grant** as shown below (Entering the port is not required)

1. Use the **Satellite** URL from which you created the Access Grant.
   * us1.storj.io (deprecated us-central-1.tardigrade.io)
   * ap1.storj.io (deprecated asia-east-1.tardigrade.io)
   * eu1.storj.io (deprecated europe-west-1.tardigrade.io)
2. For **Access Grant** please enter the Access Grant you saved above.

![](<../.gitbook/assets/image (115).png>)

After you enter the above information, hit the **Connect** button, and FileZilla will connect directly to the remote site. You should see a screen showing your local site vs. Storj DCS, like so:

![](<../.gitbook/assets/image (113).png>)

### Uploading a File

To upload a file to your local machine, simply drag it from the local to the remote site (on the decentralized cloud), as shown below:

![](../.gitbook/assets/upload.gif)

### Downloading a File

To download a file to your local machine, simply drag it from the remote site to the local site, as shown below:

![](../.gitbook/assets/download.gif)

