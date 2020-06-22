GoIngot
==

[![Codeship Status for devbackend/goingot](https://app.codeship.com/projects/40f57af0-96d1-0138-4c88-0a7d12f2e6c3/status?branch=master)](https://app.codeship.com/projects/400651)

Scaffolding for Go projects.

Follow go project layout: https://github.com/golang-standards/project-layout

Scripts
--

```shell script
# Install golangci-lint
make lint-install

# Run linters
make lint 

# Generate mocks for testing
make mock-generate

# Run tests
make test

# Run stdout code coverage
make test-coverage

# Run HTML code coverage (required web browser)
make test-coverage-html

# Build Dockder image
make docker-build

# Run Docker image
make docker-run
```