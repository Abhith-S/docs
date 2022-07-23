---
description: Storj IPFS Pinning Service (Beta)
---

# Storj IPFS Pinning Service (Beta)

## Prerequisites

{% hint style="info" %}
You should have received an email that you have been invited to the beta which will include credentials to access the service. If you have not signed up yet, you can [Join the beta](https://landing.storj.io/permanently-pin-with-storj-dcs).
{% endhint %}

The Storj IPFS Pinning service consists of an HTTP endpoint for uploading and pinning content, and an IPFS Gateway that serves the pinned content over IPFS and HTTP. Details on smart contract pinning will be made available in the future.

## How to pin with Storj IPFS

All content uploaded to the Storj IPFS service via the HTTP endpoint below is pinned. Examples are given in cURL and JavaScript, but could be done from any programming language or with existing IPFS client bindings for a given programming language, such as the [IPFS HTTP Client library](https://www.npmjs.com/package/ipfs-http-client) for npm.

### HTTP Upload endpoint

Uploading content follows the [IPFS HTTP RPC for /api/v0/add](https://docs.ipfs.io/reference/http/api/#api-v0-add) with two small differences:

1. The only optional argument supported is `wrap-with-directory`&#x20;
2. You must specify the credentials given when invited to participate in the beta as _HTTP basic authentication._

{% hint style="warning" %}
_**This is not the same as your Storj DCS username and password. Do not use your Storj DCS username and password to try and use the IPFS Pinning Service.**_
{% endhint %}

#### Example using cURL&#x20;

For example, this is how it would work with cURL and a file you wanted to pin called `/path/file.extension`. Please replace _**ipfs\_beta\_user**_ and _**ipfs\_beta\_password**_ with the beta credentials you received when accepted into the beta.

```
curl -u ipfs_beta_user:ipfs_beta_password -X POST -F file=@/path/file.extension https://www.storj-ipfs.com/api/v0/add
```

{% hint style="info" %}
The '`@`' before the file path is required for the upload to work properly. For example, if the file you wanted to upload was `/home/hello/hi.jpg`, the curl argument would be `file=@/home/hello/hi.jpg`.
{% endhint %}

#### Example using JavaScript

You'll need [Node.js](https://nodejs.org) installed on your local system for this example.

1\. Make a new JavaScript project.

Create a new directory and use `npm` to create a new project:

```shell
mkdir storj-ipfs-quickstart
cd storj-ipfs-quickstart
npm init 
```

NPM will ask a few questions about your project and create a `package.json` file.

2\. Add the `axios` HTTP client to your project dependencies.

Install the latest version of the `axios` package:

```shell
npm install axios
```

3\. Create a file called `upload.js` and open it with your code editor.

Below is the code we need to upload a file and pin it on the Storj IPFS pinning service.

Paste in the code below and read through it. Feel free to remove the comments - they're just there to highlight what's going on.

{% code title="upload.js" %}
```javascript
// The 'axios' module gives a promised-based HTTP client.
const axios = require('axios');

// The 'fs' built-in module provides us access to the file system.
const fs = require('fs');

// The 'form-data' built-in module helps us submit forms and file uploads
// to other web applications.
const FormData = require('form-data');

/**
  * Uploads a file from `filepath` and pins it to the Storj IPFS pinning service.
  * @param {string} username your username for the Storj IPFS pinning service
  * @param {string} password your password for the Storj IPFS pinning service
  * @param {string} filepath the path to the file
  */
async function pinFileToIPFS(username, password, filepath) {
    // The HTTP upload endpoint of the Storj IPFS pinning service
    const url = `https://www.storj-ipfs.com/api/v0/add`;

    // Create a form with the file to upload
    let data = new FormData();
    data.append('file', fs.createReadStream(filepath));

    // Execute the Upload request to the Storj IPFS pinning service
    return axios.post(url,
        data,
        {
            auth: {
                username: username,
                password: password
            },
            headers: {
                'Content-Type': `multipart/form-data; boundary= ${data._boundary}`,
            },
            // These arguments remove any client-side upload size restrictions
            maxContentLength: Infinity,
            maxBodyLength: Infinity,
        },
    ).then(function (response) {
        console.log(response)
    }).catch(function (error) {
        console.error(error)
    });
};

/**
 * The main entry point for the script that checks the command line arguments and
 * calls pinFileToIPFS.
 * 
 * To simplify the example, we don't do any fancy command line parsing. Just three
 * positional arguments for imagePath, name, and description
 */
 async function main() {
    const args = process.argv.slice(2)
    if (args.length !== 3) {
        console.error(`usage: ${process.argv[0]} ${process.argv[1]} <username> <password> <filepath>`)
        process.exit(1)
    }

    const [username, password, filepath] = args
    const result = await pinFileToIPFS(username, password, filepath)
    console.log(result)
}

/**
 * Don't forget to actually call the main function!
 * We can't `await` things at the top level, so this adds
 * a .catch() to grab any errors and print them to the console.
 */
main()
  .catch(err => {
      console.error(err)
      process.exit(1)
  })
```
{% endcode %}

4\. Run your script with node.

You should now be able to run the script and give it the path to a file. You'll also need to supply your username and password for the Storj IPFS pinning service.

```
node upload.js ipfs_beta_user ipfs_beta_password /path/file.extension
```

Please replace _**ipfs\_beta\_user**_ and _**ipfs\_beta\_password**_ with the beta credentials you received when accepted into the beta.

## How to retrieve pinned Objects

Any content uploaded is automatically pinned and retrievable through any software that supports IPFS natively via its CID like [IPFS Desktop](https://github.com/ipfs/ipfs-desktop) or [IPFS CLI](https://docs.ipfs.io/how-to/command-line-quick-start/). Some browsers like [Brave](https://brave.com/ipfs-support/) include support, as well as some IPFS programs.

For those applications that do not support IPFS natively, you can use any [public IPFS gateway](https://docs.ipfs.io/concepts/ipfs-gateway/), or the Storj IPFS Gateway as described below.

### HTTP via Storj IPFS Gateway

For best performance, we have provided a Storj IPFS Gateway. This gateway will only host content pinned to Storj DCS, so it is not like other public IPFS gateways.

You can construct a link like this:

```
https://www.storj-ipfs.com/ipfs/<cid>
```

In cases where the gateway is unable to retrieve a given CID (e.g., returns a 404 not found error), please double check that you are using the correct CID and that it was uploaded to the Storj IPFS service.

### Peering your IPFS node with Storj pinning nodes

If you run your own IPFS node that retrieves a lot of data pinned on the Storj IPFS Pinning Service, you may want to prioritize the connections to the Storj IPFS nodes. This will improve the download performance by bypassing the DHT lookup for the data.

Prioritizing connections to certain peers is called **Peering**, and you can tell IPFS which peers to prioritize by editing the [`Peering` configuration](https://docs.ipfs.io/how-to/configure-node/#peering) in your IPFS config file.

To _peer_ with Storj IPFS nodes, you could update the `Peering` section of your config to include their ID and addresses:

```json
{
  "Peering": {
    "Peers": [
      {
        "ID": "12D3KooWFFhc8fPYnQXdWBCowxSV21EFYin3rU27p3NVgSMjN41k",
        "Addrs": ["/ip4/5.161.92.43/tcp/4001", "/ip4/5.161.92.43/udp/4001/quic", "/ip6/2a01:4ff:f0:3b1e::1/tcp/4001", "/ip6/2a01:4ff:f0:3b1e::1/udp/4001/quic"]
      },
      {
        "ID": "12D3KooWSW4hoHmDXmY5rW7nCi9XmGTy3foFt72u86jNP53LTNBJ",
        "Addrs": ["/ip4/5.161.55.227/tcp/4001", "/ip4/5.161.55.227/udp/4001/quic", "/ip6/2a01:4ff:f0:1e5a::1/tcp/4001", "/ip6/2a01:4ff:f0:1e5a::1/udp/4001/quic"]
      },
      {
        "ID": "12D3KooWSDj6JM2JmoHwE9AUUwqAFUEg9ndd3pMA8aF2bkYckZfo",
        "Addrs": ["/ip4/5.161.92.36/tcp/4001", "/ip4/5.161.92.36/udp/4001/quic", "/ip6/2a01:4ff:f0:3764::1/tcp/4001", "/ip6/2a01:4ff:f0:3764::1/udp/4001/quic"]
      }
    ]
  }
}
```
