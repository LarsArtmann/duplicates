# duplicates

A fast and accurate code duplicate detection tool for Go codebases. Find repeated code blocks, identify technical debt, and improve code maintainability.

## Features

- **Token-sequence based detection** - Finds structural duplicates, not just textual matches
- **Configurable threshold** - Adjust sensitivity based on your needs
- **Multiple output formats** - JSON, HTML, Text, and Plumbing reports
- **Line number tracking** - Precise location of each duplicate instance
- **Scoring system** - Prioritize duplicates by impact (tokens × instances)
- **File exclusion patterns** - Ignore test files, generated code, etc.
- **Fast scanning** - Efficient AST-based analysis

## Installation

```bash
go install github.com/larsartmann/duplicates/cmd/duplicates@latest
```

Or build from source:

```bash
git clone https://github.com/larsartmann/duplicates.git
cd duplicates
go build -o duplicates ./cmd/duplicates
```

## Usage

### Basic Scan

Scan the current directory for duplicates:

```bash
duplicates
```

### Configure Threshold

Set the minimum token sequence size (default: 15):

```bash
duplicates -threshold 20
```

Lower values = more sensitive (finds smaller duplicates)
Higher values = less sensitive (only finds significant duplicates)

### Exclude Files

Ignore specific file patterns:

```bash
duplicates -exclude "*_test.go,generated.go,mock_*.go"
```

### Output Formats

Generate reports in multiple formats simultaneously:

```bash
duplicates -json report.json -html report.html -text report.txt
```

### Verbose Output

See what's happening during the scan:

```bash
duplicates -v
```

### Custom Paths

Specify output paths:

```bash
duplicates -json ./reports/duplicates.json -html ./reports/overview.html
```

## CLI Options

| Flag | Default | Description |
|------|---------|-------------|
| `-threshold` | 15 | Minimum token sequence size to consider as duplicate |
| `-exclude` | "" | Comma-separated file patterns to exclude |
| `-json` | "reports/duplicates.json" | Path to JSON report |
| `-html` | "reports/duplicates.html" | Path to HTML report |
| `-text` | "reports/duplicates.txt" | Path to Text report |
| `-plumbing` | "reports/duplicates.plumbing" | Path to Plumbing report |
| `-v` | false | Verbose output |

## Output Formats

### JSON

Machine-readable format with full details:

```json
[
  {
    "hash": "abc123...",
    "score": 150,
    "instances": [
      {
        "filename": "internal/service/user.go",
        "start_line": 45,
        "end_line": 78,
        "token_count": 50
      },
      {
        "filename": "internal/service/product.go",
        "start_line": 23,
        "end_line": 56,
        "token_count": 50
      }
    ]
  }
]
```

### HTML

Interactive HTML report with syntax highlighting and navigation. Open in your browser to explore duplicates visually.

### Text

Human-readable summary:

```
Duplicate Group #1 (Score: 150)
  internal/service/user.go:45-78
  internal/service/product.go:23-56
```

### Plumbing

Simple, parseable format for scripting:

```
hash\tfilename\tstart_line\tend_line\ttoken_count
```

## Understanding the Output

### Score Calculation

Score = Token Count × Number of Instances

Higher scores indicate:
- More duplicated code (more tokens)
- More widespread duplication (more instances)

Use scores to prioritize refactoring efforts.

### Token Threshold

The `-threshold` value determines what counts as a duplicate:
- **10-15**: Small helper functions, validation logic
- **20-30**: Business logic, medium-sized functions
- **40+**: Large functions, complex algorithms

## Use Cases

- **Code review**: Identify and eliminate duplicates before merging
- **Refactoring**: Prioritize technical debt reduction
- **CI/CD**: Fail builds if duplicate score exceeds threshold
- **Architecture review**: Find opportunities for abstraction
- **Onboarding**: Help new developers understand codebase structure

## Example Workflow

1. **Initial scan**:
   ```bash
   duplicates -threshold 20 -html report.html
   ```

2. **Review results**: Open `report.html` in your browser

3. **Prioritize**: Focus on groups with the highest scores

4. **Refactor**: Extract common code to shared utilities

5. **Verify**: Re-run to confirm duplicates are eliminated

## Development

```bash
# Run tests
go test ./...

# Build
go build ./cmd/duplicates

# Lint
golangci-lint run
```

## Roadmap

- [ ] Fail threshold (exit code 1 if score exceeds limit)
- [ ] Git ignore pattern support
- [ ] Concurrent file scanning
- [ ] stdin input support for file lists
- [ ] Config file support (.duplicatesrc)
- [ ] Language-agnostic mode

## License

[Add your license here]

## Contributing

Contributions welcome! Please read the code and submit pull requests.

---

Built with [golangci/dupl](https://github.com/golangci/dupl) algorithm
