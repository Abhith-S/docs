---
description: How to setup Comet Backup with Storj
---

# Comet Backup Integration Guide

### Introduction

Whether you’re looking for cloud backup, computer backup or data backup, [Comet](https://cometbackup.com/) provides solutions to protect and restore partitions, databases, servers, files and folders. We support Windows, MacOS and Linux backup.

Comet scales with your business. Provide backup on your terms. No contracts. Free replication, free support and free white labeling, rebrand the software with your logo. Comet’s intuitive, all-in-one platform lightens the workload for your team.

Main site: [https://cometbackup.com/](https://cometbackup.com/) - You can download a free trial [here](https://cometbackup.com/signup).

### Storj with Comet Backup

Using Storj with Comet Backup provides resilient cloud object storage with blazing performance and zero-trust security.

1. **Speed of recovery**: CDN-like performance at cold storage prices for instant recovery of your backup data
2. **Durability**: long term, your data is never going away
3. **Security**: Ransomware resistant with end to end encryption

### Before you begin&#x20;

If you do not currently have a Comet Server, please refer to Comet’s [getting started guide](https://docs.cometbackup.com/latest/). Storj functionality is available on Comet Server version 22.9.0 and above.&#x20;

## Storj setup

To begin setting up Comet with Storj, you'll need to first create an account on Storj.

Navigate to [https://storj.io/signup](https://storj.io/signup) or login to an existing Storj account.

![](https://lh5.googleusercontent.com/RhslyMTj8ubEkpt-yHlE3w3eUPF6MX5-gK5H\_QU5TzqVMQV7TL4H5H7JcW4gWeU7WyoveqwK2IEu0ADjL-q4Jy3cDyVSoNHN8tgBS\_swT2\_ob-FZ6zP\_QLhpYai\_gM5wHYU3ObIfYPrPVNg-gcABe4A\_AtZRSsg-xzVcQlJ5hhnLS2vUKTikpDPvAg)

### Create a Storj Bucket

&#x20;You need to create a Storj bucket for Comet to create Storage Vaults in.&#x20;

1. Navigate to “Buckets” on the left side menu
2.  Click “New Bucket” on the top right\


    <figure><img src="https://lh4.googleusercontent.com/8e7Gk8yvyWm_V67JUVErQo3yRb89O4R9YwiEmixsbtSQK8J43vkGvxac9XpSPWSPmISRmK2HF3JHjnKLimYN1dQTYZNcRpUcPFLsrqxnLbwMkBNB7u8hQzb9ro8RzMZZXNvcQE_kvOIO7NgF5oWO7GLPSrFKjXpUy2FbfK7U9AfdvNbRAo96QLuYDw" alt=""><figcaption></figcaption></figure>
3. Name the bucket something identifiable such as “comet-backups”\
   ![](https://lh6.googleusercontent.com/6\_SKlYLgxlTYsLJRrOKU4WPVcHOv31VOfpx79riaV1UXksquQXWP1lOqBfe667F6uZ\_GGE21DuwxoNzy\_0XORbDuSYt0wb\_f6lT9lQS2MbH1p5L4nutVVGtFxXE5OK2dWdTjY1ot-mt1sleyl11X68SeVtmmdMUq7GZ-UwRPMsvnuktFmtFMAD2TNQ)
4. Select "Continue"
5. Generate a passphrase or enter your own.\
   ![](https://lh3.googleusercontent.com/EkKbpTiLXC953KQLMbHfbCFKZ2xm3y9WLA8XoQJk0U7aLAIe1lhS3c4kXTpI1ibHMnRRWjXkE7Yacoq7DlBphVZxzwcce2PxiJU6VjKv99Sa8zZ3qOl6bpmA-08BW1VqtO14lO4SZj2ZnQIdb264zi1l9kO3Olo2caw3rQocOAylBg8PRFSIj5Mxkg)
6. Select "Continue"
7. Record the passphrase somewhere safe.

{% hint style="info" %}
Remember your passphrase as you will need it for future access of your data. Storj is unable to recover your passphrase for you.
{% endhint %}

### Create Storj Access Key

You will need to configure a Storj access key for Comet to be able to create Storage Vaults.&#x20;

To create an API Access key in the Storj web console:

1. Navigate to "Access" on the left side menu
2. Click "Create Keys for CLI" under the API key block\
   ![](<../.gitbook/assets/Screen Shot 2022-09-09 at 9.03.17 AM.png>)
3. **Type:** API Access
4. **Name:** The name of the access key (e.g. Comet)
5. **Permissions:** All
6. **Buckets:** Feel free to specify the bucket created above (e.g. comet-backups), or leave as “All”\
   ![](https://lh3.googleusercontent.com/mOYHJkcfQ7qj5BpkgICFBB35-EsRStC64m\_gAloDXsRfl3QrHw2r2kMNfidqog0l1070zvRyU7UmJLtox7mE0OfKYGp4Mok1yvS9Eqov2V\_5MDRIYethXeYYVlfe5g45\_JM-w32wuZb3H4w0BAQaL75Mim4VNQxN5HwvlBqg6L3DthzF8icKwWQvBQ)
7. Click “Create Keys” and record the API key somewhere safe.

## Comet setup

### Configuring a Storj Storage Template

1. Navigate to your Comet Server settings and find the Storage Templates section
   1. &#x20;For Comet-Hosted “Settings -> Storage -> Storage Templates”
   2. For Self-Hosted “Settings -> Authentication -> Storage Templates”
2. Click the green “Add” button to add a new Storage Template.\
   ![](https://lh4.googleusercontent.com/tQzG0Grs6Ynqzg5Y5f3OATY11m2uVP3IsVhOTsWxwbYogcrR9TOUeM8Vcl5hogaG4djrZGbs91xQAJBsaPmCfws2zfIf5ValxIbIuLTuw0207D5lpFpZVPBC5PNVS2\_JBWD03Yb4gCfGo\_IoUDliajaZA0l8ZQKKAVv7pHnemudvZEfOEj5KDGdyvw)
3. Change the “Type” to Storj DCS and then fill in the fields
   1. **Description:** What the storage will be called on the Comet Server
   2. **Satellites:** The location of the Storj satellite server. You can locate which satellite you're using by looking at the url of the Storj web console. \
      ![](<../.gitbook/assets/Screen Shot 2022-09-09 at 9.11.25 AM.png>)
   3. **API Key**: The API key created in [#create-storj-access-key](comet-backup-integration-guide.md#create-storj-access-key "mention")
   4. **Passphrase**: The passphrase created in [#create-a-storj-bucket](comet-backup-integration-guide.md#create-a-storj-bucket "mention")
   5. **Bucket**: The name of the bucket created in [#create-a-storj-bucket](comet-backup-integration-guide.md#create-a-storj-bucket "mention")\
      ![](https://lh4.googleusercontent.com/Nq8t3od-8LOpq6Ez65J4\_I9Mer3560zSWbdQw26V72x8sKoV23ILtclJ6pPMKWJfr2HqjN4vkyFGQjYWU88tEZZ83CfemCeyZfBHkOqBbIT9\_3wPgp7xVsLfE3TpxOqIgFCCj\_hTI55Dd9DG\_GtBpTtYgqyj5NWIvKpH63CWx5WI0KEwqlQN4stk3w)
4. Once complete you should now be able to request a new Storj Storage Vault for a user.\
   ![](https://lh3.googleusercontent.com/TqQOk5N0fIGOuBiDOHJWyqGhxVfxddjuZYIE3EL-IvhwbMnuO-HSgUK9-fneTNlVLkJRD6DfB9MzwYG1lmzDR9VB48NMJrUho9my-V6LqL2N4ZFScE4b5-Xas0kpTjL0XnNIA9YrqLpOAHa6g\_5WDaPBQ47MldrafK1p3O-z0grEhgq6j9H8uA4GIA)
