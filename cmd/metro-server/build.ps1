#
# USER VARIABLES
#
$imageRef = 'lokomotes/metro-server:latest-windows-amd64'
$myPkg = 'github.com/lokomotes/metro/cmd/metro-server'

#
# DO NOT TOUCH
#
function New-TemporaryDirectory {
    $parent = [System.IO.Path]::GetTempPath()
    $name = [System.IO.Path]::GetRandomFileName()
    New-Item -ItemType Directory -Path (Join-Path $parent $name)
}

$goPath = go env GOPATH

if (-Not $(Test-Path $(Join-Path $PSScriptRoot './windows.dockerfile') -PathType Leaf)) {
    Write-Error "Dockerfile is not provided"
    exit
}

$tmpPath = New-TemporaryDirectory

[String[]]$deps = @()
([String]$(go list -f '{{.Deps}}') -replace '^.|.$').Split(' ') | ForEach-Object {
    $nStd = [String]$(go list -f '{{if not .Standard}}{{.ImportPath}}{{end}}' $_)
    if ([String]::IsNullOrWhiteSpace($nStd)) { return }
    $deps += $nStd
}
$deps += $myPkg

foreach ($dep in $deps) {
    $src = [IO.Path]::Combine($goPath, "src", $dep, "*")
    $dst = [IO.Path]::Combine($tmpPath, "src", $dep)
    if ($(Test-Path $dst)) {
        continue
    }
    New-Item $dst -ItemType Directory | Out-Null
    Copy-Item $src $dst -Recurse
    Write-Host "$dep done"
}

Copy-Item ./windows.dockerfile $(Join-Path $tmpPath ./Dockerfile)

docker build `
    -t $imageRef `
    $tmpPath

Remove-Item -Recurse -Force $tmpPath
