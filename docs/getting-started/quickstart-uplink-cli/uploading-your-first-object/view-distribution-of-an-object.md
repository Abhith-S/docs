---
description: >-
  Visualize the geographic distribution of your object across the network of
  nodes.
---

# View Distribution of an Object

Check [Prerequisites](../prerequisites.md).

You can view the geographic distribution of your object and generate a shareable URL via the Link Sharing Service. Run the `uplink share --url` command below.

See [here](../../../api-reference/uplink-cli/share-command.md#link-sharing) for specifications on how to select an auth region and restrict the `uplink share --url` command.

{% tabs %}
{% tab title="Windows" %}
```
./uplink.exe share --url --not-after=+2h sj://cakes/cheesecake.jpg
```
{% endtab %}

{% tab title="Linux" %}
```
uplink share --url --not-after=+2h sj://cakes/cheesecake.jpg
```
{% endtab %}

{% tab title="macOS" %}
```
uplink share --url --not-after=+2h sj://cakes/cheesecake.jpg
```
{% endtab %}
{% endtabs %}

Copy the URL that is returned by the `uplink share --url` command and paste into your browser window.

```
=========== ACCESS RESTRICTIONS ==========================================================
Download  : Allowed
Upload    : Disallowed
Lists     : Allowed
Deletes   : Disallowed
NotBefore : No restriction
NotAfter  : 2022-03-01 09:56:13
Paths     : sj://cakes/cheesecake.jpg
=========== SERIALIZED ACCESS WITH THE ABOVE RESTRICTIONS TO SHARE WITH OTHERS ===========
Access    : 1Dv4...
========== CREDENTIALS ===================================================================
Access Key ID: jvw3fmzqyg2cvxm27qishw6y4qka
Secret Key   : ...
Endpoint     : https://gateway.storjshare.io
Public Access:  true
=========== BROWSER URL ==================================================================
REMINDER  : Object key must end in '/' when trying to share recursively
URL       : https://link.storjshare.io/s/juexo54k2db7lt5fawuqkupqkcfa/cakes/cheesecake.jpg
```

This is a real distribution of your file's pieces that you uploaded to the network. You can share this file with anyone you'd like.

![](<../../../.gitbook/assets/image (131) (1) (1).png>)
