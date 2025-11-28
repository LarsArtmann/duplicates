package report

import (
	"fmt"
	"io"

	"github.com/larsartmann/duplicates/internal/duplicates"
)

func ToPlumbing(w io.Writer, groups []duplicates.CloneGroup) error {
	for _, g := range groups {
		for _, inst := range g.Instances {
			// Format: filename:line:start-end score
			fmt.Fprintf(w, "%s:%d:%d-%d %d\n", inst.Filename, inst.StartLine, inst.StartLine, inst.EndLine, g.Score)
		}
	}
	return nil
}
