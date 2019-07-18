package levenshtein

// The Damerau–Levenshtein distance is a metric for measuring the edit distance
// between two strings. It consists of the minimum number of operations it
// would take to transpose one string to another. These operations include:
//
// - Insertion
// - Deletion
// - Substitution
// - Transposition of adjacent characters
//
// The traditional Levenshtein distance formula does not include transpositions
// like this formula does. This leads to a stricter set of rules while
// supporting a fourth rule of transposing characters.

// findMin is a helper function to get the minimum value from a slice of ints
func findMin(vals []int) int {
	var m int
	for i, e := range vals {
		if i == 0 || e < m {
			m = e
		}
	}
	return m
}

// generateRange is a helper function to make a slice of ints from 0 to the limit
func generateRange(limit int) []int {
	var r []int
	for i := 0; i < limit; i++ {
		r = append(r, i)
	}
	return r
}

// CalculateDistance calculates the optimal string alignment distance using the
// Damerau–Levenshtein formula.
func CalculateDistance(source, target string) int {
	sourceRange := generateRange(len(source) + 1)
	targetRange := generateRange(len(target) + 1)
	matrix := make([][]int, len(source)+1)

	// Build our initial matrix
	for _, i := range sourceRange {
		matrix[i] = make([]int, len(target)+1)
		for _, j := range targetRange {
			if j == 0 {
				matrix[i][j] = i
			} else {
				matrix[i][j] = j
			}
		}
	}

	// Determine the cost for insertion, deletion, substitution, and
	// transposition and use the minimum value as the true editing distance for
	// the given character position in the matrix.
	for _, i := range sourceRange[1:] {
		for _, j := range targetRange[1:] {
			var subTransCost int

			if source[i-1] == target[j-1] {
				subTransCost = 0
			} else {
				subTransCost = 1
			}

			delDist := matrix[i-1][j] + 1
			insDist := matrix[i][j-1] + 1
			subDist := matrix[i-1][j-1] + subTransCost

			matrix[i][j] = findMin([]int{delDist, insDist, subDist})

			// This is what differs from the traditional Levenshtein formula
			// and gets the transposition distance
			if i > 1 && j > 1 && source[i-1] == target[j-2] && source[i-2] == target[j-1] {
				transDist := matrix[i-2][j-2] + subTransCost
				matrix[i][j] = findMin([]int{matrix[i][j], transDist})
			}
		}
	}

	// Our resulting maximum distance is found at the len of both strings
	return matrix[len(source)][len(target)]
}
