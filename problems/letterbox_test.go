package problems

import (
	"testing"
)

func TestMakeLetterbox(t *testing.T) {
	_, err := MakeLetterBox("quadratic", "clown")

	if err != nil {
		t.Fatal("The input is expected to pass, but the function returned an error", err)
	}

	_, err = MakeLetterBox("quadratic", "down")

	if err.Error() != "Last character of word 1 must equal first character of word 2" {
		t.Fatal("Function returned the wrong error:", err)
	}

}
