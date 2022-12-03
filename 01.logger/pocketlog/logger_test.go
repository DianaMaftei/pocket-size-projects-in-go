package pocketlog_test

import (
	"pocketlog/pocketlog"
	"testing"
)

const (
	debugMessage = "Why write I still all one, ever the same,"
	infoMessage  = "And keep invention in a noted weed,"
	errorMessage = "That every word doth almost tell my name,"
)

func ExampleLogger_Debugf() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug)
	debugLogger.Debugf("Hello, %s", "world")
	// Output: [DEBUG] Hello, world
}

func TestLogger_DebugfInfofErrorf(t *testing.T) {
	type testCase struct {
		level    pocketlog.Level
		expected string
	}

	tt := map[string]testCase{
		"debug": {
			level:    pocketlog.LevelDebug,
			expected: "[DEBUG] " + debugMessage + "\n" + "[INFO] " + infoMessage + "\n" + "[ERROR] " + errorMessage + "\n",
		},
		"info": {
			level:    pocketlog.LevelInfo,
			expected: "[INFO] " + infoMessage + "\n" + "[ERROR] " + errorMessage + "\n",
		},
		"error": {
			level:    pocketlog.LevelError,
			expected: "[ERROR] " + errorMessage + "\n",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			tw := &testWriter{}

			testLogger := pocketlog.New(tc.level, pocketlog.WithOutput(tw))
			testLogger.Debugf(debugMessage)
			testLogger.Infof(infoMessage)
			testLogger.Errorf(errorMessage)

			if tw.contents != tc.expected {
				t.Errorf("invalid contents, expected %q, got %q", tc.expected, tw.contents)
			}
		})
	}
}

func TestLogger_Logf(t *testing.T) {
	type testCase struct {
		level    pocketlog.Level
		expected string
	}

	tt := map[string]testCase{
		"debug": {
			level:    pocketlog.LevelDebug,
			expected: "[DEBUG] " + debugMessage + "\n" + "[INFO] " + infoMessage + "\n" + "[ERROR] " + errorMessage + "\n",
		},
		"info": {
			level:    pocketlog.LevelInfo,
			expected: "[INFO] " + infoMessage + "\n" + "[ERROR] " + errorMessage + "\n",
		},
		"error": {
			level:    pocketlog.LevelError,
			expected: "[ERROR] " + errorMessage + "\n",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			tw := &testWriter{}

			testLogger := pocketlog.New(tc.level, pocketlog.WithOutput(tw))
			testLogger.Logf(pocketlog.LevelDebug, debugMessage)
			testLogger.Logf(pocketlog.LevelInfo, infoMessage)
			testLogger.Logf(pocketlog.LevelError, errorMessage)

			if tw.contents != tc.expected {
				t.Errorf("invalid contents, expected %q, got %q", tc.expected, tw.contents)
			}
		})
	}
}

func TestLogger_LogfMaxLength(t *testing.T) {
	type testCase struct {
		message  string
		expected string
	}

	testCases := map[string]testCase{
		"english": {
			message:  "This very long message",
			expected: "[INFO] This ve\n",
		},
		"japanese": {
			message:  "このとても長いメッセージ",
			expected: "[INFO] このとても長い\n",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			tw := &testWriter{}
			testLogger := pocketlog.New(pocketlog.LevelInfo, pocketlog.WithOutput(tw), pocketlog.WithMaxLength(7))
			testLogger.Logf(pocketlog.LevelInfo, tc.message)
			if tw.contents != tc.expected {
				t.Errorf("invalid contents, expected %q, got %q", tc.expected, tw.contents)
			}
		})
	}
}

type testWriter struct {
	contents string
}

func (tw *testWriter) Write(p []byte) (n int, err error) {
	tw.contents = tw.contents + string(p)
	return len(p), nil
}
