# mac-stats

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
git clone https://github.com/zgoldwyn/mac-dashboard.git
cd mac-dashboard
go mod tidy
```

## How to run

Must be run from a real terminal (not an IDE console):
```bash
go run .
```

Press `q`, `Escape`, or `Ctrl+C` to exit.

## Notes

- Temperature monitoring is not supported on Apple Silicon due to macOS sensor access restrictions
- Network speeds show MB/s delta per second
