Porter Stemmer for Go
=====================

[![CI](https://github.com/a2800276/porter/actions/workflows/ci.yml/badge.svg)](https://github.com/a2800276/porter/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/a2800276/porter.svg)](https://pkg.go.dev/github.com/a2800276/porter)

This is a straightforward port of Martin Porter's C implementation of
the Porter stemming algorithm. The C version this port is based on is
available for download here:
[http://tartarus.org/~martin/PorterStemmer/c_thread_safe.txt](http://tartarus.org/~martin/PorterStemmer/c_thread_safe.txt)

The original algorithm is described in the paper:

    M.F. Porter, 1980, An algorithm for suffix stripping, Program, 14(3) pp
    130-137.

## Features

- Thread-safe implementation
- Multiple APIs: simple string API and zero-allocation byte-slice API
- Command-line tool for batch processing
- Comprehensive test suite
- Benchmarked and optimized
- No external dependencies

## Installation

### Library

```bash
go get github.com/a2800276/porter
```

### CLI Tool

```bash
go install github.com/a2800276/porter/cmd/porter@latest
```

## Usage

### As a Library

```go
package main

import (
    "fmt"
    "log"
    "github.com/a2800276/porter"
)

func main() {
    // Simple string API (with allocations)
    stemmed, err := porter.Stem("running")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(stemmed) // Output: run

    // Efficient byte-slice API (zero allocations)
    word := []byte("running")
    stemmed_bytes, err := porter.StemBytes(word)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(stemmed_bytes)) // Output: run
}
```

### As a CLI Tool

Install the command-line tool:

```bash
go install github.com/a2800276/porter/cmd/porter@latest
```

Use it to stem words:

```bash
# Stem words from arguments
$ porter running jumped easily
run
jump
easili

# Stem words from stdin
$ echo -e "running\njumped\neasily" | porter
run
jump
easili

# Process a file
$ cat words.txt | porter > stemmed.txt

# Count unique stems
$ cat corpus.txt | porter | sort | uniq -c | sort -rn
```

## API

The package provides two functions for different use cases:

### `Stem(word string) (string, error)`

The simplest API that takes a string and returns a stemmed string. Handles case conversion automatically.
Returns an error if stemming fails (though this is rare in normal use).

### `StemBytes(b []byte) ([]byte, error)`

Zero-allocation API that stems the byte slice in-place and returns the stemmed portion as a slice.
The input is converted to lowercase. Best for high-performance scenarios.
Returns an error if stemming fails.

## Performance

The implementation is highly optimized:

### String API (convenient, with allocations)
```
BenchmarkStem-24          14064384    77.29 ns/op    16 B/op    2 allocs/op
```

### Byte-Slice API (fastest, zero allocations)
```
BenchmarkStemBytes-24     23443530    51.85 ns/op     0 B/op    0 allocs/op
```

The byte-slice API (`StemBytes`) is ~35% faster and performs zero allocations,
making it ideal for high-performance applications.

**Note:** Error handling adds minimal overhead (~2ns) but provides explicit feedback on failures.



## Limitations

- The algorithm operates on English words only. Input is automatically converted to lowercase.
- For the `Stem()` function, strings are converted to byte slices internally.
  For zero-copy operation, use `StemBytes()`.
- Unicode handling: The algorithm is designed for ASCII English text. Non-ASCII characters should be handled by the caller before stemming.

## Development

### Building

```bash
make build       # Build the CLI tool
make install     # Install CLI to $GOPATH/bin
```

### Running Tests

```bash
make test        # Run tests
make coverage    # Generate coverage report
make bench       # Run benchmarks
```

### Linting and Formatting

```bash
make fmt         # Format code
make vet         # Run go vet
make lint        # Run golangci-lint (requires installation)
```

## Contributing

Contributions are welcome! Please ensure:

- Tests pass: `make test`
- Code is formatted: `make fmt`
- No linting errors: `make lint`

## License

MIT licensed. See LICENSE file for details. 
