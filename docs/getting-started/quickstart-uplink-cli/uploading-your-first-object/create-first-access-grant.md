---
description: >-
  Access Grants give access to a project to create buckets, upload files to the
  network, and read them when they are needed.
---

# Create an Access Grant

[You need to have a satellite account and installed Uplink CLI](../prerequisites.md).

Navigate to the [**Access**](../../satellite-developer-account/access-grants.md) **** page within your project and then click on **Create S3 Credentials**. A modal window will pop up where you should enter a name for this access grant.

![](<../../../.gitbook/assets/image (24).png>)

![](<../../../.gitbook/assets/image (2) (2).png>)

{% hint style="info" %}
If you click **Encrypt My Access**, our client-side javascript will finalize your access grant with your encryption passphrase. Your data will remain end-to-end encrypted until you explicitly register your access grant with [Gateway MT](../../gateway-mt/) for S3 compatibility. Only then will your access grant be shared with our servers. Storj does not know or store your encryption passphrase.

However, if you are still reluctant to enter your passphrase into our web application, that's completely understandable, and you should cancel creation of Access Grant in Web UI, select **Create Keys for CLI** and follow these [instructions](../generate-access-grants-and-tokens/generate-a-token.md).

**The instructions below assume you selected **_**Encrypt My Access.**_
{% endhint %}

**Assign the permissions** you want this access grant to have, then click on **Encrypt My Access**:

![](<../../../.gitbook/assets/image (28).png>)

Select a **Passphrase** type: Either **Enter** your own _**Encryption Passphrase**_ or **Generate** a 12-Word _**Mnemonic Passphrase**_. Make sure you **save your encryption passphrase** as you'll not be able to reset this after it's created.

**Enter the Encryption Passphrase** you used for your other access grants. If this is your first access grant, we strongly encourage you to use a mnemonic phrase as your encryption passphrase (The GUI automatically generates one on the client-side for you.)

![](<../../../.gitbook/assets/image (3).png>)

{% hint style="warning" %}
**This passphrase is important!** Encryption keys derived from it are used to encrypt your data at rest, and your data will have to be re-uploaded if you want it to change!

Importantly, if you want two access grants to have access to the same data, **they must use the same passphrase**. You won't be able to access your data if the passphrase in your access grant is different than the passphrase you uploaded the data with.

Please note that **Storj does not know or store your encryption passphrase**, so if you lose it, you will not be able to recover your files.
{% endhint %}

Click either on the **Copy to clipboard** link or **Download .txt** and then confirm that you copied your Encryption Phrase to a safe place.

![](<../../../.gitbook/assets/image (1).png>)

Click the **Create my Access** link to finish generating of Access Grant.

![](<../../../.gitbook/assets/image (16).png>)

Access Grant is generated. **The Access Grant will only display once.** Save this information in a password manager or wherever you prefer to store sensitive information.&#x20;
