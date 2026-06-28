package bowling

import (
	"strconv"
	"strings"
)

// total_pins is the number of pins standing at the start of a fresh frame.
const total_pins = 10

// ParseRoll converts a single user input into the number of knocked-down pins,
// validating it against the pins still standing in the current frame.
//
//   - "X" / "x"  -> a strike, only valid when all ten pins are standing.
//   - "/"        -> a spare, completing whatever pins were left standing;
//     it is never valid as the first throw of a frame.
//   - "0".."10"  -> an explicit pin count, valid when 0 <= n <= pins_standing.
func ParseRoll(input string, pins_standing int) (int, error) {
	trimmed_input := strings.TrimSpace(input)
	if trimmed_input == "" {
		return 0, ErrEmptyInput
	}

	switch strings.ToUpper(trimmed_input) {
	case "X":
		if pins_standing != total_pins {
			return 0, ErrStrikeNotAllowed
		}
		return total_pins, nil
	case "/":
		if pins_standing == total_pins {
			return 0, ErrSpareNotAllowed
		}
		return pins_standing, nil
	}

	knocked_down, conversion_error := strconv.Atoi(trimmed_input)
	if conversion_error != nil {
		return 0, ErrUnknownSymbol
	}
	if knocked_down < 0 || knocked_down > pins_standing {
		return 0, ErrPinsOutOfRange
	}
	return knocked_down, nil
}
