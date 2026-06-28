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
