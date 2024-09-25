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

# Dependencies
You need either Chrome or Chromium browser installed for this to work
```
sudo snap install chromium
```

To use the Tor proxy option, install and enable Tor, and configure control port for circuit resets
```
sudo apt install tor
```

Start/stop/enable Tor
```
sudo systemctl start/stop tor     #To start or stop tor
sudo systemctl enable/disable tor #To enable or disable tor from running automatically
sudo systemctl status tor         #To check if its running or not
```

Configure Tor control port for circuit resets
```
sudo nano /etc/tor/torrc
```
Add these lines to the bottom of the torrc file:
```
ControlPort 9051
CookieAuthentication 0
```
Restart Tor after editing the torrc file

# Installation (deb package)

[Download the latest release here](https://github.com/sss7526/webshooter/releases/latest)
```
sudo apt install ./webshooter_1.0.0_amd64.deb
```

# Usage
```
webshooter --help (-h)
```
