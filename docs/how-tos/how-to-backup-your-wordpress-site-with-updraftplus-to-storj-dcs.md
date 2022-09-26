---
description: Backup with UpdraftPlus to Storj DCS
---

# How to backup your WordPress site with UpdraftPlus to Storj DCS

## Installation of UpdraftPlus

UpdraftPlus is a common plugin for WordPress. First, you need to [install WordPress](https://wordpress.org/support/article/how-to-install-wordpress/). You can also use`docker-compose` to [install it locally or on your remote server](https://docs.docker.com/samples/wordpress/).

To download, install and activate the UpdraftPlus plugin, please follow their guide: [https://updraftplus.com/download/](https://updraftplus.com/download/)

![](<../.gitbook/assets/image (159) (1).png>)

We will continue with configuring the UpdraftPlus plugin after we have created the Gateway MT credentials.

## Generate Gateway MT credentials

Please sign in to your Storj DCS account and **Navigate to the Access** page within your project and then click on **Create S3 Credentials**. A modal window will pop up where you should enter a name for this access grant.

![](<../.gitbook/assets/image (24).png>)

![](<../.gitbook/assets/image (17).png>)

**Assign the permissions** you want this access grant to have, then click on **Encrypt My Access.**

![](<../.gitbook/assets/image (28).png>)

**Enter the Encryption Passphrase** you used for your other access grants. If this is your first access grant, we strongly encourage you to use a mnemonic phrase as your encryption passphrase (The GUI automatically generates one on the client-side for you.)

![](<../.gitbook/assets/image (3) (1).png>)

Click either on the **Copy to clipboard** link or **Download .txt** and then confirm that you copied your Encryption Phrase to a safe place.

![](<../.gitbook/assets/image (1).png>)

Click the **Create my Access** link to finish generating of S3 credentials.

![](<../.gitbook/assets/image (41).png>)

Copy your **Access Key**, **Secret Key**, and **Endpoint** to a safe location or download them.

Now you are ready to configure **UpdraftPlus** plugin.

## Configuring UpdraftPlus plugin to work with Storj DCS

Now that we have finished generating the Gateway MT credentials, let´s go back to the UpdraftPlus configuration. Once the plugin has been activated, you should open its settings:

![](<../.gitbook/assets/image (159) (1).png>)

1\. Click the **Settings** tab at the top part of the _**Settings**_ page of the **UpdraftPlus** plugin.

![](<../.gitbook/assets/image (136) (2).png>)

2\. Specify a preferred backup schedule for files and/or databases and how many incremental backups you want to have. See [https://wordpress.org/plugins/updraftplus/](https://wordpress.org/plugins/updraftplus/) for more details.

3\. Click on the **S3-Compatible (Generic)** option. The following fields need to be filled in: **S3 access key**, **S3 secret key**, **S3 location** and **S3 end-point**.

![](<../.gitbook/assets/image (168).png>)

4\. Specify your **Access Key** from your GatewayMT credentials in the **S3 access key** field, your **Access Secret Key** from the GatewayMT credentials in the **S3 secret key** field, **Endpoint** from the GatewayMT credentials in the **S3 end-point** field (without `https://`), and your bucket access style in the **S3 location** field. Then click the **Test S3 settings** button below. If everything was specified correctly, you should see a successful message like this:

![](<../.gitbook/assets/image (135) (1).png>)

5\. Please save the **UpdraftPlus** plugin settings with the **Save changes** button at the bottom of the page.

![](<../.gitbook/assets/image (156).png>)

After you successfully configured Storj DCS as your storage back end in the **UpdraftPlus** plugin and saved your configuration, you can return back to the _**Backup / Restore**_ tab to start your backup manually right away.
