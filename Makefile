BINARY_NAME=ybeliche_telegrambot
PKG=./

version ?= latest
username ?= ybeliche
gh_token ?=

GOLANGCI_CMD=golangci-lint run
REVIVE_CMD=revive

INSTALL_GOLANGCI=go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
INSTALL_REVIVE=go install github.com/mgechev/revive@latest

.PHONY: all
all: build

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME) main.go

.PHONY: docker-build
docker-build:
	docker build -t $(BINARY_NAME) $(PKG)

.PHONY: login-github
login-github:
	echo $(gh_token) | docker login ghcr.io -u $(username) --password-stdin

.PHONY: push-github
push-github: login-github docker-build
	docker tag $(BINARY_NAME) ghcr.io/ybeliche/$(BINARY_NAME):$(version)
	docker push ghcr.io/ybeliche/$(BINARY_NAME):$(version)

.PHONY: run
run:
	go run main.go

.PHONY: lint
lint: lint-golangci lint-revive

.PHONY: lint-golangci
lint-golangci: check-golangci
	$(GOLANGCI_CMD)

.PHONY: lint-revive
lint-revive: check-revive
	$(REVIVE_CMD)

.PHONY: clean
clean:
	rm -f $(BINARY_NAME)

.PHONY: check-golangci
check-golangci:
	@if ! command -v golangci-lint >/dev/null 2>&1; then \
		echo "golangci-lint not found, installing..."; \
		$(INSTALL_GOLANGCI); \
		echo "golangci-lint installed."; \
	else \
		echo "golangci-lint is already installed."; \
	fi

.PHONY: check-revive
check-revive:
	@if ! command -v revive >/dev/null 2>&1; then \
		echo "revive not found, installing..."; \
		$(INSTALL_REVIVE); \
		echo "revive installed."; \
	else \
		echo "revive is already installed."; \
	fi
