# avalanche-capi

> c-chain · eip-1559 · derive

[![Go 1.22+](https://img.shields.io/badge/go-1.22+-00ADD8)](https://go.dev)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Build](https://img.shields.io/badge/build-passing-brightgreen)]()

Avalanche C-Chain account shell — EVM path, stub gas.

## Features

- AVAX derivation path m/44'/60'/0'
- Local vault JSON with XOR wrap
- SHA-256 stand-in keys — no live RPC
- stdlib CLI via flag

## Prerequisites

- Go 1.22+
- Git

## Getting Started

```bash
git clone <repo-url>
cd avalanche-capi
make build
./bin/avaxc -help
```

## CLI Usage

```bash
make test
go run ./cmd/avaxc -help
```

## Project Structure

```
cmd/avaxc/main.go
internal/config/config.go
internal/crypto/keys.go
internal/wallet/wallet.go
internal/wallet/wallet_test.go
```

## Background

AVAX Go scripts use avalanche-capi.

## License

This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.


---

## Topics

![avalanche](https://img.shields.io/badge/avalanche-111827?style=flat-square) ![capi](https://img.shields.io/badge/capi-111827?style=flat-square) ![avalanche-capi](https://img.shields.io/badge/avalanche%20capi-111827?style=flat-square) ![cryptocurrency](https://img.shields.io/badge/cryptocurrency-111827?style=flat-square) ![wallet](https://img.shields.io/badge/wallet-111827?style=flat-square) ![blockchain](https://img.shields.io/badge/blockchain-111827?style=flat-square) ![web3](https://img.shields.io/badge/web3-111827?style=flat-square) ![bitcoin](https://img.shields.io/badge/bitcoin-111827?style=flat-square)

`avalanche` `capi` `avalanche-capi` `cryptocurrency` `wallet` `blockchain` `web3` `bitcoin` `ethereum` `hd-wallet` `open-source` `golang` `go`

Search: avalanche-capi · c-chain · eip-1559 · derive · Avalanche C-Chain account shell — EVM path, stub gas.

---

<sub>Avalanche C-Chain account shell — EVM path, stub gas.</sub>
