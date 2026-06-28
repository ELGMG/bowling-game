<div align="center">

# 🎳 bowling-game

**Score an American ten-pin bowling game — one frame at a time.**

An interactive command-line scorer written in Go. Play frame by frame, type each
throw, and watch the cumulative scoreboard build in real time.

[![Go](https://img.shields.io/badge/Go-1.22%2B-00ADD8?logo=go&logoColor=white)](https://go.dev/dl/)
[![Platform](https://img.shields.io/badge/platform-Linux%20%7C%20macOS%20%7C%20Windows-444)](#-requirements)
[![Interface](https://img.shields.io/badge/interface-CLI-1f6feb)](#-running)
[![Tested](https://img.shields.io/badge/reference%20games-133%20%7C%20300%20%7C%20119%20%E2%9C%93-2ea043)](#-testing)

[Requirements](#-requirements) • [Running](#-running) • [How to play](#-how-to-play) • [Testing](#-testing) • [Project layout](#-project-layout) • [Technical decisions](#-technical-decisions)

</div>

---

## 🎬 Demo

```text
🎳 Ten-Pin Bowling
Enter each throw as a number (0-10), 'X' for a strike, or '/' for a spare.

Frame  1 — 10 pins standing > 1
Frame  1 — 9 pins standing > 4
  Scoreboard: F1: 5
Frame  2 — 10 pins standing > 4
Frame  2 — 6 pins standing > 5
  Scoreboard: F1: 5 | F2: 14
...
Frame 10 — 10 pins standing > X
  ★ Strike! You earned 2 bonus throws.
Frame 10 (bonus throw) — 10 pins standing > X
Frame 10 (bonus throw) — 10 pins standing > X
  Scoreboard: ... | F10: 300

Game over! Final score: 300
```

---

## 📦 Requirements

### Linux / macOS

[Go](https://go.dev/dl/) **1.22 or newer** is a required dependency. On
Debian/Ubuntu you can install it with:

```bash
sudo apt install golang-go
```

On macOS with [Homebrew](https://brew.sh/):

```bash
brew install go
```

### Windows

The program is pure Go and uses only the standard library, so it runs on Windows
**without code changes**. It is verified to cross-compile for `windows/amd64` and
`windows/arm64`.

**1. Install Go 1.22 or newer** — it is a required dependency. Either with
[winget](https://learn.microsoft.com/windows/package-manager/winget/) (PowerShell):

```powershell
winget install --id GoLang.Go -e
```

…or download the `.msi` installer for Windows from [go.dev/dl](https://go.dev/dl/)
and run it. After installing, reopen your terminal so `go` is available on the `PATH`.

**2. Use a UTF-8 capable terminal** so the icons (`🎳`, `★`, `—`) render correctly.
**Windows Terminal** or **PowerShell 7** work out of the box. If you use the
legacy `cmd.exe`, switch the console to UTF-8 first:

```bat
chcp 65001
go run .
```

> [!NOTE]
> Windows `CRLF` line endings are handled automatically, so no extra setup is
> needed for input.

---

## 🚀 Running

```bash
go run .
```

---

## 🎮 How to play

Enter one throw at a time. Each throw can be:

| Input | Meaning |
| :---: | :--- |
| `0` – `10` | Number of pins knocked down |
| `X` / `x` | **Strike** — all ten pins on the first throw |
| `/` | **Spare** — the remaining pins on the second throw |

> [!TIP]
> Invalid entries (out of range, a `/` as the first throw, an `X` when pins are
> already down…) are **rejected with an explanation and re-prompted** — you can't
> enter an impossible score.

### Scoring rules

| Frame result | Score for the frame |
| :--- | :--- |
| **Open** (≤ 9 pins in two throws) | Pins knocked down |
| **Spare** (`/`) | 10 + next **1** throw |
| **Strike** (`X`) | 10 + next **2** throws |
| **Tenth frame** | Spare → 1 bonus throw · Strike → 2 bonus throws |

---

## 🧪 Testing

The scoring engine is covered by unit tests, including the **three reference
games** from the kata specification:

```bash
go test ./...
```

| Reference game | Final score |
| :--- | :---: |
| Regular game | `133` |
| Perfect game (12 strikes) | `300` |
| All spares | `119` |

---

## 🗂️ Project layout

```text
bowling-game/
├── main.go                 # interactive CLI (input/output only)
└── internal/bowling/       # scoring engine (no I/O — fully testable)
    ├── parser.go           # turns a single input into a pin count, with validation
    ├── game.go             # frame model and frame-by-frame game flow
    ├── scorer.go           # cumulative scoring (strikes, spares, tenth-frame bonus)
    └── errors.go           # validation errors
```

---

## 🛠️ Technical decisions

> Since this project is a technical interview test for **Zollner**, the following
> technical decisions were considered and made.

- **Language: Go.** Go was chosen because it is the language Zollner uses on the
  backend, so the solution mirrors the company's real stack.

- **No graphical user interface (UI).** The original requirement only asks to
  "calculate the score during a game", not a visual interface. An **interactive
  CLI** was implemented instead, and adding a UI that was not requested was
  avoided, keeping the scope focused on the stated problem.

- **Logic decoupled from I/O.** The scoring engine lives in the `internal/bowling`
  package and is independent of the console; `main.go` only handles input/output.
  This makes the game rules **testable without stdin** and lets the CLI be replaced
  by another interface (web, API) without touching the logic.

- **Scoring algorithm over "flattened" rolls.** Instead of special-casing each
  frame, every roll is traversed in a single sequence applying the classic
  algorithm (strike = 10 + next 2 rolls, spare = 10 + next 1 roll). The tenth-frame
  bonus rolls live inside that sequence, so the look-ahead works without extra logic.

- **Input with numbers and `X` / `/` symbols.** Both numeric pins (0–10) and the
  traditional bowling notation are accepted, so the experience resembles a real
  scoresheet.

- **Explicit input validation.** Each roll is validated (range 0–10, that a frame's
  total does not exceed 10, and that `X` / `/` are used in the correct position),
  returning clear errors instead of assuming valid input.

- **Code conventions.** Code and comments in English, local variables in
  `snake_case`, and declarative function names. Exported identifiers use
  `PascalCase` only where the language requires it.

- **Cumulative score per frame.** The running score is shown after each frame (like
  the tables in the specification), not just the final total, to make the
  frame-by-frame calculation visible.

- **Test coverage.** Unit tests encode the three reference games from the kata
  (133, 300 and 119) plus validation cases and the tenth-frame bonus.
