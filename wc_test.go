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

	t.Run("counting words in a file", func(t *testing.T) {
		actual := cmd.OpenFile("/home/kittu/Desktop/kittu/Projects/app/testfiles/empty.txt", 3)
		expected := 0
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("counting lines in a file", func(t *testing.T) {
		actual := cmd.OpenFile("/home/kittu/Desktop/kittu/Projects/app/testfiles/empty.txt", 4)
		expected := 1
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("counting characters in a file", func(t *testing.T) {
		actual := cmd.OpenFile("/home/kittu/Desktop/kittu/Projects/app/testfiles/empty.txt", 1)
		expected := 13
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("counting bytes in a file", func(t *testing.T) {
		actual := cmd.PrintBytes("/home/kittu/Desktop/kittu/Projects/app/testfiles/empty.txt")
		expected := 13
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("counting the max line length in a file", func(t *testing.T) {
		actual := cmd.OpenFile("/home/kittu/Desktop/kittu/Projects/app/testfiles/empty.txt", 2)
		expected := 5
		assertCorrectMessage(t, actual, expected)
	})
}
