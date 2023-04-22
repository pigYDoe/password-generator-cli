# ⚿ password-generator-cli

 [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/pigYDoe/password-generator-cli/blob/main/LICENSE)

 It's a command-line program designed to create strong, customizable passwords, enhancing account security. With options for password length and character exclusion.

## Usage
```
pwgen.exe [options]
```
Options:
```
-L password lenght (default 7)

-d exclude digits

-lc exclude lowercase letters

-sc exclude special characters

-uc exclude uppercase letters
```
## How to compile statically?

Windows:
```
$env:CGO_ENABLED = "0"
$env:GOOS = "windows"
$env:GOARCH = "amd64"

go build -a -ldflags '-extldflags "-static"' -o output_binary_name.exe your_main_package.go
```
Linux:
```
export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64

go build -a -ldflags '-extldflags "-static"' -o output_binary_nameyour_main_package.go
```
