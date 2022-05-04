# access revoke

This command allows you to revoke the access.

## Usage

{% tabs %}
{% tab title="Windows" %}
```
./uplink.exe access [flags] <revokee>
```
{% endtab %}

{% tab title="Linux" %}
```
uplink access [flags] <revokee>
```
{% endtab %}

{% tab title="macOS" %}
```
uplink access [flags] <revokee>
```
{% endtab %}
{% endtabs %}

## Arguments

| Argument    | Description                    |
| ----------- | ------------------------------ |
| `<revokee>` | Access name or value to revoke |

### Flags

| Flag              | Description                                |
| ----------------- | ------------------------------------------ |
| `--access string` | Access name or value performing the revoke |

### Global flags

| Global flags          | Description                                   |
| --------------------- | --------------------------------------------- |
| `--config-dir string` | Directory that stores the configuration       |
| `--help`, `-h`        | prints help for the command                   |
| `--advanced`          | when used with -h, prints advanced flags help |

## Examples

### Revoke a stored access

{% hint style="warning" %}
If you want to revoke the current access, you need to switch to a different access with the [`uplink access use`](access-use.md) command before proceeding.
{% endhint %}

{% tabs %}
{% tab title="Windows" %}
```
./uplink.exe access revoke us1-ro
```
{% endtab %}

{% tab title="Linux" %}
```
uplink access revoke us1-ro
```
{% endtab %}

{% tab title="macOS" %}
```
uplink access revoke us1-ro
```
{% endtab %}
{% endtabs %}

```
Revoked access "us1-ro"
```

### Revoke an access grant

{% tabs %}
{% tab title="Windows" %}
```
./uplink.exe access revoke 19hFrjmsi...
```
{% endtab %}

{% tab title="Linux" %}
```
uplink access revoke 19hFrjmsi...
```
{% endtab %}

{% tab title="macOS" %}
```
uplink access revoke 19hFrjmsi...
```
{% endtab %}
{% endtabs %}

```
Revoked access "19hFrjmsi..."
```
