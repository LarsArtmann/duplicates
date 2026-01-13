# ⚠️ ARCHIVED - Please Migrate to art-dupl

> **Status**: This project is **ARCHIVED** as of January 14, 2026
> **Action**: Please migrate to **[art-dupl](https://github.com/LarsArtmann/art-dupl)** - all functionality is now available there

---

## Migration Information

This repository has been archived because **all functionality has been successfully merged into the more feature-rich `art-dupl` project**.

### Why Archive?

- **Complete feature parity**: Every feature from `duplicates` is now in `art-dupl`
- **Enhanced capabilities**: `art-dupl` offers many additional features (config files, sorting, filtering, etc.)
- **Better performance**: Improved LineIndex implementation with binary search
- **More mature**: `art-dupl` has professional CLI (cobra/fang), auto-completion, and comprehensive testing
- **Single tooling**: Consolidates duplicate detection into one maintained project

### What's in art-dupl?

✅ **All duplicates features**:
- Token-sequence detection
- Configurable threshold
- Multiple output formats (JSON, HTML, Text, Plumbing)
- Line number tracking
- Scoring system (tokens × instances)
- File exclusion patterns
- Fast AST-based scanning

✅ **Plus enhancements**:
- **Simple JSON format** (`--simple-json`) - Exact format match with duplicates
- **Enhanced JSON format** (`--json`) - Rich metadata and statistics
- **Multiple sorting options** - Sort by size, occurrence, hash
- **Config file support** - Team consistency with JSON/YAML config
- **Auto-generated code filtering** - Ignore sqlc, templ, etc.
- **Professional CLI** - Auto-completion, version info, man pages
- **Performance profiling** - Identify bottlenecks
- **Multiple detection methods** - Suffix tree and hash-based

### Quick Migration Guide

#### Step 1: Install art-dupl

```bash
# Install from source
git clone https://github.com/LarsArtmann/art-dupl.git
cd art-dupl
make build

# Or via Go
go install github.com/LarsArtmann/art-dupl@latest
```

#### Step 2: Replace Commands

```bash
# Old (duplicates):
duplicates -threshold 20 -json report.json -html report.html

# New (art-dupl):
art-dupl --simple-json > report.json
art-dupl --html > report.html

# Or generate all at once:
art-dupl --all --output-dir ./reports --threshold 20
```

#### Step 3: Flag Mapping

| duplicates Flag | art-dupl Flag | Notes |
|----------------|----------------|--------|
| `-threshold N` | `--threshold N` or `-t N` | Same functionality |
| `-json report.json` | `--simple-json > report.json` | Use redirection or `--output-dir` |
| `-html report.html` | `--html > report.html` | Use redirection or `--output-dir` |
| `-text report.txt` | default (no flag) | Text is default, use `>` to redirect |
| `-plumbing` | `--plumbing` | Same functionality |
| `-v` | `--verbose` or `-v` | Same functionality |
| `-exclude "pattern"` | `--exclude-pattern "pattern"` | More powerful patterns |

#### Step 4: Update CI/CD Scripts

Replace these patterns in your pipelines:

```yaml
# Old:
- run: duplicates -threshold 30 -json report.json

# New:
- run: art-dupl --simple-json --threshold 30 > report.json
```

### Documentation

- **art-dupl**: [https://github.com/LarsArtmann/art-dupl](https://github.com/LarsArtmann/art-dupl)
- **Migration Guide**: See [art-dupl/MIGRATION_QUICK_START.md](https://github.com/LarsArtmann/art-dupl/blob/main/MIGRATION_QUICK_START.md)
- **Technical Report**: See [art-dupl/MIGRATION_REPORT_duplicates.md](https://github.com/LarsArtmann/art-dupl/blob/main/MIGRATION_REPORT_duplicates.md)

---

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

> **IMPORTANT**: This project is archived. Please install [art-dupl](https://github.com/LarsArtmann/art-dupl) instead.

### Install art-dupl (Recommended)

```bash
# Install from source
git clone https://github.com/LarsArtmann/art-dupl.git
cd art-dupl
make build

# Or via Go
go install github.com/LarsArtmann/art-dupl@latest
```

### Install duplicates (Archived - Not Recommended)

For legacy reasons only. All features now in art-dupl.

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

> **IMPORTANT**: This repository is archived. Development continues in [art-dupl](https://github.com/LarsArtmann/art-dupl).

For legacy maintenance of this archived project:

```bash
# Run tests
go test ./...

# Build
go build ./cmd/duplicates

# Lint
golangci-lint run
```

**For new features and bug fixes**, please contribute to [art-dupl](https://github.com/LarsArtmann/art-dupl).

## Roadmap

> **Status**: Development has moved to [art-dupl](https://github.com/LarsArtmann/art-dupl)

The following features are now available in art-dupl:

- [x] Fail threshold (exit code 1 if score exceeds limit) - Available in art-dupl
- [x] Config file support - Available in art-dupl (`--config dupl.json`)
- [x] Concurrent file scanning - Available in art-dupl
- [x] stdin input support for file lists - Available in art-dupl (`--files`)
- [x] Multiple sorting options - Available in art-dupl (`--sort`)
- [x] Filter auto-generated code - Available in art-dupl (`--filter-generated`)
- [x] Performance profiling - Available in art-dupl (`--profile`)
- [ ] Language-agnostic mode - Not yet planned

**Recommendation**: Use [art-dupl](https://github.com/LarsArtmann/art-dupl) for all future features.

## License

[Add your license here]

## Contributing

> **This repository is archived** - Please contribute to [art-dupl](https://github.com/LarsArtmann/art-dupl) instead.

This repository is maintained in read-only mode for historical purposes. All new features, bug fixes, and improvements should be made to [art-dupl](https://github.com/LarsArtmann/art-dupl), which contains all functionality from this project plus many enhancements.

---

**Archived**: January 14, 2026
**Successor**: [art-dupl](https://github.com/LarsArtmann/art-dupl)
**Reason**: Complete feature parity achieved

Built with [golangci/dupl](https://github.com/golangci/dupl) algorithm
