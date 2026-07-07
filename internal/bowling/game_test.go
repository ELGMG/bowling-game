package bowling

import (
	"reflect"
	"testing"
)

// play_all records every throw in order, failing the test on the first error.
func play_all(t *testing.T, game *Game, throws []string) {
	t.Helper()
	for throw_index, throw := range throws {
		if record_error := game.RecordThrow(throw); record_error != nil {
			t.Fatalf("throw %d (%q) was unexpectedly rejected: %v", throw_index, throw, record_error)
		}
	}
}

func TestExampleGames(t *testing.T) {
	test_cases := []struct {
		name              string
		throws            []string
		cumulative_scores []int
	}{
		{
			name:              "regular game",
			throws:            []string{"1", "4", "4", "5", "6", "/", "5", "/", "X", "0", "1", "7", "/", "6", "/", "X", "2", "/", "6"},
			cumulative_scores: []int{5, 14, 29, 49, 60, 61, 77, 97, 117, 133},
		},
		{
			name:              "perfect game",
			throws:            []string{"X", "X", "X", "X", "X", "X", "X", "X", "X", "X", "X", "X"},
			cumulative_scores: []int{30, 60, 90, 120, 150, 180, 210, 240, 270, 300},
		},
		{
			name:              "all spares",
			throws:            []string{"1", "/", "1", "/", "1", "/", "1", "/", "1", "/", "1", "/", "1", "/", "1", "/", "1", "/", "1", "/", "X"},
			cumulative_scores: []int{11, 22, 33, 44, 55, 66, 77, 88, 99, 119},
		},
		{
			name:              "half game",
			throws:            []string{"1", "/", "1", "/", "1", "5"},
			cumulative_scores: []int{11, 22, 28},
		},
		{
			name:              "gutter game (all zeros)",
			throws:            []string{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0"},
			cumulative_scores: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name:              "nine and miss every frame",
			throws:            []string{"9", "0", "9", "0", "9", "0", "9", "0", "9", "0", "9", "0", "9", "0", "9", "0", "9", "0", "9", "0"},
			cumulative_scores: []int{9, 18, 27, 36, 45, 54, 63, 72, 81, 90},
		},
		{
			name:              "all fives with bonus (150)",
			throws:            []string{"5", "/", "5", "/", "5", "/", "5", "/", "5", "/", "5", "/", "5", "/", "5", "/", "5", "/", "5", "/", "5"},
			cumulative_scores: []int{15, 30, 45, 60, 75, 90, 105, 120, 135, 150},
		},
		{
			name:              "strike then open alternating",
			throws:            []string{"X", "4", "5", "X", "4", "5", "X", "4", "5", "X", "4", "5", "X", "4", "5"},
			cumulative_scores: []int{19, 28, 47, 56, 75, 84, 103, 112, 131, 140},
		},
		{
			name:              "single opening strike, rest gutter",
			throws:            []string{"X", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0"},
			cumulative_scores: []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		},
		{
			name:              "single opening spare, rest gutter",
			throws:            []string{"5", "/", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0"},
			cumulative_scores: []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		},
		{
			name:              "tenth-frame strike with two bonus strikes",
			throws:            []string{"4", "5", "4", "5", "4", "5", "4", "5", "4", "5", "4", "5", "4", "5", "4", "5", "4", "5", "X", "X", "X"},
			cumulative_scores: []int{9, 18, 27, 36, 45, 54, 63, 72, 81, 111},
		},
		{
			name:              "tenth-frame spare with bonus",
			throws:            []string{"4", "5", "4", "5", "4", "5", "4", "5", "4", "5", "4", "5", "4", "5", "4", "5", "4", "5", "5", "/", "5"},
			cumulative_scores: []int{9, 18, 27, 36, 45, 54, 63, 72, 81, 96},
		},
		{
			name:              "dutch 200 (alternating strike and spare)",
			throws:            []string{"X", "9", "/", "X", "9", "/", "X", "9", "/", "X", "9", "/", "X", "9", "/", "X"},
			cumulative_scores: []int{20, 40, 60, 80, 100, 120, 140, 160, 180, 200},
		},
		{
			name:              "nine strikes then open tenth",
			throws:            []string{"X", "X", "X", "X", "X", "X", "X", "X", "X", "8", "1"},
			cumulative_scores: []int{30, 60, 90, 120, 150, 180, 210, 238, 257, 266},
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.name, func(t *testing.T) {
			game := NewGame()
			play_all(t, game, test_case.throws)

			if !game.IsGameOver() {
				t.Fatalf("expected the game to be over after %d throws", len(test_case.throws))
			}
			got_scores := game.Scores()
			if !reflect.DeepEqual(got_scores, test_case.cumulative_scores) {
				t.Errorf("cumulative scores = %v, want %v", got_scores, test_case.cumulative_scores)
			}
			final_total := test_case.cumulative_scores[len(test_case.cumulative_scores)-1]
			if game.Total() != final_total {
				t.Errorf("total = %d, want %d", game.Total(), final_total)
			}
		})
	}
}

// TestRecordThrowValidation feeds illegal inputs so that every validation flag
// is activated, and checks the matching error is returned while the game state
// is left untouched.
func TestRecordThrowValidation(t *testing.T) {
	test_cases := []struct {
		name       string
		setup      []string // legal throws played before the illegal one
		bad_throw  string
		want_error error
	}{
		{
			name:       "empty input is rejected",
			bad_throw:  "",
			want_error: ErrEmptyInput,
		},
		{
			name:       "whitespace-only input is rejected",
			bad_throw:  "   ",
			want_error: ErrEmptyInput,
		},
		{
			name:       "unknown symbol is rejected",
			bad_throw:  "hello",
			want_error: ErrUnknownSymbol,
		},
		{
			name:       "strike after a partial frame is rejected",
			setup:      []string{"4"},
			bad_throw:  "X",
			want_error: ErrStrikeNotAllowed,
		},
		{
			name:       "spare as the first throw is rejected",
			bad_throw:  "/",
			want_error: ErrSpareNotAllowed,
		},
		{
			name:       "pins above those standing are rejected",
			setup:      []string{"4"},
			bad_throw:  "7",
			want_error: ErrPinsOutOfRange,
		},
		{
			name:       "more than ten pins on the first throw is rejected",
			bad_throw:  "11",
			want_error: ErrPinsOutOfRange,
		},
		{
			name:       "negative pins are rejected",
			bad_throw:  "-1",
			want_error: ErrPinsOutOfRange,
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.name, func(t *testing.T) {
			game := NewGame()
			play_all(t, game, test_case.setup)

			pins_before := game.PinsStanding()
			frame_before := game.CurrentFrameNumber()

			if got_error := game.RecordThrow(test_case.bad_throw); got_error != test_case.want_error {
				t.Fatalf("RecordThrow(%q) error = %v, want %v", test_case.bad_throw, got_error, test_case.want_error)
			}

			// A rejected throw must record nothing and leave the game where it was.
			if game.PinsStanding() != pins_before {
				t.Errorf("pins standing changed after a rejected throw: got %d, want %d", game.PinsStanding(), pins_before)
			}
			if game.CurrentFrameNumber() != frame_before {
				t.Errorf("frame advanced after a rejected throw: got %d, want %d", game.CurrentFrameNumber(), frame_before)
			}
		})
	}
}

// TestGameInProgress checks the state a game reports while it is still being
// played: it is not over, it advances through frames, and only the frames
// played so far are scored.
func TestGameInProgress(t *testing.T) {
	game := NewGame()

	if game.IsGameOver() {
		t.Fatal("a brand-new game should not be over")
	}
	if game.CurrentFrameNumber() != 1 {
		t.Errorf("a new game should start on frame 1, got %d", game.CurrentFrameNumber())
	}

	// Three complete open frames: 1,4 | 4,5 | 6,3.
	play_all(t, game, []string{"1", "4", "4", "5", "6", "3"})

	if game.IsGameOver() {
		t.Fatal("the game should still be in progress after three of ten frames")
	}
	if game.CurrentFrameNumber() != 4 {
		t.Errorf("expected to have advanced to frame 4, got %d", game.CurrentFrameNumber())
	}
	want_scores := []int{5, 14, 23}
	if got_scores := game.Scores(); !reflect.DeepEqual(got_scores, want_scores) {
		t.Errorf("scores while in progress = %v, want %v", got_scores, want_scores)
	}

	// A single throw into the fourth frame: still that frame, still not over.
	play_all(t, game, []string{"7"})

	if game.IsGameOver() {
		t.Fatal("the game should not be over mid-way through the fourth frame")
	}
	if game.CurrentFrameNumber() != 4 {
		t.Errorf("expected to still be playing frame 4, got %d", game.CurrentFrameNumber())
	}
	if game.PinsStanding() != 3 {
		t.Errorf("expected 3 pins standing after a 7 on the first throw, got %d", game.PinsStanding())
	}
}

func TestRecordThrowRejectsThrowsAfterGameOver(t *testing.T) {
	game := NewGame()
	play_all(t, game, []string{"X", "X", "X", "X", "X", "X", "X", "X", "X", "X", "X", "X"})

	if record_error := game.RecordThrow("5"); record_error != ErrGameOver {
		t.Errorf("expected ErrGameOver after a finished game, got %v", record_error)
	}
}

func TestTenthFrameBonusThrows(t *testing.T) {
	open_nine := []string{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0"}

	t.Run("strike grants two bonus throws", func(t *testing.T) {
		game := NewGame()
		play_all(t, game, open_nine)

		if game.ExtraThrowsGranted() != 0 || game.IsBonusThrow() {
			t.Fatal("no bonus should be granted before the tenth-frame strike")
		}
		play_all(t, game, []string{"X"})
		if game.ExtraThrowsGranted() != 2 {
			t.Errorf("strike should grant 2 bonus throws, got %d", game.ExtraThrowsGranted())
		}
		if !game.IsBonusThrow() {
			t.Error("the throw after a tenth-frame strike should be a bonus throw")
		}
	})

	t.Run("spare grants one bonus throw", func(t *testing.T) {
		game := NewGame()
		play_all(t, game, open_nine)
		play_all(t, game, []string{"4"})
		if game.ExtraThrowsGranted() != 0 || game.IsBonusThrow() {
			t.Fatal("no bonus after only the first tenth-frame throw of an open attempt")
		}
		play_all(t, game, []string{"/"})
		if game.ExtraThrowsGranted() != 1 {
			t.Errorf("spare should grant 1 bonus throw, got %d", game.ExtraThrowsGranted())
		}
		if !game.IsBonusThrow() {
			t.Error("the throw after a tenth-frame spare should be a bonus throw")
		}
	})

	t.Run("open tenth frame grants no bonus", func(t *testing.T) {
		game := NewGame()
		play_all(t, game, open_nine)
		play_all(t, game, []string{"3", "4"})
		if game.ExtraThrowsGranted() != 0 {
			t.Errorf("an open tenth frame should grant no bonus, got %d", game.ExtraThrowsGranted())
		}
		if !game.IsGameOver() {
			t.Error("an open tenth frame should end the game after two throws")
		}
	})
}

func TestPinsStandingResetsOnTenthFrameStrike(t *testing.T) {
	game := NewGame()
	// Play nine open frames so the game is on the tenth frame.
	play_all(t, game, []string{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0"})

	if game.CurrentFrameNumber() != 10 {
		t.Fatalf("expected to be on frame 10, got %d", game.CurrentFrameNumber())
	}
	if game.PinsStanding() != total_pins {
		t.Fatalf("expected 10 pins standing at the start of frame 10, got %d", game.PinsStanding())
	}
	if record_error := game.RecordThrow("X"); record_error != nil {
		t.Fatalf("strike on the tenth frame was rejected: %v", record_error)
	}
	if game.PinsStanding() != total_pins {
		t.Errorf("expected the pins to reset to 10 after a tenth-frame strike, got %d", game.PinsStanding())
	}
}
