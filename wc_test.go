package main

import (
	"testing"

	//"app/cmd"
    W "app/cmd/workers"
)

func TestWc(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, actual, expected int) {
		t.Helper()
		if actual != expected {
			t.Errorf("Actual:%d Expected:%d", actual, expected)
		}
	}

    //Tesing all the functions against an empty file.
	t.Run("counting words in a file", func(t *testing.T) {
		actual := W.WordsInLines("./testfiles/empty.txt")
		expected := 0
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("counting lines in a file", func(t *testing.T) {
		actual := W.CountLines("./testfiles/empty.txt")
		expected := 0
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("counting characters in a file", func(t *testing.T) {
		actual := W.CountChars("./testfiles/empty.txt")
		expected := 0
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("counting bytes in a file", func(t *testing.T) {
		actual := W.PrintBytes("./testfiles/empty.txt")
		expected := 0
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("counting the max line length in a file", func(t *testing.T) {
		actual := W.MaxLine("./testfiles/empty.txt")
		expected := 0
		assertCorrectMessage(t, actual, expected)
	})
}
