---
description: >-
  Copies a local file or Storj object to another location locally or in Storj
  DCS
---

# cp

## Usage

{% tabs %}
{% tab title="Windows" %}
```
./uplink.exe cp [flags] SOURCE DESTINATION
```
{% endtab %}

{% tab title="Linux" %}
```
uplink cp [flags] SOURCE DESTINATION
```
{% endtab %}

{% tab title="macOS" %}
```
uplink cp [flags] SOURCE DESTINATION
```
{% endtab %}
{% endtabs %}

The `cp` command is used to upload and download objects. The `cp` command abstracts the complexity of encryption, erasure coding, and distributing pieces of a file to storage nodes.

## Flags

| Flag                            | Description                                                                                                                                                              |
| ------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `--access string`               | the serialized access, or name of the access to use                                                                                                                      |
| `-r, --recursive`               | Perform a recursive copy                                                                                                                                                 |
| `-t, --transfers int`           | Controls how many uploads/downloads to perform in parallel (default 1)                                                                                                   |
| `--dry-run`                     | Print what operations would happen but don't execute them                                                                                                                |
| `--progress`                    | Show a progress bar when possible (default true)                                                                                                                         |
| `--range string`                | Downloads the specified range bytes of an object. For more information about the HTTP Range header, see https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.35 |
| `-p, --parallelism int`         | Controls how many parallel chunks to upload/download from a file (default 1)                                                                                             |
| `--parallelism-chunk-size Size` | Controls the size of the chunks for parallelism (default 64.0 MiB)                                                                                                       |
| `--expires relative_date`       | Schedule removal after this time (e.g. '+2h', 'now', '2020-01-02T15:04:05Z0700')                                                                                         |
| `--help`, `-h`                  | help for cp                                                                                                                                                              |

## Examples

### Copy a local file into an existing bucket

When the `cp` command is used to copy a file to Storj DCS (upload), the CLI first encrypts the file client-side, then splits it into a minimum of x erasure-coded pieces, and finally, the x pieces are uploaded in parallel to x different storage nodes. x currently equals 80 but is subject to change depending on continuous optimization.&#x20;

To copy `cheesecake.jpg` into the existing bucket `cakes`, use the following command:

{% tabs %}
{% tab title="Windows" %}
```
./uplink.exe cp cheesecake.jpg sj://cakes 
```
{% endtab %}

{% tab title="Linux" %}
```
uplink cp cheesecake.jpg sj://cakes
```
{% endtab %}

{% tab title="macOS" %}
```
uplink cp cheesecake.jpg sj://cakes
```
{% endtab %}
{% endtabs %}

{% hint style="info" %}
You cannot use regular expressions to specify which files to copy (e.g. `uplink cp cheese* sj://cakes` will not work). Also, you can only specify one source at a time (no ~~`uplink cp cheesecake.jpg cheesecake.png sj://cakes`~~)
{% endhint %}

Output:

![](../../.gitbook/assets/upload\_file.png)

### Copy a file from a bucket to a local drive

When the `cp` command is used to copy a file from Storj DCS (download), the CLI first downloads the minimum number of pieces to reconstitute a file (typically 29 pieces), then re-encodes the pieces into a single file, and finally decrypts the file client-side.&#x20;

To copy a file from a project to a local drive, use:

{% tabs %}
{% tab title="Windows" %}
```
./uplink.exe cp sj://cakes/cheesecake.jpg ~/Downloads/
```
{% endtab %}

{% tab title="Linux" %}
```
uplink cp sj://cakes/cheesecake.jpg ~/Downloads/
```
{% endtab %}

{% tab title="macOS" %}
```
uplink cp sj://cakes/cheesecake.jpg ~/Downloads/
```
{% endtab %}
{% endtabs %}

![](../../.gitbook/assets/download\_object.png)

### Copy a local file into a bucket with an expiration date

The uploaded object can be set to expire at a certain time. After the expiration date, the file is no longer available and no longer will generate usage charges. To set an expiration date for a file when uploading it, you should use the `cp` command with the `--expires` flag:

{% tabs %}
{% tab title="Windows" %}
```
./uplink.exe cp --expires 2021-12-31T13:00:00+02:00 cheesecake.jpg sj://cakes
```
{% endtab %}

