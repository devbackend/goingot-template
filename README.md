GoIngot
==

[![pipeline status](https://gitlab.com/devbackend/goingot/badges/master/pipeline.svg)](https://gitlab.com/devbackend/goingot/-/pipelines)

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