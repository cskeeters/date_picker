# AGENTS.md - Development Guide for date-picker

This document provides guidelines and commands for developing the date-picker CLI tool, a terminal-based date picker built with Go and Bubble Tea.

## Build Commands

### Primary Build Commands
- `go build` - Compile the main binary
- `go build -v` - Compile with verbose output
- `go build -o date-picker-custom` - Compile with custom output name

### Usage Examples
- `./date-picker` - Interactive date picker starting with today's date
- `./date-picker 2024-01-15` - Interactive date picker starting with the specified date
- `./date-picker -f "%Y/%m/%d" 2024-01-15` - Interactive date picker with custom output format
- `./date-picker -f "%A, %B %d, %Y"` - Interactive date picker with full date output format

### Clean Builds
- `go clean` - Remove object files and cached files
- `go clean -cache` - Remove the entire go build cache
- `go clean -modcache` - Remove the entire module cache

### Cross-Platform Builds
- `GOOS=linux GOARCH=amd64 go build` - Build for Linux AMD64
- `GOOS=darwin GOARCH=arm64 go build` - Build for macOS ARM64
- `GOOS=windows GOARCH=amd64 go build` - Build for Windows AMD64

## Test Commands

### Basic Testing
- `go test` - Run tests in current directory
- `go test ./...` - Run tests recursively in all subdirectories
- `go test -v` - Run tests with verbose output
- `go test -short` - Run tests excluding those that take a long time

### Advanced Testing Options
- `go test -race` - Run tests with race detector enabled
- `go test -cover` - Run tests with coverage analysis
- `go test -coverprofile=coverage.out` - Generate coverage profile
- `go tool cover -html=coverage.out` - View coverage report in browser

### Running a Single Test
- `go test -run TestFunctionName` - Run specific test function
- `go test -run "Test.*Date.*"` - Run tests matching regex pattern
- `go test -run TestFunctionName -v` - Run specific test with verbose output

### Benchmarking
- `go test -bench=.` - Run all benchmarks
- `go test -bench=BenchmarkFunctionName` - Run specific benchmark
- `go test -bench=. -benchmem` - Include memory allocation statistics

## Lint and Format Commands

### Code Formatting
- `gofmt -d .` - Show diff of formatting changes needed
- `gofmt -w .` - Apply formatting changes in place

### Code Analysis
- `go vet` - Report suspicious constructs
- `go vet ./...` - Vet all packages recursively
- `golint ./...` - Run golint for style issues (if installed)

### Dependency Management
- `go mod tidy` - Clean up go.mod and go.sum
- `go mod download` - Download modules to local cache
- `go mod verify` - Verify dependencies have expected content
- `go list -m all` - List all dependencies

## Code Style Guidelines

### General Principles
- Follow the [Effective Go](https://golang.org/doc/effective_go.html) guidelines
- Write idiomatic Go code that feels natural to experienced Go developers
- Prefer clarity over cleverness

### Naming Conventions
- Use `MixedCaps` or `mixedCaps` for all names (no underscores)
- Exported identifiers start with capital letters
- Keep names short but descriptive
- Acronyms should have consistent casing: `someEksCluster`, not `someEKSCluster`

### Package Organization
- Keep packages small and focused on single responsibilities
- Use descriptive package names
- Avoid circular dependencies
- Group related functionality together

### Function Design
- Functions should be small and focused
- Use clear, descriptive names
- Return errors as the last return value
- Handle errors appropriately (don't ignore them)

### Error Handling
- Always check for errors
- Don't panic (except in truly exceptional circumstances)
- Return errors to caller rather than handling them internally unless appropriate
- Use custom error types when additional context is needed

### Comments
- Write comments for all exported functions, types, and constants
- Comments should explain *why* not just *what*
- Use complete sentences starting with the name being described
- Keep comments up to date with code changes

### Imports
- Group imports into three sections with blank lines:
  1. Standard library packages
  2. Third-party packages
  3. Local packages
- Remove unused imports
- Use meaningful import aliases when needed

### Control Structures
- Use `for` loops for iteration
- Use `if`/`else` for conditional logic
- Avoid deeply nested structures
- Use early returns to reduce nesting
- Use `switch` for multiple conditions on same value

### Pointers
- Prefer value types over pointers unless necessary
- Use pointers when:
  - Mutating a value passed to a function
  - Dealing with large structs to avoid copying
  - Need to distinguish between zero value and nil

### Constants and Variables
- Use constants for values that don't change
- Group related constants together
- Use meaningful variable names
- Declare variables at smallest scope possible

### Structs and Interfaces
- Define small, focused interfaces
- Use struct embedding for composition
- Keep structs simple and focused
- Use pointer receivers for methods that modify the receiver

### Testing
- Write table-driven tests when possible
- Use descriptive test names: `TestFunctionName_DescribesBehavior`
- Test both success and failure cases
- Use `t.Parallel()` for independent tests
- Include benchmarks for performance-critical code

### Type Assertions and Conversions
- Use type assertions sparingly
- Check for success when using type assertions
- Prefer type switches over repeated type assertions

### Goroutines and Concurrency
- Use channels for communication between goroutines
- Avoid global state
- Use sync.WaitGroup for coordinating goroutines
- Use context.Context for cancellation

### File Organization
- Keep main.go focused on application startup
- Separate business logic into packages
- Use internal/ directory for private packages
- Use pkg/ directory for public packages if needed

### Security Considerations
- Validate all inputs
- Use safe defaults
- Avoid exposing sensitive information in logs
- Keep dependencies updated

### Performance
- Avoid premature optimization
- Profile before optimizing
- Use efficient data structures
- Minimize allocations in hot paths

## Development Workflow

1. Make changes to code
2. Run `go vet` to check for issues
3. Run `go test` to verify functionality
4. Run `go build` to ensure compilation
5. Test manually if needed

## Tool Installation

### Recommended Tools
- `go install golang.org/x/tools/cmd/goimports@latest` - Enhanced import management
- `go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest` - Advanced linting
- `go install github.com/cosmtrek/air@latest` - Live reloading for development

### IDE Integration
- Use Go extension for VS Code
- Configure go vet and golint integration
</content>
<parameter name="filePath">/Users/chad/working/date-picker/AGENTS.md