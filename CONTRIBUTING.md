# Contributing to Neko 🐾

First off, thank you for considering contributing to Neko! It's people like you that make open source such a great community.

## 🐛 Bug Reports & Feature Requests
If you find a bug or have a suggestion for a new feature (like a new cat breed or toy!), please open an issue in the GitHub repository. Provide as much detail as possible.

## 🛠️ Local Development

1. **Fork and Clone**: Fork the repository and clone it to your local machine.
2. **Setup Environment**: Ensure you have Go 1.21+ installed.
3. **Run tests**: Make sure all tests pass before making your changes by running `make test` or `go test ./...`.
4. **Create a branch**: Make your changes in a new git branch:
   ```bash
   git checkout -b my-feature-branch
   ```
5. **Commit your changes**: Write clear and descriptive commit messages.
6. **Push to your fork**:
   ```bash
   git push origin my-feature-branch
   ```
7. **Submit a Pull Request**: Go to the original repository and open a Pull Request. Provide a clear description of the problem your PR solves or the feature it adds.

## 🎨 Adding New ASCII Art
If you are adding new Breeds, Toys, or Environments, please ensure they fit within the existing 26-character width rendering boundaries in `cmd/start.go`.

Happy coding, and stay cozy!
