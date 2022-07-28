---
description: >-
  Create an Access Grant in the Uplink CLI with Satellite and API Key info from
  the Satellite Admin Console
---

# Create Access Grant in CLI

1. [You need to have a satellite account and installed Uplink CLI](../prerequisites.md).
2. To start, proceed through the initial steps of creating a new Access Grant.
3. Navigate to [**Access**](../../satellite-developer-account/access-grants.md) page and click the **Create Keys for CLI** link (rightmost option).

![](<../../../.gitbook/assets/image (24).png>)

4\. Provide name, permissions and optionally buckets, select **Create Keys**.

![](<../../../.gitbook/assets/image (12).png>)

5\. Copy and save the **Satellite Address** and **API Key** in a safe place or download them as they will only appear once.

![](<../../../.gitbook/assets/image (21).png>)

6\. Make sure you've already [downloaded the Uplink CLI](../../../downloads/download-uplink-cli.md) and run `uplink setup`.

{% tabs %}
{% tab title="Windows" %}
```
./uplink.exe setup
```
{% endtab %}

{% tab title="Linux" %}
```
uplink setup
```
{% endtab %}

{% tab title="macOS" %}
```
uplink setup
```
{% endtab %}
{% endtabs %}

{% hint style="info" %}
For anyone who has previously configured an Uplink, please use a named access. If you want to replace the default access, you need to either [Create an Access Grant](../uploading-your-first-object/create-first-access-grant.md) and use the [`uplink access import`](../../../api-reference/uplink-cli/access-command/access-import.md#import-access-grant-and-replace-the-existing-access) command with `--force` flag to import it, or use the [`uplink access create --import-to <name>`](../../../api-reference/uplink-cli/access-command/access-create.md#create-an-access-grant-and-replace-the-existing-access) command with `--force` flag to create an Access Grant in CLI and import it to the specified access in the local store of Uplink.
{% endhint %}

7\. Follow the prompts. When asked for your API Key, enter it (you should have saved it in step 4 above).

8\. Generate the Access Grant by running `uplink share` with no restrictions.

{% hint style="info" %}
&#x20;If you chose an access name, you'll need to specify it in the following command as `--access=name`
{% endhint %}

{% tabs %}
{% tab title="Windows" %}
```
./uplink.exe share --readonly=false
```
{% endtab %}

{% tab title="Linux" %}
```
uplink share --readonly=false
```
{% endtab %}

{% tab title="macOS" %}
```
uplink share --readonly=false
```
{% endtab %}
{% endtabs %}

{% hint style="danger" %}
Keep your full-rights Access Grant secret, it contains the encryption key and will enable uploading, downloading or deleting your data from the entire project!
{% endhint %}

9\. Your Access Grant should have been output.&#x20;

{% hint style="success" %}
The alternative for using the `uplink setup` command and then `uplink share` is to use the [`uplink access create`](../../../api-reference/uplink-cli/access-command/access-create.md) command instead, it will print the Access Grant right away.
{% endhint %}
