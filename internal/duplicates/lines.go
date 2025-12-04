package duplicates

import "sort"

// LineIndex helps convert byte offsets to line numbers.
type LineIndex struct {
	newlines []int
}

// NewLineIndex creates an index from file content.
func NewLineIndex(content []byte) *LineIndex {
	newlines := make([]int, 0, len(content)/40) // estimate 40 chars per line
	newlines = append(newlines, 0)              // Line 1 starts at 0
	for i, b := range content {
		if b == '\n' {
			newlines = append(newlines, i+1)
		}
	}
	return &LineIndex{newlines: newlines}
}

// Line returns the 1-based line number for a byte offset.
func (li *LineIndex) Line(offset int) int {
	// Find the first newline index that is > offset.
	// The line number is the index of that newline in our array.
	// actually, we want the index i such that newlines[i] <= offset < newlines[i+1]

	// Search returns the first index where f(i) is true.
	idx := sort.Search(len(li.newlines), func(i int) bool {
		return li.newlines[i] > offset
	})
	return idx
}
