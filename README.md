# linew

A simple and efficient line wrapping tool for text files written in Go.

## Features

- Wrap long lines at a specified width (default: 80 characters)
- Preserve indentation when wrapping
- Process files or standard input
- Output to files or standard output
- Handle Unicode text properly

## Installation

```bash
go install github.com/rRateLimit/linew/cmd/linew@latest
```

Or build from source:

```bash
git clone https://github.com/rRateLimit/linew.git
cd linew
go build -o linew cmd/linew/main.go
```

## Usage

```bash
linew [options] [file]
```

### Options

- `-w, --width`: Maximum width for line wrapping (default: 80)
- `-i, --indent`: Preserve indentation (default: true)
- `--no-indent`: Do not preserve indentation
- `-o, --output`: Output file (default: stdout)
- `-h, --help`: Show help message

### Examples

Wrap lines at 80 characters (default):
```bash
linew input.txt
```

Wrap lines at 100 characters:
```bash
linew -w 100 input.txt
```

Process from stdin and output to file:
```bash
cat long_text.txt | linew -o wrapped.txt
```

Wrap without preserving indentation:
```bash
linew --no-indent input.txt
```

## How it Works

The tool reads text line by line and:
1. Splits lines that exceed the specified width
2. Preserves word boundaries when possible
3. Maintains indentation if enabled
4. Breaks very long words that exceed the width limit

## Development

Run tests:
```bash
go test ./...
```

Build:
```bash
go build -o linew cmd/linew/main.go
```