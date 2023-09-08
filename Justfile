set positional-arguments

[private]
alias outdated := list-outdated
[private]
alias align := check-structalign

alias gen-go := generate-go-apis

_default:
  @just --list

# Run golangci-lint for all packages
lint *ARGS:
  golangci-lint run {{ARGS}}

# Generate all Go related APIs and files
generate-go-apis:
  go run k8s.io/code-generator/cmd/deepcopy-gen@v0.28.1 \
    --go-header-file ./hack/boilerplate.go.txt \
    --input-dirs ./feature,./instance,./instance/button,./auth \
    -O zz_generated.deepcopy \
    -o .

# Update a specific dependency or all
[no-cd]
update dep:
  GOPRIVATE=github.com/loft-sh/* go run github.com/icholy/gomajor@latest get {{dep}}

# Upgrade a specific dependency
[no-cd]
upgrade dep:
  GOPRIVATE=github.com/loft-sh/* go run github.com/marwan-at-work/mod/cmd/mod@atest upgrade --mod-name={{dep}}

# List outdated dependencies for the current module
[no-cd]
list-outdated:
  go run github.com/icholy/gomajor@latest list

# Check struct memory alignment and print potential improvements
[no-exit-message]
check-structalign *ARGS:
  go run github.com/dkorunic/betteralign/cmd/betteralign@latest {{ARGS}} ./...
