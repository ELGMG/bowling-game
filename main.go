package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ELGMG/bowling-game/internal/bowling"
)

func main() {
	input_reader := bufio.NewScanner(os.Stdin)
	game := bowling.NewGame()

	print_intro()

	for !game.IsGameOver() {
		frame_before := game.CurrentFrameNumber()
		extra_throws_before := game.ExtraThrowsGranted()
		prompt_for_throw(game)

		line, has_line := read_line(input_reader)
		if !has_line {
			fmt.Println("\nNo more input. Ending the game early.")
			return
		}

		if record_error := game.RecordThrow(line); record_error != nil {
			fmt.Printf("  ✗ %v. Try again.\n", record_error)
			continue
		}

		announce_bonus_throws(extra_throws_before, game.ExtraThrowsGranted())

		if game.CurrentFrameNumber() != frame_before || game.IsGameOver() {
			print_scoreboard(game)
		}
	}

	fmt.Printf("\nGame over! Final score: %d\n", game.Total())
}

// print_intro explains how to enter rolls.
func print_intro() {
	fmt.Println("🎳 Ten-Pin Bowling")
	fmt.Println("Enter each throw as a number (0-10), 'X' for a strike, or '/' for a spare.")
	fmt.Println()
}

// prompt_for_throw shows the current frame and the pins still standing, marking
// the tenth-frame bonus throws so the player knows they are extra.
func prompt_for_throw(game *bowling.Game) {
	frame_label := fmt.Sprintf("Frame %2d", game.CurrentFrameNumber())
	if game.IsBonusThrow() {
		frame_label += " (bonus throw)"
	}
	fmt.Printf("%s — %d pins standing > ", frame_label, game.PinsStanding())
}

// announce_bonus_throws reports the extra throws the moment they are earned in
// the tenth frame: two bonus throws for a strike, one for a spare.
func announce_bonus_throws(extra_throws_before, extra_throws_after int) {
	if extra_throws_after <= extra_throws_before {
		return
	}
	switch extra_throws_after {
	case 2:
		fmt.Println("  ★ Strike! You earned 2 bonus throws.")
	case 1:
		fmt.Println("  ★ Spare! You earned 1 bonus throw.")
	}
}

// read_line reads the next line of input, reporting whether one was available.
func read_line(input_reader *bufio.Scanner) (string, bool) {
	if !input_reader.Scan() {
		return "", false
	}
	return strings.TrimSpace(input_reader.Text()), true
}

// print_scoreboard prints the cumulative score after each played frame.
func print_scoreboard(game *bowling.Game) {
	cumulative_scores := game.Scores()
	score_cells := make([]string, 0, len(cumulative_scores))
	for frame_index, score := range cumulative_scores {
		score_cells = append(score_cells, fmt.Sprintf("F%d: %d", frame_index+1, score))
	}
	fmt.Printf("  Scoreboard: %s\n", strings.Join(score_cells, " | "))
}
