package internal

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// GitIntegration handles Git repository initialization
type GitIntegration struct {
	ProjectPath string
}

// NewGitIntegration creates a new Git integration instance
func NewGitIntegration(projectPath string) *GitIntegration {
	return &GitIntegration{
		ProjectPath: projectPath,
	}
}

// InitGit initializes a Git repository in the project directory
func (g *GitIntegration) InitGit() error {
	// Check if git is available
	if !g.isGitAvailable() {
		return fmt.Errorf("git is not available on the system")
	}

	// Change to project directory
	if err := os.Chdir(g.ProjectPath); err != nil {
		return fmt.Errorf("failed to change to project directory: %w", err)
	}

	// Initialize git repository
	cmd := exec.Command("git", "init")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to initialize git repository: %w", err)
	}

	return nil
}

// CreateGitignore creates a .gitignore file with common patterns
func (g *GitIntegration) CreateGitignore() error {
	gitignorePath := filepath.Join(g.ProjectPath, ".gitignore")
	
	// Check if .gitignore already exists
	if _, err := os.Stat(gitignorePath); err == nil {
		return nil // .gitignore already exists
	}

	content := `# Dependencies
node_modules/
vendor/

# Build outputs
dist/
build/
*.exe
*.dll
*.so
*.dylib

# Environment variables
.env
.env.local
.env.*.local

# IDE files
.vscode/
.idea/
*.swp
*.swo
*~

# OS files
.DS_Store
Thumbs.db

# Logs
*.log
npm-debug.log*
yarn-debug.log*
yarn-error.log*

# Runtime data
pids
*.pid
*.seed
*.pid.lock

# Coverage directory used by tools like istanbul
coverage/

# Temporary folders
tmp/
temp/
`

	return os.WriteFile(gitignorePath, []byte(content), 0644)
}

// MakeInitialCommit creates an initial commit
func (g *GitIntegration) MakeInitialCommit() error {
	// Change to project directory
	if err := os.Chdir(g.ProjectPath); err != nil {
		return fmt.Errorf("failed to change to project directory: %w", err)
	}

	// Add all files
	cmd := exec.Command("git", "add", ".")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to add files to git: %w", err)
	}

	// Make initial commit
	cmd = exec.Command("git", "commit", "-m", "Initial commit")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to make initial commit: %w", err)
	}

	return nil
}

// SetupGit performs complete Git setup
func (g *GitIntegration) SetupGit() error {
	// Initialize git repository
	if err := g.InitGit(); err != nil {
		return err
	}

	// Create .gitignore file
	if err := g.CreateGitignore(); err != nil {
		return fmt.Errorf("failed to create .gitignore: %w", err)
	}

	// Make initial commit
	if err := g.MakeInitialCommit(); err != nil {
		return fmt.Errorf("failed to make initial commit: %w", err)
	}

	return nil
}

// isGitAvailable checks if git is available on the system
func (g *GitIntegration) isGitAvailable() bool {
	cmd := exec.Command("git", "--version")
	return cmd.Run() == nil
} 