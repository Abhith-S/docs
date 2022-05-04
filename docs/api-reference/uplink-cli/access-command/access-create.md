# access create

This command allow you to create and print the Access Grant to `stdout`, export it to a file or import it as an access to Uplink.

## Usage

{% tabs %}
{% tab title="Windows" %}
```
./uplink.exe access create [flags]
```
{% endtab %}

{% tab title="Linux" %}
```
uplink access create [flags]
```
{% endtab %}

{% tab title="macOS" %}
```
uplink access create [flags]
```
{% endtab %}
{% endtabs %}

## Flags

| Flag                         | Description                                                                        |
| ---------------------------- | ---------------------------------------------------------------------------------- |
| `--passphrase-stdin`         | If set, the passphrase is read from `stdin`, and all other values must be provided |
| `--satellite-address string` | Satellite address from satellite UI (prompted if unspecified)                      |
| `--api-key string`           | API key from satellite UI (prompted if unspecified)                                |
| `--import-as string`         | Import the access as this name                                                     |
| `--export-to string`         | Export the access to this file path                                                |
| `-f, --force`                | Force overwrite an existing saved access                                           |
| `--use`                      | Switch the default access to the newly created one                                 |

### Global flags

| Global flags          | Description                                   |
| --------------------- | --------------------------------------------- |
| `--config-dir string` | Directory that stores the configuration       |
| `--help`, `-h`        | prints help for the command                   |
| `--advanced`          | when used with -h, prints advanced flags help |

## Examples

### Create an Access Grant without prompts

As result it will print the created access grant to `stdout`.

{% tabs %}
{% tab title="Windows" %}
```
echo "SuperPassword" | ./uplink.exe access create --satellite-address 12tRQrMTWUWwzwGh18i7Fqs67kmdhH9t6aToeiwbo5mfS2rUmo@us2.storj.io:7777 --api-key 1dfJ354T.... --passphrase-stdin
```
{% endtab %}

{% tab title="Linux" %}
```
echo "SuperPassword" | uplink access create --satellite-address 12tRQrMTWUWwzwGh18i7Fqs67kmdhH9t6aToeiwbo5mfS2rUmo@us2.storj.io:7777 --api-key 1dfJ354T.... --passphrase-stdin
```
{% endtab %}

{% tab title="macOS" %}
```
echo "SuperPassword" | uplink access create --satellite-address 12tRQrMTWUWwzwGh18i7Fqs67kmdhH9t6aToeiwbo5mfS2rUmo@us2.storj.io:7777 --api-key 1dfJ354T.... --passphrase-stdin
```
{% endtab %}
{% endtabs %}

```
18yMsZpg6ZQdz........
```

### Create an Access Grant and export to the file

You will export the created access grant to the file.

{% tabs %}
{% tab title="Windows" %}
```
./uplink.exe access create --export-to access.txt
```
{% endtab %}

{% tab title="Linux" %}
```
uplink access create --export-to access.txt
```
{% endtab %}

{% tab title="macOS" %}
```
uplink access create --export-to access.txt
```
{% endtab %}
{% endtabs %}

```
Exported access to: /home/user/access.txt
```

### Create an Access Grant and import it to Uplink

You will import the created access grant to Uplink as a named access.

{% tabs %}
{% tab title="Windows" %}
```
./uplink.exe access create --import-as us2
```
{% endtab %}

{% tab title="Linux" %}
```
uplink access create --import-as us2
```
{% endtab %}

{% tab title="macOS" %}
```
uplink access create --import-as us2
```
{% endtab %}
{% endtabs %}

```
Imported access "us2" to "/home/user/.config/storj/uplink/access.json"
```

### Create an Access Grant and replace the existing access

You will import the created access grant to uplink as a named access and replace it if it exists.

{% tabs %}
{% tab title="Windows" %}
```
./uplink.exe access create --import-as us2 --force
```
{% endtab %}

{% tab title="Linux" %}
```
uplink access create --import-as us2 --force
```
{% endtab %}

{% tab title="macOS" %}
```
uplink access create --import-as us2 --force
```
{% endtab %}
{% endtabs %}

```
Imported access "us2" to "/home/user/.config/storj/uplink/access.json"
```
