package duplicates

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/golangci/dupl/job"
	"github.com/golangci/dupl/syntax"
)

// Scanner handles the duplicate detection logic.
type Scanner struct {
	Threshold int
	Exclude   []string // TODO: Implement exclusion logic
}

func NewScanner(threshold int) *Scanner {
	if threshold <= 0 {
		threshold = 15
	}
	return &Scanner{Threshold: threshold}
}

// Scan runs the duplicate detection on the given paths.
func (s *Scanner) Scan(root string) ([]CloneGroup, error) {
	// 1. Find files
	files := make(chan string)
	go func() {
		filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil || info.IsDir() {
				return nil
			}
			// Check exclusions
			for _, pattern := range s.Exclude {
				matched, err := filepath.Match(pattern, info.Name())
				if err == nil && matched {
					return nil
				}
				// Also check full path matching if needed?
				// filepath.Match checks against the name.
				// If we want to exclude directories like "internal/duplicates/*", we need to check path.
				// But standard filepath.Match is simple.
			}

			if strings.HasSuffix(path, ".go") && !strings.Contains(path, "vendor/") {
				files <- path
			}
			return nil
		})
		close(files)
	}()

	// 2. Build AST Tree
	schan := job.Parse(files)
	t, data, done := job.BuildTree(schan)
	<-done
	t.Update(&syntax.Node{Type: -1})

	// 3. Find Duplicates
	mchan := t.FindDuplOver(s.Threshold)
	duplGroups := make(map[string][][]*syntax.Node)
	for m := range mchan {
		match := syntax.FindSyntaxUnits(*data, m, s.Threshold)
		if len(match.Frags) > 0 {
			duplGroups[match.Hash] = append(duplGroups[match.Hash], match.Frags...)
		}
	}

	// 4. Process Groups
	var result []CloneGroup

	// Cache file content to avoid re-reading
	fileCache := make(map[string]*LineIndex)

	for hash, rawClones := range duplGroups {
		clones := unique(rawClones)
		if len(clones) <= 1 {
			continue
		}

		// Calculate Score
		// = tokens * instances
		tokens := 0
		for _, n := range clones[0] {
			tokens += n.Owns + 1
		}
		score := tokens * len(clones)

		group := CloneGroup{
			Hash:  hash,
			Score: score,
		}

		for _, c := range clones {
			if len(c) == 0 {
				continue
			}
			startNode := c[0]
			endNode := c[len(c)-1]

			filename := startNode.Filename

			// Get Line Numbers
			idx, ok := fileCache[filename]
			if !ok {
				content, err := os.ReadFile(filename)
				if err != nil {
					// If we can't read the file, default to 0
					// In production we might want to log this
					idx = &LineIndex{newlines: []int{0}}
				} else {
					idx = NewLineIndex(content)
				}
				fileCache[filename] = idx
			}

			group.Instances = append(group.Instances, Clone{
				Filename:   filename,
				StartLine:  idx.Line(startNode.Pos),
				EndLine:    idx.Line(endNode.End),
				TokenCount: tokens,
			})
		}
		result = append(result, group)
	}

	return result, nil
}

// unique filters out subset duplicates, same as original main.go.
func unique(group [][]*syntax.Node) [][]*syntax.Node {
	fileMap := make(map[string]map[int]struct{})
	var newGroup [][]*syntax.Node
	for _, seq := range group {
		node := seq[0]
		file, ok := fileMap[node.Filename]
		if !ok {
			file = make(map[int]struct{})
			fileMap[node.Filename] = file
		}
		if _, ok := file[node.Pos]; !ok {
			file[node.Pos] = struct{}{}
			newGroup = append(newGroup, seq)
		}
	}
	return newGroup
}
