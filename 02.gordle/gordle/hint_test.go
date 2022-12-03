package gordle

import "testing"

func TestFeedback_String(t *testing.T) {
	testCases := map[string]struct {
		feedback feedback
		expected string
	}{
		"all correct hints": {
			feedback: []hint{correctPosition, correctPosition, correctPosition},
			expected: "💚💚💚",
		}, "various hints": {
			feedback: []hint{wrongPosition, correctPosition, absentCharacter},
			expected: "🟡💚⬜️",
		},
		"invalid hint": {
			feedback: []hint{42},
			expected: "💔",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := tc.feedback.String()

			if actual != tc.expected {
				t.Errorf("invalid result, expected %s but got %s", tc.expected, actual)
			}
		})
	}
}
