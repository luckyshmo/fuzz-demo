package main

import (
	"testing"
	"unicode/utf8"
)

func TestReverse(t *testing.T) {
	rev := Reverse("test")
	if Reverse("test") != "tset" {
		t.Errorf("Reverse: %q, want %q", rev, "test")
	}
}

func FuzzReverse(f *testing.F) {
	f.Fuzz(func(t *testing.T, orig string) {
		rev := Reverse(orig)
		doubleRev := Reverse(rev)

		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", rev, doubleRev)
		}

		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produces invalid UTF-8 string %q", rev)
		}
	})
}
