package internal

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

// CreateStructure creates the folder structure and files based on the template
func CreateStructure(tmpl *Template) error {
	// Create the main project directory
	if err := os.Mkdir(tmpl.Name, 0755); err != nil {
		return fmt.Errorf("failed to create project directory: %w", err)
	}

	// Create folders
	for _, folder := range tmpl.Folders {
		path := filepath.Join(tmpl.Name, folder)
		if err := os.MkdirAll(path, 0755); err != nil {
			return fmt.Errorf("failed to create folder %s: %w", folder, err)
		}
	}

	// Create files
	for _, file := range tmpl.Files {
		path := filepath.Join(tmpl.Name, file)
		
		// Ensure parent directories exist for nested files
		dir := filepath.Dir(path)
		if dir != "." && dir != tmpl.Name {
			if err := os.MkdirAll(dir, 0755); err != nil {
				return fmt.Errorf("failed to create parent directory for %s: %w", file, err)
			}
		}

		// Create empty file
		f, err := os.Create(path)
		if err != nil {
			return fmt.Errorf("failed to create file %s: %w", file, err)
		}
		f.Close()
	}

	return nil
}

// PrintStructure prints what would be created in a tree-like format
func PrintStructure(tmpl *Template) {
	color.Cyan("📁 %s/", tmpl.Name)
	
	// Print folders
	for _, folder := range tmpl.Folders {
		color.Blue("   📁 %s/", folder)
	}
	
	// Print files
	for _, file := range tmpl.Files {
		color.Green("   📄 %s", file)
	}
} 