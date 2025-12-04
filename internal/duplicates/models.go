package duplicates

// Clone represents a single instance of a duplicated code block.
type Clone struct {
	Filename   string `json:"filename"`
	StartLine  int    `json:"start_line"`
	EndLine    int    `json:"end_line"`
	TokenCount int    `json:"token_count"`
	// We could add Snippet here later if needed
}

// CloneGroup represents a set of identical code blocks found in the codebase.
type CloneGroup struct {
	Hash      string  `json:"hash"`
	Score     int     `json:"score"`
	Instances []Clone `json:"instances"`
}
