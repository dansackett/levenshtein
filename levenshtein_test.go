package levenshtein

import "testing"

func TestLevenShtein(t *testing.T) {
	if CalculateDistance("CA", "ABC") != 3 {
		t.Errorf("Distance should be 3, got %d", CalculateDistance("CA", "ABC"))
	}

	if CalculateDistance("", "ABC") != 3 {
		t.Errorf("Distance should be 3, got %d", CalculateDistance("", "ABC"))
	}

	if CalculateDistance("CA", "") != 2 {
		t.Errorf("Distance should be 2, got %d", CalculateDistance("CA", ""))
	}

	if CalculateDistance("DOG", "CAT") != 3 {
		t.Errorf("Distance should be 3, got %d", CalculateDistance("DOG", "CAT"))
	}
}
