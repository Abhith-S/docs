---
description: How to manage production media assets in Iconik using Storj DCS.
---

# Iconik Integration Guide

## Introduction

**Iconik is a cloud-native solution that gathers and organizes media securely from any storage.**

**Main site**: [https://www.iconik.io/](https://www.iconik.io/)

App link: [https://app.iconik.io/](https://app.iconik.io/)

## Configure Iconik to use Storj DCS

To set up Storj on Iconik, first click the “Admin” option at the top left of the screen.

![](https://lh3.googleusercontent.com/LYemothVK\_w8Wmo6vf\_P0rNw83g-3ZH2ooU7innQZ1t-YZgWSGswcOstfMAHBmSJfNzbxNDURwl5iJ38d3tOG5XatkVDXfe0\_6tG1ICWmkSDdx8TdYmiGBwKst7tTHqPNJe2lsaUgTvlaM9ULQ)\


Clicking `Admin` should change the icons which run vertically on the left side of the screen.  Click the `Storages` icon from these icons.

![](https://lh5.googleusercontent.com/efNVbSw7wB4Xq\_UmBv2BzZn16\_6qfdMg-0dskhdzacNPk5YuBoIp1oJ3zCGySFA\_I\_-TUywK4CRgbQbL4iG9HEYg1L0LoyG2zkXkzdQGXRpPlILCk2osKXwsLkqq5v6H4ZzRvy3xO3lsScxkyQ)



In the upper right corner of the screen, click the `New Storage` option.

![](https://lh5.googleusercontent.com/q\_os0a7mgrwfVuao-NrJNt\_hWWQF\_6tURZkPedN--YxzvHtOvqb6uiZQqaf9WiawGKZ9lBwxCdf9F6ggoYp\_djAHs4ma3Jehq\_q-asdeNNP7NRxZO6-\_HIx82XnVTcXzNmpJswgxceQnA4fRrg)



An `Add new storage` pop-up should appear. &#x20;

<img src="https://lh6.googleusercontent.com/P-UDXei2L-mcx1kgV3YvlGzIYRQMqYGEW5w92GFCyhDSyzXkbEhIEPqZc2e-PSSvGT_E5kTjDBv9-9cabirKPtopDrJ8MGjMPLfsZDArevZtEKi4aRWjkZYqisLjrvO8_Jd1cJlY-DFJqgkJzw" alt="" data-size="original">



Selecting `Amazon S3` from the `Storage Type` drop-down list will change the pop-up dialog to include many more options.  Enter the appropriate values into the corresponding inputs.

Name:  \<user chosen>

Storage Purpose:  \<user chosen, often “Files”>

Storage Type:  “Amazon S3”

Access Key:  \<value from Storj DCS website or “uplink” command>

Secret Key:  \<value from Storj DCS website or “uplink” command>

Bucket:  \<user chosen, but the bucket must already exist>

Path:  \<user chosen, optional, often empty>

Region:  “storj” \<any non-empty value works>

Endpoint:  “https://gateway.storjshare.io/”

Use Acceleration: \<unselected>

Add unique id to the file names:  \<user preference>

Use Acceleration: \<unselected>

Add unique id to the file names:  \<user preference>

<img src="https://lh4.googleusercontent.com/BYgovgW7Q6wYm3I5PG2HCNCxFQ_I8FEJDkDfGXW41CHyH1D415zF4sZgrNJVU9hZVPQi53bjIwYfQb9jia1xLibyjMYvDsHcvu6n7VAjzZRtRHtzlN3dPNkVTbKEYCMdkjU7M_y4SfgZyzK7kw" alt="" data-size="original">

\


Click `Close` on the subsequent pop-up to continue.

![](https://lh6.googleusercontent.com/pRWFa24CEzsz5INVSSu6Ck-I8YoQ-QMTSyH81qOXDuEST\_IIFBjTmv2d1hLiVhrMhkhCtJ9kAxf6RnjJdwFMOetnB3IjLeHDHvVC98iYxwyEVprL5C9pOB0zHDk8L1PFXN2GF-yVUXs8IBFlCQ)\


Your Iconik Storage configuration should now appear in the `Storages` list view.

![](https://lh5.googleusercontent.com/cbr7-YQwJNBdIZr5UMuMpHuPObjJ59-ME9va1DBmOz4AzT23WEYMY1esgEHk6R\_m9f9bcoxNn5FIPSznkz4o\_YCJ0obsEwllQp2jPBdRm93d01vIU0bdwvghRcKxKuNVcYh-9wlA5paT5\_7w-A)



Clicking on the newly created Storage option will allow you to configure it.  In addition, this page contains a `Browser` option.  Checking that video assets are visible in the Browser is an easy way to verify that your connection to Storj is set up correctly.

![](https://lh6.googleusercontent.com/VolImslMr\_QzfBxLNfm7fZCsLDJyps1lFVmJP0vmRKfsG3TmMAVwP8aN4YQgDjLy-urt7TuPwwg9RafNqImXNCD05VX\_1QjZJlJpYD\_x9dj2i47eRSE8D7Ofupjjj5qokB7aaJR7HN3Fobye-Q)

\
\


\
