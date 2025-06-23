package main

import (
	"bytes"
	"strings"
	"testing"

	"go.yaml.in/yaml/v3"
)

func TestYAMLCommentPreservation(t *testing.T) {
	// Test input YAML with various types of comments
	input := `# This is a header comment
name: Alice  # This is an inline comment
age: 30

# This is a section comment
address:
  street: 123 Main St  # Street address
  city: Wonderland     # City name
`
	// Create a buffer to capture the output
	var buf bytes.Buffer

	// Decode the input YAML
	decoder := yaml.NewDecoder(strings.NewReader(input))
	var node yaml.Node
	err := decoder.Decode(&node)
	if err != nil {
		t.Fatalf("Failed to decode YAML: %v", err)
	}

	// Encode back to YAML
	encoder := yaml.NewEncoder(&buf)
	err = encoder.Encode(&node)
	if err != nil {
		t.Fatalf("Failed to encode YAML: %v", err)
	}

	// Get the output as string
	output := buf.String()

	// Verify all comments are preserved
	expectedComments := []string{
		"# This is a header comment",
		"# This is an inline comment",
		"# This is a section comment",
		"# Street address",
		"# City name",
	}

	for _, comment := range expectedComments {
		if !strings.Contains(output, comment) {
			t.Errorf("Expected comment not found in output: %s", comment)
		}
	}

	// Verify the data structure is preserved
	var originalData, outputData struct {
		Name    string `yaml:"name"`
		Age     int    `yaml:"age"`
		Address struct {
			Street string `yaml:"street"`
			City   string `yaml:"city"`
		} `yaml:"address"`
	}

	// Decode original input
	err = yaml.Unmarshal([]byte(input), &originalData)
	if err != nil {
		t.Fatalf("Failed to unmarshal original YAML: %v", err)
	}

	// Decode output
	err = yaml.Unmarshal([]byte(output), &outputData)
	if err != nil {
		t.Fatalf("Failed to unmarshal output YAML: %v", err)
	}

	// Compare the data structures
	if originalData.Name != outputData.Name {
		t.Errorf("Name mismatch: got %s, want %s",
			outputData.Name, originalData.Name)
	}
	if originalData.Age != outputData.Age {
		t.Errorf("Age mismatch: got %d, want %d",
			outputData.Age, originalData.Age)
	}
	if originalData.Address.Street != outputData.Address.Street {
		t.Errorf("Street mismatch: got %s, want %s",
			outputData.Address.Street, originalData.Address.Street)
	}
	if originalData.Address.City != outputData.Address.City {
		t.Errorf("City mismatch: got %s, want %s",
			outputData.Address.City, originalData.Address.City)
	}
}
