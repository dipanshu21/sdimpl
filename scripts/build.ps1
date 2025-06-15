param(
    [Parameter(Position=0)]
    [string]$ProjectName
)

function Get-ProjectType {
    param([string]$ProjectPath)
    
    if (Test-Path "$ProjectPath/package.json") {
        return "nodejs"
    }
    elseif (Test-Path "$ProjectPath/requirements.txt" -or Test-Path "$ProjectPath/setup.py" -or Test-Path "$ProjectPath/pyproject.toml") {
        return "python"
    }
    elseif (Test-Path "$ProjectPath/go.mod") {
        return "golang"
    }
    elseif (Test-Path "$ProjectPath/*.csproj") {
        return "dotnet"
    }
    elseif (Test-Path "$ProjectPath/pom.xml" -or Test-Path "$ProjectPath/build.gradle") {
        return "java"
    }
    return "unknown"
}

function Build-Project {
    param(
        [string]$ProjectPath,
        [string]$ProjectType
    )

    Write-Host "Building $ProjectPath ($ProjectType)..."
    
    switch ($ProjectType) {
        "nodejs" {
            Push-Location $ProjectPath
            npm install
            npm run build
            Pop-Location
        }
        "python" {
            Push-Location $ProjectPath
            python -m pip install -r requirements.txt
            if (Test-Path "setup.py") {
                python setup.py build
            }
            Pop-Location
        }
        "golang" {
            Push-Location $ProjectPath
            go build ./...
            Pop-Location
        }
        "dotnet" {
            Push-Location $ProjectPath
            dotnet build
            Pop-Location
        }
        "java" {
            if (Test-Path "$ProjectPath/pom.xml") {
                Push-Location $ProjectPath
                mvn clean install
                Pop-Location
            }
            elseif (Test-Path "$ProjectPath/build.gradle") {
                Push-Location $ProjectPath
                gradle build
                Pop-Location
            }
        }
        default {
            Write-Host "Unknown project type for $ProjectPath"
        }
    }
}

# Create directories if they don't exist
$directories = @("projects", "shared/lib")
foreach ($dir in $directories) {
    if (-not (Test-Path $dir)) {
        New-Item -ItemType Directory -Path $dir
        Write-Host "Created directory: $dir"
    }
}

if ($ProjectName) {
    $projectPath = "projects/$ProjectName"
    if (Test-Path $projectPath) {
        $projectType = Get-ProjectType $projectPath
        Build-Project $projectPath $projectType
    }
    else {
        Write-Host "Project '$ProjectName' not found in projects directory"
        exit 1
    }
}
else {
    # Build all projects
    Get-ChildItem -Path "projects" -Directory | ForEach-Object {
        $projectPath = $_.FullName
        $projectType = Get-ProjectType $projectPath
        Build-Project $projectPath $projectType
    }
}
