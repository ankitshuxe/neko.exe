# Neko 🐾

> A cozy, gamified, terminal-based Pomodoro pet.

**Neko** transforms your standard terminal into a relaxing, highly-customizable Pomodoro timer. It blends productivity tracking with "Tamagotchi-style" virtual pet mechanics, ambient background sounds, and gorgeous ASCII art to help you maintain focus and prevent burnout.

---

## ✨ Features

- **Gamified Focus**: Earn 1 "Fish Coin" for every minute of uninterrupted focus time. But beware, canceling your session early will startle the cat and cost you a coin!
- **The Neko Shop**: Spend your hard-earned Fish Coins to unlock up to 30 unique items, including new Cat Breeds (Sphynx, Persian, Galactic), Environments (Cardboard Castle, Window Sill), and Toys (Laser Pointer, Yarn Ball).
- **Dynamic ASCII Art**: Your equipped breed, toy, and environment dynamically change how your timer looks. Watch your cat play while you work and sleep when you take a break!
- **Built-in Ambience**: Play perfectly-looped offline white noise (`purr`, `rain`, `lofi`) natively in the terminal without opening a browser.
- **Deep Analytics**: Track your productivity through an organized daily diary, a 30-day percentage-based bar chart, and a weekly "GitHub-style" contribution grid.

---

## 🚀 Installation

Ensure you have [Go](https://go.dev/) installed (1.16+ required for `//go:embed` support).

Clone the repository and build the binary:
```bash
git clone https://github.com/yourusername/neko.git
cd neko
go build -o neko .
```
*(Tip: Move the `neko` binary into your system's PATH to run it from anywhere!)*

---

## 🎮 How to Play (Usage)

Simply run `neko` with no arguments to see the beautiful welcome dashboard!

### ⏱️ Timers
- **Start a Focus Session**: `neko start [minutes]`
  - Example: `neko start 45 -t coding -s rain`
  - Flags: `-t` to tag your session, `-s` to play ambient sound (`purr`, `rain`, `lofi`).
- **Start a Break**: `neko break [minutes]`
  - Let your cat take a nap.

*(During any session, press `p` to pause/resume, or `q` to quit).*

### 🛍️ Economy & Customization
- **View the Shop**: `neko shop`
  - See your Fish Coin balance and browse the available items.
- **Buy an Item**: `neko shop buy [item_id]`
  - Example: `neko shop buy breed_siamese`
- **Equip an Item**: `neko shop equip [item_id]`
  - Swap out your current breed, toy, or environment to customize your active timer!

### 📊 Analytics
- **Daily Log**: `neko diary`
  - View today's focus blocks grouped by your custom tags.
- **Monthly Stats**: `neko stats`
  - See an ASCII bar chart of your most-used tags over the last 30 days and a 7-day attendance grid.

---

## 💾 Data Persistence
Your coins, shop inventory, and focus history are automatically saved securely in your computer's native user configuration directory (e.g., `~/.config/neko/history.json` or `%AppData%\neko\history.json`). You will never lose your progress even if you move the binary or restart your machine!
