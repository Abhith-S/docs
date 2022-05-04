# access import

This command allows you to import the Access Grant to Uplink.

## Usage

{% tabs %}
{% tab title="Windows" %}
```
./uplink.exe access import [flags] <name> <access|filename>
```
{% endtab %}

{% tab title="Linux" %}
```
uplink access import [flags] <name> <access|filename>
```
{% endtab %}

{% tab title="macOS" %}
```
uplink access import [flags] <name> <access|filename>
```
{% endtab %}
{% endtabs %}

## Arguments

| Argument             | Description                                  |
| -------------------- | -------------------------------------------- |
| `<name>`             | Name to save the access as                   |
| `<access\|filename>` | Serialized access value or file path to save |

## Flags

| Flag          | Description                                        |
| ------------- | -------------------------------------------------- |
| `-f, --force` | Force overwrite an existing saved access           |
| `--use`       | Switch the default access to the newly created one |

### Global flags

| Global flags          | Description                                   |
| --------------------- | --------------------------------------------- |
| `--config-dir string` | Directory that stores the configuration       |
| `--help`, `-h`        | prints help for the command                   |
| `--advanced`          | when used with -h, prints advanced flags help |

## Examples

Please [Create an Access Grant from satellite UI](../../../getting-started/quickstart-uplink-cli/uploading-your-first-object/create-first-access-grant.md) or [Create an Access Grant from CLI](../../../getting-started/quickstart-uplink-cli/generate-access-grants-and-tokens/generate-a-token.md) before proceeding.

### Import Access Grant from the file

Save the created Access Grant to the file `access.txt`. As result, this command will import the Access Grant from the file to the access with the specified name into Uplink.

{% tabs %}
{% tab title="Windows" %}
```
./uplink.exe access import main access.txt
```
{% endtab %}

{% tab title="Linux" %}
```
uplink access import main access.txt
```
{% endtab %}

{% tab title="macOS" %}
```
uplink access import main access.txt
```
{% endtab %}
{% endtabs %}

```
Imported access "main" to "/home/user/.config/storj/uplink/access.json"
```

### Import Access Grant from the console

As result, the Access Grant will be imported from the console to the access with the specified name.

{% tabs %}
{% tab title="Windows" %}
```
./uplink.exe access import main 18fglgkoitmfvkogmoitr....
```
{% endtab %}

{% tab title="Linux" %}
```
uplink access import main 18fglgkoitmfvkogmoitr....
```
{% endtab %}

{% tab title="macOS" %}
```
uplink access import main 18fglgkoitmfvkogmoitr....
```
{% endtab %}
{% endtabs %}

```
Imported access "main" to "/home/user/.config/storj/uplink/access.json"
```

### Import Access Grant and replace the existing access

You will import the created access grant to uplink as a named access

{% tabs %}
{% tab title="Windows" %}
```
./uplink.exe access import main access.txt --force
```
{% endtab %}

{% tab title="Linux" %}
```
uplink access import main access.txt --force
```
{% endtab %}

{% tab title="macOS" %}
```
uplink access import main access.txt --force
```
{% endtab %}
{% endtabs %}

```
Imported access "main" to "/home/user/.config/storj/uplink/access.json
```
