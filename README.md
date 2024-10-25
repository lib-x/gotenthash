# gotenthash

[![GoDoc](https://godoc.org/github.com/lib-x/gotenthash?status.svg)](https://godoc.org/github.com/lib-x/gotenthash)

gotenthash is a Go implementation of [TentHash](https://github.com/cessen/tenthash), a 160-bit non-cryptographic hash function.

## Features

- Fast and simple non-cryptographic hash function
- Support for incremental hashing and streaming input
- Compatible with the original C implementation of TentHash

## Installation

Using Go modules, you can install gotenthash with the following command:

```bash
go get github.com/lib-x/gotenthash
```

## Usage

### Basic Usage

```go
import "github.com/lib-x/gotenthash"

h := gotenthash.New()
h.Write([]byte("Hello, world!"))
fmt.Printf("%x\n", h.Sum(nil))
```
### Incremental Hashing

```go
h := gotenthash.New()
h.Write([]byte("Hello, "))
h.Write([]byte("world!"))
fmt.Printf("%x\n", h.Sum(nil))
```

### Streaming Hashing

```go
file, := os.Open("example.txt")
defer file.Close()
hash, err := gotenthash.HashReader(file)
if err != nil {
// Handle error
}
```



## API Documentation

- `func New() *TentHasher`: Create a new TentHasher instance
- `func (t *TentHasher) Write(data []byte) (int, error)`: Write data to the hasher
- `func (t *TentHasher) Sum(b []byte) []byte`: Compute the current hash value
- `func (t *TentHasher) Reset()`: Reset the hasher state
- `func (t *TentHasher) WriteReader(r io.Reader) (int64, error)`: Write data from an io.Reader
- `func (t *TentHasher) SumReader(r io.Reader) ([]byte, error)`: Compute hash value from an io.Reader
- `func Hash(data []byte) [DigestSize]byte`: Compute hash value of a byte slice
- `func HashReader(reader io.Reader) ([DigestSize]byte, error)`: Compute hash value from an io.Reader

## Performance

(You can add some performance test results here, such as comparisons with other hash functions)

## Contributing

Contributions are welcome! Please submit a pull request or create an issue to discuss the changes you want to make.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgements

This project is based on [TentHash](https://github.com/cessen/tenthash) by [Casey Rodarmor](https://github.com/cessen).
