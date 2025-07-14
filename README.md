# Directory Scaffolder

A CLI tool written in Go that creates folder structures and files based on YAML templates.

## Features

- ✅ Create folder structures from YAML templates
- ✅ Support for nested folders
- ✅ Dry-run mode to preview changes
- ✅ Colorful output with emojis
- ✅ Built-in templates included

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

### Built-in Templates

The project includes several sample templates in the `templates/` directory:

- `react-app.yaml` - React application structure
- `go-api.yaml` - Go API project structure
- `node-express.yaml` - Node.js Express application

## Project Structure

```
scaffold/
├── cmd/                # cobra CLI commands
│   └── root.go
├── internal/
│   ├── parser.go       # load & parse YAML
│   └── creator.go      # create dirs & files
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
```

## License

MIT 