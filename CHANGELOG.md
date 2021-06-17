# Changelog

## [Unreleased]

## [1.0.0] - 2021-06-17
### Added
- initial support for script running
- support for tilde expansion in the `cd` command
- support for home dir compression to tilde for the shell prompt

### Changed
- `gohome` command renamed to `home` for simplicity sake
- Full refactor of the codebase flow and design based on my improved understanding of Go
- Shell now has a concept of per-session variables rather than using Environment for all variables

### Removed
- `redo` command removed and will be refactored for a future release.

### Fixed
- ctrl+c signal now correctly clears the line
- ctrl+d signal now correctly cleanly exits the shell
- dependency spiderweb has been completely removed from shell libraries

## [0.1.0] - 2020-08-19
### Added
- Initial release of Turtleshell

[Unreleased]: https://github.com/jgrancell/turtleshell/compare/v1.0.0...HEAD
[1.0.0]: https://github.com/jgrancell/turtleshell/compare/v0.1.0...v1.0.0
[0.1.0]: https://github.com/jgrancell/turtleshell/releases/tag/v0.1.0