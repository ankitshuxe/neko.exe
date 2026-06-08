# [neko.exe] // CONTRIBUTION PROTOCOLS

Access to the `[neko.exe]` source code implies adherence to strict engineering standards. We welcome patches, tactical upgrades, and bug resolutions that align with our terminal-native, zero-bloat philosophy.

---

## // ANOMALY REPORTS & FEATURE PROPOSALS

If you identify a bug or wish to propose a new ASCII integration (such as a new feline archetype or tactical environment), please open an issue in the repository. Provide comprehensive reproduction steps and execution logs. Ambiguity is not tolerated.

---

## // LOCAL DEVELOPMENT SEQUENCE

1. **Fork and Clone**: Secure your own branch of the repository locally.
2. **Setup Environment**: Ensure Go 1.21+ is active in your PATH.
3. **Run Diagnostics**: Before pushing changes, verify integrity by executing `make test` or `go test ./...`. All tests must pass.
4. **Create a Vector**: Branch off main:
   ```bash
   git checkout -b feature/your-tactical-upgrade
   ```
5. **Commit Modifications**: Use strict, imperative commit messages (e.g., "Add syntax highlighting to diary command").
6. **Push Vector**:
   ```bash
   git push origin feature/your-tactical-upgrade
   ```
7. **Initialize Pull Request**: Submit your PR with a thorough explanation of the payload.

---

## // ASCII ART ASSET INTEGRATION

If you are expanding the ASCII Marketplace (adding new Breeds, Toys, or Environments), you must strictly adhere to the layout geometry. All assets must fit within a fixed 26-character width boundary to maintain cross-terminal alignment in `cmd/start.go`.

**Execution is everything. Stay focused.**
