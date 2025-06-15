# Simple Development Projects

A multi-language development workspace for creating and managing multiple projects. This repository provides a simple structure for organizing projects in different programming languages with shared components.

## Structure

```
.
├── projects/           # All projects live here
│   ├── project1/      # Individual project folders
│   └── project2/
├── shared/            # Shared components
│   └── lib/          # Reusable libraries/modules
└── scripts/           # Build system scripts
    ├── build.ps1     # Build projects
    ├── run.ps1       # Run projects
    └── test.ps1      # Run tests
```

## Usage

### Building Projects

```powershell
# Build all projects
.\scripts\build.ps1

# Build specific project
.\scripts\build.ps1 project1
```

### Running Projects

```powershell
# Run a specific project
.\scripts\run.ps1 project1
```

### Testing Projects

```powershell
# Test all projects
.\scripts\test.ps1

# Test specific project
.\scripts\test.ps1 project1
```

## Adding New Projects

1. Create a new directory in `projects/`:
   ```
   projects/
   └── my-new-project/
       ├── src/
       └── tests/
   ```

2. Add appropriate project files based on your language:
   - Node.js: `package.json`
   - Python: `requirements.txt` or `setup.py`
   - Go: `go.mod`
   - .NET: `.csproj`
   - Java: `pom.xml` or `build.gradle`

## Sharing Code Between Projects

### Node.js Projects
```json
{
  "dependencies": {
    "shared-auth": "file:../../shared/lib/auth"
  }
}
```

### Python Projects
```
# requirements.txt
-e ../../shared/lib/data-access
```

### Go Projects
```go
// go.mod
replace example.com/shared/cache => ../../shared/lib/cache
```

## Supported Languages

- Node.js/TypeScript
- Python
- Go
- .NET
- Java

The build system automatically detects project types based on the presence of standard configuration files (package.json, requirements.txt, go.mod, etc.).
