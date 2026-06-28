package bowling

import "errors"

// Validation errors returned while parsing a roll or recording a throw.
var (
	// ErrEmptyInput is returned when no value was provided for a roll.
	ErrEmptyInput = errors.New("no roll was provided")
	// ErrUnknownSymbol is returned when the input is neither a number nor a known symbol.
	ErrUnknownSymbol = errors.New("roll must be a number from 0 to 10, 'X' (strike) or '/' (spare)")
	// ErrStrikeNotAllowed is returned when 'X' is used while pins are already partially knocked down.
	ErrStrikeNotAllowed = errors.New("a strike ('X') is only valid when all ten pins are standing")
	// ErrSpareNotAllowed is returned when '/' is used as the first throw of a frame.
	ErrSpareNotAllowed = errors.New("a spare ('/') is only valid as a second throw")
	// ErrPinsOutOfRange is returned when the number of pins is negative or above what is standing.
	ErrPinsOutOfRange = errors.New("knocked-down pins must be between 0 and the pins still standing")
	// ErrGameOver is returned when a throw is recorded after the game has finished.
	ErrGameOver = errors.New("the game is already over")
)
