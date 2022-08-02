---
description: Host snapshots at reduced cost when spinning up new nodes
---

# Create and host Chainstate Snapshots

### Create Storj account

After [creating an account](https://storj.io/signup) on Storj DCS, you’ll need to setup the uplink cli.

### Setup uplink CLI

The uplink cli is a tool similar to aws cli that allows creation of buckets and uploading/downloading snapshots directly from the 15,000+ storage nodes. We also have an aws s3 compatible api as an alternative.

#### Install uplink

Linux AMD64

```
curl -L https://github.com/storj/storj/releases/latest/download/uplink_linux_amd64.zip -o uplink_linux_amd64.zip
unzip -o uplink_linux_amd64.zip
sudo install uplink /usr/local/bin/uplink
```

For different uplink binaries see [Download Uplink CLI](../downloads/download-uplink-cli.md)

#### Create access grant and setup uplink

Uplink requires an access grant in order to access Storj DCS. You'll need to use the Storj web console to create one.

Click "Create Access Grant" in the Storj web console.

![Creating an access token in the Storj web console](<../.gitbook/assets/Screen Shot 2022-07-01 at 10.23.52 AM.png>)

Click "Continue in CLI" after giving your access grant a name

{% hint style="info" %}
Keep the credential window open until you have completed the `uplink setup` command below
{% endhint %}

![](<../.gitbook/assets/Screen Shot 2022-07-01 at 10.33.37 AM.png>)

Run `uplink setup` to start the credentials prompt&#x20;

```
uplink setup
```

Enter a name for the credential (default is "main")

Copy your "API Key" from the web console to the uplink cli "Enter API key or Access grant" prompt

![](<../.gitbook/assets/Screen Shot 2022-07-01 at 10.44.39 AM.png>)

Copy your "Satellite Address" from the web console to the uplink cli "Satellite address" prompt

![](<../.gitbook/assets/Screen Shot 2022-07-01 at 10.45.58 AM.png>)

Enter a passphrase to complete the setup

{% hint style="info" %}
Remember your **Passphrase** you will need it for future access of the data
{% endhint %}

```
$ uplink setup
Enter name to import as [default: main]: 
Enter API key or Access grant: <access grant>
Satellite address: <satellite address>
Passphrase: 
```

### **Create bucket**

Create a bucket called `snapshots`

```
uplink mb sj://snapshots
```

### **Upload snapshots**

Compress small files/directories to a single compressed file (e.g use `tar`).&#x20;

```
tar cf snapshot.tar /path/to/snapshot 
```

Use `uplink cp` to upload your snapshot to Storj DCS. Scale parallelism to at most 2x your thread count (16 threads = 32 parallelism)

```
uplink cp --parallelism 8 snapshot.tar sj://snapshots/snapshot.tar 
```

### Create d**ownload** access grant

For node operators in your community, you'll need to generate another access grant with limited permissions.

Create another access grant with the following limitations

1. Check only 'Download'  (leave others deselected)
2. Select `snapshots` as the bucket from the dropdown

![](<../.gitbook/assets/Screen Shot 2022-07-01 at 2.56.04 PM (1).png>)

Select "Continue in Browser"

Copy the access grant in file. You'll use it later in the template below.

![](<../.gitbook/assets/Screen Shot 2022-07-01 at 3.06.19 PM.png>)

{% hint style="warning" %}
When populating the template, only use the access grant that is limited to downloads. Do not use the access grant that was used to upload snapshots.
{% endhint %}

{% hint style="danger" %}
We recommend not publicly publishing access grants.
{% endhint %}

Adapt the [Download blockchain snapshots template](https://github.com/storj/chainstate-snapshots/blob/main/download-chainstate-template.md), changing `<your_access_grant>` fields to the Download-only access grant. \
\
The commands from the template can be published to your communities.
