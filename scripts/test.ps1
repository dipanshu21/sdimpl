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

function Test-Project {
    param(
        [string]$ProjectPath,
        [string]$ProjectType
    )

    Write-Host "Testing $ProjectPath ($ProjectType)..."
    
    switch ($ProjectType) {
        "nodejs" {
            Push-Location $ProjectPath
            npm test
            Pop-Location
        }
        "python" {
            Push-Location $ProjectPath
            if (Test-Path "pytest.ini" -or (Test-Path "tests/")) {
                pytest
            }
            elseif (Test-Path "setup.py") {
                python setup.py test
            }
            else {
                python -m unittest discover
            }
            Pop-Location
        }
        "golang" {
            Push-Location $ProjectPath
            go test ./...
            Pop-Location
        }
        "dotnet" {
            Push-Location $ProjectPath
            dotnet test
            Pop-Location
        }
        "java" {
            if (Test-Path "$ProjectPath/pom.xml") {
                Push-Location $ProjectPath
                mvn test
                Pop-Location
            }
            elseif (Test-Path "$ProjectPath/build.gradle") {
                Push-Location $ProjectPath
                gradle test
                Pop-Location
            }
        }
        default {
            Write-Host "Unknown project type for $ProjectPath"
        }
    }
}

if ($ProjectName) {
    $projectPath = "projects/$ProjectName"
    if (Test-Path $projectPath) {
        $projectType = Get-ProjectType $projectPath
        Test-Project $projectPath $projectType
    }
    else {
        Write-Host "Project '$ProjectName' not found in projects directory"
        exit 1
    }
}
else {
    # Test all projects
    Get-ChildItem -Path "projects" -Directory | ForEach-Object {
        $projectPath = $_.FullName
        $projectType = Get-ProjectType $projectPath
        Test-Project $projectPath $projectType
    }
}
