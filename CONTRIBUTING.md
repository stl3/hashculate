# Contributing to Hashculate

Thank you for your interest in contributing to Hashculate! This document provides guidelines and information for contributors.

## Code of Conduct

By participating in this project, you agree to abide by our Code of Conduct. Please be respectful and constructive in all interactions.

## How to Contribute

### Reporting Bugs

1. Check if the bug has already been reported in the [Issues](https://github.com/stl3/hashculate/issues)
2. If not, create a new issue with:
   - Clear description of the problem
   - Steps to reproduce
   - Expected vs actual behavior
   - Your environment (OS, Go version)
   - Sample files or commands that trigger the issue

### Suggesting Features

1. Check existing [Issues](https://github.com/stl3/hashculate/issues) for similar requests
2. Create a new issue with:
   - Clear description of the feature
   - Use case and motivation
   - Proposed implementation (if you have ideas)

### Contributing Code

1. **Fork the repository**
2. **Create a feature branch**
   ```bash
   git checkout -b feature/your-feature-name
   ```
3. **Make your changes**
   - Follow the existing code style
   - Add tests for new functionality
   - Update documentation if needed
4. **Test your changes**
   ```bash
   go test -v ./...
   go test -race -v ./...
   ```
5. **Commit your changes**
   ```bash
   git commit -m "Add feature: your feature description"
   ```
6. **Push to your fork**
   ```bash
   git push origin feature/your-feature-name
   ```
7. **Create a Pull Request**

## Development Guidelines

### Code Style

- Follow standard Go conventions
- Use `gofmt` to format your code
- Use meaningful variable and function names
- Add comments for exported functions and complex logic
- Keep functions focused and reasonably sized

### Testing

- Write tests for new functionality
- Ensure all tests pass before submitting
- Include both positive and negative test cases
- Test edge cases and error conditions

### Documentation

- Update README.md if you add new features
- Add inline comments for complex code
- Update help text if you modify command-line options

## Project Structure

```
hashculate/
├── main.go           # Main application code
├── main_test.go      # Tests
├── README.md         # Project documentation
├── LICENSE           # License file
├── go.mod            # Go module file
└── .github/
    └── workflows/
        └── ci.yml    # GitHub Actions CI
```

## Building and Testing

### Prerequisites

- Go 1.18 or later
- Git

### Local Development

```bash
# Clone your fork
git clone https://github.com/stl3/hashculate.git
cd hashculate

# Build the project
go build -o hashculate .

# Run tests
go test -v ./...

# Test with race detector
go test -race -v ./...

# Run the application
./hashculate test_file.txt
```

### Cross-Platform Testing

Test on multiple platforms when possible:

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o hashculate-linux .

# Windows
GOOS=windows GOARCH=amd64 go build -o hashculate-windows.exe .

# macOS
GOOS=darwin GOARCH=amd64 go build -o hashculate-macos .
```

## Pull Request Guidelines

### Before Submitting

- [ ] Code follows project style guidelines
- [ ] All tests pass
- [ ] New functionality includes tests
- [ ] Documentation is updated
- [ ] Commit messages are clear and descriptive

### Pull Request Description

Include in your PR description:

- Summary of changes
- Motivation for the changes
- Any breaking changes
- Testing performed
- Screenshots (if applicable)

## Release Process

Releases are managed by maintainers and follow semantic versioning:

- **Major** (x.0.0): Breaking changes
- **Minor** (0.x.0): New features, backward compatible
- **Patch** (0.0.x): Bug fixes, backward compatible

## Getting Help

- Check the [README.md](README.md) for usage information
- Look through existing [Issues](https://github.com/stl3/hashculate/issues)
- Create a new issue for questions or problems

## Recognition

Contributors will be recognized in the project documentation and release notes.

Thank you for contributing to Hashculate!
