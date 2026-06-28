package bowling

// frame_count is the number of frames in a standard game.
const frame_count = 10

// CalculateScores returns the cumulative score after each played frame.
//
// It flattens every throw into a single slice of pin counts and applies the
// classic ten-pin algorithm: a strike scores ten plus the next two rolls, a
// spare scores ten plus the next roll, and an open frame scores its two rolls.
// The tenth-frame bonus throws live in the flattened slice, so the look-ahead
// works without any special casing.
func CalculateScores(frames []frame) []int {
	all_rolls := flatten_rolls(frames)

	cumulative_scores := make([]int, 0, frame_count)
	running_total := 0
	roll_index := 0

	for frame_number := 0; frame_number < frame_count && roll_index < len(all_rolls); frame_number++ {
		switch {
		case is_strike(all_rolls, roll_index):
			running_total += total_pins + strike_bonus(all_rolls, roll_index)
			roll_index++
		case is_spare(all_rolls, roll_index):
			running_total += total_pins + spare_bonus(all_rolls, roll_index)
			roll_index += 2
		default:
			running_total += roll_at(all_rolls, roll_index) + roll_at(all_rolls, roll_index+1)
			roll_index += 2
		}
		cumulative_scores = append(cumulative_scores, running_total)
	}

	return cumulative_scores
}

// flatten_rolls concatenates the pin counts of every frame, in playing order.
func flatten_rolls(frames []frame) []int {
	all_rolls := make([]int, 0, len(frames)*2)
	for _, current_frame := range frames {
		all_rolls = append(all_rolls, current_frame.rolls...)
	}
	return all_rolls
}

// is_strike reports whether the roll at roll_index knocked down all ten pins.
func is_strike(all_rolls []int, roll_index int) bool {
	return roll_at(all_rolls, roll_index) == total_pins
}

// is_spare reports whether the two rolls starting at roll_index total ten pins.
func is_spare(all_rolls []int, roll_index int) bool {
	return roll_at(all_rolls, roll_index)+roll_at(all_rolls, roll_index+1) == total_pins
}

// strike_bonus returns the pins of the two rolls following a strike.
func strike_bonus(all_rolls []int, roll_index int) int {
	return roll_at(all_rolls, roll_index+1) + roll_at(all_rolls, roll_index+2)
}

// spare_bonus returns the pins of the single roll following a spare.
func spare_bonus(all_rolls []int, roll_index int) int {
	return roll_at(all_rolls, roll_index+2)
}

// roll_at returns the pins at roll_index, or 0 when the index is out of range.
func roll_at(all_rolls []int, roll_index int) int {
	if roll_index < 0 || roll_index >= len(all_rolls) {
		return 0
	}
	return all_rolls[roll_index]
}
