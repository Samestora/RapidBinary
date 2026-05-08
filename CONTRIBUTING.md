# Contributing to RapidBinary

## Tools
- **Go 1.24.4** (Core logic)
- **Shared Package**: Branding and UI are managed in `internal/shared/`.
- **Architecture**: Keep binaries in `cmd/` and shared logic in `internal/`.

## Project Structure
- `cmd/rbserver`: The web-based file server.
- `cmd/rbhash`: Standalone integrity verification tool.
- `internal/shared`: UI, Headers, and common constants.

## Pull Request Process
1. Fork the repo and create your branch.
2. Ensure your code is formatted with `go fmt`.
3. Submit your PR with a brief description of the performance gain or feature added.
