<div align="center">
  <h1>[neko.exe]</h1>
  <p><b>TERMINAL-NATIVE POMODORO COMPANION</b></p>
  
  [![Go Report Card](https://goreportcard.com/badge/github.com/ankitshuxe/neko.exe)](https://goreportcard.com/report/github.com/ankitshuxe/neko.exe)
  [![License: MIT](https://img.shields.io/badge/License-MIT-F05522.svg)](https://opensource.org/licenses/MIT)
  [![Build Status](https://github.com/ankitshuxe/neko.exe/actions/workflows/ci.yml/badge.svg)](https://github.com/ankitshuxe/neko.exe/actions)
  [![Winget](https://img.shields.io/badge/Winget-Available-1A1A1A)](https://github.com/microsoft/winget-pkgs)
</div>

<br/>

**Abandon the browser.** `neko.exe` is a strictly CLI-based productivity companion engineered for developers who demand zero electron bloat. Execute flawless 25-minute focus intervals directly from your shell, earn Fish Coins, and unlock ASCII environments. 

---

## // STRICT INTERVAL PROTOCOLS

- **Gamified Accountability**: Earn 1 "Fish Coin" for every minute of uninterrupted deep work. But beware: interrupting your session (Ctrl+C) triggers a "startled" penalty, immediately deducting hard-earned currency. Focus is mandatory.
- **The ASCII Marketplace**: Spend your Fish Coins to unlock up to 30 unique tactical upgrades. Exchange currency for new Cat Breeds (Sphynx, Persian, Galactic), Environments (Cardboard Castle, Window Sill), and Toys (Laser Pointer, Yarn Ball).
- **Dynamic Terminal UI**: Your equipped breed, toy, and environment dynamically change the ASCII layout. Watch your terminal cat play while you work, and sleep when you take a break.
- **Offline Audio Ambience**: Play perfectly-looped offline white noise (`purr`, `rain`, `lofi`) natively in the terminal. No browser tabs required.
- **Deep Analytics**: Track your productivity through an organized daily diary, a 30-day percentage-based bar chart, and a weekly GitHub-style contribution grid.

---

## // DEPLOYMENT VECTORS

Select your preferred package manager to initiate the installation sequence.

### 01. GOLANG
```bash
go install github.com/ankitshuxe/neko.exe@latest
```

### 02. WINDOWS
```bash
winget install ankitshuxe.Neko
```

### 03. MACOS / LINUX
```bash
brew install ankitshuxe/tap/neko
```

### BUILD FROM SOURCE
Ensure you have [Go](https://go.dev/) 1.21+ installed.
```bash
git clone https://github.com/ankitshuxe/neko.exe.git
cd neko.exe
make build
```

---

## // EXECUTION COMMANDS

Simply run `neko` with no arguments to see the ASCII welcome dashboard.

### TIMERS
- **Initiate Focus Session**: `neko start [minutes]`
  - *Example:* `neko start 45 -t coding -s rain`
  - *Flags:* `-t` to tag your session, `-s` to play ambient sound (`purr`, `rain`, `lofi`).
- **Initiate Break Protocol**: `neko break [minutes]`
  - Let your cat take a nap.

*(During any session, press `p` to pause/resume, or `q` to quit).*

### ECONOMY & CUSTOMIZATION
- **View The Shop**: `neko shop`
  - See your Fish Coin balance and browse the available tactical upgrades.
- **Execute Purchase**: `neko shop buy [item_id]`
  - *Example:* `neko shop buy breed_siamese`
- **Equip Item**: `neko shop equip [item_id]`
  - Swap out your current breed, toy, or environment to customize your active timer.

### ANALYTICS
- **Daily Log**: `neko diary`
  - View today's focus blocks grouped by your custom tags.
- **Monthly Stats**: `neko stats`
  - See an ASCII bar chart of your most-used tags over the last 30 days and a 7-day attendance grid.

---

## // SECURE DATA PERSISTENCE

Your coins, shop inventory, and focus history are automatically saved securely in your computer's native user configuration directory (e.g., `~/.config/neko/history.json` or `%AppData%\neko\history.json`).

*Neko is fully audited for security. It employs atomic file-writes to prevent corruption, strict `0600` file permissions to prevent unauthorized access on shared machines, and bounds-checking to prevent memory bloat over years of usage.*

---

## // CONTRIBUTING
Contributions, issues, and feature requests are welcome! Feel free to check the [issues page](../../issues).
Please read our [Contributing Guide](CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.
