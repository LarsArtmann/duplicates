package report

import (
	"fmt"
	"html"
	"io"
	"os"

	"github.com/larsartmann/duplicates/internal/duplicates"
)

func ToHTML(w io.Writer, groups []duplicates.CloneGroup) error {
	fmt.Fprint(w, `<!DOCTYPE html>
<html>
<head>
<style>
body { font-family: sans-serif; padding: 20px; }
.group { border: 1px solid #ccc; margin-bottom: 20px; padding: 10px; }
.header { background: #eee; padding: 5px; font-weight: bold; }
.clone { margin: 10px 0; border-left: 3px solid #007bff; padding-left: 10px; }
.code { background: #f8f8f8; padding: 10px; overflow-x: auto; font-family: monospace; white-space: pre; }
</style>
</head>
<body>
<h1>Duplicates Report</h1>
`)

	for _, g := range groups {
		fmt.Fprintf(w, `<div class="group">
<div class="header">Score: %d | Instances: %d</div>
`, g.Score, len(g.Instances))
		
		for _, inst := range g.Instances {
			fmt.Fprintf(w, `<div class="clone">
<div>%s:%d-%d</div>
`, inst.Filename, inst.StartLine, inst.EndLine)
			
			// Optional: Read file snippet (this is slow for many duplicates, but useful)
			// For now, let's just show file/line to keep it simple and fast.
			// If we want snippets, we need to pass a FileProvider.
			
			// Let's try to read the file content for the snippet
			content, err := os.ReadFile(inst.Filename)
			if err == nil {
				lines := splitLines(content)
				start := inst.StartLine - 1
				end := inst.EndLine
				if start < 0 { start = 0 }
				if end > len(lines) { end = len(lines) }
				
				fmt.Fprintln(w, `<div class="code">`)
				for i := start; i < end; i++ {
					fmt.Fprintf(w, "%4d | %s\n", i+1, html.EscapeString(string(lines[i])))
				}
				fmt.Fprintln(w, `</div>`)
			}
			
			fmt.Fprintln(w, `</div>`)
		}
		fmt.Fprintln(w, `</div>`)
	}

	fmt.Fprint(w, `</body></html>`)
	return nil
}

func splitLines(content []byte) [][]byte {
	var lines [][]byte
	start := 0
	for i, b := range content {
		if b == '\n' {
			lines = append(lines, content[start:i])
			start = i + 1
		}
	}
	if start < len(content) {
		lines = append(lines, content[start:])
	}
	return lines
}
