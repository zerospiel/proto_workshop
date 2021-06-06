.PHONY: lint
lint:
	find . -name "*.go" -type file -mindepth 2 -maxdepth 2 -exec dirname {} \; | xargs -n1 -I {} -- bash -c 'cd {} && golangci-lint run --verbose ./...'

.PHONY: deps
deps:
	find . -name go.mod -exec bash -c 'pushd "$${1%go.mod}" && go mod tidy && popd' _ {} \;