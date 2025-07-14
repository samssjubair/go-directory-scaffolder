# Directory Scaffolder

A CLI tool written in Go that creates folder structures and files based on YAML templates.

## Features

- ✅ Create folder structures from YAML templates
- ✅ Support for nested folders
- ✅ Dry-run mode to preview changes
- ✅ Colorful output with emojis
- ✅ Built-in templates included
- ✅ Built-in template selection

## Installation

```bash
go mod tidy
go build -o scaffold .
```

## Usage

### Basic Usage

```bash
# Create structure from a YAML template
./scaffold project.yaml

# Preview what would be created (dry-run)
./scaffold --dry-run project.yaml
```

### Built-in Templates

```bash
# List all available built-in templates
./scaffold --list-templates

# Use a built-in template
./scaffold --template react-app
./scaffold --template go-api
./scaffold --template node-express

# Use built-in template with custom project name
./scaffold --template react-app my-custom-app

# Preview built-in template
./scaffold --template go-api --dry-run
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
│   └── root.go
├── internal/
│   ├── parser.go       # load & parse YAML
│   ├── creator.go      # create dirs & files
│   └── templates.go    # built-in templates
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

# Run tests
go test ./...

# Build
go build -o scaffold .

# Run
./scaffold templates/react-app.yaml
./scaffold --template react-app
```
