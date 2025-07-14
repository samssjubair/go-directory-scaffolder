package internal

import (
	"io/ioutil"

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
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var tmpl Template
	if err := yaml.Unmarshal(data, &tmpl); err != nil {
		return nil, err
	}
	return &tmpl, nil
} 