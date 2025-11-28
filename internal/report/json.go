package report

import (
	"encoding/json"
	"io"

	"github.com/larsartmann/duplicates/internal/duplicates"
)

func ToJSON(w io.Writer, groups []duplicates.CloneGroup) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(groups)
}
