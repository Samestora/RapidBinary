# Changelog
All notable changes to the **RapidBinary** toolset will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]
- *No changes yet.*

## [1.1.0] - 2026-05-08
### Added
- **Toolsuite Model**: Transitioned from a single binary to a multi-tool architecture.
- **rbhash**: New standalone utility for forensic SHA256 integrity verification.
- **Shared UI**: Centralized branding and headers in `internal/shared/header.go` for consistency.
- **CONTRIBUTING.md**: Guidelines for high-speed, low-memory forensic development.
- **Automated CI/CD**: Updated GitHub Actions matrix to support multi-binary releases across Windows, Linux, and macOS.

### Changed
- Refactored project structure to use `cmd/rbserver` and `cmd/rbhash` subdirectories.
- Optimized buffer management to stabilize transfers over 5GHz hotspots.

### Fixed
- Corrected routing issues in bridged VM environments by prioritizing valid gateways

## [1.0.0] - 2026-05-08
### Added
- **Core Server**: High-performance, concurrent file serving built with Go 1.24.4.
- **Upload UI**: Modern web interface for browser-based file uploads (activated via `--upload`).
- **Embedded Assets**: UI templates and CSS embedded directly into the binary using `//go:embed`.
- **Multi-Tool Architecture**: Project structured to support future RapidBinary tools.
- **Automated CI/CD**: GitHub Actions matrix workflow for Windows (zip), Linux (tar.gz), and macOS (tar.gz) builds.
- **CLI Flags**: Support for `--port`, `--dir`, and `--upload` configuration.
