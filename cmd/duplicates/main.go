package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/larsartmann/duplicates/internal/duplicates"
	"github.com/larsartmann/duplicates/internal/report"
)

func main() {
	var (
		threshold    = flag.Int("threshold", 15, "Minimum token sequence size to consider as a duplicate")
		jsonPath     = flag.String("json", "reports/duplicates.json", "Path to output JSON report")
		textPath     = flag.String("text", "reports/duplicates.txt", "Path to output Text report")
		htmlPath     = flag.String("html", "reports/duplicates.html", "Path to output HTML report")
		plumbingPath = flag.String("plumbing", "reports/duplicates.plumbing", "Path to output Plumbing report")
		exclude      = flag.String("exclude", "", "Comma-separated list of file patterns to exclude (e.g. '*_test.go,generated.go')")
		verbose      = flag.Bool("v", false, "Verbose output")
	)
	flag.Parse()

	if *verbose {
		log.Printf("Scanning with threshold %d...", *threshold)
	}

	scanner := duplicates.NewScanner(*threshold)
	if *exclude != "" {
		// Split by comma and trim spaces
		patterns := strings.Split(*exclude, ",")
		for i, p := range patterns {
			patterns[i] = strings.TrimSpace(p)
		}
		scanner.Exclude = patterns
	}

	groups, err := scanner.Scan(".")
	if err != nil {
		log.Fatal(err)
	}

	// Sort by Score
	sort.Slice(groups, func(i, j int) bool {
		return groups[i].Score > groups[j].Score
	})

	// Ensure reports directory exists
	if err := os.MkdirAll("reports", 0o755); err != nil {
		log.Fatal(err)
	}

	// Generate JSON
	if *jsonPath != "" {
		f, err := os.Create(*jsonPath)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		if err := report.ToJSON(f, groups); err != nil {
			log.Fatal(err)
		}
		if *verbose {
			log.Printf("JSON report written to %s", *jsonPath)
		}
	}

	// Generate Text
	if *textPath != "" {
		f, err := os.Create(*textPath)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		if err := report.ToText(f, groups); err != nil {
			log.Fatal(err)
		}
		if *verbose {
			log.Printf("Text report written to %s", *textPath)
		}
	}

	// Generate HTML
	if *htmlPath != "" {
		f, err := os.Create(*htmlPath)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		if err := report.ToHTML(f, groups); err != nil {
			log.Fatal(err)
		}
		if *verbose {
			log.Printf("HTML report written to %s", *htmlPath)
		}
	}

	// Generate Plumbing
	if *plumbingPath != "" {
		f, err := os.Create(*plumbingPath)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		if err := report.ToPlumbing(f, groups); err != nil {
			log.Fatal(err)
		}
		if *verbose {
			log.Printf("Plumbing report written to %s", *plumbingPath)
		}
	}

	fmt.Printf("Scan complete. Found %d duplicate groups.\n", len(groups))
}
