# Justfile for duplicates project
# https://github.com/casey/just

default: help

# Show help
help:
    @echo "duplicates - Code duplicate detection tool"
    @echo ""
    @echo "Available commands:"
    @echo "  just build          - Build the binary"
    @echo "  just test           - Run tests"
    @echo "  just lint           - Run linters"
    @echo "  just scan           - Run duplicate scan on current directory"
    @echo "  just clean          - Remove reports directory"
    @echo "  just run            - Build and scan"
    @echo ""

# Build the binary
build:
    @echo "Building duplicates..."
    go build -o duplicates ./cmd/duplicates

# Run tests
test:
    @echo "Running tests..."
    go test ./... -v

# Run linters
lint:
    @echo "Running linters..."
    @if command -v golangci-lint >/dev/null 2>&1; then \
        golangci-lint run; \
    else \
        echo "golangci-lint not installed. Install with: go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest"; \
    fi

# Run duplicate scan
scan:
    @echo "Running duplicate scan..."
    go run ./cmd/duplicates

# Run scan with custom threshold
scan-threshold threshold:
    @echo "Running duplicate scan with threshold {{threshold}}..."
    go run ./cmd/duplicates -threshold {{threshold}}

# Run scan with verbose output
scan-v:
    @echo "Running duplicate scan (verbose)..."
    go run ./cmd/duplicates -v

# Run scan excluding files
scan-exclude patterns:
    @echo "Running duplicate scan excluding: {{patterns}}..."
    go run ./cmd/duplicates -exclude {{patterns}}

# Remove reports
clean:
    @echo "Cleaning reports..."
    rm -rf reports/

# Build and run
run: build scan

# Format code
fmt:
    @echo "Formatting code..."
    gofumpt -w . || gofmt -w .
    goimports -w . || true

# Format and lint
check: fmt lint test
    @echo "All checks passed!"

# Install binary
install:
    @echo "Installing duplicates..."
    go install ./cmd/duplicates
