package internal

import (
	"fmt"
)

// BuiltinTemplate represents a built-in template
type BuiltinTemplate struct {
	Name        string
	Description string
	Path        string
}

// BuiltinTemplates holds all available built-in templates
var BuiltinTemplates = map[string]BuiltinTemplate{
	"react-app": {
		Name:        "react-app",
		Description: "React application with components, pages, and utils",
		Path:        "templates/react-app.yaml",
	},
	"go-api": {
		Name:        "go-api",
		Description: "Go API project with handlers, models, and database",
		Path:        "templates/go-api.yaml",
	},
	"node-express": {
		Name:        "node-express",
		Description: "Node.js Express application with routes and middleware",
		Path:        "templates/node-express.yaml",
	},
}

// ListBuiltinTemplates returns a list of all available built-in templates
func ListBuiltinTemplates() []BuiltinTemplate {
	templates := make([]BuiltinTemplate, 0, len(BuiltinTemplates))
	for _, tmpl := range BuiltinTemplates {
		templates = append(templates, tmpl)
	}
	return templates
}

// GetBuiltinTemplate returns a built-in template by name
func GetBuiltinTemplate(name string) (*BuiltinTemplate, error) {
	tmpl, exists := BuiltinTemplates[name]
	if !exists {
		return nil, fmt.Errorf("template '%s' not found", name)
	}
	return &tmpl, nil
}

// LoadBuiltinTemplate loads a built-in template by name
func LoadBuiltinTemplate(name string) (*Template, error) {
	tmpl, err := GetBuiltinTemplate(name)
	if err != nil {
		return nil, err
	}
	return LoadTemplate(tmpl.Path)
} 