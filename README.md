# Git Commit Mocker

A Go command-line tool that generates mock Git commits between specified dates with configurable frequency and commit messages.

## Features

- Creates mock Git commits with random timestamps between specified dates
- Configurable commit frequency using "temperature" (0-10) for weekdays and weekends
- Customizable minimum and maximum commits per day
- Optional custom commit messages from a file
- Option to repeat commit messages or generate random ones
- Automatic Git repository initialization if none exists
- Validates Git installation as a prerequisite

## Installation

### Prerequisites

- Go 1.21 or higher (for building from source)
- Git installed on your system

### From Source

1. Clone the repository:

```bash
git clone https://github.com/johnmerga/git-commit-mocker.git
cd git-commit-mocker
```

2. Build the binary:

```bash
go build -o git-commit-mocker ./cmd/git-commit-mocker
```

3. (Optional) Move to your PATH:

```bash
mv git-commit-mocker /usr/local/bin/
```

### Binary Download

Download the pre-built binary for your platform from the Releases page.

## Usage

Basic syntax:

```
git-commit-mocker [flags]
```

### Required Flags

- `-start string`: Start date in YYYY-MM-DD format
- `-end string`: End date in YYYY-MM-DD format

### Optional Flags

- `-weekday-temp float`: Weekday commit frequency (0-10, default 5.0)
- `-weekend-temp float`: Weekend commit frequency (0-10, default 5.0)
- `-max-commits int`: Maximum commits per day (default 5)
- `-min-commits int`: Minimum commits per day (default 1)
- `-msg-file string`: Path to .txt file with commit messages (one per line)
- `-repeat-msgs`: Repeat commit messages if exhausted (required if msg-file is provided)

## Examples

### Basic usage with default settings:

```bash
git-commit-mocker -start 2025-01-01 -end 2025-01-31
```

- Commits ~50% of weekdays and weekends
- 1-5 commits per day
- Random default messages

### High frequency weekdays, low weekends:

```bash
git-commit-mocker -start 2025-01-01 -end 2025-01-31 -weekday-temp 8.0 -weekend-temp 2.0 -max-commits 3 -min-commits 1
```

- Commits ~80% of weekdays, ~20% of weekends
- 1-3 commits per day

### With custom commit messages:

```bash
git-commit-mocker -start 2025-01-01 -end 2025-01-31 -msg-file messages.txt -repeat-msgs
```

Example `messages.txt`:

```
feat: add new functionality
fix: resolve critical bug
docs: update readme
```

## How It Works

- Validates Git installation and initializes a repository if needed
- Calculates commit probability based on temperature (0-10 scale):
  - 10 = commit every day
  - 5 = commit ~50% of days
  - 0 = no commits
- For each day between start and end dates:
  - Determines if commits should occur based on temperature
  - Generates random number of commits between min and max
  - Creates commits with random timestamps within the day
  - Uses custom messages if provided, or random defaults

## Directory Structure

```
git-commit-mocker/
├── cmd/                # Main application entry point
│   └── git-commit-mocker/
│       └── main.go
├── internal/          # Private application code
│   ├── config/       # Configuration handling
│   ├── git/         # Git operations
│   ├── messages/   # Commit message management
│   └── scheduler/  # Commit scheduling logic
├── go.mod            # Go module definition
└── README.md        # This file
```

## Building for Distribution

To create binaries for different platforms:

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o git-commit-mocker-linux ./cmd/git-commit-mocker

# Windows
GOOS=windows GOARCH=amd64 go build -o git-commit-mocker-windows.exe ./cmd/git-commit-mocker

# macOS
GOOS=darwin GOARCH=amd64 go build -o git-commit-mocker-darwin ./cmd/git-commit-mocker
```

## Contributing

- Fork the repository
- Create a feature branch
- Submit a pull request with your changes

## License

MIT License

## Acknowledgments

Built with Go and love ❤️
