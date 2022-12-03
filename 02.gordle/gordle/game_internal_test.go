package gordle

import (
	"errors"
	"golang.org/x/exp/slices"
	"strings"
	"testing"
)

func TestGameAsk(t *testing.T) {
	testCases :=
		map[string]struct {
			input string
			want  []rune
		}{
			"5 characters in english": {
				input: "HELLO",
				want:  []rune("HELLO"),
			},
			"5 characters in arabic": {
				input: "مرحبا",
				want:  []rune("مرحبا"),
			},
			"5 characters in japanese": {
				input: "こんにちは",
				want:  []rune("こんにちは"),
			},
			"3 characters in japanese": {
				input: "こんに\nこんにちは",
				want:  []rune("こんにちは"),
			},
		}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			g, _ := New(strings.NewReader(tc.input), []string{"HELLO", "SALUT", "ПРИВЕТ", "ΧΑΙΡΕ"}, 0)

			got := g.ask()

			if !slices.Equal(got, tc.want) {
				t.Errorf("ask() got = %v, expected %v", string(got), string(tc.want))
			}
		})
	}
}

func TestGameValidateGuess(t *testing.T) {
	game, _ := New(strings.NewReader("smth"), []string{"HELLO", "SALUT", "ПРИВЕТ", "ΧΑΙΡΕ"}, 5)

	testCases := map[string]struct {
		guess []rune
		want  error
	}{
		"guess is desired length": {
			guess: []rune("hello"),
			want:  nil,
		},
		"guess is shorter than desired length": {
			guess: []rune("hi"),
			want:  errInvalidWordLength,
		},
		"guess is longer than desired length": {
			guess: []rune("greetings"),
			want:  errInvalidWordLength,
		},
		"guess is empty": {
			guess: []rune(""),
			want:  errInvalidWordLength,
		},
		"guess is nil": {
			guess: nil,
			want:  errInvalidWordLength,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := game.validateGuess(tc.guess)

			if !errors.Is(err, tc.want) {
				t.Errorf("error does not match expected type, got %v, expected %v", err, tc.want)
			}
		})
	}

}

func Test_splitToUppercaseCharacters(t *testing.T) {
	testCases := map[string]struct {
		word     string
		expected []rune
	}{
		"string is a mix of upper and lowercase characters": {
			word:     "AppLe",
			expected: []rune("APPLE"),
		},
		"string is all uppercase characters": {
			word:     "APPLE",
			expected: []rune("APPLE"),
		},
		"string is empty": {
			word:     "",
			expected: []rune(""),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			result := splitToUppercaseCharacters(tc.word)

			if !slices.Equal(result, tc.expected) {
				t.Errorf("splitToUppercaseCharacters() got = %v, expected %v", string(result), string(tc.expected))
			}
		})
	}
}

func Test_computeFeedback(t *testing.T) {
	testCases := map[string]struct {
		guess            string
		solution         string
		expectedFeedback feedback
	}{
		"nominal": {
			guess:            "HERTZ",
			solution:         "HERTZ",
			expectedFeedback: feedback{correctPosition, correctPosition, correctPosition, correctPosition, correctPosition},
		},
		"double character": {
			guess:            "HELLO",
			solution:         "HELLO",
			expectedFeedback: feedback{correctPosition, correctPosition, correctPosition, correctPosition, correctPosition},
		},
		"double character with wrong answer": {
			guess:            "HELLL",
			solution:         "HELLO",
			expectedFeedback: feedback{correctPosition, correctPosition, correctPosition, correctPosition, absentCharacter},
		},
		"five identical, but only two are there": {
			guess:            "LLLLL",
			solution:         "HELLO",
			expectedFeedback: feedback{absentCharacter, absentCharacter, correctPosition, correctPosition, absentCharacter},
		},
		"two identical, but not in the right position (from left to right)": {
			guess:            "HLLEO",
			solution:         "HELLO",
			expectedFeedback: feedback{correctPosition, wrongPosition, correctPosition, wrongPosition, correctPosition},
		},
		"three identical, but not in the right position (from left to right)": {
			guess:            "HLLLO",
			solution:         "HELLO",
			expectedFeedback: feedback{correctPosition, absentCharacter, correctPosition, correctPosition, correctPosition},
		},
		"one correct, one incorrect, one absent (left of the correct)": {
			guess:            "LLLWW",
			solution:         "HELLO",
			expectedFeedback: feedback{wrongPosition, absentCharacter, correctPosition, absentCharacter, absentCharacter},
		},
		"swapped characters": {
			guess:            "HOLLE",
			solution:         "HELLO",
			expectedFeedback: feedback{correctPosition, wrongPosition, correctPosition, correctPosition, wrongPosition},
		},
		"absent character": {
			guess:            "HULFO",
			solution:         "HELFO",
			expectedFeedback: feedback{correctPosition, absentCharacter, correctPosition, correctPosition, correctPosition},
		},
		"absent character and incorrect": {
			guess:            "HULPP",
			solution:         "HELPO",
			expectedFeedback: feedback{correctPosition, absentCharacter, correctPosition, correctPosition, absentCharacter},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			result := computeFeedback([]rune(tc.guess), []rune(tc.solution))

			if !tc.expectedFeedback.Equal(result) {
				t.Errorf("computeFeedback() got = %v, expected %v", string(result), string(tc.expectedFeedback))
			}
		})
	}
}
