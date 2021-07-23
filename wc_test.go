package main

import (
	"testing"

	"app/cmd"
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
		actual := cmd.WordsInLines("./testfiles/empty.txt")
		expected := 0
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("counting lines in a file", func(t *testing.T) {
		actual := cmd.CountLines("./testfiles/empty.txt")
		expected := 0
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("counting characters in a file", func(t *testing.T) {
		actual := cmd.CountChars("./testfiles/empty.txt")
		expected := 0
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("counting bytes in a file", func(t *testing.T) {
		actual := cmd.PrintBytes("./testfiles/empty.txt")
		expected := 0
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("counting the max line length in a file", func(t *testing.T) {
		actual := cmd.MaxLine("./testfiles/empty.txt")
		expected := 0
		assertCorrectMessage(t, actual, expected)
	})
}
