# Contributing Guide

Thank you for considering contributing to our project! Here are some guidelines about how you can help:

## How to Contribute

1. Report issues or work on existing ones
2. Add new features or enhance existing ones
3. Fix bugs
4. Improve documentation

## Development Setup

1. Fork the repository
2. Clone your fork: `git clone https://github.com/yourusername/gitctx.git`
3. Install dependencies: `go mod tidy`

## Coding Standards

- Format your code using `gofmt`
- Follow Go best practices and idioms
- Write meaningful commit messages
- Add comments for complex logic
- Maintain test coverage

## Pull Request Process

1. Create a new branch for your feature: `git checkout -b feature-name`
2. Make your changes and commit them: `git commit -m 'Description of changes'`
3. Push to your fork: `git push origin feature-name`
4. Open a Pull Request with a clear title and description
5. Wait for review and address any feedback

## Testing

Please ensure all tests pass before submitting your contribution. If you're adding new features, include appropriate tests:

```bash
go test ./...
```

## Code Review

- All submissions require review
- We use GitHub pull requests for this purpose
- Follow up on review comments and make requested changes
- Be patient and respectful during the review process

## License

By contributing, you agree that your contributions will be licensed under the MIT License.
