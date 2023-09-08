set positional-arguments

[private]
alias align := check-structalign

_default:
  @just --list

# Run golangci-lint for all packages
lint *ARGS:
  golangci-lint run {{ARGS}}

# Generate all Go related APIs and files
gen:
  rm -rf licenseapi
  mkdir -p licenseapi

  cp $(find . -type f -name "*.go") licenseapi

  go run k8s.io/code-generator/cmd/deepcopy-gen@v0.28.1 \
    --go-header-file ./hack/boilerplate.go.txt \
    --input-dirs ./licenseapi \
    -O zz_generated.deepcopy

  cp licenseapi/zz_generated.deepcopy.go .

  rm -rf licenseapi

# Check struct memory alignment and print potential improvements
[no-exit-message]
check-structalign *ARGS:
  go run github.com/dkorunic/betteralign/cmd/betteralign@latest {{ARGS}} ./...
