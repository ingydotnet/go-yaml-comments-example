YAML Comments Example
=====================

This Go project demonstrates how to preserve comments when reading and writing
YAML files using the `go.yaml.in/yaml/v3` package.


## Overview

The program showcases `Node`, `NewDecoder`, and `NewEncoder` functions from the
`go.yaml.in/yaml/v3` package.

The program will process a sample YAML document that includes various types of
comments and output the result to stdout.


## Usage

You can use `make` to run the program or the tests:

```bash
make help
make run
make test
make test v=1
make shell
```

Note: The Makefile will automatically install the Go language locally within
this project, so it should be dependency free and just work.


## Example Input

```yaml
# This is a header comment
name: Alice  # This is an inline comment
age: 30

# This is a section comment
address:
  street: 123 Main St  # Street address
  city: Wonderland     # City name
```


## License

This project is open source and available under the MIT License.
