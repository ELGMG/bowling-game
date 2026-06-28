# A Bowling Game

**Date:** 1/31/2020  
**Source Document:** bowling-game.md

---

## 🇬🇧 English Specification (Original Text)

Write a program to calculate the score during an American Ten-Pin Bowling game.

We can briefly summarize the scoring for this form of bowling:

* **Frames:** Each game, or "line" of bowling, includes ten turns, or "frames" for the bowler.
* **Tries:** In each frame, the bowler gets up to two tries to knock down all the pins.
* **Open Frame:** If in two tries, he fails to knock them all down, his score for that frame is the total number of pins knocked down in his two tries.
* **Spare (`/`):** If in two tries he knocks them all down, this is called a "spare" and his score for the frame is ten plus the number of pins knocked down on his next throw (in his next turn).
* **Strike (`X`):** If on his first try in the frame he knocks down all the pins, this is called a "strike". His turn is over, and his score for the frame is ten plus the simple total of the pins knocked down in his next two rolls.
* **Tenth Frame Bonus:** If he gets a spare or strike in the last (tenth) frame, the bowler gets to throw one or two more bonus balls, respectively. These bonus throws are taken as part of the same turn. If the bonus throws knock down all the pins, the process does not repeat: the bonus throws are only used to calculate the score of the final frame.
* **Game Score:** The game score is the total of all frame scores.

### Suggested Test Cases & Examples

#### Example 1: Regular Game

| Frame | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 |
| :--- | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: |
| **Rolls** | `1, 4` | `4, 5` | `6, /` | `5, /` | `X` | `0, 1` | `7, /` | `6, /` | `X` | `2, /, 6` |
| **Score** | 5 | 14 | 29 | 49 | 60 | 61 | 77 | 97 | 117 | 133 |

#### Example 2: Perfect Game (12 Strikes)

| Frame | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 |
| :--- | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: |
| **Rolls** | `X` | `X` | `X` | `X` | `X` | `X` | `X` | `X` | `X` | `X, X, X` |
| **Score** | 30 | 60 | 90 | 120 | 150 | 180 | 210 | 240 | 270 | 300 |

#### Example 3: All Spares

| Frame | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 |
| :--- | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: |
| **Rolls** | `1, /` | `1, /` | `1, /` | `1, /` | `1, /` | `1, /` | `1, /` | `1, /` | `1, /` | `1, /, X` |
| **Score** | 11 | 22 | 33 | 44 | 55 | 66 | 77 | 88 | 99 | 119 |

---

## 🇪🇸 Versión en Español (Traducción de la Kata)

Escribe un programa para calcular la puntuación durante una partida de bolos americanos de diez bolos (*Ten-Pin Bowling*).

Podemos resumir brevemente el sistema de puntuación de la siguiente manera:

* **Frames (Entradas):** Cada partida o "línea" de bolos incluye diez turnos o *frames* para el jugador.
* **Intentos:** En cada *frame*, el jugador tiene hasta dos lanzamientos para derribar todos los bolos.
* **Frame Abierto:** Si en sus dos lanzamientos no consigue derribarlos todos, su puntuación para ese *frame* es el número total de bolos derribados en ambos tiros.
* **Spare (`/`):** Si derriba todos los bolos en sus dos lanzamientos, se denomina *spare*. Su puntuación en ese *frame* será de **10 más el número de bolos que derribe en su siguiente lanzamiento** (en su próximo turno).
* **Strike (`X`):** Si derriba todos los bolos en su primer lanzamiento del *frame*, se denomina *strike*. Su turno termina inmediatamente, y su puntuación para ese *frame* será de **10 más la suma simple de los bolos derribados en sus dos siguientes lanzamientos**.
* **Bonificación del 10.º Frame:** Si el jugador consigue un *spare* o un *strike* en el último (décimo) *frame*, obtiene el derecho a lanzar una o dos bolas extra de bonificación, respectivamente. Estos tiros extra se realizan como parte del mismo turno. Si en estos tiros extra se derriban todos los bolos, el proceso no se repite: únicamente sirven para calcular la puntuación final del décimo *frame*.
* **Puntuación Total:** La puntuación de la partida es la suma total de las puntuaciones de todos los *frames*.

*(Las tablas de puntuaciones de los casos de prueba aplican exactamente igual para ambas versiones).*