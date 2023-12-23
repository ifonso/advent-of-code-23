package main

import "testing"

func TestReverse(t *testing.T) {
	word := "lana"
	reversedAssert := "anal"

	reversed := Reverse(word)

	if reversedAssert != reversed {
		t.Errorf("Expected reversed 'lana' to be 'anal', but it is not.")
	}
}

func TestContains(t *testing.T) {
	phrase := "onwwayagainafjsbndj"
	containingWord := "way"
	notContainingWord := "two"

	containsFirst, firstIndex := Contains(phrase, containingWord)
	containsSecond, secondIndex := Contains(phrase, notContainingWord)

	if !containsFirst || firstIndex != 3 {
		t.Errorf("Expected phase to contain 'way' and the index of 'w' to be 3.")
	}

	if containsSecond || secondIndex != -1 {
		t.Errorf("Expected phase to not contain 'two' and the returned index to be -1.")
	}
}
