package pathbuilder

import "testing"

func TestCustom(t *testing.T) {
	for _, test := range customTests {
		actual := Custom(test.separator, test.components)
		if actual != test.expected {
			t.Errorf("Custom test %q %q, expected %q, actual %q",
				test.separator, test.components, test.expected, actual)
		}
	}
}

func BenchmarkCustom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range customTests {
			Custom(test.separator, test.components)
		}
	}
}
