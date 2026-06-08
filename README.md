<div align="center">
  <h1>[neko.exe]</h1>
  <p><b>A cozy, gamified, terminal-based Pomodoro pet.</b></p>
  
  [![Go Report Card](https://goreportcard.com/badge/github.com/ankitshuxe/neko.exe)](https://goreportcard.com/report/github.com/ankitshuxe/neko.exe)
  [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
  [![Build Status](https://github.com/ankitshuxe/neko.exe/actions/workflows/ci.yml/badge.svg)](https://github.com/ankitshuxe/neko.exe/actions)
  [![Winget](https://img.shields.io/badge/Winget-Available-blue)](https://github.com/microsoft/winget-pkgs)
</div>

<br/>

**Neko** transforms your standard terminal into a relaxing, highly-customizable Pomodoro timer. It blends productivity tracking with "Tamagotchi-style" virtual pet mechanics, ambient background sounds, and gorgeous ASCII art to help you maintain focus and prevent burnout.

---

## Features

- **Gamified Focus**: Earn 1 "Fish Coin" for every minute of uninterrupted focus time. But beware, canceling your session early will startle the cat and cost you a coin!
- **The Neko Shop**: Spend your hard-earned Fish Coins to unlock up to 30 unique items, including new Cat Breeds (Sphynx, Persian, Galactic), Environments (Cardboard Castle, Window Sill), and Toys (Laser Pointer, Yarn Ball).
- **Dynamic ASCII Art**: Your equipped breed, toy, and environment dynamically change how your timer looks. Watch your cat play while you work and sleep when you take a break!
- **Built-in Ambience**: Play perfectly-looped offline white noise (`purr`, `rain`, `lofi`) natively in the terminal without opening a browser.
- **Deep Analytics**: Track your productivity through an organized daily diary, a 30-day percentage-based bar chart, and a weekly "GitHub-style" contribution grid.

---

## Installation

There are multiple ways to install Neko depending on your platform and preferences.

### Windows (Winget) - *Recommended*
The easiest way to install Neko on Windows is via the official Windows Package Manager:
```bash
winget install ankitshuxe.Neko
```

### From GitHub Releases (macOS, Linux, Windows)
Pre-compiled binaries are available for all major operating systems.
1. Navigate to the [Releases page](../../releases/latest).
2. Download the binary for your OS and architecture.
3. Extract it and move it to a folder in your system's `PATH`.

### Build from Source (Go)
Ensure you have [Go](https://go.dev/) 1.21+ installed.
```bash
git clone https://github.com/ankitshuxe/neko.exe.git
cd neko.exe
make build
# or if you don't have make: go build -o neko .
```

---

## How to Play (Usage)

Simply run `neko` with no arguments to see the beautiful welcome dashboard!

### Timers
- **Start a Focus Session**: `neko start [minutes]`
  - Example: `neko start 45 -t coding -s rain`
  - Flags: `-t` to tag your session, `-s` to play ambient sound (`purr`, `rain`, `lofi`).
- **Start a Break**: `neko break [minutes]`
  - Let your cat take a nap.

*(During any session, press `p` to pause/resume, or `q` to quit).*

### Economy & Customization
- **View the Shop**: `neko shop`
  - See your Fish Coin balance and browse the available items.
- **Buy an Item**: `neko shop buy [item_id]`
  - Example: `neko shop buy breed_siamese`
- **Equip an Item**: `neko shop equip [item_id]`
  - Swap out your current breed, toy, or environment to customize your active timer!

### Analytics
- **Daily Log**: `neko diary`
  - View today's focus blocks grouped by your custom tags.
- **Monthly Stats**: `neko stats`
  - See an ASCII bar chart of your most-used tags over the last 30 days and a 7-day attendance grid.

---

## Data Persistence & Security

Your coins, shop inventory, and focus history are automatically saved securely in your computer's native user configuration directory (e.g., `~/.config/neko/history.json` or `%AppData%\neko\history.json`).

*Neko is fully audited for security. It employs atomic file-writes to prevent corruption, strict `0600` file permissions to prevent unauthorized access on shared machines, and bounds-checking to prevent memory bloat over years of usage.*

---

## Contributing
Contributions, issues, and feature requests are welcome! Feel free to check the [issues page](../../issues).
Please read our [Contributing Guide](CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.
