package main

import "testing"

func ExampleMain() {
	main()
	// Output:
	// Hello world
}

func TestGreet_English(t *testing.T) {
	lang := "en"
	expectedGreeting := "Hello world"

	greeting := greet(language(lang))

	if greeting != expectedGreeting {
		// mark this test as failed
		t.Errorf("expected: %q, got: %q", expectedGreeting, greeting)
	}
}

func TestGreet_French(t *testing.T) {
	lang := "fr"
	expectedGreeting := "Bonjour le monde"

	greeting := greet(language(lang))

	if greeting != expectedGreeting {
		// mark this test as failed
		t.Errorf("expected: %q, got: %q", expectedGreeting, greeting)
	}
}

func TestGreet(t *testing.T) {
	type testCase struct {
		lang             language
		expectedGreeting string
	}

	var tests = map[string]testCase{
		"English": {
			lang:             "en",
			expectedGreeting: "Hello world",
		},
		"French": {
			lang:             "fr",
			expectedGreeting: "Bonjour le monde",
		},
		"Akkadian, not supported": {
			lang:             "akk",
			expectedGreeting: `unsupported language: "akk"`,
		},
		"Greek": {
			lang:             "el",
			expectedGreeting: "Χαίρετε Κόσμε",
		},
		"Hebrew": {
			lang:             "he",
			expectedGreeting: "שלום עולם",
		},
		"Urdu": {
			lang:             "ur",
			expectedGreeting: "ہیلو دنیا",
		},
		"Vietnamese": {
			lang:             "vi",
			expectedGreeting: "Xin ch→o Thế Giới",
		},
		"Empty": {
			lang:             "",
			expectedGreeting: `unsupported language: ""`,
		},
	}

	// range over all the scenarios
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			greeting := greet(tc.lang)

			if greeting != tc.expectedGreeting {
				t.Errorf("expected: %q, got: %q", tc.expectedGreeting, greeting)
			}
		})
	}
}
