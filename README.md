# Directory Scaffolder

A CLI tool written in Go that creates folder structures and files based on YAML templates.

![Typ0 Demo](https://img.shields.io/badge/Go-1.21+-blue.svg)
![License](https://img.shields.io/badge/License-MIT-green.svg)

## Features

- ✅ Create folder structures from YAML templates
- ✅ Support for nested folders
- ✅ Dry-run mode to preview changes
- ✅ Colorful output with emojis
- ✅ Built-in templates included
- ✅ Built-in template selection
- ✅ Template validation
- ✅ Git integration

## Installation

### Homebrew (macOS/Linux)

```bash
brew install samssjubair/scaffold/scaffold
```

### Manual Installation

Download the latest binary from the [releases page](https://github.com/samssjubair/go-directory-scaffolder/releases).

### Build from Source

```bash
git clone https://github.com/samssjubair/go-directory-scaffolder.git
cd go-directory-scaffolder
go build -o scaffold .
```

<img width="1680" height="671" alt="screenshot" src="https://github.com/user-attachments/assets/597aed5a-fc13-4316-9071-f17176ae8afc" />

## Usage

### Basic Usage

```bash
# Create structure from a YAML template
scaffold project.yaml

# Preview what would be created (dry-run)
scaffold --dry-run project.yaml
```

### Built-in Templates

```bash
# List all available built-in templates
scaffold --list-templates

# Use a built-in template
scaffold --template react-app
scaffold --template go-api
scaffold --template node-express

# Use built-in template with custom project name
scaffold --template react-app my-custom-app

# Preview built-in template
scaffold --template go-api --dry-run
```

### Git Integration

```bash
# Create project and initialize Git repository
scaffold --init-git project.yaml
scaffold --template react-app --init-git

# This will:
# - Create the project structure
# - Initialize a Git repository
# - Create a comprehensive .gitignore file
# - Make an initial commit
```

### Template Validation

```bash
# Skip template validation (not recommended)
scaffold --skip-validation project.yaml

# Templates are automatically validated for:
# - Empty names
# - Invalid characters
# - Duplicate folders/files
# - Conflicts between folders and files
```

### Available Built-in Templates

- **react-app** - React application with components, pages, and utils
- **go-api** - Go API project with handlers, models, and database  
- **node-express** - Node.js Express application with routes and middleware

### Example YAML Template

```yaml
name: my-project
folders:
  - src
  - public
  - src/components
files:
  - README.md
  - package.json
  - .gitignore
  - src/index.js
```

This will create:

```
my-project/
├── src/
│   └── components/
├── public/
├── README.md
├── package.json
├── .gitignore
└── src/index.js
```

## Project Structure

```
scaffold/
├── cmd/                # cobra CLI commands
│   ├── root.go
│   └── version.go
├── internal/
│   ├── parser.go       # load & parse YAML
│   ├── creator.go      # create dirs & files
│   ├── templates.go    # built-in templates
│   ├── validator.go    # template validation
│   └── git.go         # Git integration
├── templates/          # sample templates
├── go.mod
├── main.go
└── README.md
```

## Dependencies

- **cobra** - CLI framework
- **gopkg.in/yaml.v2** - YAML parsing
- **fatih/color** - Colored output

## Development

```bash
# Install dependencies
go mod tidy

# Build
go build -o scaffold .

# Run
./scaffold templates/react-app.yaml
./scaffold --template react-app
```

## Template Validation

The scaffolder automatically validates templates for:

- **Empty names** - Project name cannot be empty
- **Invalid characters** - Names cannot contain `<>:"/\|?*`
- **Duplicates** - No duplicate folders or files
- **Conflicts** - Files and folders cannot have the same name

## Git Integration

When using `--init-git`, the scaffolder will:

1. **Initialize Git repository** in the project directory
2. **Create .gitignore** with common patterns for various languages
3. **Make initial commit** with all created files

The generated `.gitignore` includes patterns for:
- Dependencies (node_modules, vendor)
- Build outputs (dist, build)
- Environment files (.env)
- IDE files (.vscode, .idea)
- OS files (.DS_Store, Thumbs.db)
- Logs and temporary files

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

### Development Setup

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/feature_name`)
3. Commit your changes (`git commit -m 'add new feature'`)
4. Push to the branch (`git push origin feature/feature_name`)
5. Open a Pull Request

## License

MIT License - see [LICENSE](LICENSE) file for details.

## Acknowledgments

- [cobra](https://github.com/spf13/cobra)
- [gopkg.in/yaml.v2](https://github.com/go-yaml/yaml)
- [fatih/color](https://github.com/fatih/color)
