package template

import (
	"testing"
	"time"
)

func TestToDuration(t *testing.T) {
	from := time.Date(2017, time.November, 15, 23, 0, 0, 0, time.UTC).Unix()

	vals := map[int64]string{
		time.Date(2018, time.November, 15, 23, 0, 0, 0, time.UTC).Unix():   "8760h0m0s",
		time.Date(2017, time.November, 16, 23, 0, 0, 0, time.UTC).Unix():   "24h0m0s",
		time.Date(2017, time.November, 15, 23, 30, 0, 0, time.UTC).Unix():  "30m0s",
		time.Date(2017, time.November, 15, 23, 10, 15, 0, time.UTC).Unix(): "10m15s",
		time.Date(2017, time.October, 15, 23, 0, 0, 0, time.UTC).Unix():    "-744h0m0s",
	}

	for input, want := range vals {
		if got := toDuration(from, input); got != want {
			t.Errorf("Want transform %d-%d to %s, got %s", from, input, want, got)
		}
	}
}

func TestTruncate(t *testing.T) {
	vals := map[string]string{
		"foobarz": "fooba",
		"foöäüüu": "foöäü",
		"üpsßßßk": "üpsßß",
		"1234567": "12345",
		"!'§$%&/": "!'§$%",
	}

	for input, want := range vals {
		if got := truncate(input, 5); got != want {
			t.Errorf("Want transform %s to %s, got %s", input, want, got)
		}
	}
}

func TestNegativeTruncate(t *testing.T) {
	vals := map[string]string{
		"foobarz": "rz",
		"foöäüüu": "üu",
		"üpsßßßk": "ßk",
		"1234567": "67",
		"!'§$%&/": "&/",
	}

	for input, want := range vals {
		if got := truncate(input, -5); got != want {
			t.Errorf("Want transform %s to %s, got %s", input, want, got)
		}
	}
}

func TestSince(t *testing.T) {
	t.Skip()
}

func TestUppercaseFirst(t *testing.T) {
	vals := map[string]string{
		"hello":  "Hello",
		"ßqwert": "ßqwert",
		"üps":    "Üps",
		"12345":  "12345",
		"Foobar": "Foobar",
	}

	for input, want := range vals {
		if got := uppercaseFirst(input); got != want {
			t.Errorf("Want transform %s to %s, got %s", input, want, got)
		}
	}
}

func TestRegexReplace(t *testing.T) {
	expected := "hello-my-String-123"
	actual := regexReplace("(.*?)\\/(.*)", "hello/my-String-123", "$1-$2")
	if actual != "hello-my-String-123" {
		t.Errorf("error, expected %s, got %s", expected, actual)
	}
}
