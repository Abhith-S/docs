---
description: How to backup your files with Arq and Storj DCS.
---

# Arq Integration Guide

## Introduction

**Arq** is a backup software that allows you to use your own cloud storage account on a Mac or Windows based system. You can use **Arq** with **Storj DCS** S3-compatible storage.

**Main site**: [https://www.arqbackup.com/](https://www.arqbackup.com/) - you can download and follow instructions to install Arq [here](https://www.arqbackup.com/download/).

## Configure Arq to use Storj DCS



**Your Storj Account**

1. You’ll need a Storj account. Sign up [here](https://us1.storj.io/signup), or log into your existing account [here](https://us1.storj.io/login).
2. Click on “Buckets” in the Storj DCS console and create a bucket for your Arq backups.
3. Click on “Access” in the Storj DCS console and click “Create Access Grant”.
4. Give it a name and click “Continue”.
5. Accept the default permissions and click “Continue in Browser”.
6. Enter an encryption passphrase and click “Next” (Storj uses this password to encrypt your data within their system)
7. Click “Generate S3 Gateway Credentials” at the bottom.
8. Click “Generate Credentials”. **Leave the resulting web page open** in your browser while you configure Arq.

{% hint style="warning" %}
New Users should be presented with the option to “Create a backup plan”.  Existing users may need to create a backup plan from a menu.
{% endhint %}



Pick “New Backup Plan” from Arq’s File menu. Click “Add Storage Location”, choose “Storj”, and click Continue:

![](<../.gitbook/assets/image (34).png>)

Copy and paste the “Access Key” and “Secret Key” values from your web browser into the “Storj Access Key ID” and “Storj Secret Access Key” fields in Arq and click “Continue”:

![](<../.gitbook/assets/image (42).png>)

Check “Use existing bucket”, choose your bucket, and click “Continue”:

![](<../.gitbook/assets/image (32).png>)

Click “Continue” to use the storage location you just added:

![](https://www.arqbackup.com/blog/wp-content/uploads/2022/05/Screen-Shot-2022-05-27-at-9.45.48-AM-1024x814.png)

Choose an encryption password for Arq to encrypt your data **before** transmitting it (this password will never leave your computer):

{% hint style="warning" %}
Files are stored encrypted within the Storj network.  Using Arq's encryption would add a second layer of encryption.  Users may want to uncheck `Encrypt with password` when given the option.  This is optional.&#x20;
{% endhint %}

![](https://www.arqbackup.com/blog/wp-content/uploads/2022/05/Screen-Shot-2022-05-27-at-9.46.38-AM-1024x814.png)

Choose which files you’d like to back up, and click “Create Backup Plan”:

{% hint style="warning" %}
&#x20;To change the schedule, the files being backed up, and many other options, click on your backup plan on the left and click “Edit…”.
{% endhint %}

![](https://www.arqbackup.com/blog/wp-content/uploads/2022/05/Screen-Shot-2022-05-27-at-9.47.12-AM-1024x814.png)

Congratulations, you have successfully configured Arq to back up your data to Storj DCS!

\
