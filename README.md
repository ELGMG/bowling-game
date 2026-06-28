# bowling-game

A program to calculate the score during an American ten-pin bowling game.

It runs as an interactive CLI: you play frame by frame, entering each throw,
and after every frame the cumulative scoreboard is printed.

## Requirements

- [Go](https://go.dev/dl/) 1.22 or newer (`sudo apt install golang-go` on Debian/Ubuntu).

## Running

```bash
go run .
```

Enter each throw as:

- a number from `0` to `10` (pins knocked down),
- `X` for a strike (all ten pins on the first throw),
- `/` for a spare (the remaining pins on the second throw).

Invalid entries are rejected with an explanation and re-prompted.

### Example session

```
Frame  1 — 10 pins standing > 1
Frame  1 — 9 pins standing > 4
  Scoreboard: F1: 5
Frame  2 — 10 pins standing > 4
...
```

## Testing

The scoring engine is covered by unit tests, including the three reference
games from the kata specification (final scores 133, 300 and 119):

```bash
go test ./...
```

## Layout

- `main.go` — interactive CLI (input/output only).
- `internal/bowling/` — scoring engine:
  - `parser.go` — turns a single input into a pin count, with validation.
  - `game.go` — frame model and frame-by-frame game flow.
  - `scorer.go` — cumulative scoring (strikes, spares and the tenth-frame bonus).
  - `errors.go` — validation errors.
