# mac-usage

A terminal dashboard for monitoring macOS system stats in real time, built in Go.

## What it does

Displays live system metrics in a terminal UI that updates every second:

- **CPU Usage** — overall usage percentage
- **Memory Usage** — RAM used as a percentage
- **Battery** — current charge percentage
- **Network I/O** — upload and download speed in MB/s

## How it works

Stats are collected using `gopsutil` and polled every second via a goroutine. The stats are sent over a buffered channel to the UI, which is built with `termui`. A `select` loop handles both incoming stats and keyboard events simultaneously.

## Installation

Requires Go 1.18+
```bash
