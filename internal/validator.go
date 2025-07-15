package internal

import (
	"fmt"
	"strings"
)

// ValidationError represents a validation error
type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// ValidateTemplate validates a template structure
func ValidateTemplate(tmpl *Template) error {
	var errors []ValidationError

	// Validate name
	if tmpl.Name == "" {
		errors = append(errors, ValidationError{
			Field:   "name",
			Message: "project name cannot be empty",
		})
	}

	// Check for invalid characters in name
	if strings.ContainsAny(tmpl.Name, `<>:"/\|?*`) {
		errors = append(errors, ValidationError{
			Field:   "name",
			Message: "project name contains invalid characters",
		})
	}

	// Validate folders
	seenFolders := make(map[string]bool)
	for i, folder := range tmpl.Folders {
		if folder == "" {
			errors = append(errors, ValidationError{
				Field:   fmt.Sprintf("folders[%d]", i),
				Message: "folder name cannot be empty",
			})
			continue
		}

		// Check for invalid characters
		if strings.ContainsAny(folder, `<>:"|?*`) {
			errors = append(errors, ValidationError{
				Field:   fmt.Sprintf("folders[%d]", i),
				Message: "folder name contains invalid characters",
			})
		}

		// Check for duplicates
		if seenFolders[folder] {
			errors = append(errors, ValidationError{
				Field:   fmt.Sprintf("folders[%d]", i),
				Message: fmt.Sprintf("duplicate folder: %s", folder),
			})
		}
		seenFolders[folder] = true
	}

	// Validate files
	seenFiles := make(map[string]bool)
	for i, file := range tmpl.Files {
		if file == "" {
			errors = append(errors, ValidationError{
				Field:   fmt.Sprintf("files[%d]", i),
				Message: "file name cannot be empty",
			})
			continue
		}

		// Check for invalid characters
		if strings.ContainsAny(file, `<>:"|?*`) {
			errors = append(errors, ValidationError{
				Field:   fmt.Sprintf("files[%d]", i),
				Message: "file name contains invalid characters",
			})
		}

		// Check for duplicates
		if seenFiles[file] {
			errors = append(errors, ValidationError{
				Field:   fmt.Sprintf("files[%d]", i),
				Message: fmt.Sprintf("duplicate file: %s", file),
			})
		}
		seenFiles[file] = true

		// Check for conflicts between files and folders
		if seenFolders[file] {
			errors = append(errors, ValidationError{
				Field:   fmt.Sprintf("files[%d]", i),
				Message: fmt.Sprintf("file conflicts with folder: %s", file),
			})
		}
	}

	// Check for conflicts between folders and files
	for folder := range seenFolders {
		if seenFiles[folder] {
			errors = append(errors, ValidationError{
				Field:   "folders",
				Message: fmt.Sprintf("folder conflicts with file: %s", folder),
			})
		}
	}

	if len(errors) > 0 {
		return ValidationErrors(errors)
	}

	return nil
}

// ValidationErrors represents multiple validation errors
type ValidationErrors []ValidationError

func (e ValidationErrors) Error() string {
	if len(e) == 0 {
		return ""
	}

	var messages []string
	for _, err := range e {
		messages = append(messages, err.Error())
	}
	return strings.Join(messages, "; ")
} 