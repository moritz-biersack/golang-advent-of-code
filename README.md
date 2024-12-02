# Golang Advent of Code 2024

Learning go by doing the Advent of Code 2024.

## Usage

```bash
go run day<day_number>/main.go --part <1 or 2> --input <input_file>
```

## Testing

```bash
go test ./day<day_number>  # consider the dot
```

## Learnings

### Day 1

#### Flags

The flag package can be used to parse command line arguments.

```go
var part int
var inputpath string
flag.IntVar(&part, "part", 1, "part 1 or 2")
flag.StringVar(&inputpath, "input", "example.txt", "file path of input")
flag.Parse()
```

#### Reading files

To read a whole file, use `os.ReadFile`.

```go
content, err := os.ReadFile(inputpath)
```

#### Splitting strings

To split a string, use `strings.Split`.

```go
lines := strings.Split(string(content), "\n")
```

To decompose a string with space separated values, use `fmt.Sscanf`.

```go
var a, b int
_, err := fmt.Sscanf(line, "%d %d", &a, &b)
```

### Day 2

#### Parsing strings to int

To parse a string to an int, use `strconv.Atoi`.

```go
strconv.Atoi("123")
```

#### Popping elements from a slice

To pop an element from a slice, use the following pattern.
Without copying the original slice, it the original slice is modified.

```go
c := make([]string, len(original))
copy(c, levels)
removedOne := append(c[:i], c[i+1:]...)
```

## References

Project structure based on [alexchao26/advent-of-code-go](https://github.com/alexchao26/advent-of-code-go).
