package asciify

import (
	"testing"
)

func TestAsciiConvert(t *testing.T) {
	if Convert("Test") != "Test" {
	    t.Fatal("Error in converting plain ASCII string")
    }
}
