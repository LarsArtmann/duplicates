package report

import (
	"fmt"
	"io"

	"github.com/larsartmann/duplicates/internal/duplicates"
)

func ToText(w io.Writer, groups []duplicates.CloneGroup) error {
	for _, g := range groups {
		fmt.Fprintf(w, "Found %d clones (score: %d):\n", len(g.Instances), g.Score)
		for _, inst := range g.Instances {
			fmt.Fprintf(w, "  %s:%d-%d\n", inst.Filename, inst.StartLine, inst.EndLine)
		}
		fmt.Fprintln(w)
	}
	return nil
}
