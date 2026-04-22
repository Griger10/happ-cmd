# happcmd

[![CI](https://img.shields.io/badge/CI-passing-brightgreen)](https://github.com/Griger10/happcmd/actions)
[![codecov](https://codecov.io/gh/Griger10/happ-cmd/graph/badge.svg?token=FQSCBFFR8A)](https://codecov.io/gh/Griger10/happ-cmd)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)
[![Go Version](https://img.shields.io/badge/go-1.26-blue)](go.mod)

CLI tool for generating routing profile import links for the [Happ](https://happ.su) application.

## Installation

### Download binary

Download the latest release from the [Releases](https://github.com/Griger10/happcmd/releases) page.

### Build from source

```bash
git clone https://github.com/Griger10/happcmd
cd happcmd
go build -o happcmd .
```

## Usage

### Interactive mode

Run without arguments to open the interactive menu:

```bash
happcmd
```

### CLI mode

```bash
# Default Russia routing profile
happcmd generate

# With a custom profile name
happcmd generate -n "My Profile"

# Auto-activate on import
happcmd generate -m onadd

# Add sites to direct routing
happcmd generate --add-direct-site "domain:github.com" --add-direct-site "domain:notion.so"

# Add sites to block list
happcmd generate --add-block-site "geosite:gambling"

# Add IPs to direct routing
happcmd generate --add-direct-ip "1.2.3.4/32"

# Combined example
happcmd generate -n "Work" -m onadd \
  --add-direct-site "domain:github.com" \
  --add-block-site "geosite:gambling"
```

## Flags

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--name` | `-n` | `DefaultProfile` | Profile name in Happ |
| `--mode` | `-m` | `add` | Import mode: `add` or `onadd` |
| `--add-direct-site` | — | — | Add domain to direct routing |
| `--add-block-site` | — | — | Add domain to block list |
| `--add-direct-ip` | — | — | Add IP/CIDR to direct routing |

## Import modes

| Mode | Description |
|------|-------------|
| `add` | Adds the profile to the list. The first added profile becomes active after geo files are loaded |
| `onadd` | Adds and immediately activates the profile |

## Default profile

**Direct routing (no tunnel):**
- Russian domains and IPs (`geosite:ru`, `geoip:ru`)
- VK, Yandex, Mail.ru and related CDNs
- Government sites (`geosite:category-gov-ru`)
- Local networks

**Blocked:**
- Ad networks (`geosite:category-ads-all`)

**DNS:**
- Remote: Google DoH (`https://dns.google/dns-query`)
- Domestic: Yandex DoU

## Requirements

- Go 1.21+
- [Happ](https://happ.su) app on iOS

## License

MIT
