package pathbuilder

import "testing"

func TestCustom(t *testing.T) {
	for _, testCase := range customTestCases {
		actual := Custom(testCase.separator, testCase.components)
		if actual != testCase.expected {
			t.Errorf("Custom test %q %q, expected %q, actual %q",
				testCase.separator, testCase.components, testCase.expected, actual)
		}
	}
}

func BenchmarkCustom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testCase := range customTestCases {
			Custom(testCase.separator, testCase.components)
		}
	}
}

func TestURL(t *testing.T) {
	for _, testCase := range urlTestCases {
		actual := URL(testCase.components)
		if actual != testCase.expected {
			t.Errorf("URL test %q, expected %q, actual %q",
				testCase.components, testCase.expected, actual)
		}
	}
}

func BenchmarkURL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testCase := range urlTestCases {
			URL(testCase.components)
		}
	}
}

func TestSystem(t *testing.T) {
	for _, testCase := range systemTestCases {
		actual := System(testCase.components)
		if actual != testCase.expected {
			t.Errorf("System test %q, expected %q, actual %q",
				testCase.components, testCase.expected, actual)
		}
	}
}

func BenchmarkSystem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testCase := range systemTestCases {
			System(testCase.components)
		}
	}
}

func TestUnix(t *testing.T) {
	for _, testCase := range unixTestCases {
		actual := Unix(testCase.components)
		if actual != testCase.expected {
			t.Errorf("Unix test %q, expected %q, actual %q",
				testCase.components, testCase.expected, actual)
		}
	}
}

func BenchmarkUnix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testCase := range unixTestCases {
			Unix(testCase.components)
		}
	}
}

func TestWindows(t *testing.T) {
	for _, testCase := range windowsTestCases {
		actual := Windows(testCase.components)
		if actual != testCase.expected {
			t.Errorf("Windows test %q, expected %q, actual %q",
				testCase.components, testCase.expected, actual)
		}
	}
}

func BenchmarkWindows(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testCase := range windowsTestCases {
			Windows(testCase.components)
		}
	}
}
