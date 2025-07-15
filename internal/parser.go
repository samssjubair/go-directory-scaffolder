package internal

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// Template represents the structure defined in the YAML file
type Template struct {
	Name    string   `yaml:"name"`
	Folders []string `yaml:"folders"`
	Files   []string `yaml:"files"`
}

// LoadTemplate reads and parses a YAML template file
func LoadTemplate(path string) (*Template, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var tmpl Template
	if err := yaml.Unmarshal(data, &tmpl); err != nil {
		return nil, err
	}

	// Validate the template
	if err := ValidateTemplate(&tmpl); err != nil {
		return nil, fmt.Errorf("template validation failed: %w", err)
	}

	return &tmpl, nil
} 