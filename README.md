# BatSecrets

BatSecrets is a stealthy detective in the shadows of the web, specialized in revealing concealed secrets within URLs. Its keen algorithm meticulously scans for sensitive information like API keys, access tokens, and other clandestine data, empowering you to fortify your applications and infrastructure proactively. With BatSecrets by your side, stay one step ahead in the battle against security threats, safeguarding your digital assets with precision and foresight.

# Installation

To install BatSecrets, simply use go install:

```
go install github.com/MrBrooks89/BatSecrets/cmd/BatSecrets@latest
```

This command will download and install the latest version of BatSecrets to your Go bin directory.

# Usage

Use the following command to check a single URL for secrets inside JavaScript file:

```
BatSecrets -u https://example.com/js/app.js
```

To check a list of URLs with JavaScript files, provide a file containing the list:

```
BatSecrets -l path/to/url-jslist.txt
```

Optional flags:

```
-c int
        Limit the concurrency of goroutines (default 100)
-l string
        File containing a list of URLs to check for secrets
-u string
        URL to check for secrets
-v      Print verbose output
```

Example:
```
BatSecrets -l urls-jslist.txt -c 50
```

BatSecrets will scan the provided URLs for secrets and display any findings. Each secret is reported with its name, description, URL, and the matched string.

# Secrets Detected

BatSecrets includes detection for a wide range of secrets, including but not limited to:

    AWS access keys
    Facebook client IDs and secret keys
    Google API keys and OAuth access keys
    Github personal access tokens and OAuth access tokens
    Slack tokens and webhooks
    Twitter client IDs and secret keys
    And many more...

Please note that BatSecrets is intended for educational and testing purposes only. Use it responsibly and with permission.

# License

This project is licensed under the MIT License.
# BatSecrets
# BatSecrets
