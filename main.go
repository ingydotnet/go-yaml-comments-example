package main

import (
	"fmt"
	"os"
	"strings"

	"go.yaml.in/yaml/v3"
)

func main() {
	// Sample YAML with comments
	yamlInput := `# This is a header comment
name: Alice  # This is an inline comment
age: 30

# This is a section comment
address:
  street: 123 Main St  # Street address
  city: Wonderland     # City name
`

	// Create a decoder from the YAML string
	decoder := yaml.NewDecoder(strings.NewReader(yamlInput))

	// Decode into a Node
	var node yaml.Node
	err := decoder.Decode(&node)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error decoding YAML: %v\n", err)
		os.Exit(1)
	}

	// Create an encoder to write to standard output
	encoder := yaml.NewEncoder(os.Stdout)

	// Encode the Node back to YAML
	err = encoder.Encode(&node)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error encoding YAML: %v\n", err)
		os.Exit(1)
	}
} 
