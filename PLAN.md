# Duplicate Finder Enhancement Plan

## 1. Foundation & Cleanup (High Priority)
- [x] Add .gitignore
- [x] Create `internal/duplicates` package structure.
- [x] Move `dupl` dependency logic into `internal/duplicates/scanner.go`.
- [x] Define Domain Models (`CloneGroup`, `Clone`) in `internal/duplicates/models.go` to decouple from `dupl` internals.

## 2. Line Number Resolution (Critical)
- [x] Implement `FileReader` interface to facilitate testing (Implicitly via `os.ReadFile` usage, could be strictly interfaced).
- [x] Create `internal/duplicates/lines.go` with `func OffsetToLine(content []byte, offset int) int`.
- [x] Integrate Line Resolution into the Scanner to populate `StartLine` and `EndLine` in models.

## 3. Reporting Engine
- [x] Create `internal/report` package.
- [x] Implement `JSON` reporter with full line info.
- [x] Implement `HTML` reporter (porting existing logic but using new models).
- [x] Implement `Text` and `Plumbing` reporters.

## 4. CLI & Configuration
- [x] Refactor `cmd/duplicates/main.go` to use `flag`.
- [x] Add flags:
    - [x] `-threshold` (default 15)
    - [x] `-exclude` (glob patterns)
    - [x] `-json`, `-html`, `-text` (output paths)
    - [ ] `-fail-threshold` (exit code 1 if score > X)

## 5. Testing & Quality
- [x] Add unit tests for `OffsetToLine` (`lines_test.go`).
- [ ] Add integration test with a sample Go file.
- [ ] Add `Makefile` or `Justfile` for easy running.
- [ ] Add CI workflow (GitHub Actions).

## 6. Performance (Future)
- [ ] Concurrent file walking.
- [ ] Stream processing for large codebases.
