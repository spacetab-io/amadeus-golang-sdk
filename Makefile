deps:
	go mod vendor
.PHONY: deps

lint:
	golint $(PKGS)
.PHONY: lint

test-unit:
	go test ./... --race --cover -count=1 -timeout 1s -coverprofile=c.out -v
.PHONY: test-unit

coverage-html:
	go tool cover -html=c.out -o coverage.html

test: deps test-unit coverage-html
.PHONY: test