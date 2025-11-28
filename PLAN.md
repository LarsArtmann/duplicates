# Duplicate Finder Enhancement Plan

## 1. Foundation & Cleanup (High Priority)
- [x] Add .gitignore
- [ ] Create `internal/duplicates` package structure.
- [ ] Move `dupl` dependency logic into `internal/duplicates/scanner.go`.
- [ ] Define Domain Models (`CloneGroup`, `Clone`) in `internal/duplicates/models.go` to decouple from `dupl` internals.

## 2. Line Number Resolution (Critical)
- [ ] Implement `FileReader` interface to facilitate testing.
- [ ] Create `internal/duplicates/lines.go` with `func OffsetToLine(content []byte, offset int) int`.
- [ ] Integrate Line Resolution into the Scanner to populate `StartLine` and `EndLine` in models.

## 3. Reporting Engine
- [ ] Create `internal/report` package.
- [ ] Implement `JSON` reporter with full line info.
- [ ] Implement `HTML` reporter (porting existing logic but using new models).
- [ ] Implement `Text` and `Plumbing` reporters.

## 4. CLI & Configuration
- [ ] Refactor `cmd/duplicates/main.go` to use `flag` or `cobra`.
- [ ] Add flags:
    - `-threshold` (default 15)
    - `-exclude` (glob patterns)
    - `-json`, `-html`, `-text` (output paths)
    - `-fail-threshold` (exit code 1 if score > X)

## 5. Testing & Quality
- [ ] Add unit tests for `OffsetToLine`.
- [ ] Add integration test with a sample Go file.
- [ ] Add `Makefile` or `Justfile` for easy running.
- [ ] Add CI workflow (GitHub Actions).

## 6. Performance (Future)
- [ ] Concurrent file walking.
- [ ] Stream processing for large codebases.
