# test-litecov

Test repository for verifying litecov coverage reporting across multiple languages.

## Languages

- **Go**: Basic calculator with partial test coverage
- **Python**: Calculator module with partial test coverage
- **JavaScript**: Calculator with Jest and partial test coverage

## Purpose

This repo tests that litecov:
1. Correctly parses LCOV coverage files from different languages
2. Shows partial coverage (e.g., `Divide` function is uncovered in tests)
3. Shows files with NO coverage at all (e.g., `utils.*` files)
4. Generates proper GitHub annotations for uncovered lines
5. Posts accurate coverage comments on PRs
# Trigger rebuild
# Test uncovered lines
