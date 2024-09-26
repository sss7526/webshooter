# webshooter
CLI utility to save full screenshots and/or PDFs of target webpages using a headless browser.

- Validates target URLs for valid URL syntax, prepends https:// if target is passed without protocol prefix (http:// for localhost connection)
- Supports query strings and subdirectories/files in target URLs
- Verbose mode to print network requests and responses and allowed/blocked request statuses to terminal
- Blocks residual requests from being sent with wildcard keyword matches in URLs (plan to expand and support custom keyword blocks)
- Passes generic UA string and referer in network request headers
- Options to either save screenshot of rendered page or save to PDF (or both) with timestamped filenames
- Supports stacking short option flags (ex: -ipPv for image, pdf, tor proxy, and verbose modes respectively)
- Optionally read target list from file
- Optionally connect to target (including .onion sites) over Tor proxy (Requires Tor to be installed and running)

- Alternative query mode to send query strings to specified search engine, returns and deduplicates search results from first 4 pages

# Dependencies
You need either Chrome or Chromium browser installed for this to work
```bash
sudo snap install chromium
```

To use the Tor proxy option, install and enable Tor, and configure control port for circuit resets
```bash
sudo apt install tor
```

Start/stop/enable Tor
```bash
sudo systemctl start/stop tor     #To start or stop tor
sudo systemctl enable/disable tor #To enable or disable tor from running automatically
sudo systemctl status tor         #To check if its running or not
```

Configure Tor control port for circuit resets
```bash
sudo nano /etc/tor/torrc
```
Add these lines to the bottom of the torrc file:
```plaintext
ControlPort 9051
CookieAuthentication 0
```
Restart Tor after editing the torrc file

# Installation (deb package)

[Download the latest release here](https://github.com/sss7526/webshooter/releases/latest)
```bash
sudo apt install ./webshooter_x.x.x_amd64.deb
```

# Usage
```
webshooter --help (-h)
```

output
```plaintext

Webshooter
Version: 1.2.0

CLI utility to take screenshots and save PDFs of target web pages.
Also can run queries against search engines and return the search result URLs.
Blocks residual network requests to undesired URLs (trackers, ads, etc).

Usage:
    -e, --engine: search engine to query against (google and local instances of whoogle currently supported)
    -f, --file: Reads in target URLs from file. Cannot be used with --targets (-t) flag
    -i, --image: If specified, saves screenshot of target webpage as a PNG
    -p, --pdf: If specified, saves PDF copy of target webpage
    -q, --query: Sends query string to specified search engine
    -t, --targets: Space separated list of one or more target URLs
    -P, --proxy: If specified, connect to target over Tor (Tor must be installed and running)
    -T, --translate: If specified, translates the target webpage before capture
    -v, --verbose: Increase verbosity, shows http requests/responses and allowed/blocked status
```