# 💤 sleepy-time

A simple Go library that randomly moves the mouse and presses keys to keep your machine from going idle—ideal for remote work setups, virtual meetings, or just staying “active.”

## 🚀 Features
- Random mouse movements across the screen
- Random keypresses from a configurable list
- Customizable run duration and delay between actions
- Lightweight and easy to use

## 🧰 Installation

```bash
go get github.com/Zeref010122/sleepy-time


## ⚙️ Configuration Options

- **RunDuration**: How long to keep the program running (default: 10 minutes)
- **Delay**: Delay between actions (default: 1 second)
- **Keys**: List of keys to randomly press (default: `["a", "b", "enter", "space"]`)
- **MoveMouse**: Enable/disable random mouse movement (default: `true`)
- **TerminationKeys**: Hotkey(s) to stop the program (default: `["ctrl", "alt", "q"]`)

## 🛑 Stopping

Press the configured termination hotkey(s) (default: `Ctrl+Alt+Q`) to stop the program at any time.