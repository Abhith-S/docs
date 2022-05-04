# access use

This command allows you to switch the current access for Uplink.

## Usage

{% tabs %}
{% tab title="Windows" %}
```
./uplink.exe access use <access>
```
{% endtab %}

{% tab title="Linux" %}
```
uplink access use <access>
```
{% endtab %}

{% tab title="macOS" %}
```
uplink access use <access>
```
{% endtab %}
{% endtabs %}

## Arguments

| Argument   | Description        |
| ---------- | ------------------ |
| `<access>` | Access name to use |

### Global flags

| Global flags          | Description                                   |
| --------------------- | --------------------------------------------- |
| `--config-dir string` | Directory that stores the configuration       |
| `--help`, `-h`        | prints help for the command                   |
| `--advanced`          | when used with -h, prints advanced flags help |

## Example

You need to have more than one accesses in the local store of Uplink before proceeding. See [`uplink access create`](access-create.md), [`uplink access import`](access-import.md) and [`uplink setup`](../setup-command.md) commands for information how to create/import/setup an access.

{% tabs %}
{% tab title="Windows" %}
```
./uplink.exe access use us1
```
{% endtab %}

{% tab title="Linux" %}
```
uplink access use us1
```
{% endtab %}

{% tab title="macOS" %}
```
uplink access use us1
```
{% endtab %}
{% endtabs %}

```
Switched default access to "us1"
```
