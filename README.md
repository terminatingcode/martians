# Martian Robots [WIP]

Martian Robots is a simple, barebones [Go](https://go.dev/doc/tutorial/getting-started) app that reads an input file with instructions for robot movement commands and outputs the results.

The executable can be run by inputting the input file as a command line argument:

```
./bin/martians path/to/input.txt
```

### For Mac:

```
./bin/martians

```

### For Windows:

```
.\bin\martians.exe
```

## Build

To build the project from source code and for your specific OS architecture:

```
go build
```

## Test

## Unit Tests

```
go test ./...
```

To test the test fixture against the executable:

```
./bin/main test-input.txt
```
