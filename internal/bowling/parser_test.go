package bowling

import (
	"errors"
	"testing"
)

func TestParseRoll(t *testing.T) {
	test_cases := []struct {
		name          string
		input         string
		pins_standing int
		want_pins     int
		want_error    error
	}{
		{name: "strike with all pins standing", input: "X", pins_standing: 10, want_pins: 10},
		{name: "lowercase strike", input: "x", pins_standing: 10, want_pins: 10},
		{name: "spare completes the frame", input: "/", pins_standing: 4, want_pins: 4},
		{name: "plain number", input: "7", pins_standing: 10, want_pins: 7},
		{name: "number equal to pins standing", input: "6", pins_standing: 6, want_pins: 6},
		{name: "input is trimmed", input: "  3 ", pins_standing: 10, want_pins: 3},
		{name: "empty input", input: "", pins_standing: 10, want_error: ErrEmptyInput},
		{name: "unknown symbol", input: "?", pins_standing: 10, want_error: ErrUnknownSymbol},
		{name: "strike without all pins", input: "X", pins_standing: 4, want_error: ErrStrikeNotAllowed},
		{name: "spare as first throw", input: "/", pins_standing: 10, want_error: ErrSpareNotAllowed},
		{name: "negative number", input: "-1", pins_standing: 10, want_error: ErrPinsOutOfRange},
		{name: "more pins than standing", input: "8", pins_standing: 5, want_error: ErrPinsOutOfRange},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.name, func(t *testing.T) {
			got_pins, got_error := ParseRoll(test_case.input, test_case.pins_standing)
			if !errors.Is(got_error, test_case.want_error) {
				t.Fatalf("error = %v, want %v", got_error, test_case.want_error)
			}
			if got_error == nil && got_pins != test_case.want_pins {
				t.Errorf("pins = %d, want %d", got_pins, test_case.want_pins)
			}
		})
	}
}
