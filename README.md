<div align="center">

# 🎳 bowling-game

**Score an American ten-pin bowling game — one frame at a time.**

An interactive command-line scorer written in Go. Play frame by frame, type each
throw, and watch the cumulative scoreboard build in real time.

[![Go](https://img.shields.io/badge/Go-1.22%2B-00ADD8?logo=go&logoColor=white)](https://go.dev/dl/)
[![Platform](https://img.shields.io/badge/platform-Linux%20%7C%20macOS%20%7C%20Windows-444)](#-requirements)
[![Interface](https://img.shields.io/badge/interface-CLI-1f6feb)](#-running)
[![Tested](https://img.shields.io/badge/reference%20games-133%20%7C%20300%20%7C%20119%20%E2%9C%93-2ea043)](#-testing)

[Requirements](#-requirements) • [Running](#-running) • [How to play](#-how-to-play) • [Testing](#-testing) • [Project layout](#-project-layout) • [Technical decisions](#-decisiones-técnicas-tomadas-para-la-implementación)

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

## 🛠️ Decisiones técnicas tomadas para la implementación

> Ya que este proyecto es una prueba técnica para **Zollner** se plantearon y tomaron las siguientes decisiones tecnicas.

- **Lenguaje: Go.** Se eligió Go porque es el lenguaje que Zollner utiliza en el
  backend, de modo que la solución refleja el stack real de la empresa.

- **Sin interfaz gráfica (UI).** El requerimiento inicial solo pide "calcular la
  puntuación durante una partida", no una interfaz visual. Se implementó por
  tanto una **CLI interactiva** y se evitó añadir una UI que no estaba solicitada,
  manteniendo el alcance acotado al problema planteado sin desvíos.

- **Lógica separada del I/O.** El motor de puntuación vive en el paquete
  `internal/bowling` y es independiente de la consola; `main.go` solo se encarga
  de la entrada/salida. Esto hace que las reglas del juego sean **testeables sin
  stdin** y que la CLI pudiera reemplazarse por otra interfaz (web, API) sin tocar
  la lógica.

- **Algoritmo de puntuación sobre los tiros "aplanados".** En lugar de casos
  especiales por frame, todos los tiros se recorren en una sola secuencia
  aplicando el algoritmo clásico (strike = 10 + 2 tiros siguientes, spare = 10 +
  1 tiro siguiente). Los tiros de bonificación del 10.º frame quedan dentro de esa
  secuencia, por lo que el *look-ahead* funciona sin lógica adicional.

- **Entrada con números y símbolos `X` / `/`.** Se aceptan tanto pinos numéricos
  (0–10) como la notación tradicional de bowling, para que la experiencia se
  parezca a una hoja de anotación real.

- **Validación explícita de entradas.** Cada tiro se valida (rango 0–10, que la
  suma de un frame no exceda 10, y que `X` / `/` se usen en la posición correcta)
  devolviendo errores claros, en lugar de asumir entradas correctas.

- **Convenciones de código.** Código y comentarios en inglés, variables locales en
  `snake_case` y funciones con nombres declarativos. Los identificadores
  exportados usan `PascalCase` únicamente donde el lenguaje lo exige.

- **Marcador acumulado por frame.** Se muestra el puntaje acumulado tras cada
  frame (como las tablas de la especificación), no solo el total final, para hacer
  visible el cálculo frame a frame.

- **Cobertura de pruebas.** Se incluyen tests unitarios que codifican los tres
  juegos de referencia de la kata (133, 300 y 119) además de casos de validación y
  del bonus del 10.º frame.
