package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/golangci/dupl/job"
	"github.com/golangci/dupl/printer"
	"github.com/golangci/dupl/syntax"
)

func main() {
	// 1. Find all Go files
	files := make(chan string)
	go func() {
		filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
			if err == nil && !info.IsDir() && strings.HasSuffix(path, ".go") && !strings.Contains(path, "vendor/") {
				files <- path
			}
			return nil
		})
		close(files)
	}()

	// 2. Run analysis
	schan := job.Parse(files)
	t, data, done := job.BuildTree(schan)
	<-done
	t.Update(&syntax.Node{Type: -1})

	// 3. Find duplicates (threshold 15 tokens)
	mchan := t.FindDuplOver(15)
	groups := make(map[string][][]*syntax.Node)
	for m := range mchan {
		match := syntax.FindSyntaxUnits(*data, m, 15)
		if len(match.Frags) > 0 {
			groups[match.Hash] = append(groups[match.Hash], match.Frags...)
		}
	}

	// 4. Sort by severity (Score = tokens * instances)
	type Group struct {
		Clones [][]*syntax.Node
		Score  int
	}
	var sorted []Group
	for _, rawClones := range groups {
		clones := unique(rawClones)
		if len(clones) > 1 {
			// Score: tokens * instances.
			// Each node has 'Owns' descendants. We sum them to get total token mass.
			tokens := 0
			for _, n := range clones[0] {
				tokens += n.Owns + 1
			}
			score := tokens * len(clones)
			sorted = append(sorted, Group{Clones: clones, Score: score})
		}
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Score > sorted[j].Score // Descending
	})

	// 5. Output reports
	outputs := []struct {
		Name string
		Fn   func(io.Writer, printer.ReadFile) printer.Printer
	}{
		{"reports/duplicates.txt", printer.NewText},
		{"reports/duplicates.html", printer.NewHTML},
		{"reports/duplicates.plumbing", printer.NewPlumbing},
	}

	for _, out := range outputs {
		f, err := os.Create(out.Name)
		if err != nil {
			log.Fatal(err)
		}
		p := out.Fn(f, os.ReadFile)
		p.PrintHeader()
		for _, g := range sorted {
			p.PrintClones(g.Clones)
		}
		p.PrintFooter()
		f.Close()
	}

	// JSON Output
	f, _ := os.Create("reports/duplicates.json")
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	type JsonClone struct {
		Filename string `json:"filename"`
		Tokens   int    `json:"tokens"`
	}
	type JsonGroup struct {
		Score  int         `json:"score"`
		Tokens int         `json:"tokens"`
		Clones []JsonClone `json:"clones"`
	}
	var jsonGroups []JsonGroup
	for _, g := range sorted {
		var jg JsonGroup
		jg.Score = g.Score
		
		tokens := 0
		for _, n := range g.Clones[0] {
			tokens += n.Owns + 1
		}
		jg.Tokens = tokens

		for _, c := range g.Clones {
			startNode := c[0]
			jg.Clones = append(jg.Clones, JsonClone{
				Filename: startNode.Filename,
				Tokens:   jg.Tokens,
			}) 
		}
		jsonGroups = append(jsonGroups, jg)
	}
	enc.Encode(jsonGroups)
	f.Close()
	
	fmt.Println("Reports generated in reports/")
}

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
