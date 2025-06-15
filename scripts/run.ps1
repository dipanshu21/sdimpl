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

function Run-Project {
    param(
        [string]$ProjectPath,
        [string]$ProjectType
    )

    Write-Host "Running $ProjectPath ($ProjectType)..."
    
    switch ($ProjectType) {
        "nodejs" {
            Push-Location $ProjectPath
            npm start
            Pop-Location
        }
        "python" {
            Push-Location $ProjectPath
            if (Test-Path "main.py") {
                python main.py
            }
            elseif (Test-Path "app.py") {
                python app.py
            }
            else {
                Write-Host "No main Python file found (tried main.py, app.py)"
            }
            Pop-Location
        }
        "golang" {
            Push-Location $ProjectPath
            go run .
            Pop-Location
        }
        "dotnet" {
            Push-Location $ProjectPath
            dotnet run
            Pop-Location
        }
        "java" {
            if (Test-Path "$ProjectPath/target") {
                $jarFile = Get-ChildItem -Path "$ProjectPath/target" -Filter "*-with-dependencies.jar" | Select-Object -First 1
                if ($jarFile) {
                    java -jar $jarFile.FullName
                }
                else {
                    Write-Host "No executable JAR found in target directory"
                }
            }
            elseif (Test-Path "$ProjectPath/build/libs") {
                $jarFile = Get-ChildItem -Path "$ProjectPath/build/libs" -Filter "*.jar" | Select-Object -First 1
                if ($jarFile) {
                    java -jar $jarFile.FullName
                }
                else {
                    Write-Host "No executable JAR found in build/libs directory"
                }
            }
        }
        default {
            Write-Host "Unknown project type for $ProjectPath"
        }
    }
}

if (-not $ProjectName) {
    Write-Host "Please specify a project name"
    exit 1
}

$projectPath = "projects/$ProjectName"
if (Test-Path $projectPath) {
    $projectType = Get-ProjectType $projectPath
    Run-Project $projectPath $projectType
}
else {
    Write-Host "Project '$ProjectName' not found in projects directory"
    exit 1
}
