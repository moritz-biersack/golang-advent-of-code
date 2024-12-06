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

### Day 3

#### Regex

To use regex, use the `regexp` package.

```go
r := regexp.MustCompile(`(?s)(?:do\(\)|^)(.*?)(?:don't\(\)|$)`)
matches := r.FindAllStringSubmatch(input, -1)
```

The `(?s)` flag is used to make `.` match newlines.
`All` in `FindAllStringSubmatch` is used to find all matches.
`Submatch` is used to get the matched groups.

### Day 4

#### Slicing strings

If the string is containing ASCII only, it can be treated as a slice of bytes.

```go
reversed := make([]byte, len(subString))
copy(reversed, subString)
slices.Reverse(reversed)
```

In case it would contain unicode, one can convert to `[]rune` first.

```go
runes := []rune(someString)
```

#### Double-check off-by-one errors

Always make sure to double-check your lengths and limits.

#### Constant array

An array can not be defined as a constant.

## References

Project structure based on [alexchao26/advent-of-code-go](https://github.com/alexchao26/advent-of-code-go).
