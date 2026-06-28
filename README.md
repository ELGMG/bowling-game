# bowling-game

A program to calculate the score during an American ten-pin bowling game.

It runs as an interactive CLI: you play frame by frame, entering each throw,
and after every frame the cumulative scoreboard is printed.

## Requirements

### Linux / macOS

- [Go](https://go.dev/dl/) 1.22 or newer (`sudo apt install golang-go` on Debian/Ubuntu).

### Windows

The program is pure Go and uses only the standard library, so it runs on Windows
without code changes. It is verified to cross-compile for `windows/amd64` and
`windows/arm64`. To run it you need:

- [Go](https://go.dev/dl/) 1.22 or newer for Windows (the `.msi` installer from go.dev).
- A UTF-8 capable terminal so the icons (`🎳`, `★`, `—`) render correctly.
  **Windows Terminal** or **PowerShell 7** work out of the box. If you use the
  legacy `cmd.exe`, switch the console to UTF-8 first:

  ```bat
  chcp 65001
  go run .
  ```

  Windows `CRLF` line endings are handled automatically, so no extra setup is
  needed for input.

To build a standalone Windows executable (from any OS):

```bash
GOOS=windows GOARCH=amd64 go build -o bowling.exe .
```

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

## Decisiones técnicas tomadas para la implementación

> Ya que este proyecto es una prueba técnica para **Zollner** se plantearon y tomaron las sigueintes decisiones tecnicas.

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