{% tab title="Linux" %}
```
uplink cp  --expires 2021-12-31T13:00:00+02:00 cheesecake.jpg sj://cakes
```
{% endtab %}

{% tab title="macOS" %}
```
uplink cp  --expires 2021-12-31T13:00:00+02:00 cheesecake.jpg sj://cakes
```
{% endtab %}
{% endtabs %}

The date is given in the `yyyy-mm-ddThh:mm:ssZhh:mm` format defined in ISO 8601.  `2021-12-31T13:00:00+02:00` reads "December, 31st at 1pm UTC+2". A date ending with "Z", such as `2021-12-31T13:00:00Z`, is in UTC.

The command above gives the following output:

![](../../.gitbook/assets/upload\_file.png)

### Copy an object with parallelism

If you have enough upstream bandwidth, you can use the multipart functionality to upload objects faster.

To increase upload speed, you can use the `cp` command with the `--parallelism 10`  flag (the number you can set according to your preferences and available upstream bandwidth):

{% tabs %}
{% tab title="Windows" %}
```
./uplink.exe cp --parallelism 10 cheesecake.jpg sj://cakes
```
{% endtab %}

{% tab title="Linux" %}
```
uplink cp --parallelism 10 cheesecake.jpg sj://cakes
```
{% endtab %}

{% tab title="macOS" %}
```
uplink cp --parallelism 10 cheesecake.jpg sj://cakes
```
{% endtab %}
{% endtabs %}

Since our sample object is small, you likely will not notice a difference.

![](../../.gitbook/assets/upload\_file.png)

It would be significantly different with big objects like videos or OS images etc. and for upstream bandwidth much greater than 100Mbps.

### Recursive copy of objects from local location to the bucket

You can recursively copy files:

{% tabs %}
{% tab title="Windows" %}
```
./uplink.exe cp --recursive ~/receipts sj://cakes/
```
{% endtab %}

{% tab title="Linux" %}
```
uplink cp --recursive ~/receipts sj://cakes/
```
{% endtab %}

{% tab title="macOS" %}
```
uplink cp --recursive ~/receipts sj://cakes/
```
{% endtab %}
{% endtabs %}

Sample output:

```
upload /home/user/receipts/cheescake.jpg to sj://cakes/cheescake.jpg
upload /home/user/receipts/pancake.jpg to sj://cakes/pancake.jpg
```

### Copy an object from one location to another within Storj DCS

It is possible to copy a file from one Storj DCS location to another Storj DCS location within the same project.

When the `cp` command is used to copy a file from one Storj DCS location to another Storj DCS location, the object will be copied entirely on the "server" side - **this will not count against your egress limits, as the object is not being downloaded**.

{% hint style="info" %}
You need to have at least version 1.54.1 of Uplink installed to support server-side copy
{% endhint %}

First, to create a new bucket, we will use the `mb` command, as copying is possible only to an existing bucket.

{% tabs %}
{% tab title="Windows" %}
```
./uplink.exe mb sj://new-recipes
```
{% endtab %}

{% tab title="Linux" %}
```
uplink mb sj://new-recipes
```
{% endtab %}

{% tab title="macOS" %}
```
uplink mb sj://new-recipes
```
{% endtab %}
{% endtabs %}

```powershell
Bucket new-recipes created
```

{% hint style="info" %}
Nested buckets are not supported, but you can use prefixes, as they would act almost like subfolders.
{% endhint %}

Now, to copy a file from a bucket within a project to another bucket in the same project with prefix `cakes`, use:

{% tabs %}
{% tab title="Windows" %}
```
./uplink.exe cp sj://cakes/cheesecake.jpg sj://new-recipes/cakes/cheesecake.jpg
```
{% endtab %}

{% tab title="Linux" %}
```
uplink cp sj://cakes/cheesecake.jpg sj://new-recipes/cakes/cheesecake.jpg
```
{% endtab %}

{% tab title="macOS" %}
```
uplink cp sj://cakes/cheesecake.jpg sj://new-recipes/cakes/cheesecake.jpg
```
{% endtab %}
{% endtabs %}

Sample Output:

```
upload sj://cakes/cheesecake.jpg sj://new-recipes/cakes/cheesecake.jpg
```

There is no progress bar shown since nothing was downloaded or uploaded, as the copying happens on the "server" side (within a Storj DCS project.)
