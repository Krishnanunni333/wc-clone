package main

import (
	"testing"
    "app/workers"
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
		actual := workers.WordsInLines(workers.OpenFile("./testfiles/empty.txt"))
		expected := 0
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("counting lines in a file", func(t *testing.T) {
		actual := workers.CountLines(workers.OpenFile("./testfiles/empty.txt"))
		expected := 0
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("counting characters in a file", func(t *testing.T) {
		actual := workers.CountChars(workers.OpenFile("./testfiles/empty.txt"))
		expected := 0
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("counting bytes in a file", func(t *testing.T) {
		actual := workers.PrintBytes(workers.OpenFile("./testfiles/empty.txt"))
		expected := 0
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("counting the max line length in a file", func(t *testing.T) {
		actual := workers.MaxLine(workers.OpenFile("./testfiles/empty.txt"))
		expected := 0
		assertCorrectMessage(t, actual, expected)
	})
}
