# access remove

This command allows you to remove the access from local store of Uplink.

## Usage

{% tabs %}
{% tab title="Windows" %}
```
./uplink.exe access remove <name>
```
{% endtab %}

{% tab title="Linux" %}
```
uplink access remove <name>
```
{% endtab %}

{% tab title="macOS" %}
```
uplink access remove <name>
```
{% endtab %}
{% endtabs %}

## Arguments

| Argument | Description           |
| -------- | --------------------- |
| `<name>` | Access name to delete |

### Global flags

| Global flags          | Description                                   |
| --------------------- | --------------------------------------------- |
| `--config-dir string` | Directory that stores the configuration       |
| `--help`, `-h`        | prints help for the command                   |
| `--advanced`          | when used with -h, prints advanced flags help |

## Example

You need to have an access in the local store of Uplink before proceeding. See [`uplink access create`](access-create.md), [`uplink access import`](access-import.md) and [`uplink setup`](../setup-command.md) commands for information how to create/import/setup an access.

{% hint style="info" %}
If you want to remove the current access, you need to switch to another before proceeding, using the [`uplink access use`](access-use.md) command.
{% endhint %}

{% tabs %}
{% tab title="Windows" %}
```
./uplink.exe access remove us2
```
{% endtab %}

{% tab title="Linux" %}
```
uplink access remove us2
```
{% endtab %}

{% tab title="macOS" %}
```
uplink access remove us2
```
{% endtab %}
{% endtabs %}

```
Removed access "us2" from "/home/user/.config/storj/uplink/access.json"
```
