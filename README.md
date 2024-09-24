# webshooter
CLI utility to save full screenshots and/or PDFs of target webpages using a headless browser.

- Validates target URLs for valid URL syntax, prepends https:// if target is passed without protocol prefix (http:// for localhost connection)
- Supports query strings and subdirectories/files in target URLs
- Verbose mode to print network requests and responses and allowed/blocked request statuses to terminal
- Blocks residual requests from being sent with wildcard keyword matches in URLs (plan to expand and support custom keyword blocks)
- Passes generic UA string and referer in network request headers
- Options to either save screenshot of rendered page or save to PDF (or both) with timestamped filenames

# Dependencies
You need either Chrome or Chromium browser installed for this to work
```
sudo snap install chromium
```

# Installation (deb package)

```
wget https://github.com/sss7526/webshooter/releases/latest/download/1.0.0/webshooter_1.0.0_amd64.deb
sudo apt install ./webshooter_1.0.0_amd64.deb
```

# Usage
```
webshooter --help (-h)
```
