---
description: How to backup your files with Arq and Storj DCS.
---

# Arq Integration Guide

## Introduction

**Arq** is a backup software that allows you to use your own cloud storage account on a Mac or Windows based system. You can use **Arq** with **Storj DCS** S3-compatible storage.

**Main site**: [https://www.arqbackup.com/](https://www.arqbackup.com/) - you can download and follow instructions to install Arq [here](https://www.arqbackup.com/download/).

## Configure Arq to use Storj DCS



**Your Storj Account**

1. You’ll need a Storj account. Sign up [here](https://storj.io/signup), or log into your existing account [here](https://storj.io/login).
2. Click on [**Buckets**](../getting-started/satellite-developer-account/objects.md) in the Storj DCS console and [create a bucket](../getting-started/quickstart-objectbrowser.md#creating-buckets) for your Arq backups.
3. Click on [**Access**](../getting-started/satellite-developer-account/access-grants.md) in the Storj DCS console and click [**Create S3 Credentials**](../getting-started/satellite-developer-account/access-grants.md#create-s3-credentials).
4. Give it a name, select _**All**_ permissions and click **Encrypt My Access**.
5. Enter the encryption passphrase and click either on the **Copy to clipboard** link or **Download .txt** to copy or download your encryption phrase.
6. Confirm that you copied your Encryption Phrase to a safe place and click the **Create my Access** link.
7. **Leave the resulting web page open** in your browser while you configure Arq.

{% hint style="warning" %}
New Users should be presented with the option to _**Create a backup plan**_.  Existing users may need to create a backup plan from a menu.
{% endhint %}

Pick **New Backup Plan** from Arq’s File menu. Click **Add Storage Location**, choose _**Storj**_, and click **Continue**:

![](<../.gitbook/assets/image (34) (2).png>)

Copy and paste the _**Access Key**_ and _**Secret Key**_ values from your web browser into the _**Storj Access Key ID**_ and _**Storj Secret Access Key**_ fields in Arq and click **Continue**:

![](<../.gitbook/assets/image (42).png>)

Check **Use existing bucket**, choose your bucket, and click **Continue**:

![](<../.gitbook/assets/image (32) (2).png>)

Click **Continue** to use the storage location you just added:

![](https://www.arqbackup.com/blog/wp-content/uploads/2022/05/Screen-Shot-2022-05-27-at-9.45.48-AM-1024x814.png)

Choose an encryption password for Arq to encrypt your data **before** transmitting it (this password will never leave your computer):

{% hint style="warning" %}
Files are stored encrypted within the Storj network.  Using Arq's encryption would add a second layer of encryption.  Users may want to uncheck `Encrypt with password` when given the option.  This is optional.&#x20;
{% endhint %}

![](https://www.arqbackup.com/blog/wp-content/uploads/2022/05/Screen-Shot-2022-05-27-at-9.46.38-AM-1024x814.png)

Choose which files you’d like to back up, and click **Create Backup Plan**:

{% hint style="warning" %}
&#x20;To change the schedule, the files being backed up, and many other options, click on your backup plan on the left and click “**Edit…**”.
{% endhint %}

![](https://www.arqbackup.com/blog/wp-content/uploads/2022/05/Screen-Shot-2022-05-27-at-9.47.12-AM-1024x814.png)

Congratulations, you have successfully configured Arq to back up your data to Storj DCS!

\
