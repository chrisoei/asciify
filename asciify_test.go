package asciify

import (
	"testing"
)

func TestAsciiConvert(t *testing.T) {
	if Convert("Test") != "Test" {
		t.Fatal("Error in converting plain ASCII string")
	}
}

func TestIsAscii(t *testing.T) {
	if !IsAscii("Test") {
		t.Fatal("Test string should be detected as plain ASCII")
	}
}

func TestNonAscii(t *testing.T) {
	if IsAscii("\u2019") {
		t.Fatal("Test string should not be detected as plain ASCII")
	}
}
