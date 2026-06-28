package bowling

// frame holds the pins knocked down on each throw within a single frame.
type frame struct {
	rolls []int
}

// Game models a single line of ten-pin bowling, advancing frame by frame as
// throws are recorded.
type Game struct {
	frames        []frame
	current_frame int
}

// NewGame returns an empty game positioned on the first frame.
func NewGame() *Game {
	return &Game{
		frames:        make([]frame, frame_count),
		current_frame: 0,
	}
}

// RecordThrow parses a single user input and adds it to the current frame,
// advancing to the next frame once the current one is complete. It returns a
// validation error (and records nothing) when the input is not a legal throw.
func (g *Game) RecordThrow(input string) error {
	if g.IsGameOver() {
		return ErrGameOver
	}

	knocked_down, parse_error := ParseRoll(input, g.PinsStanding())
	if parse_error != nil {
		return parse_error
	}

	current := &g.frames[g.current_frame]
	current.rolls = append(current.rolls, knocked_down)

	if g.is_current_frame_complete() {
		g.advance_frame()
	}
	return nil
}

// PinsStanding returns how many pins are available for the next throw, taking
// the tenth-frame pin resets after a strike or spare into account.
func (g *Game) PinsStanding() int {
	return pins_standing_in_frame(g.frames[g.current_frame], g.is_tenth_frame())
}

// IsGameOver reports whether all ten frames (and any tenth-frame bonus throws)
// have been played.
func (g *Game) IsGameOver() bool {
	return g.current_frame == frame_count-1 && g.is_current_frame_complete()
}

// CurrentFrameNumber returns the 1-based number of the frame being played.
func (g *Game) CurrentFrameNumber() int {
	return g.current_frame + 1
}

// IsBonusThrow reports whether the upcoming throw is an extra (bonus) throw of
// the tenth frame, earned by a strike or a spare. A leading strike grants two
// bonus throws (the second and third), while a spare grants one (the third).
func (g *Game) IsBonusThrow() bool {
	if !g.is_tenth_frame() {
		return false
	}
	rolls := g.frames[g.current_frame].rolls
	throw_count := len(rolls)
	if throw_count >= 1 && rolls[0] == total_pins {
		return true // any throw following a leading strike is a bonus throw
	}
	return throw_count == 2 && rolls[0]+rolls[1] == total_pins // third throw after a spare
}

// ExtraThrowsGranted returns how many bonus throws the tenth frame has earned so
// far: two for a leading strike, one for a spare, and zero otherwise.
func (g *Game) ExtraThrowsGranted() int {
	if !g.is_tenth_frame() {
		return 0
	}
	rolls := g.frames[g.current_frame].rolls
	if len(rolls) >= 1 && rolls[0] == total_pins {
		return 2
	}
	if len(rolls) >= 2 && rolls[0]+rolls[1] == total_pins {
		return 1
	}
	return 0
}

// Scores returns the cumulative score after each played frame.
func (g *Game) Scores() []int {
	return CalculateScores(g.frames)
}

// Total returns the final game score.
func (g *Game) Total() int {
	scores := g.Scores()
	if len(scores) == 0 {
		return 0
	}
	return scores[len(scores)-1]
}

// advance_frame moves to the next frame, staying on the tenth once it is reached.
func (g *Game) advance_frame() {
	if g.current_frame < frame_count-1 {
		g.current_frame++
	}
}

// is_tenth_frame reports whether the current frame is the final one.
func (g *Game) is_tenth_frame() bool {
	return g.current_frame == frame_count-1
}

// is_current_frame_complete reports whether the current frame accepts no more throws.
func (g *Game) is_current_frame_complete() bool {
	return is_frame_complete(g.frames[g.current_frame], g.is_tenth_frame())
}

// is_frame_complete reports whether a frame can accept no further throws.
func is_frame_complete(f frame, is_tenth bool) bool {
	throw_count := len(f.rolls)
	if is_tenth {
		if throw_count < 2 {
			return false
		}
		if throw_count >= 3 {
			return true
		}
		// Two throws so far: the frame is only finished without a bonus throw
		// when it stayed open (neither a strike nor a spare).
		return f.rolls[0]+f.rolls[1] < total_pins
	}

	if throw_count == 1 && f.rolls[0] == total_pins {
		return true // strike ends a regular frame after one throw
	}
	return throw_count >= 2
}

// pins_standing_in_frame returns the pins available for the next throw of a frame.
func pins_standing_in_frame(f frame, is_tenth bool) int {
	throw_count := len(f.rolls)
	if throw_count == 0 {
		return total_pins
	}

	if !is_tenth {
		return total_pins - f.rolls[0]
	}

	// Tenth frame: pins are reset after every strike or completed spare.
	if throw_count == 1 {
		if f.rolls[0] == total_pins {
			return total_pins
		}
		return total_pins - f.rolls[0]
	}

	// Second bonus throw. Either the first throw was a strike, or the first two
	// throws made a spare (which always resets the pins).
	if f.rolls[0] == total_pins {
		if f.rolls[1] == total_pins {
			return total_pins
		}
		return total_pins - f.rolls[1]
	}
	return total_pins
}
